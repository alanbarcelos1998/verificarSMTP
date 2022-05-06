package smtp

import (
	"net/smtp"
)

var resposta string

func VerificaSmtp(enderecoSmtp string) (string, error) {
	// Recebe o endereço smtp e faz a checagem
	c, err := smtp.Dial(enderecoSmtp)
	if err != nil {
		resposta = "Endereço SMTP não existe"
		return resposta, err
	} else {
		resposta = "Endereço SMTP existe"
		return resposta, err
	}

	err = c.Quit()
	if err != nil {
		resposta = "Erro na função"
		return resposta, err
	}

	return "Sem resposta", err
}
