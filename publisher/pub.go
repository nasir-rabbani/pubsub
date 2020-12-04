package main

import (
	"os"
	"pub/helper"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{"Client": "Publisher"}).Info("This is the Publisher...")

	filePath := "../input.json"
	ok := helper.Publish(filePath) // Publishing message to MQ

	if ok {
		log.WithFields(log.Fields{"Client": "Publisher"}).Info("Message published")
	}
}
