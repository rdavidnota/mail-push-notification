package mail

import (
	"github.com/rdavidnota/mail-push-notification/source/commands/mail"
	"github.com/rdavidnota/mail-push-notification/source/controllers/authorization"
	"net/http"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	if authorization.Authorization(w, r) {
		w.Header().Set("Content-Type", "application/json")

		to := r.FormValue("to")
		cc := r.FormValue("cc")
		message := r.FormValue("message")

		if !mail.SaveMail(to, cc, message) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		w.WriteHeader(http.StatusOK)
	}
}
