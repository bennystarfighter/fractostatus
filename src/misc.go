package main

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"os/exec"
)

func Contains(a []string, s string) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}
	return false
}

func runBashCommand(command string) (string, error) {
	out, err := exec.Command("bash", "-c", command).Output()
	return string(out), err
}

func encodeToGob(in interface{}) (io.Reader, error) {
	var b bytes.Buffer
	enGob := gob.NewEncoder(&b)
	
	return &b, enGob.Encode(in)
}

func (s *State) updateClientListDB() error {
	return s.localDB.Save("client-list", s.clients)
}
