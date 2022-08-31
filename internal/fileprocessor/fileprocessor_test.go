package fileprocessor

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileProcessor_ReadCustomers_Success(t *testing.T) {
	customersFromFile := []entity.Customer{
		{"Mr", "John", "Smith", "john.smith@example.com"},
		{"Mrs", "Michelle", "Smith", ""},
	}

	fileProcessor := FileProcessor{customerCsvFileName: "../../test/customers.csv"}

	customers, err := fileProcessor.ReadCustomers()

	assert.Nil(t, err)
	assert.Equal(t, len(customersFromFile), len(customers))
	for i := range customers {
		assert.Equal(t, customersFromFile[i].Title, customers[i].Title)
		assert.Equal(t, customersFromFile[i].FirstName, customers[i].FirstName)
		assert.Equal(t, customersFromFile[i].LastName, customers[i].LastName)
		assert.Equal(t, customersFromFile[i].Email, customers[i].Email)
	}
}

func TestFileProcessor_ReadCustomers_FileNotFound(t *testing.T) {
	fileProcessor := FileProcessor{customerCsvFileName: "./test/alo123?!.csv"}

	_, err := fileProcessor.ReadCustomers()

	assert.NotNil(t, err)
}

func TestFileProcessor_ReadTemplate_Success(t *testing.T) {
	templateFromFile := entity.Template{
		From:     "The marketing team <marketing@example.com",
		Subject:  "A new product is ...",
		MimeType: "text/plain",
		Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell that ...\n",
	}

	fileProcessor := FileProcessor{templateJsonFileName: "../../test/email_template.json"}

	template, err := fileProcessor.ReadTemplate()

	assert.Nil(t, err)
	assert.Equal(t, templateFromFile.From, template.From)
	assert.Equal(t, templateFromFile.To, template.To)
	assert.Equal(t, templateFromFile.Subject, template.Subject)
	assert.Equal(t, templateFromFile.MimeType, template.MimeType)
	assert.Equal(t, templateFromFile.Body, template.Body)
}

func TestFileProcessor_ReadTemplate_FileNotFound(t *testing.T) {
	fileProcessor := FileProcessor{templateJsonFileName: "./test/alo123?!.json"}

	_, err := fileProcessor.ReadTemplate()

	assert.NotNil(t, err)
}

func TestFileProcessor_WriteValidEmails(t *testing.T) {
	outputEmails := []entity.Template{
		{
			"The marketing team \u003cmarketing@example.com",
			"john.smith@example.com",
			"A new product is ...",
			"text/plain",
			"Hi Mr John Smith, \nToday, 31 Aug 2022, we would like to tell that ...\n",
		},
	}

	outputEmailsFileName := "../../test/output_emails.json"
	fileProcessor := FileProcessor{outputEmailsJsonFileName: outputEmailsFileName}

	err := fileProcessor.WriteValidEmails(outputEmails)

	assert.Nil(t, err)
	assert.FileExists(t, outputEmailsFileName)
}

func TestFileProcessor_WriteInvalidCustomers(t *testing.T) {
	invalidCustomers := []entity.Customer{
		{
			"Mrs",
			"Michelle",
			"Smith",
			"",
		},
	}

	errorFileName := "../../test/errors.csv"
	fileProcessor := FileProcessor{errorFileName: errorFileName}

	err := fileProcessor.WriteInvalidCustomers(invalidCustomers)

	assert.Nil(t, err)
	assert.FileExists(t, errorFileName)
}
