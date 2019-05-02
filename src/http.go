package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Server struct {
	address  string
	password string
}

type Content struct {
	lastalive    time.Time
	hostname     string
	processwatch []process
}

type process struct {
	alive bool
	name  string
	pid   int
}

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
	for i := 0; i < len(processes); i++ {
		cmd := `ps -e | grep ` + processes[i]
		_, err := runBashCommand(cmd)
		if err != nil {
			return err
		}
		fmt.Println(content.hostname, content.lastalive.String(), content.processwatch)
	}

	return nil
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
