package utils

import "os"

func GetDynamicHost() string {
	if envHost := os.Getenv("SWAGGER_HOST"); envHost != "" {
		return envHost
	}
	return "localhost:8080" // default
}
