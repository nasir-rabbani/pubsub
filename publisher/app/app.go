package app

import (
	"io/ioutil"
	"pub/app/helpers/loghelper"
	"pub/app/helpers/rabbithelper"
	"pub/app/models"
	"pub/app/utils"

	"github.com/streadway/amqp"
)

// Publish - to publish message to queue
func Publish(ch *amqp.Channel) {

	// declare queue
	q, err := rabbithelper.CreateNewQueue(ch, "my-queue")
	utils.HandelException(err, false)

	// reading data for json file
	msg, err := ioutil.ReadFile(models.AppConfig.FilePath)
	if err != nil {
		loghelper.LogError("Error reading input file", err)
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: msg})
	loghelper.LogFatal(err, "Failed to publish a message")
	loghelper.LogInfo("Message Published")
}
