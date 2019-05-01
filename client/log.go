package main

import (
	"fmt"
	"log"
	"os"
)

func (s *State) initlog() {
	var err error
	s.logfile, err = os.Open("log.txt")
	if err != nil {
		log.Println(err)
		os.Create("log.txt")
		s.logfile, err = os.Open("log.txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
	}
}

func (s *State) writeLog(input string) {
	_, err := s.logfile.WriteString(input)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *State) writeLogPrint(input string) {
	_, err := s.logfile.WriteString(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(input)
}
