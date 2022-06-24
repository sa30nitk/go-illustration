package config

import (
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Load() error {
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("resources")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("config file not found")
		} else {
			return errors.New(fmt.Sprintf("reading config file failed with error: %s\n", err))
		}
	}

	// for transforming app.host to app_host
	repl := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(repl)
	return nil
}
