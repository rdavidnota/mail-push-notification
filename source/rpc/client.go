package rpc

import (
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/rpc/config"
	"github.com/rdavidnota/mail-push-notification/source/utils"
	"github.com/streadway/amqp"
)

func mailRPC() (err error) {
	conn, err := amqp.Dial(config.GetUrl())
	utils.FailOnErrorFatal(err, "Conexión fallida con el servidor RabbitMQ")
	defer conn.Close()

	ch, errChannel := conn.Channel()
	utils.FailOnErrorFatal(errChannel, "Error al abrir un canal con el servidor RabbitMQ")
	defer ch.Close()

	_, errQueue := ch.QueueDeclare(
		config.GetMailRequest(), // name
		true,                    // durable
		false,                   // delete when usused
		false,                   // exclusive
		false,                   // noWait
		nil,                     // arguments
	)
	utils.FailOnErrorFatal(errQueue, "Failed to declare a queue")

	messages, errConsume := ch.Consume(
		config.GetMailResponse(), // queue
		"",                       // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	utils.FailOnErrorFatal(errConsume, "Failed to register a consumer")

	corrId := utils.RandomString(32)

	request := "test"

	err = ch.Publish(
		"",                      // exchange
		config.GetMailRequest(), // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       config.GetMailResponse(),
			Body:          []byte(request),
		})
	utils.FailOnErrorFatal(err, "Failed to publish a message")

	for message := range messages {
		fmt.Println(message.ReplyTo)
		fmt.Println(message.CorrelationId)
		if corrId == message.CorrelationId {
			fmt.Printf("Respuesta: %s", message.Body)
			break
		}
	}

	return
}

func client() {

	for i := 0; i < 100; i++ {
		err := mailRPC()
		utils.FailOnErrorNormal(err, "Failed to handle RPC request")
	}

}
