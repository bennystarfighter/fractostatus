package main

import (
	"errors"
	"net/http"
	"strconv"
)

type Server struct {
	address  string
	password string
}

func httpSendToServer(input []byte, se Server) error {
	r, err := http.Post(se.address, "application/x-gob", nil)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New(strconv.Itoa(r.StatusCode) + r.Status)
	}
	return nil
}

func httpHandleIncomingData(w http.ResponseWriter, r *http.Request) {
}
