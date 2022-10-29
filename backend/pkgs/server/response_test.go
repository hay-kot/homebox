package server

import (
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

	err := Respond(recorder, http.StatusNoContent, dummystruct)
	assert.NoError(t, err)

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

	err := Respond(recorder, http.StatusCreated, dummystruct)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.JSONEq(t, recorder.Body.String(), `{"name":"dummy"}`)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

}
