package repo

import (
	"context"
	"testing"

	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func labelFactory() types.LabelCreate {
	return types.LabelCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
}

func useLabels(t *testing.T, len int) ([]*ent.Label, func()) {
	t.Helper()

	labels := make([]*ent.Label, len)
	for i := 0; i < len; i++ {
		itm := labelFactory()

		item, err := tRepos.Labels.Create(context.Background(), tGroup.ID, itm)
		assert.NoError(t, err)
		labels[i] = item
	}

	return labels, func() {
		for _, item := range labels {
			err := tRepos.Labels.Delete(context.Background(), item.ID)
			assert.NoError(t, err)
		}
	}
}

func TestLabelRepository_Get(t *testing.T) {
	labels, cleanup := useLabels(t, 1)
	defer cleanup()
	label := labels[0]

	// Get by ID
	foundLoc, err := tRepos.Labels.Get(context.Background(), label.ID)
	assert.NoError(t, err)
	assert.Equal(t, label.ID, foundLoc.ID)
}

func TestLabelRepositoryGetAll(t *testing.T) {
	_, cleanup := useLabels(t, 10)
	defer cleanup()

	all, err := tRepos.Labels.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	assert.Len(t, all, 10)
}

func TestLabelRepository_Create(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	assert.NoError(t, err)

	// Get by ID
	foundLoc, err := tRepos.Labels.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = tRepos.Labels.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLabelRepository_Update(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	assert.NoError(t, err)

	updateData := types.LabelUpdate{
		ID:          loc.ID,
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}

	update, err := tRepos.Labels.Update(context.Background(), updateData)
	assert.NoError(t, err)

	foundLoc, err := tRepos.Labels.Get(context.Background(), loc.ID)
	assert.NoError(t, err)

	assert.Equal(t, update.ID, foundLoc.ID)
	assert.Equal(t, update.Name, foundLoc.Name)
	assert.Equal(t, update.Description, foundLoc.Description)

	err = tRepos.Labels.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLabelRepository_Delete(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	assert.NoError(t, err)

	err = tRepos.Labels.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)

	_, err = tRepos.Labels.Get(context.Background(), loc.ID)
	assert.Error(t, err)
}
