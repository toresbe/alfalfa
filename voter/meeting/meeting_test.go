package meeting

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	Meetings = make(map[uuid.UUID]Meeting)
	code := m.Run()
	os.Exit(code)
}

func TestCreateMeeting(t *testing.T) {
	if len(Meetings) != 0 {
		t.Error("Assumed empty meetings list, but not the case")
	}
	UUID := Create("Hello", []int{})
	if len(Meetings) != 1 {
		t.Errorf(`Meeting %s is not in Meetings list`, UUID)
	}
}