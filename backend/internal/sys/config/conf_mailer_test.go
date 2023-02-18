package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MailerReady_Success(t *testing.T) {
	mc := &MailerConf{
		Host:     "host",
		Port:     1,
		Username: "username",
		Password: "password",
		From:     "from",
	}

	assert.True(t, mc.Ready())
}

func Test_MailerReady_Failure(t *testing.T) {
	mc := &MailerConf{}
	assert.False(t, mc.Ready())

	mc.Host = "host"
	assert.False(t, mc.Ready())

	mc.Port = 1
	assert.False(t, mc.Ready())

	mc.Username = "username"
	assert.False(t, mc.Ready())

	mc.Password = "password"
	assert.False(t, mc.Ready())

	mc.From = "from"
	assert.True(t, mc.Ready())
}
