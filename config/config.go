package config

import "github.com/spf13/viper"

type Config struct {
	App App
	NR  NewRelic
}

func NewConfig() Config {
	return Config{
		App: NewApp(),
		NR: NewRelic{
			License: getStringWithDefault("newrelic.license", ""),
			App:     getStringWithDefault("newrelic.app.name", ""),
			Enabled: viper.GetBool("newrelic.enabled")},
	}
}
