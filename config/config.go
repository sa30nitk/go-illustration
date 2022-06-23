package config

import "github.com/spf13/viper"

type Config struct {
	App App
}

type App struct {
	Port int64
}

func NewConfig() Config {
	return Config{
		App: NewApp(),
	}
}

func NewApp() App {
	return App{Port: viper.GetInt64("app.port")}
}
