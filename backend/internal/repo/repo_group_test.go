package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Group_Create(t *testing.T) {
	g, err := testRepos.Groups.Create(context.Background(), "test")

	assert.NoError(t, err)
	assert.Equal(t, "test", g.Name)

	// Get by ID
	foundGroup, err := testRepos.Groups.GetOneId(context.Background(), g.ID)
	assert.NoError(t, err)
	assert.Equal(t, g.ID, foundGroup.ID)
}
