package loghelper

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logFields = log.Fields{"Client": "Subscriber"}

// Init - to make one time initial setup for logrus
func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// LogInfo logs a message at level Info.
func LogInfo(args ...interface{}) {
	log.WithFields(logFields).Info(args...)
}

// LogWarn logs a message at level Warn.
func LogWarn(args ...interface{}) {
	log.WithFields(logFields).Warn(args...)
}

// LogError logs a message at level Error.
func LogError(args ...interface{}) {
	log.WithFields(logFields).Error(args...)
}

// LogFatal logs a message at level Fatal.
func LogFatal(args ...interface{}) {
	log.WithFields(logFields).Info(args...)
}
