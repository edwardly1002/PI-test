package ioprocessor

import "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"

type IOProcessor interface {
	ReadCustomers() ([]entity.Customer, error)
	ReadTemplate() (*entity.Template, error)
	WriteValidEmails([]entity.Template) error
	WriteInvalidCustomers([]entity.Customer) error
}
