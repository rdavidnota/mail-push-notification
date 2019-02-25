package mongo

import (
	"context"
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/entities"
	"github.com/rdavidnota/mail-push-notification/source/repository/mongo/config"
	"github.com/rdavidnota/mail-push-notification/source/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnection(mail entities.Mail) {

	client, err := mongo.Connect(context.TODO(), config.GetUrl())

	if err != nil {
		utils.FailOnErrorFatal(err, "Error al conectar con el servidor de mongo")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		utils.FailOnErrorFatal(err, "No se tiene conexion con el servidor de mongo")
	}

	fmt.Println("Connected to MongoDB!")
}
