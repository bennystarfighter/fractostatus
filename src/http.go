package main

import (
	"net/http"
	"os"
	"os/exec"
	"time"

	ps "github.com/mitchellh/go-ps"
)

type Server struct {
	address  string
	password string
}

type Content struct {
	lastalive time.Time
	hostname  string
	//	processwatch []process
}

/*
type process struct {
	alive bool
	name  string
	pid   int
}
*/
type Hardware struct {
	cpuUsage int
	ramUsage int
}

func prepareData(processes []string) error {
	var content Content
	var err error
	content.lastalive = time.Now()
	content.hostname, err = os.Hostname()
	if err != nil {
		return err
	}

	output, err := ps.Processes()
	if err != nil {
		return err
	}
	for i := 0; i < len(output); i++ {
		output[i].Executable()
	}
	return nil
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
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

func (se *Server) httpSender(input []byte) {
	http.Post(se.address, "application/x-gob", nil)
}
