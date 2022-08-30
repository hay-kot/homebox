package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageBuilder(t *testing.T) {
	t.Parallel()

	mb := NewMessageBuilder().
		SetBody("Hello World!").
		SetSubject("Hello").
		SetTo("John Doe", "john@doe.com").
		SetFrom("Jane Doe", "jane@doe.com")

	msg := mb.Build()

	assert.Equal(t, "Hello", msg.Subject)
	assert.Equal(t, "Hello World!", msg.Body)
	assert.Equal(t, "John Doe", msg.To.Name)
	assert.Equal(t, "john@doe.com", msg.To.Address)
	assert.Equal(t, "Jane Doe", msg.From.Name)
	assert.Equal(t, "jane@doe.com", msg.From.Address)
}
