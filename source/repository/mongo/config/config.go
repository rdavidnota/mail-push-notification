package config

import (
	"fmt"
	"github.com/rdavidnota/mail-push-notification/source/configuration"
)

func GetUrl() string {
	mongodb := configuration.GetConfiguration().Mongodb
	url := fmt.Sprintf("mongodb://%s:%s", mongodb.Host, mongodb.Port)

	return url
}
