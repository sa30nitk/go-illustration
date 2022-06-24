package config

type Config struct {
	App App
}

func NewConfig() Config {
	return Config{
		App: NewApp(),
	}
}
