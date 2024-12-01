package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		// Канал для завершения обработки
		finished := make(chan struct{})
		panicChan := make(chan interface{})

		// Обработка запроса в горутине
		go func() {
			defer func() {
				if r := recover(); r != nil {
					panicChan <- r
				}
			}()
			c.Next() // Выполняем основной обработчик
			close(finished)
		}()

		select {
		case <-ctx.Done():
			// Если таймаут истёк
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"error": "Request timed out",
			})
		case <-finished:
			// Если обработка завершилась до истечения таймаута
		case p := <-panicChan:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Internal server error: %v", p),
			})
		}
	}
}
