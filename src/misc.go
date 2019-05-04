package main

import (
	"bytes"
	"encoding/gob"
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
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func encodeToGob(in interface{}) ([]byte, error) {
	var b bytes.Buffer
	var err error
	enGob := gob.NewEncoder(&b)
	if err = enGob.Encode(in); err != nil {
		log.Fatal(err)
	}
	return b.Bytes(), err
}

func (s *State) updateClientListDB() error {
	err := s.localDB.Save("client-list", s.clients)
	if err != nil {
		return err
	}
	return nil
}
