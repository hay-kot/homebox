package base

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/backend/internal/mocks"
)

func GetTestHandler(t *testing.T) *BaseController {
	return NewBaseController(mocks.GetStructLogger(), nil)
}

func TestHandlersv1_HandleBase(t *testing.T) {
	// Setup
	hdlrFunc := GetTestHandler(t).HandleBase(func() bool { return true }, "v1")

	// Call Handler Func
	rr := httptest.NewRecorder()
	hdlrFunc(rr, nil)

	// Validate Status Code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, rr.Code)
	}

	// Validate Json Payload
	expected := `{"item":{"health":true,"versions":["v1"],"title":"Go API Template","message":"Welcome to the Go API Template Application!"}}`

	if rr.Body.String() != expected {
		t.Errorf("Expected json to be %s, got %s", expected, rr.Body.String())
	}

}
