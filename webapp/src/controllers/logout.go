package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Remove os dados de autentificação salvo nos brower
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)

}
