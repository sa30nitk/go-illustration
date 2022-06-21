package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Printf("reading config file failed with error: %s\n", err)
		}
		return
	}

	fmt.Println("config file loaded")
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.AllSettings())
}
