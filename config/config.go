package config

import "github.com/spf13/viper"

type Config struct {
	App    App
	NR     NewRelic
	StatsD StatsD
}

func NewConfig() Config {
	return Config{
		App: NewApp(),
		NR: NewRelic{
			License: getStringWithDefault("newrelic.license", ""),
			App:     getStringWithDefault("newrelic.app.name", ""),
			Enabled: viper.GetBool("newrelic.enabled")},
		StatsD: StatsD{
			Host:   getStringWithDefault("statsd.host", ""),
			Prefix: getStringWithDefault("statsd.Prefix", "gl"),
		},
	}
}
