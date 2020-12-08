package rabbithelper

import "github.com/streadway/amqp"

// CreateNewConsumer creates new consumer with queue name
func CreateNewConsumer(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	ctx, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		return nil, err
	}

	return ctx, nil
}

// CreateNewQueue for message storing
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
