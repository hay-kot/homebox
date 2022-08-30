package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_SetAuthContext(t *testing.T) {
	user := &types.UserOut{
		ID: uuid.New(),
	}

	token := uuid.New().String()

	ctx := SetUserCtx(context.Background(), user, token)

	ctxUser := UseUserCtx(ctx)

	assert.NotNil(t, ctxUser)
	assert.Equal(t, user.ID, ctxUser.ID)

	ctxUserToken := UseTokenCtx(ctx)
	assert.NotEmpty(t, ctxUserToken)
}

func Test_SetAuthContext_Nulls(t *testing.T) {
	ctx := SetUserCtx(context.Background(), nil, "")

	ctxUser := UseUserCtx(ctx)

	assert.Nil(t, ctxUser)

	ctxUserToken := UseTokenCtx(ctx)
	assert.Empty(t, ctxUserToken)
}
