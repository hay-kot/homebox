package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Respond_NoContent(t *testing.T) {
	recorder := httptest.NewRecorder()
	dummystruct := struct {
		Name string
	}{
		Name: "dummy",
	}

	Respond(recorder, http.StatusNoContent, dummystruct)

	assert.Equal(t, http.StatusNoContent, recorder.Code)
	assert.Empty(t, recorder.Body.String())
}

func Test_Respond_JSON(t *testing.T) {
	recorder := httptest.NewRecorder()
	dummystruct := struct {
		Name string `json:"name"`
	}{
		Name: "dummy",
	}

	Respond(recorder, http.StatusCreated, dummystruct)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"name":"dummy"}`)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

}

func Test_RespondError(t *testing.T) {
	recorder := httptest.NewRecorder()
	var customError = errors.New("custom error")

	RespondError(recorder, http.StatusBadRequest, customError)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"details":["custom error"], "message":"Bad Request", "error":true}`)

}
func Test_RespondInternalServerError(t *testing.T) {
	recorder := httptest.NewRecorder()

	RespondInternalServerError(recorder)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"details":["internal server error"], "message":"Internal Server Error", "error":true}`)

}
func Test_RespondUnauthorized(t *testing.T) {
	recorder := httptest.NewRecorder()

	RespondUnauthorized(recorder)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"details":["unauthorized"], "message":"Unauthorized", "error":true}`)

}
func Test_RespondForbidden(t *testing.T) {
	recorder := httptest.NewRecorder()

	RespondForbidden(recorder)

	assert.Equal(t, http.StatusForbidden, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"details":["forbidden"], "message":"Forbidden", "error":true}`)

}
