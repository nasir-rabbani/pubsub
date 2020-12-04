package helper

import (
	log "github.com/sirupsen/logrus"
)

// Check - Function to haandle error
func Check(err error, msg string) {
	if err != nil {
		log.WithFields(log.Fields{"Client": "Publisher"}).Error(msg)
		panic(err)
	}

}
