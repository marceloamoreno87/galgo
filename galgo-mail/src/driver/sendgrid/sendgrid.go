package sendgrid

import (
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailMessageSendgrid struct {
	To          []string `json:"to[]" example:"test@test.com, test2@test2.com" binding:"required"`
	Cc          []string `json:"cc[]" example:"test@test.com, test2@test2.com"`
	Attachments []string `json:"attachments[]" example:"/path/file/" binding:"required"`
	Subject     string   `json:"subject" example:"testing" binding:"required"`
	From        string   `json:"from" example:"marceloamoreno87@gmail.com" binding:"required"`
	Body        string   `json:"body" example:"<h1>Hello, world!</h1>" binding:"required"`
}

func (MailMessageSendgrid *MailMessageSendgrid) GetTo() []string {
	return MailMessageSendgrid.To
}

func (MailMessageSendgrid *MailMessageSendgrid) GetCc() []string {
	return MailMessageSendgrid.Cc
}

func (MailMessageSendgrid *MailMessageSendgrid) GetSubject() string {
	return MailMessageSendgrid.Subject
}

func (MailMessageSendgrid *MailMessageSendgrid) GetFrom() string {
	return MailMessageSendgrid.From
}

func (MailMessageSendgrid *MailMessageSendgrid) GetBody() string {
	return MailMessageSendgrid.Body
}

func (MailMessageSendgrid *MailMessageSendgrid) SetTo(To []string) *MailMessageSendgrid {
	MailMessageSendgrid.To = To
	return MailMessageSendgrid
}

func (MailMessageSendgrid *MailMessageSendgrid) SetCc(Cc []string) *MailMessageSendgrid {
	MailMessageSendgrid.Cc = Cc
	return MailMessageSendgrid
}

func (MailMessageSendgrid *MailMessageSendgrid) SetSubject(Subject string) *MailMessageSendgrid {
	MailMessageSendgrid.Subject = Subject
	return MailMessageSendgrid
}

func (MailMessageSendgrid *MailMessageSendgrid) SetFrom(From string) *MailMessageSendgrid {
	MailMessageSendgrid.From = From
	return MailMessageSendgrid
}

func (MailMessageSendgrid *MailMessageSendgrid) SetBody(Body string) *MailMessageSendgrid {
	MailMessageSendgrid.Body = Body
	return MailMessageSendgrid
}

func (MailMessageSendgrid *MailMessageSendgrid) Send() (err error) {
	m := setMail(
		MailMessageSendgrid.GetFrom(),
		MailMessageSendgrid.GetTo(),
		MailMessageSendgrid.GetCc(),
		MailMessageSendgrid.GetSubject(),
		MailMessageSendgrid.GetBody(),
	)
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"),
		"/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err = sendgrid.API(request)
	if err != nil {
		return
	}
	return
}

func setMail(from string, to []string, cc []string, subject string, body string) (m *mail.SGMailV3) {
	m = mail.NewV3Mail()
	mail_from := mail.NewEmail(from, from)
	content := mail.NewContent("text/html", body)
	m.SetFrom(mail_from)
	m.AddContent(content)
	personalization := mail.NewPersonalization()
	tos := getTos(to)
	ccs := getCcs(cc)
	personalization.AddTos(tos...)
	personalization.AddCCs(ccs...)
	personalization.Subject = subject
	m.AddPersonalizations(personalization)
	return
}

func getTos(to []string) (tos []*mail.Email) {
	for _, mail_to := range to {
		m := mail.NewEmail(mail_to, mail_to)
		tos = append(tos, m)
	}
	return
}

func getCcs(cc []string) (ccs []*mail.Email) {
	for _, mail_to := range cc {
		m := mail.NewEmail(mail_to, mail_to)
		ccs = append(ccs, m)
	}
	return
}
