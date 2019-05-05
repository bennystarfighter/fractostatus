package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	db "github.com/AlmightyFloppyFish/sfsdb-go"
)

type State struct {
	// Client
	processlist []string
	pollrate    int
	server      Server
	identifier  string

	// Server
	port           int
	clientPassword string
	aliveTimeout   int64
	TLSmode        bool
	certFilePath   string
	keyFilePath    string
	localDB        db.Database

	// Printmode
	clients []string
}

func main() {
	var s State
	err := s.initConfig()
	if err != nil {
		log.Fatal("Config ERROR:", err)
		return
	}
	var clientMode bool
	flag.BoolVar(&clientMode, "client", false, "Start fractostatus in client-mode")
	var serverMode bool
	flag.BoolVar(&serverMode, "server", false, "Start fractostatus in server-mode")
	flag.Parse()

	// If both modes are selected
	if clientMode && serverMode {
		fmt.Println("Cannot start fractostatus in both modes. Choose one.")
		return
		// Printing mode
	} else if !clientMode && !serverMode {
		err := s.printerRun()
		if err != nil {
			log.Fatal(err)
			return
		}
		// Client mode
	} else if clientMode {
		log.Println("Starting Client!")
		err = s.clientRun()
		if err != nil {
			log.Fatal(err)
			return
		}
		// Server mode
	} else if serverMode {
		log.Println("Starting Server!")
		err = s.serverRun()
		if err != nil {
			log.Fatal(err)
			return
		}
		// Unknown
	} else {
		fmt.Println("Runmode is neither print, client or server.")
	}
}
