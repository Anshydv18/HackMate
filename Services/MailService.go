package services

import (
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	dto "Hackmate/Model/Dto"
	requests "Hackmate/Model/Requests"
	templates "Hackmate/Templates"
	"context"
	"errors"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMailService(ctx *context.Context, request *requests.MailRequest) error {

	return SendMail(ctx, &dto.MailInfo{
		To:     request.Mail,
		Header: "HackMate Connect",
	})

}

func SendMail(ctx *context.Context, request *dto.MailInfo) error {

	if len(request.To) == 0 {
		errors.New("enter to addresses")
	}

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

	emailContent := templates.GetCustomisedMessage(ctx, "Ansh", "Ankit", "phone number : 9569816212")

	email.AddTo(request.To...)
	email.AddCc(request.CC...)
	email.AddBcc(request.BCC...)
	email.SetSubject(request.Header)
	email.SetBody(mail.TextPlain, emailContent)

	go email.Send(smtpClient)
	return nil
}
