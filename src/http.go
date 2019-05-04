package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
		return errors.New(strconv.Itoa(r.StatusCode) + r.Status)
	}
	return nil
}

func (s *State) httpHandleIncomingData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	var clientContent Content
	d := gob.NewDecoder(r.Body)
	if err := d.Decode(&clientContent); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("2")
	if clientContent.Password != s.clientPassword {
		w.Write([]byte("Access denied, wrong password!"))
		return
	}
	fmt.Println("3")
	if !Contains(s.clients, clientContent.Identifier) {
		s.clients = append(s.clients, clientContent.Identifier)
	}
	fmt.Println("4")
	err := s.updateClientListDB()
	if err != nil {
		w.Write([]byte("Failed to update client-list"))
		return
	}
	fmt.Println("5")
	err = s.localDB.Save(clientContent.Identifier, clientContent)
	if err != nil {
		w.Write([]byte("Failed to update client Content"))
		return
	}
	fmt.Println("6")
	w.Write([]byte("Success"))
}
