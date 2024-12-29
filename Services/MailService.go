package services

import (
	constants "Hackmate/Constants"
	database "Hackmate/Database"
	env "Hackmate/Env"
	dto "Hackmate/Model/Dto"
	requests "Hackmate/Model/Requests"
	templates "Hackmate/Templates"
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMailService(ctx *context.Context, request *requests.MailRequest) error {

	MailDetails := &dto.MailInfo{
		SenderName:     request.SenderName,
		ContactDetails: request.ContactDetails,
		Header:         constants.MailHeader[request.Status],
		To:             request.Mail,
		ReciverName:    request.TeamName,
	}

	if MailDetails.ReciverName == "" {
		UserDetail, err := GetUserByEmail(ctx, request.Mail[0])
		if err != nil {
			return errors.New("email fetching failed")
		}
		MailDetails.ReciverName = UserDetail.Name
	}

	go SendMail(ctx, MailDetails)
	return nil
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

	emailContent := templates.GetCustomisedMessage(ctx, request.ReciverName, request.SenderName, request.ContactDetails)

	email.AddTo(request.To...)
	email.AddCc(request.CC...)
	email.AddBcc(request.BCC...)
	email.SetSubject(request.Header)
	email.SetBody(mail.TextPlain, emailContent)

	go email.Send(smtpClient)
	return nil
}

func UploadPhoto(ctx *context.Context, request *requests.ImageRequest) (string, error) {
	cld, err := database.CloudinaryConnect(ctx)
	if err != nil {
		return "", err
	}
	src, err := request.Image.Open()
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(*ctx, src, uploader.UploadParams{
		PublicID: "profile" + request.Image.Filename,
	})

	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
