package services

import (
	base "Hackmate/Base"
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	dto "Hackmate/Model/Dto"
	hmerrors "Hackmate/Model/Errors"
	requests "Hackmate/Model/Requests"
	templates "Hackmate/Templates"
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMailService(ctx *context.Context, request *requests.MailRequest) *hmerrors.Bderror {

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
			return err
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

func SendOtpMail(ctx *context.Context, to string, otp int64) error {

	if len(to) == 0 {
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

	emailContent := templates.OtpVerificationTemplate(ctx, otp)

	email.AddTo(to)
	email.SetSubject("Verification Mail")
	email.SetBody(mail.TextPlain, emailContent)

	go email.Send(smtpClient)
	return nil
}

func UploadMedia(ctx *context.Context, request *requests.ImageRequest) (string, *hmerrors.Bderror) {
	cld := base.CloudinaryInstance
	if cld == nil {
		return "", hmerrors.InvalidInputError(ctx, "connection failed with cloudinary", request)
	}

	src, err := request.Image.Open()
	if err != nil {
		return "", hmerrors.InvalidInputError(ctx, err.Error(), request)
	}

	resp, err := cld.Upload.Upload(*ctx, src, uploader.UploadParams{
		PublicID: "profile" + request.Image.Filename,
	})
	if err != nil {
		return "", hmerrors.InvalidInputError(ctx, err.Error(), request)
	}

	return resp.SecureURL, nil
}
