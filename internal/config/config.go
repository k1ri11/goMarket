package config

// Config содержит параметры конфигурации приложения.
type Config struct {
	ServerPort string
}

// Load загружает конфигурацию из переменных окружения или устанавливает значения по умолчанию.
func Load() *Config {
	cfg := &Config{
		ServerPort: "8080",
	}

	return cfg
}
