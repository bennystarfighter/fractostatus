package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type State struct {
	logfile       *os.File
	serveraddress string
	processlist   []string
	pollrate      int
}

func main() {
	var s State
	//var err error
	log.Println("Starting client!")
	viper.SetConfigName("config")                     // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/fractostatus") // call multiple times to add many search paths
	viper.AddConfigPath(".")                          // optionally look for config in the working directory
	err := viper.ReadInConfig()                       // Find and read the config file
	if err != nil {                                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	s.serveraddress = viper.GetString("server")
	s.processlist = viper.GetStringSlice("processes-watch")
	s.pollrate = viper.GetInt("pollrate")
}
