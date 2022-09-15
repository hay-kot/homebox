package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_NewHandlerV1(t *testing.T) {

	v1Base := BaseUrlFunc("/testing/v1")
	ctrl := NewControllerV1(mockHandler.svc)

	assert.NotNil(t, ctrl)

	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
}

func TestHandlersv1_HandleBase(t *testing.T) {
	// Setup
	hdlrFunc := mockHandler.HandleBase(func() bool { return true }, types.Build{
		Version:   "0.1.0",
		Commit:    "HEAD",
		BuildTime: "now",
	})

	// Call Handler Func
	rr := httptest.NewRecorder()
	hdlrFunc(rr, nil)

	// Validate Status Code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, rr.Code)
	}

	// Validate Json Payload
	expected := `{"health":true,"versions":null,"title":"Go API Template","message":"Welcome to the Go API Template Application!","Build":{"version":"0.1.0","commit":"HEAD","buildTime":"now"}}`

	if rr.Body.String() != expected {
		t.Errorf("Expected json to be %s, got %s", expected, rr.Body.String())
	}

}
