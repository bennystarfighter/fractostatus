package main

import (
	"encoding/gob"
	"errors"
	"io"
	"log"
	"net/http"
)

type Server struct {
	address  string
	password string
}

func httpSendToServer(input io.Reader, se Server) error {
	r, err := http.Post(se.address, "application/x-gob", input)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}
	return nil
}

func (s *State) httpHandleIncomingData(w http.ResponseWriter, r *http.Request) {
	var clientContent Content
	d := gob.NewDecoder(r.Body)
	if err := d.Decode(&clientContent); err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(clientContent)
	if clientContent.Password != s.clientPassword {
		w.WriteHeader(401)
		w.Write([]byte("Access denied, wrong password!"))
		return
	}
	if !Contains(s.clients, clientContent.Identifier) {
		s.clients = append(s.clients, clientContent.Identifier)
	}
	err := s.updateClientListDB()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to update client-list"))
		return
	}
	err = s.localDB.Save(clientContent.Identifier, clientContent)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to update client Content"))
		return
	}
	w.Write([]byte("Success"))
	log.Println("Processed data from: " + clientContent.Identifier)
}
