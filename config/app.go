package config

type App struct {
	Port     int64
	LogLevel string
}

func NewApp() App {
	return App{
		Port:     getInt64WithDefault("App.port", int64(8090)),
		LogLevel: getStringWithDefault("App.loglevel", "info"),
	}
}
