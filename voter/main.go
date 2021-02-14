package main

import (
    log "github.com/sirupsen/logrus"

    "time"
    "./api"
)

type Attendant int

type Vote struct {
    choice      int
    msisdn      int
    timestamp   time.Time
}

type Decision struct {
    votes       []Vote
    description string
    options     map[int]string
    started     time.Time
    ended       time.Time
}

func (d *Decision) GetConclusion() {
}

func (d *Decision) AddVote(v *Vote) {
}

func (d *Decision) GetStatus() {
}



func main() {
    log.SetFormatter(&log.TextFormatter{})

    a := api.App{}
    a.Initialize()
    a.Run(":10000")
}