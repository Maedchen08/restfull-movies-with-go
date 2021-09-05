package config

import (
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var appConfig *Config

type Config struct {
	AppName     string
	AppPort     string
	LogLevel    string
	Environment string
}

func Init() *Config {
	defaultEnv := ".env"

	err := gotenv.Load(defaultEnv)
	if err != nil {
		log.Fatalf("Could not load environment")
	}

	log.SetOutput(os.Stdout)

	appConfig = &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetString("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
	}

	return appConfig

}
