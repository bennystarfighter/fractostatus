package main

import (
	"log"
	"os"
)

type State struct {
	mode          bool
	logfile       *os.File
	serveraddress string
	processlist   []string
	pollrate      int
}

func main() {
	var s State
	//var err error
	log.Println("Starting client!")
	err := s.initConfig()
	if err != nil {
		log.Fatal("Config ERROR:", err)
		return
	}
	prepareData([]string{"bash"})
}
