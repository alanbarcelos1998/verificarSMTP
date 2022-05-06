package smtp_test

import (
	"fmt"
	"smtp-api/smtp"
	"testing"
)

type enderecoSmtpParaTeste struct {
	enderecos string
}

func TestVerificaSmtp(t *testing.T) {
	enderecoSmtp := []enderecoSmtpParaTeste{
		{"smtp-mail.outlook.com:587"},
		{"smtp.mail.yahoo.com:587"},
	}

	for _, enderecos := range enderecoSmtp {
		res, err := smtp.VerificaSmtp(enderecos.enderecos)
		if err != nil {
			t.Errorf("O erro recebido: %s\n", err)
		}
		fmt.Println(res)
	}
}
