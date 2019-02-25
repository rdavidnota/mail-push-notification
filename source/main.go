package main

import (
	"github.com/gorilla/mux"
	"github.com/rdavidnota/mail-push-notification/source/controllers/mail"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/send-mail", mail.SendMail).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":9000", router))
}
