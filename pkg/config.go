package pkg

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Config содержит параметры конфигурации.
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `yaml:"server"`
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"db"`
}

// LoadConfig загружает конфигурацию из файла.
func LoadConfig(filePath string) (*Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
		return nil, err
	}

	return &config, nil
}

func GetDBUrl(config *Config) string {
	return "host=" + config.DB.Host +
		" user=" + config.DB.User +
		" password=" + config.DB.Password +
		" dbname=" + config.DB.Name +
		" port=" + config.DB.Port +
		" sslmode=disable"

}
