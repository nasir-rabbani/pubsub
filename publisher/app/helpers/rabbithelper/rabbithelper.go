package rabbithelper

import "github.com/streadway/amqp"

// CreateNewQueue for message publishing
func CreateNewQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}
