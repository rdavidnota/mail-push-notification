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

	client, errConnect := mongo.Connect(context.TODO(), config.GetUrl())
	utils.FailOnErrorFatal(errConnect, "Error al conectar con el servidor de mongo")

	// Check the connection
	errPing := client.Ping(context.TODO(), nil)
	utils.FailOnErrorFatal(errPing, "No se tiene conexion con el servidor de mongo")

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	insertResult, errInsert := collection.InsertOne(context.TODO(), mail)

	utils.FailOnErrorNormal(errInsert, "Error al insertar el correo")
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	errDisconnect := client.Disconnect(context.TODO())
	utils.FailOnErrorFatal(errDisconnect, "Error al desconectar con el servidor de mongo")

	fmt.Println("Connection to MongoDB closed.")
}
