package placeholdercontentmanager

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaceholderContentManager_SetCustomers(t *testing.T) {
	customers := []entity.Customer{
		{
			"Mrs",
			"Michelle",
			"Smith",
			"",
		},
	}

	placeholderContentManager := PlaceholderContentManager{}

	err := placeholderContentManager.SetCustomers(customers)

	assert.Nil(t, err)
	assert.ObjectsAreEqual(customers, placeholderContentManager.customers)
}

func TestPlaceholderContentManager_SetTemplate(t *testing.T) {
	template := entity.Template{
		"The marketing team \u003cmarketing@example.com",
		"john.smith@example.com",
		"A new product is ...",
		"text/plain",
		"Hi Mr John Smith, \nToday, 31 Aug 2022, we would like to tell that ...\n",
	}

	placeholderContentManager := PlaceholderContentManager{}

	err := placeholderContentManager.SetTemplate(template)

	assert.Nil(t, err)
	assert.ObjectsAreEqual(template, placeholderContentManager.template)
}

func TestPlaceholderContentManager_PrepareEmailContents(t *testing.T) {
	customers := []entity.Customer{
		{
			"Mr",
			"John",
			"Smith",
			"",
		},
		{
			"Mrs",
			"Michelle",
			"Smith",
			"michelle.smith@example.com",
		},
	}
	template := entity.Template{
		From:     "The marketing team \u003cmarketing@example.com",
		Subject:  "Hello {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}} A new product is ...",
		MimeType: "text/plain",
		Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, 31 Aug 2022, we would like to tell that ...\n",
	}
	validEmails := []entity.Template{
		{
			From:     "The marketing team \u003cmarketing@example.com",
			Subject:  "Hello {{Mr}} {{John}} {{Smith}} A new product is ...",
			MimeType: "text/plain",
			Body:     "Hi {{Mr}} {{John}} {{Smith}}, \nToday, 31 Aug 2022, we would like to tell that ...\n",
		},
	}
	invalidCustomers := []entity.Customer{
		{
			"Mrs",
			"Michelle",
			"Smith",
			"michelle.smith@example.com",
		},
	}

	placeholderContentManager := PlaceholderContentManager{
		customers,
		template,
	}

	actualValidEmails, actualInvalidCustomers, err := placeholderContentManager.PrepareEmailContents()

	assert.Nil(t, err)
	assert.ObjectsAreEqual(validEmails, actualValidEmails)
	assert.ObjectsAreEqual(invalidCustomers, actualInvalidCustomers)
}
