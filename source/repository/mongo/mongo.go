package mongo

import (
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/entities"
)

func TestConnection(mail entities.Mail) {
	/*client, err := mongo.Connect(context.TODO(), "mongodb://db:27017")

	if err != nil {
		//utils.FailOnErrorFatal(err, "Error al conectar con el servidor de mongo")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		//utils.FailOnErrorFatal(err, "No se tiene conexion con el servidor de mongo")
	}*/

	fmt.Println("Connected to MongoDB!")
}
