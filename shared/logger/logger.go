package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type (
	// MywishesLogger return implementation logger methods
	MywishesLogger interface {
		Info(args ...interface{})
		Infof(format string, args ...interface{})
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
	}

	// MyLogger return object implementation of MywishesLogger interface
	MyLogger struct {
		MywishesLogger
	}
)

var (
	log            *logrus.Logger
	once           sync.Once
	instanceLogger *MyLogger
)

// Info is a logrus log message at level info on the standard logger
func (q *MyLogger) Info(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"prefix": "go-ses",
	}).Info(args...)
}

// Infof is a logrus log message at level info with style format on the standard logger
func (q *MyLogger) Infof(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"prefix": "go-ses",
	}).Infof(format, args...)
}

// Error is a logrus log message at level Error on the standard logger
func (q *MyLogger) Error(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"prefix": "go-ses",
	}).Error(args...)
}

// Errorf is a logrus log message at level error with style format on the standard logger
func (q *MyLogger) Errorf(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"prefix": "go-ses",
	}).Errorf(format, args...)
}

// NewMywishesLogger is a factory return implementation object methods in logger
func NewMywishesLogger() *MyLogger {
	once.Do(func() {
		log = logrus.New()
		log.Formatter = &prefixed.TextFormatter{
			FullTimestamp: true,
		}
		instanceLogger = &MyLogger{}
	})

	return instanceLogger
}
