package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/backend/pkgs/faker"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorBuilder_HasErrors_NilList(t *testing.T) {
	t.Parallel()

	var ebNilList = ErrorBuilder{}
	assert.False(t, ebNilList.HasErrors(), "ErrorBuilder.HasErrors() should return false when list is nil")

}

func Test_ErrorBuilder_HasErrors_EmptyList(t *testing.T) {
	t.Parallel()

	var ebEmptyList = ErrorBuilder{
		errs: []string{},
	}
	assert.False(t, ebEmptyList.HasErrors(), "ErrorBuilder.HasErrors() should return false when list is empty")

}

func Test_ErrorBuilder_HasErrors_WithError(t *testing.T) {
	t.Parallel()

	var ebList = ErrorBuilder{}
	ebList.AddError(errors.New("test error"))

	assert.True(t, ebList.HasErrors(), "ErrorBuilder.HasErrors() should return true when list is not empty")

}

func Test_ErrorBuilder_AddError(t *testing.T) {
	t.Parallel()

	randomError := make([]error, 10)

	f := faker.NewFaker()

	errorStrings := make([]string, 10)

	for i := 0; i < 10; i++ {
		err := errors.New(f.RandomString(10))
		randomError[i] = err
		errorStrings[i] = err.Error()
	}

	// Check Results
	var ebList = ErrorBuilder{}

	for _, err := range randomError {
		ebList.AddError(err)
	}

	assert.Equal(t, errorStrings, ebList.errs, "ErrorBuilder.AddError() should add an error to the list")
}

func Test_ErrorBuilder_Respond(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()

	randomError := make([]error, 5)

	for i := 0; i < 5; i++ {
		err := errors.New(f.RandomString(5))
		randomError[i] = err
	}

	// Check Results
	var ebList = ErrorBuilder{}

	for _, err := range randomError {
		ebList.AddError(err)
	}

	fakeWriter := httptest.NewRecorder()

	ebList.Respond(fakeWriter, 422)

	assert.Equal(t, 422, fakeWriter.Code, "ErrorBuilder.Respond() should return a status code of 422")

	// Check errors payload is correct

	errorsStruct := struct {
		Errors  []string `json:"details"`
		Message string   `json:"message"`
		Error   bool     `json:"error"`
	}{
		Errors:  ebList.errs,
		Message: http.StatusText(http.StatusUnprocessableEntity),
		Error:   true,
	}

	asJson, _ := json.Marshal(errorsStruct)
	assert.JSONEq(t, string(asJson), fakeWriter.Body.String(), "ErrorBuilder.Respond() should return a JSON response with the errors")

}
