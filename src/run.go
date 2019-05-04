package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "github.com/AlmightyFloppyFish/sfsdb-go"
	"github.com/fatih/color"
)

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

func (s *State) printRun() error {
	s.localDB = db.New("db", 10, 0)
	s.localDB.Load("client-list", &s.clients)
	var clientsData []Content
	for i := range s.clients {
		var clientData Content
		err := s.localDB.Load(s.clients[i], &clientData)
		if err != nil {
			continue
		}
	}

	for i := range clientsData {
		lastalive := clientsData[i].Lastalive
		identifier := clientsData[i].Identifier
		hostname := clientsData[i].Hostname
		var processNameList []string
		for d := range clientsData[i].ConfirmedProcesses {
			processNameList = append(processNameList, clientsData[i].ConfirmedProcesses[d].Executable())
		}
		aliveProcesses := strings.Join(processNameList, ",")

		aliveTime := lastalive.Add(time.Duration(int64(s.aliveTimeout)) * time.Second)
		if aliveTime.After(time.Now()) {
			color.Set(color.FgGreen)
		} else {
			color.Set(color.FgRed)
		}
		fmt.Println(`ID: ` + identifier + ` | ` + hostname + `Processes: ` + aliveProcesses)
	}
	return nil
}
