package main

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

type Content struct {
	Identifier string
	Lastalive  time.Time
	Hostname   string
	Processes  []Process
	Password   string
}

type Process struct {
	Name    string
	running bool
}

func (s *State) prepData(processesWatch []string) (Content, error) {
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
	for i := range processesWatch {
		var process Process
		process.Name = processesWatch[i]
		process_is_running := strings.Contains(string(out), processesWatch[i])
		if process_is_running {
			process.running = true
		} else {
			process.running = false
		}
		content.Processes = append(content.Processes, process)
	}
	content.Password = s.server.password
	return content, nil
}
