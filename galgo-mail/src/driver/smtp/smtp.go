package smtp

import (
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

type MailMessageSmtp struct {
	To          []string `json:"to[]" example:"test@test.com, test2@test2.com" binding:"required"`
	Cc          []string `json:"cc[]" example:"test@test.com, test2@test2.com"`
	Attachments []string `json:"attachments[]" example:"/path/file/" binding:"required"`
	Subject     string   `json:"subject" example:"testing" binding:"required"`
	From        string   `json:"from" example:"marceloamoreno87@gmail.com" binding:"required"`
	Body        string   `json:"body" example:"<h1>Hello, world!</h1>" binding:"required"`
}

func (MailMessageSmtp *MailMessageSmtp) GetTo() []string {
	return MailMessageSmtp.To
}

func (MailMessageSmtp *MailMessageSmtp) GetCc() []string {
	return MailMessageSmtp.Cc
}

func (MailMessageSmtp *MailMessageSmtp) GetSubject() string {
	return MailMessageSmtp.Subject
}

func (MailMessageSmtp *MailMessageSmtp) GetFrom() string {
	return MailMessageSmtp.From
}

func (MailMessageSmtp *MailMessageSmtp) GetBody() string {
	return MailMessageSmtp.Body
}

func (MailMessageSmtp *MailMessageSmtp) SetTo(To []string) *MailMessageSmtp {
	MailMessageSmtp.To = To
	return MailMessageSmtp
}

func (MailMessageSmtp *MailMessageSmtp) SetCc(Cc []string) *MailMessageSmtp {
	MailMessageSmtp.Cc = Cc
	return MailMessageSmtp
}

func (MailMessageSmtp *MailMessageSmtp) SetSubject(Subject string) *MailMessageSmtp {
	MailMessageSmtp.Subject = Subject
	return MailMessageSmtp
}

func (MailMessageSmtp *MailMessageSmtp) SetFrom(From string) *MailMessageSmtp {
	MailMessageSmtp.From = From
	return MailMessageSmtp
}

func (MailMessageSmtp *MailMessageSmtp) SetBody(Body string) *MailMessageSmtp {
	MailMessageSmtp.Body = Body
	return MailMessageSmtp
}

func (MailMessageSmtp *MailMessageSmtp) Send() (err error) {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", MailMessageSmtp.GetFrom())
	mailer.SetHeader("To", MailMessageSmtp.GetTo()...)
	mailer.SetHeader("Cc", MailMessageSmtp.GetCc()...)
	mailer.SetHeader("Subject", MailMessageSmtp.GetSubject())
	mailer.SetBody("text/html", MailMessageSmtp.GetBody())

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return
	}
	return
}
