package config

type App struct {
	Port     int64
	LogLevel string
}

func NewApp() App {
	return App{
		Port:     getInt64WithDefault("app.port", int64(8090)),
		LogLevel: getStringWithDefault("app.loglevel", "info"),
	}
}
