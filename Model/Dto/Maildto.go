package dto

type MailInfo struct {
	SenderName     string
	ContactDetails string
	ReciverName    string
	To             []string
	CC             []string
	BCC            []string
	Header         string
	ContentType    string
	Mailbody       string
}
