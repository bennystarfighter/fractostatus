package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type State struct {
	clientMode bool
	serverMode bool
	// Client
	processlist []string
	pollrate    int
	server      Server
	identifier  string

	// Server
	port           int
	clientPassword string
	TLSmode        bool
	certFilePath   string
	keyFilePath    string
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
	} else if s.serverMode {
		log.Println("Starting Server!")
		err = s.serverRun()
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		err = s.printRun()
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
		time.Sleep(time.Duration(s.pollrate) * time.Second)
	}
}

func (s *State) serverRun() error {
	http.HandleFunc("/", httpHandleIncomingData)
	if s.TLSmode {
		return http.ListenAndServeTLS(":"+strconv.Itoa(s.port), s.certFilePath, s.keyFilePath, nil)
	} else {
		return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
	}
}

func (s *State) printRun() error {

	return nil
}
