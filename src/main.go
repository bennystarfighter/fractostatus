package main

import (
	"log"
	"os"
	"time"
)

type State struct {
	clientMode  bool
	logfile     *os.File
	processlist []string
	pollrate    int
	server      Server
}

func main() {
	var s State
	//var err error
	err := s.initConfig()
	if err != nil {
		log.Fatal("Config ERROR:", err)
		return
	}
	if s.clientMode {
		log.Println("Starting Client!")
		err = s.clientRun()
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		log.Println("Starting Server!")
		err = serverRun()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	prepData(s.processlist)
}

func (s *State) clientRun() error {
	for {
		content, err := prepData(s.processlist)
		if err != nil {
			return err
		}
		encoded, err := encodeToGob(content)
		if err != nil {
			return err
		}
		err = httpSendToServer(encoded, s.server)
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(10 * time.Second)
	}
}

func serverRun() error {
	for {

	}
}
