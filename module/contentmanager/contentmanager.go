package contentmanager

import "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"

type ContentManager interface {
	PrepareEmailContents() ([]entity.Template, []entity.Customer, error)
	SetCustomers([]entity.Customer) error
	SetTemplate(entity.Template) error
}
