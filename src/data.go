package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Content struct {
	Identifier         string
	Lastalive          time.Time
	Hostname           string
	ConfirmedProcesses []string
	Password           string
}

type Hardware struct {
	cpuUsage int
	ramUsage int
}

func (s *State) prepData(processes []string) (Content, error) {
	var content Content
	var err error
	content.Identifier = s.identifier
	content.Lastalive = time.Now()
	content.Hostname, err = os.Hostname()
	if err != nil {
		return content, err
	}

	out, err := exec.Command("ps", "-e").Output()
	if err != nil {
		return content, err
	}
	for i := range processes {
		process_is_running := strings.Contains(string(out), processes[i])
		if process_is_running {
			content.ConfirmedProcesses = append(content.ConfirmedProcesses, processes[i])
		}
	}
	fmt.Println(content.ConfirmedProcesses)
	content.Password = s.server.password
	return content, nil
}
