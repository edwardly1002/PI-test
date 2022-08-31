package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/fileprocessor"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/nullemailsender"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/placeholdercontentmanager"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/module/emailmanager"
)

func main() {
	fileProcessor := fileprocessor.NewFileProcessor(
		"asset/customers.csv",
		"asset/email_template.json",
		"asset/output_emails.json",
		"asset/errors.csv",
	)
	placeholderContentManager := placeholdercontentmanager.NewPlaceholderContentManager()
	nullEmailSender := nullemailsender.NewNullEmailSender()

	emailManager := emailmanager.NewEmailManagerImp(
		fileProcessor,
		placeholderContentManager,
		nullEmailSender,
	)

	err := emailManager.ProcessAndSend()
	if err != nil {
		fmt.Println("main.fail_process_and_send")
	}
}
