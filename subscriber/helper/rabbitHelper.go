package helper

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// ConnectRabbit -
func ConnectRabbit() *amqp.Connection {
	// connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Check(err, "Failed to connect to RabbitMQ")
	log.WithFields(log.Fields{"Client": "Subscriber"}).Info("Connected to RabbitMQ")
	return conn
}

// ReadMessages - To read messages on MQ
func ReadMessages() {
	conn := ConnectRabbit() // initiate rabbitMQ connection
	defer conn.Close()

	// opening channel
	ch, err := conn.Channel()
	Check(err, "Failed to open channel")

	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare("my-queue", false, false, false, false, nil)
	Check(err, "failed to create queue")

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	Check(err, "Error reading mesages from queue")

	go func() {
		for msg := range messages {
			log.WithFields(log.Fields{"Client": "Subscriber"}).Info("Message received")
			ok := SaveMessage(msg.Body)
			if ok {
				log.WithFields(log.Fields{"Client": "Subscriber"}).Info("Message saved")
			}
		}
	}()
	waiting := make(chan bool)
	fmt.Println("The subscriber will wait till exited (CTRL+C)")
	<-waiting
}
