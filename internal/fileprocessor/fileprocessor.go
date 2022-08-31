package fileprocessor

import (
	"encoding/json"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"os"
)

type FileProcessor struct {
	customerCsvFileName      string
	templateJsonFileName     string
	outputEmailsJsonFileName string
	errorFileName            string
}

func NewFileProcessor(
	customerCsvFileName string,
	templateJsonFileName string,
	outputEmailsJsonFileName string,
	errorCsvFileName string,
) *FileProcessor {
	return &FileProcessor{
		customerCsvFileName,
		templateJsonFileName,
		outputEmailsJsonFileName,
		errorCsvFileName,
	}
}

func (fp *FileProcessor) ReadCustomers() ([]entity.Customer, error) {
	in, err := os.Open(fp.customerCsvFileName)
	if err != nil {
		fmt.Println("file_processor.read_customers.fail_open_csv_file", err)
		return nil, err
	}
	defer in.Close()

	var customers []entity.Customer
	err = gocsv.UnmarshalFile(in, &customers)
	if err != nil {
		fmt.Println("file_processor.read_customers.fail_unmarshal_csv", err)
		return nil, err
	}

	return customers, nil
}

func (fp *FileProcessor) ReadTemplate() (*entity.Template, error) {
	in, err := os.Open(fp.templateJsonFileName)
	if err != nil {
		fmt.Println("file_processor.read_template.fail_open_json_file", err)
		return nil, err
	}
	defer in.Close()

	var template entity.Template
	byteValue, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Println("file_processor.read_template.fail_read_json", err)
		return nil, err
	}
	err = json.Unmarshal(byteValue, &template)
	if err != nil {
		fmt.Println("file_processor.read_template.fail_unmarshal_json", err)
		return nil, err
	}

	return &template, nil
}

func (fp *FileProcessor) WriteValidEmails(emails []entity.Template) error {
	bytes, err := json.MarshalIndent(emails, "", "\t")
	if err != nil {
		fmt.Println("file_processor.write_valid_emails.fail_marshal_emails", err)
		return err
	}

	err = ioutil.WriteFile(fp.outputEmailsJsonFileName, bytes, 0644)
	if err != nil {
		fmt.Println("file_processor.write_valid_emails.fail_write_file", err)
		return err
	}

	return nil
}

func (fp *FileProcessor) WriteInvalidCustomers(customers []entity.Customer) error {
	f, err := os.OpenFile(fp.errorFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("file_processor.write_invalid_customers.fail_open_file", err)
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(customers, f)
	if err != nil {
		fmt.Println("file_processor.write_invalid_customers.fail_marshal_file", err)
		return err
	}
	return nil
}
