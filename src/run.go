package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/AlmightyFloppyFish/sfsdb-go"
	"github.com/fatih/color"
)

type PrintContent struct {
	Identifier string
	Hostname   string
	Alive      bool
	Lastalive  time.Time
	Processes  []Process
}

func (s *State) clientRun() error {
	for {
		time.Sleep(time.Duration(int64(s.pollrate)) * time.Second)
		content, err := s.prepData(s.processlist)
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
	}
}

func (s *State) serverRun() error {
	s.localDB = db.New("db", 10, 0)
	clientListExist := s.localDB.Exists("client-list")
	// sfsdb already creates if they don't exist anyway. 
	if !clientListExist {
		s.localDB.Save("client-list", []Content{})
	}
	http.HandleFunc("/", s.httpHandleIncomingData)
	if s.TLSmode {
		return http.ListenAndServeTLS(":"+strconv.Itoa(s.port), s.certFilePath, s.keyFilePath, nil)
	} else {
		return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
	}
}

func (s *State) printerRun() error {
	err := s.initConfig()
	if err != nil {
		return err
	}
	s.localDB = db.New("db", 10, 0)
	s.localDB.Load("client-list", &s.clients)
	var clientsData []Content
	for i := range s.clients {
		var clientData Content
		err := s.localDB.Load(s.clients[i], &clientData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		clientsData = append(clientsData, clientData)
	}
	//fmt.Println(clientsData)
	var p PrintContent
	for _, cd := range clientsData {
		p.Lastalive = cd.Lastalive
		p.Identifier = cd.Identifier
		p.Hostname = cd.Hostname
		p.Processes = cd.Processes
		//fmt.Println(clientsData[i].Processes, p.Processes)
		if (p.Lastalive.Unix() + s.aliveTimeout) > time.Now().Unix() {
			p.Alive = true
		} else {
			p.Alive = false
		}
		p.Print()
		p = PrintContent{}
	}
	return nil
}

func (p *PrintContent) Print() {
	w := color.New(color.FgWhite)
	g := color.New(color.FgHiGreen)
	r := color.New(color.FgRed)
	if p.Alive {
		g.Print(p.Identifier)
		w.Print(`/`)
		g.Print(p.Hostname)
		w.Print(` | `)
	} else {
		r.Print(p.Identifier)
		w.Print(`/`)
		r.Print(p.Hostname)
		w.Print(` | `)
	}
	for _, process := range p.Processes {
		if process.Running {
			g.Print(process.Name)
			w.Print(`,`)
		} else {
			r.Print(process.Name)
			w.Print(`,`)
		}
	}
	fmt.Println()
}
