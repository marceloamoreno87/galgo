package ses

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type MailMessageSes struct {
	To          []string `json:"to[]" example:"test@test.com, test2@test2.com" binding:"required"`
	Cc          []string `json:"cc[]" example:"test@test.com, test2@test2.com"`
	Attachments []string `json:"attachments[]" example:"/path/file/" binding:"required"`
	Subject     string   `json:"subject" example:"testing" binding:"required"`
	From        string   `json:"from" example:"marceloamoreno87@gmail.com" binding:"required"`
	Body        string   `json:"body" example:"<h1>Hello, world!</h1>" binding:"required"`
}

func (MailMessageSes *MailMessageSes) GetTo() []string {
	return MailMessageSes.To
}

func (MailMessageSes *MailMessageSes) GetCc() []string {
	return MailMessageSes.Cc
}

func (MailMessageSes *MailMessageSes) GetSubject() string {
	return MailMessageSes.Subject
}

func (MailMessageSes *MailMessageSes) GetFrom() string {
	return MailMessageSes.From
}

func (MailMessageSes *MailMessageSes) GetBody() string {
	return MailMessageSes.Body
}

func (MailMessageSes *MailMessageSes) SetTo(To []string) *MailMessageSes {
	MailMessageSes.To = To
	return MailMessageSes
}

func (MailMessageSes *MailMessageSes) SetCc(Cc []string) *MailMessageSes {
	MailMessageSes.Cc = Cc
	return MailMessageSes
}

func (MailMessageSes *MailMessageSes) SetSubject(Subject string) *MailMessageSes {
	MailMessageSes.Subject = Subject
	return MailMessageSes
}

func (MailMessageSes *MailMessageSes) SetFrom(From string) *MailMessageSes {
	MailMessageSes.From = From
	return MailMessageSes
}

func (MailMessageSes *MailMessageSes) SetBody(Body string) *MailMessageSes {
	MailMessageSes.Body = Body
	return MailMessageSes
}

func (MailMessageSes *MailMessageSes) Send() (err error) {
	email_template := generateSESTemplate(
		MailMessageSes.GetFrom(),
		MailMessageSes.GetTo(),
		MailMessageSes.GetCc(),
		MailMessageSes.GetSubject(),
		MailMessageSes.GetBody(),
	)
	sess, err := getSessionSES()
	service := ses.New(sess)
	_, err = service.SendEmail(email_template)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return aerr
		} else {
			return
		}
	}
	return
}

func getSessionSES() (sess *session.Session, error error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("SES_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("SES_ACCESS_KEY_ID"), os.Getenv("SES_SECRET_ACCESS_KEY"), ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	return
}

func generateSESTemplate(from string, to []string, cc []string, subject string, body string) (template *ses.SendEmailInput) {
	template = &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: aws.StringSlice(cc),
			ToAddresses: aws.StringSlice(to),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("utf-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
	}
	return
}
