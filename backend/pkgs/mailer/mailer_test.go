package mailer

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TestMailerConfig = "test-mailer.json"
)

func GetTestMailer() (*Mailer, error) {
	// Read JSON File
	bytes, err := os.ReadFile(TestMailerConfig)

	mailer := &Mailer{}

	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	err = json.Unmarshal(bytes, mailer)

	if err != nil {
		return nil, err
	}

	return mailer, nil
}

func Test_Mailer(t *testing.T) {
	t.Parallel()

	mailer, err := GetTestMailer()
	if err != nil {
		t.Skip("Error Reading Test Mailer Config - Skipping")
	}

	if !mailer.Ready() {
		t.Skip("Mailer not ready - Skipping")
	}

	message, err := RenderWelcome()
	if err != nil {
		t.Error(err)
	}

	mb := NewMessageBuilder().
		SetBody(message).
		SetSubject("Hello").
		SetTo("John Doe", "john@doe.com").
		SetFrom("Jane Doe", "jane@doe.com")

	msg := mb.Build()

	err = mailer.Send(msg)

	require.NoError(t, err)
}
