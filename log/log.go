package log

import (
	"github.com/sirupsen/logrus"
)

// Info ...
func Info(eventID, message string) {
	logrus.WithFields(logrus.Fields{
		"eventId": eventID,
	}).Info(message)
}
