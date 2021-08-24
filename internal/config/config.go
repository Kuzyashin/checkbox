package config

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type AppConfig struct {
	TomTom struct {
		ApiKey string `required:"true"`
	}
	Database struct {
		Username string `required:"false"`
		Password string `required:"false"`
		Host     string `required:"false"`
		Port     int    `required:"false"`
		DBName   string `required:"false" default:"checkbox"`
	}
	Server struct {
		Host string `required:"false" default:"0.0.0.0"`
		Port string `required:"false" default:"3000"`
	}
}

var cfg AppConfig
var once sync.Once

func GetConfig() (*AppConfig, error) {
	var err error
	once.Do(func() { err = envconfig.Process("", &cfg) })
	return &cfg, err
}
