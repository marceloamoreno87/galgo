package mail

type MailMessageInterface interface {
	Send() (err error)
}
