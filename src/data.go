package main

import (
	"os"
	"time"

	ps "github.com/mitchellh/go-ps"
)

type Content struct {
	Lastalive          time.Time
	Hostname           string
	ConfirmedProcesses []ps.Process
}

type Hardware struct {
	cpuUsage int
	ramUsage int
}

func (s *State) sendData() error {
	content, err := prepData(s.processlist)
	if err != nil {
		return err
	}
	encoded, err := encodeToGob(content)
	if err != nil {
		return err
	}
	httpSendToServer(encoded, s.server)
	return nil
}

func prepData(processes []string) (Content, error) {
	var content Content
	var err error
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
	return content, nil
}
