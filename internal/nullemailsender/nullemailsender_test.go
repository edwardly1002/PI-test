package nullemailsender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNullEmailSender_Send(t *testing.T) {
	nullEmailSender := NullEmailSender{}

	err := nullEmailSender.Send(nil)

	assert.Nil(t, err)
}
