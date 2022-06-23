package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load() {
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("resources")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Printf("reading config file failed with error: %s\n", err)
		}
		return
	}

	// for transforming app.host to app_host
	repl := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(repl)

	fmt.Println("config file loaded")
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.AllSettings())
}
