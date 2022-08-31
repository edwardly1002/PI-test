package emailsender

type EmailSender interface {
	Send(interface{}) error
}
