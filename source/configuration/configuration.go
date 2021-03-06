package configuration

import (
	"errors"
	"github.com/rdavidnota/mail-push-notification/source/utils"
	"os"
)

type RabbitMQ struct {
	User         string
	Pass         string
	Host         string
	Port         string
	MailRequest  string
	MailResponse string
}

type MongoDB struct {
	User string
	Pass string
	Host string
	Port string
}

type Configuration struct {
	Rabbitmq RabbitMQ
	Mongodb  MongoDB
}

func validationEnvVar(env_var string) bool {
	if os.Getenv(env_var) == "" {
		return false
	}

	return true
}

func validateEnvVars() error {
	sw := false

	if !validationEnvVar("RABBITMQ_HOST") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_HOST")
		sw = true
	}

	if !validationEnvVar("RABBITMQ_PORT") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_PORT")
		sw = true
	}

	if !validationEnvVar("RABBITMQ_USER") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_USER")
		sw = true
	}

	if !validationEnvVar("RABBITMQ_PASS") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_PASS")
		sw = true
	}

	if !validationEnvVar("RABBITMQ_MAIL_RESPONSE") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_MAIL_RESPONSE")
		sw = true
	}

	if !validationEnvVar("RABBITMQ_MAIL_REQUEST") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada RABBITMQ_MAIL_REQUEST")
		sw = true
	}

	if !validationEnvVar("MONGO_HOST") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada MONGO_HOST")
		sw = true
	}

	if !validationEnvVar("MONGO_PORT") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada MONGO_PORT")
		sw = true
	}

	if !validationEnvVar("MONGO_USER") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada MONGO_USER")
		sw = true
	}

	if !validationEnvVar("MONGO_PASS") {
		err := errors.New("Variable de entorno no configurada")
		utils.FailOnErrorNormal(err, "Variable de entorno no configurada MONGO_PASS")
		sw = true
	}

	if sw {
		err := errors.New("Variable de entorno no configurada")
		return err
	}
	return nil

}

func GetConfiguration() Configuration {
	config := Configuration{}

	err := validateEnvVars()
	utils.FailOnErrorFatal(err, "Variable de entorno no configuradas RabbitMQ")

	config.Rabbitmq.Host = os.Getenv("RABBITMQ_HOST")
	config.Rabbitmq.Port = os.Getenv("RABBITMQ_PORT")
	config.Rabbitmq.User = os.Getenv("RABBITMQ_USER")
	config.Rabbitmq.Pass = os.Getenv("RABBITMQ_PASS")
	config.Rabbitmq.MailRequest = os.Getenv("RABBITMQ_MAIL_REQUEST")
	config.Rabbitmq.MailResponse = os.Getenv("RABBITMQ_MAIL_RESPONSE")

	config.Mongodb.User = os.Getenv("MONGO_PASS")
	config.Mongodb.Pass = os.Getenv("MONGO_USER")
	config.Mongodb.Host = os.Getenv("MONGO_HOST")
	config.Mongodb.Port = os.Getenv("MONGO_PORT")

	return config
}
