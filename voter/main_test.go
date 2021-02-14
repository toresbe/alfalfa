package main

import (
	"./meeting"
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App
func TestMain(m *testing.M) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected int, res *httptest.ResponseRecorder) {
	if expected != res.Code {
		t.Errorf("Expected response code %d. Got %d\n", expected, res.Code)
	}

	if t.Failed() {
		t.Logf(`Body: "%s"`, res.Body)
	}
}

func checkErr(t *testing.T, desc string, err error) {
	if err != nil {
		t.Errorf("%s: %s", desc, err.Error())
	}
}

func TestCreateMeeting(t *testing.T) {
	var m = meeting.Meeting{Name: "Foo", Attendants: []int{91859508, 22225555}}

	var jsonStr, err = json.Marshal(m)
	checkErr(t, "marshaling meeting", err)

	req, _ := http.NewRequest("POST", "/api/v1/meetings", bytes.NewBuffer(jsonStr))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response)
}