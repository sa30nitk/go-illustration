package config

import (
	"strings"

	"github.com/spf13/viper"
)

func getStringWithDefault(key, defaultValue string) string {
	v := viper.GetString(key)
	v = strings.TrimSpace(v)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

func getInt64WithDefault(key string, defaultValue int64) int64 {
	v := viper.GetInt64(key)
	if v != 0 {
		return v
	}
	return defaultValue
}
