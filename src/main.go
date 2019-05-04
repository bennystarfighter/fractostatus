package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	db "github.com/AlmightyFloppyFish/sfsdb-go"
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
	//var err error
	var print bool
	flag.BoolVar(&print, "p", false, "")
	flag.Parse()
	if print {
		err := s.printRun()
		if err != nil {
			log.Fatal(err)
			return
		}
		os.Exit(0)
	}
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
		fmt.Println("nothing!")
	}
}
