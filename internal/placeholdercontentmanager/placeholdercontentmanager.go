package placeholdercontentmanager

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"
	"strings"
	"time"
)

type PlaceholderContentManager struct {
	customers []entity.Customer
	template  entity.Template
}

func NewPlaceholderContentManager() *PlaceholderContentManager {
	return &PlaceholderContentManager{}
}

func (o *PlaceholderContentManager) PrepareEmailContents() ([]entity.Template, []entity.Customer, error) {
	var validEmailTemplates []entity.Template
	var invalidCustomers []entity.Customer

	for _, customer := range o.customers {
		if !isValidEmailAddress(customer.Email) {
			invalidCustomers = append(invalidCustomers, customer)
			continue
		}

		var template entity.Template
		template.From = replacePlaceholders(o.template.From, customer)
		template.Subject = replacePlaceholders(o.template.Subject, customer)
		template.MimeType = replacePlaceholders(o.template.MimeType, customer)
		template.Body = replacePlaceholders(o.template.Body, customer)
		template.To = customer.Email

		validEmailTemplates = append(validEmailTemplates, template)
	}

	return validEmailTemplates, invalidCustomers, nil
}

func isValidEmailAddress(email string) bool {
	if email == "" {
		return false
	}
	return true
}

func replacePlaceholders(s string, customer entity.Customer) string {
	s = strings.Replace(s, "{{TITLE}}", customer.Title, -1)
	s = strings.Replace(s, "{{FIRST_NAME}}", customer.FirstName, -1)
	s = strings.Replace(s, "{{LAST_NAME}}", customer.LastName, -1)
	s = strings.Replace(s, "{{TODAY}}", time.Now().Format("02 Jan 2006"), -1)
	return s
}

func (o *PlaceholderContentManager) SetCustomers(customers []entity.Customer) (err error) {
	o.customers = customers
	return nil
}

func (o *PlaceholderContentManager) SetTemplate(template entity.Template) (err error) {
	o.template = template
	return nil
}
