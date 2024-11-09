package middleware

import (
	"log"
	"net/http"
	"time"
)

// ResponseWriterWithStatus оборачивает http.ResponseWriter для захвата статус-кода ответа.
type ResponseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader записывает заголовок ответа с указанным статус-кодом.
func (w *ResponseWriterWithStatus) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// RequestLogger логирует метод, путь и статус-код для каждого обработанного запроса.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем ResponseWriter, чтобы отслеживать статус-код.
		rw := &ResponseWriterWithStatus{ResponseWriter: w, statusCode: http.StatusOK}

		// Вызываем следующий обработчик.
		next.ServeHTTP(rw, r)

		// Логируем данные о запросе и ответе.
		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			time.Since(start),
		)
	})
}
