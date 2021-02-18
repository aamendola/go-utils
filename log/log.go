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

// Panic ...
func Panic(eventID string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			logrus.WithFields(logrus.Fields{
				"eventId": eventID,
			}).Error(args...)
		}
		logrus.WithFields(logrus.Fields{
			"eventId": eventID,
		}).Error(err)
	}
}

// Fatal ...
func Fatal(eventID string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			logrus.WithFields(logrus.Fields{
				"eventId": eventID,
			}).Error(args...)
		}
		logrus.WithFields(logrus.Fields{
			"eventId": eventID,
		}).Fatal(err)
	}
}
