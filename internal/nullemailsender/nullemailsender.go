package nullemailsender

type NullEmailSender struct {
}

func NewNullEmailSender() *NullEmailSender {
	return &NullEmailSender{}
}

func (*NullEmailSender) Send(interface{}) error {
	return nil
}
