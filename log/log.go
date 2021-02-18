package log

import (
	"github.com/sirupsen/logrus"
)

// Info ...
func Info(eventID string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"eventId": eventID,
	}).Info(args...)
}
