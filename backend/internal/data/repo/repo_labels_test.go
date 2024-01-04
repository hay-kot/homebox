package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func labelFactory() LabelCreate {
	return LabelCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
}

func useLabels(t *testing.T, len int) []LabelOut {
	t.Helper()

	labels := make([]LabelOut, len)
	for i := 0; i < len; i++ {
		itm := labelFactory()

		item, err := tRepos.Labels.Create(context.Background(), tGroup.ID, itm)
		require.NoError(t, err)
		labels[i] = item
	}

	t.Cleanup(func() {
		for _, item := range labels {
			_ = tRepos.Labels.delete(context.Background(), item.ID)
		}
	})

	return labels
}

func TestLabelRepository_Get(t *testing.T) {
	labels := useLabels(t, 1)
	label := labels[0]

	// Get by ID
	foundLoc, err := tRepos.Labels.GetOne(context.Background(), label.ID)
	require.NoError(t, err)
	assert.Equal(t, label.ID, foundLoc.ID)
}

func TestLabelRepositoryGetAll(t *testing.T) {
	useLabels(t, 10)

	all, err := tRepos.Labels.GetAll(context.Background(), tGroup.ID)
	require.NoError(t, err)
	assert.Len(t, all, 10)
}

func TestLabelRepository_Create(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	require.NoError(t, err)

	// Get by ID
	foundLoc, err := tRepos.Labels.GetOne(context.Background(), loc.ID)
	require.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = tRepos.Labels.delete(context.Background(), loc.ID)
	require.NoError(t, err)
}

func TestLabelRepository_Update(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	require.NoError(t, err)

	updateData := LabelUpdate{
		ID:          loc.ID,
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}

	update, err := tRepos.Labels.UpdateByGroup(context.Background(), tGroup.ID, updateData)
	require.NoError(t, err)

	foundLoc, err := tRepos.Labels.GetOne(context.Background(), loc.ID)
	require.NoError(t, err)

	assert.Equal(t, update.ID, foundLoc.ID)
	assert.Equal(t, update.Name, foundLoc.Name)
	assert.Equal(t, update.Description, foundLoc.Description)

	err = tRepos.Labels.delete(context.Background(), loc.ID)
	require.NoError(t, err)
}

func TestLabelRepository_Delete(t *testing.T) {
	loc, err := tRepos.Labels.Create(context.Background(), tGroup.ID, labelFactory())
	require.NoError(t, err)

	err = tRepos.Labels.delete(context.Background(), loc.ID)
	require.NoError(t, err)

	_, err = tRepos.Labels.GetOne(context.Background(), loc.ID)
	require.Error(t, err)
}
