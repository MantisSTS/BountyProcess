package helpers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQHelper struct{}

func (rmq *RabbitMQHelper) Fetch(channel string, queue string, results chan string) {

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	defer conn.Close()

	// Open a channel
	rmqChan, err := conn.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	defer rmqChan.Close()

	// Declare a queue
	_, err = rmqChan.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Consume messages
	msgs, err := rmqChan.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	for msg := range msgs {
		results <- string(msg.Body)
	}

}

func (rmq *RabbitMQHelper) Publish(channel string, queue string, message string) error {

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	defer conn.Close()

	// Open a channel
	rmqChan, err := conn.Channel()

	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}

	defer rmqChan.Close()

	// Declare a queue
	_, err = rmqChan.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Publish a message
	err = rmqChan.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	return err
}
