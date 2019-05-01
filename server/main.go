package main

import (
	"github.com/spf13/viper"
)

func main() {
	err := initconfig()
	if err != nil {

		return
	}
}

func initconfig() error {
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	err := viper.ReadInConfig()                  // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		return err
	}
	return nil
}
