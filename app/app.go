package app

import (
	"github.com/caarlos0/env"
	"sync"
)

type config struct {
	PostgresDsn   string `env:"POSTGRES_DSN" envDefault:"host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"`
	ClickhouseDsn string `env:"CLICKHOUSE_DSN" envDefault:"http://localhost:8123"`
}

type App struct {
	Config config
}

var once sync.Once
var instance *App

func GetInstance() *App {
	once.Do(func() {
		instance = &App{}
		env.Parse(&instance.Config)
	})
	return instance
}
