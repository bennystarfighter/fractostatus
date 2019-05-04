package main

import (
	"os"
	"time"

	ps "github.com/mitchellh/go-ps"
)

type Content struct {
	Identifier         string
	Lastalive          time.Time
	Hostname           string
	ConfirmedProcesses []ps.Process
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

	output, err := ps.Processes()
	if err != nil {
		return content, err
	}

	for i := range processes {
		outputexec := output[i].Executable()
		isExist := Contains(processes, outputexec)
		if isExist {
			content.ConfirmedProcesses = append(content.ConfirmedProcesses, output[i])
		}
	}
	content.Password = s.server.password
	return content, nil
}
