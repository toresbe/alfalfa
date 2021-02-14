package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
    "time"
    "./meeting"
)

type App struct {
    Router *mux.Router
}

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

func (a *App) meetingCreate(w http.ResponseWriter, r *http.Request) {
    var p meeting.Meeting
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, "JSON decode error: " + err.Error(), http.StatusBadRequest)
        return
    }

    log.WithFields(log.Fields{
        "meetingName": p.Name,
        "attendants": p.Attendants,
    }).Info(`meeting created`)
}

func main() {
    log.SetFormatter(&log.TextFormatter{})

    a := App{}
    a.Initialize()
    a.Run(":10000")
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/api/v1/meetings", a.meetingCreate).Methods("POST")
}

func (a *App) Initialize() {
    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}