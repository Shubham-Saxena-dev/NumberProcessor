package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	EnvConfigs *config
	once       sync.Once
)

type config struct {
	App   app
	Redis redisConfiguration
}

type app struct {
	Environment    string `env:"APP_ENV" envDefault:"dev"`
	AppName        string `env:"APP_NAME" envDefault:"cariad"`
	AppPort        string `env:"APP_PORT" envDefault:"8080"`
	AppHost        string `env:"APP_HOST" envDefault:"localhost"`
	RequestTimeout int    `env:"APP_REQ_TIMEOUT" envDefault:"500"`
}

type redisConfiguration struct {
	Host     string `env:"REDIS_HOST" envDefault:"cache"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASS"`
	DB       int    `env:"REDIS_DB" envDefault:"1"`
}

func InitFromFile(configName string) {
	if EnvConfigs == nil {
		once.Do(
			func() {
				EnvConfigs = loadEnvVariables(configName)
			})
	}
}

func loadEnvVariables(fileName string) *config {

	_, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("File: %v not found.", fileName)
	}

	err = godotenv.Load(fileName)

	if err != nil {
		log.Fatalf("unable to load .env file %v: %v", fileName, err)
	}

	config := config{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Failed to parse env variables:%v", err)
	}
	return &config
}
