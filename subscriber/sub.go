package main

import (
	// "fmt"
	"os"
	"sub/helper"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{"Client": "Subscriber"}).Info("This is the subscriber...")
	helper.ReadMessages()
}
