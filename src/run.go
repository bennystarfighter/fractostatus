package main

import (
	"encoding/gob"
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
		fmt.Println(content)
		encoded, err := encodeToGob(content)
		if err != nil {
			return err
		}
		fmt.Println("*********************")
		var clientContent Content
		d := gob.NewDecoder(encoded)
		if err := d.Decode(&clientContent); err != nil {
			log.Println(err)
			return err
		}

		fmt.Println(clientContent)

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
	for i := range clientsData {
		p.Lastalive = clientsData[i].Lastalive
		p.Identifier = clientsData[i].Identifier
		p.Hostname = clientsData[i].Hostname
		p.Processes = clientsData[i].Processes
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
	g := color.New(color.FgGreen)
	r := color.New(color.FgRed)
	if p.Alive {
		g.Print(p.Identifier + `/` + p.Hostname + ` | `)
	} else {
		g.Print(p.Identifier + `/` + p.Hostname + ` | `)
	}
	for i := range p.Processes {
		if p.Processes[i].running {
			g.Print(p.Processes[i].Name + `,`)
		} else if !p.Processes[i].running {
			r.Print(p.Processes[i].Name + `,`)
		}
	}
	fmt.Println()
}
