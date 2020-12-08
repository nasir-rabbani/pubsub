package app

import (
	"encoding/json"
	"fmt"
	"sub/app/helpers/loghelper"
	"sub/app/helpers/rabbithelper"
	"sub/app/models"
	"sub/app/services"
	"sub/app/utils"

	"github.com/streadway/amqp"
)

// Init -
func Init(ch *amqp.Channel) {

	// declare queue
	queue, err := rabbithelper.CreateNewQueue(ch, "my-queue")
	utils.HandelException(err, false)

	// create new consumer
	messages, err := rabbithelper.CreateNewConsumer(ch, queue.Name)
	utils.HandelException(err, false)

	// Read messages in routines
	go func() {
		for msg := range messages {

			// Unmarshal JSON to struct
			msgData := &models.MsgData{}
			err := json.Unmarshal(msg.Body, msgData)
			if err != nil {
				loghelper.LogError(err)
				continue
			}

			// Store the data to database
			err = services.SaveMsg(msgData)
			if err != nil {
				loghelper.LogError(err)
				continue
			}
		}
	}()

	waiting := make(chan bool)
	fmt.Println("The subscriber will wait till exited (CTRL+C)")
	<-waiting
}
