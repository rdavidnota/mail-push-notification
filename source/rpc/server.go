package main

import (
	"encoding/json"
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/commands/mail"
	"github.com/rdavidnota/mail-push-notification/source/rpc/config"
	"github.com/rdavidnota/mail-push-notification/source/utils"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial(config.GetUrl())
	utils.FailOnErrorFatal(err, "Conexión fallida con el servidor RabbitMQ")
	defer conn.Close()

	ch, errChannel := conn.Channel()
	utils.FailOnErrorFatal(errChannel, "Error al abrir un canal con el servidor RabbitMQ")
	defer ch.Close()

	_, errQueue := ch.QueueDeclare(
		config.GetMailResponse(), // name
		true,                     // durable
		false,                    // delete when usused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	utils.FailOnErrorFatal(errQueue, "Error al declarar la cola de respuesta con el servidor RabbitMQ")

	errQos := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnErrorFatal(errQos, "Error al configurar QoS con el servidor RabbitMQ")

	messages, errConsume := ch.Consume(
		config.GetMailRequest(), // queue
		"",                      // consumer
		false,                   // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	utils.FailOnErrorFatal(errConsume, "Error al registrar un consumidor con el servidor RabbitMQ")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			result := mail.SendMail("", "", "")

			fmt.Println(message.ReplyTo)
			fmt.Println(message.CorrelationId)
			fmt.Printf("Respuesta: %t \n", result)

			response, err := json.Marshal(result)
			utils.FailOnErrorNormal(err, "Error al convertir el resultado a una respuesta valida")

			if err == nil {
				errPublish := ch.Publish(
					"",              // exchange
					message.ReplyTo, // routing key
					false,           // mandatory
					false,           // immediate
					amqp.Publishing{
						ContentType:   "text/plain",
						CorrelationId: message.CorrelationId,
						Body:          response,
					})

				utils.FailOnErrorNormal(errPublish, "Error al publicar un mensaje con el servidor RabbitMQ")
				if errPublish == nil {
					message.Ack(false)
				}
			}

		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
