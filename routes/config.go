package routes

import "github.com/gorilla/mux"

func Config(router *mux.Router) {
	router.HandleFunc("/checksmtp/{endereco}", checaSmtpValido).Methods("GET")
}
