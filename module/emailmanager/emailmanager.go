package emailmanager

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/module/contentmanager"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/module/emailsender"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/module/ioprocessor"
)

type EmailManager interface {
	ProcessAndSend() error
}

type EmailManagerImp struct {
	ioProcessor    ioprocessor.IOProcessor
	contentManager contentmanager.ContentManager
	emailSender    emailsender.EmailSender
}

func NewEmailManagerImp(
	ioProcessor ioprocessor.IOProcessor,
	contentManager contentmanager.ContentManager,
	emailSender emailsender.EmailSender,
) *EmailManagerImp {
	return &EmailManagerImp{
		ioProcessor,
		contentManager,
		emailSender,
	}
}

func (e *EmailManagerImp) ProcessAndSend() (err error) {
	customers, err := e.ioProcessor.ReadCustomers()
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_read_inp")
		return err
	}
	err = e.contentManager.SetCustomers(customers)
	if err != nil {
		return err
	}

	template, err := e.ioProcessor.ReadTemplate()
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_read_inp")
		return err
	}
	err = e.contentManager.SetTemplate(template)
	if err != nil {
		return err
	}

	validEmails, invalidCustomers, err := e.contentManager.PrepareEmailContents()
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_prepare_content")
		return err
	}

	err = e.ioProcessor.WriteValidEmails(validEmails)
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_write_valid_emails")
		return err
	}

	err = e.ioProcessor.WriteInvalidCustomers(invalidCustomers)
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_write_invalid_customer")
		return err
	}

	err = e.emailSender.Send(validEmails)
	if err != nil {
		fmt.Println("email_manager.process_and_send.fail_send_emails")
		return err
	}

	return nil
}
