package routes

import (
	"encoding/json"
	"net/http"
	smtpVerifica "smtp-api/smtp"

	"github.com/gorilla/mux"
)

func checaSmtpValido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	endereco, _ := vars["endereco"]

	resultado, _ := smtpVerifica.VerificaSmtp(endereco)

	encoder := json.NewEncoder(w)
	encoder.Encode(resultado)
}
