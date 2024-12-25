package services

import (
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	"context"
	"errors"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMail(ctx *context.Context, to string) error {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = env.Get(constants.OWNER_MAIL)
	server.Password = env.Get(constants.GMAIL_PASS)
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return errors.New("error while connecting")
	}

	email := mail.NewMSG()
	email.SetFrom(env.Get(constants.OWNER_MAIL))
	email.AddTo(to)
	//email.AddCc([]string("another_you@example.com"))
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextPlain, "hello guys ,i am with ur ")
	if er := email.Send(smtpClient); er != nil {
		return errors.New("send mail failed")
	}
	return nil
}
