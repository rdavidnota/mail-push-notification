package config

import (
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/configuration"
)

func GetUrl() string {
	rabbitmq := configuration.GetConfiguration().Rabbitmq
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmq.User, rabbitmq.Pass, rabbitmq.Host, rabbitmq.Port)

	return url
}

func GetMailRequest() string {
	rabbitmq := configuration.GetConfiguration().Rabbitmq
	return rabbitmq.MailRequest
}

func GetMailResponse() string {
	rabbitmq := configuration.GetConfiguration().Rabbitmq
	return rabbitmq.MailResponse
}
