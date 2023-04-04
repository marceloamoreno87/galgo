package mail

func Send(m MailMessageInterface) (err error) {
	err = m.Send()

	return
}
