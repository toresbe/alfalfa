package meeting

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Meeting struct {
	Name        string  `json:meetingName`
	Attendants  []int   `json:attendantList`
}

var Meetings map[uuid.UUID]Meeting

func Create(name string, attendants []int) (NewMeetingUUID uuid.UUID) {
	var m = Meeting{name, attendants}
	var UUID = uuid.New()
	Meetings[UUID] = m
	log.WithFields(log.Fields{
		"UUID": UUID,
		"name": name,
		"attendants": attendants,
	}).Info("created meeting")
	return
}