package helper

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// ConnectRabbit -
func ConnectRabbit() *amqp.Connection {
	// connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Check(err, "Failed to connect to RabbitMQ")

	log.WithFields(log.Fields{"Client": "Publisher"}).Info("Connected to RabbitMQ")

	return conn
}

// Publish - To publish message to MQ
func Publish(path string) bool {
	// creating the RabbitMQ connection
	conn := ConnectRabbit()

	defer conn.Close()

	// opening a channel
	ch, err := conn.Channel()
	Check(err, "Failed to open channel")

	defer ch.Close()

	// creating a Queue to publish messages on
	q, err := ch.QueueDeclare("my-queue", false, false, false, false, nil)
	Check(err, "Failed to create Queue")

	// reading data for json file
	msg, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{"Client": "Publisher"}).Error("Error reading input file")
		return false
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: msg})
	Check(err, "Error publishing message on queue.")
	return true
}
