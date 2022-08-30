package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewHandlerV1(t *testing.T) {

	v1Base := BaseUrlFunc("/testing/v1")
	ctrl := NewControllerV1(mockHandler.log, mockHandler.svc)

	assert.NotNil(t, ctrl)

	assert.Equal(t, ctrl.log, mockHandler.log)

	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
}
