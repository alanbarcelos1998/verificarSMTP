package smtp

import (
	"net/smtp"
	"strings"
)

var resposta string

func VerificaSmtp(enderecoSmtp string) (string, error) {
	portas := [3]string{"587", "465", "25"}

	var err error

	if strings.Contains(enderecoSmtp, portas[0]) || strings.Contains(enderecoSmtp, portas[1]) || strings.Contains(enderecoSmtp, portas[2]) {
		// Recebe o endereço smtp e faz a checagem
		c, err := smtp.Dial(enderecoSmtp)
		if err != nil {
			resposta = "Endereço SMTP não existe"
			err = c.Quit()
			return resposta, err
		} else {
			resposta = "Endereço SMTP existe"
			err = c.Quit()
			return resposta, err
		}
	}

	return "Porta informada não condiz com o padrão", err

}
