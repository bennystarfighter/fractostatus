package main

import (
	"bytes"
	"encoding/gob"
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
	enGob := gob.NewEncoder(&b)
	return b.Bytes(), enGob.Encode(in)
}
