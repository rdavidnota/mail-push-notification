package authorization

import "net/http"

func Authorization(w http.ResponseWriter, r *http.Request) bool {

	username, password, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if username != "rnota" || password != "mercado.nota" {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}
