package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/fileprocessor"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/nullemailsender"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/placeholdercontentmanager"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/module/emailmanager"
	"os"
)

func main() {
	templateFileName, ok := os.LookupEnv("TEMPLATE_FILE")
	if !ok {
		panic("The ENV VARIABLE <TEMPLATE_FILE> is not set")
	}

	customersFileName, ok := os.LookupEnv("CUSTOMERS_FILE")
	if !ok {
		panic("The ENV VARIABLE <CUSTOMERS_FILE> is not set")
	}

	outputEmailsFileName, ok := os.LookupEnv("OUTPUT_FILE")
	if !ok {
		panic("The ENV VARIABLE <OUTPUT_FILE> is not set")
	}

	errorsFileName, ok := os.LookupEnv("ERRORS_FILE")
	if !ok {
		panic("The ENV VARIABLE <ERRORS_FILE> is not set")
	}

	fileProcessor := fileprocessor.NewFileProcessor(
		customersFileName,
		templateFileName,
		outputEmailsFileName,
		errorsFileName,
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
