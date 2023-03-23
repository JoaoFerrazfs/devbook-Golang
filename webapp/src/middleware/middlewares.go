package middleware

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve informações da requisição no terminal
func Logger(proximFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximFuncao(w, r)
	}
}

// Verifica a existencia de cookies
func Autenticacao(proximFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		proximFuncao(w, r)
	}
}
