package mail

import (
	"github.com/rdavidnota/mail-push-notification/source/entities"
	"github.com/rdavidnota/mail-push-notification/source/repository/mongo"
	"github.com/rdavidnota/mail-push-notification/source/utils"
)

func SaveMail(to string, cc string, message string) bool {

	mail := entities.Mail{To: to, Cc: cc, Message: message}
	mail.Id = utils.RandomString(32)

	mongo.TestConnection(mail)

	return true
}

func SendMail(to string, cc string, message string) bool {
	return true
}
