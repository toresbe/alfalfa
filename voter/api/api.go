package api

import (
	"../meeting"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)
type App struct {
	Router *mux.Router
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