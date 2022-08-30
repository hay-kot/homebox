package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/backend/internal/mocks/chimocker"
	"github.com/hay-kot/git-web-template/backend/internal/mocks/factories"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
	"github.com/stretchr/testify/assert"
)

const (
	UrlUser      = "/api/v1/admin/users"
	UrlUserId    = "/api/v1/admin/users/%v"
	UrlUserIdChi = "/api/v1/admin/users/{id}"
)

type usersResponse struct {
	Users []types.UserOut `json:"item"`
}

type userResponse struct {
	User types.UserOut `json:"item"`
}

func Test_HandleAdminUserGetAll_Success(t *testing.T) {
	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, UrlUser, nil)

	mockHandler.HandleAdminUserGetAll()(r, req)

	response := usersResponse{
		Users: []types.UserOut{},
	}

	_ = json.Unmarshal(r.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, len(users), len(response.Users))

	knowEmail := []string{
		users[0].Email,
		users[1].Email,
		users[2].Email,
		users[3].Email,
	}

	for _, user := range users {
		assert.Contains(t, knowEmail, user.Email)
	}

}

func Test_HandleAdminUserGet_Success(t *testing.T) {
	targetUser := users[2]
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(UrlUserId, targetUser.ID), nil)

	req = chimocker.WithUrlParam(req, "id", fmt.Sprintf("%v", targetUser.ID))

	mockHandler.HandleAdminUserGet()(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	response := userResponse{
		User: types.UserOut{},
	}

	_ = json.Unmarshal(res.Body.Bytes(), &response)
	assert.Equal(t, targetUser.ID, response.User.ID)
}

func Test_HandleAdminUserCreate_Success(t *testing.T) {
	payload := factories.UserFactory()

	r := httptest.NewRecorder()

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, UrlUser, bytes.NewBuffer(body))
	req.Header.Set(server.ContentType, server.ContentJSON)

	mockHandler.HandleAdminUserCreate()(r, req)

	assert.Equal(t, http.StatusCreated, r.Code)

	usr, err := mockHandler.svc.Admin.GetByEmail(context.Background(), payload.Email)

	assert.NoError(t, err)
	assert.Equal(t, payload.Email, usr.Email)
	assert.Equal(t, payload.Name, usr.Name)
	assert.NotEqual(t, payload.Password, usr.Password) // smoke test - check password is hashed

	_ = mockHandler.svc.Admin.Delete(context.Background(), usr.ID)
}

func Test_HandleAdminUserUpdate_Success(t *testing.T) {
	t.Skip()
}

func Test_HandleAdminUserUpdate_Delete(t *testing.T) {
	t.Skip()
}
