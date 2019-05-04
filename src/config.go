package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func (s *State) initConfig() error {
	viper.SetConfigName("config")                     // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/fractostatus") // call multiple times to add many search paths
	viper.AddConfigPath(".")                          // optionally look for config in the working directory
	err := viper.ReadInConfig()                       // Find and read the config file
	if err != nil {                                   // Handle errors reading the config file
		return err
	}
	s.getConfigValues()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		s.getConfigValues()
	})
	return nil
}

func (s *State) getConfigValues() {
	s.clientMode = viper.GetBool("client-mode")
	s.identifier = viper.GetString("identifier")
	s.server.address = viper.GetString("server-address")
	s.server.password = viper.GetString("server-password")
	s.processlist = viper.GetStringSlice("process-watch")
	s.pollrate = viper.GetInt("pollrate")
	// ********
	s.serverMode = viper.GetBool("server-mode")
	s.port = viper.GetInt("port")
	s.aliveTimeout = viper.GetInt64("alive-timeout")
	s.clientPassword = viper.GetString("password")
	s.TLSmode = viper.GetBool("TLS")
	s.certFilePath = viper.GetString("certfile-path")
	s.keyFilePath = viper.GetString("keyfile-path")
	s.clientPassword = viper.GetString("password")
}
