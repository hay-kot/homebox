package repo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestDocumentRepository_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		gid uuid.UUID
		doc types.DocumentCreate
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Document
		wantErr bool
	}{
		{
			name: "create document",
			args: args{
				ctx: context.Background(),
				gid: tGroup.ID,
				doc: types.DocumentCreate{
					Title: "test document",
					Path:  "/test/document",
				},
			},
			want: &ent.Document{
				Title: "test document",
				Path:  "/test/document",
			},
			wantErr: false,
		},
		{
			name: "create document with empty title",
			args: args{
				ctx: context.Background(),
				gid: tGroup.ID,
				doc: types.DocumentCreate{
					Title: "",
					Path:  "/test/document",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create document with empty path",
			args: args{
				ctx: context.Background(),
				gid: tGroup.ID,
				doc: types.DocumentCreate{
					Title: "test document",
					Path:  "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	ids := make([]uuid.UUID, 0, len(tests))

	t.Cleanup(func() {
		for _, id := range ids {
			err := tRepos.Docs.Delete(context.Background(), id)
			assert.NoError(t, err)
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tRepos.Docs.Create(tt.args.ctx, tt.args.gid, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
				return
			}

			assert.Equal(t, tt.want.Title, got.Title)
			assert.Equal(t, tt.want.Path, got.Path)
			ids = append(ids, got.ID)
		})
	}
}

func useDocs(t *testing.T, num int) []*ent.Document {
	t.Helper()

	results := make([]*ent.Document, 0, num)
	ids := make([]uuid.UUID, 0, num)

	for i := 0; i < num; i++ {
		doc, err := tRepos.Docs.Create(context.Background(), tGroup.ID, types.DocumentCreate{
			Title: fk.Str(10),
			Path:  fk.Path(),
		})

		assert.NoError(t, err)
		assert.NotNil(t, doc)
		results = append(results, doc)
		ids = append(ids, doc.ID)
	}

	t.Cleanup(func() {
		for _, id := range ids {
			err := tRepos.Docs.Delete(context.Background(), id)

			if err != nil {
				assert.True(t, ent.IsNotFound(err))
			}
		}
	})

	return results
}

func TestDocumentRepository_GetAll(t *testing.T) {
	entities := useDocs(t, 10)

	for _, entity := range entities {
		assert.NotNil(t, entity)
	}

	all, err := tRepos.Docs.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)

	assert.Len(t, all, 10)
	for _, entity := range all {
		assert.NotNil(t, entity)

		for _, e := range entities {
			if e.ID == entity.ID {
				assert.Equal(t, e.Title, entity.Title)
				assert.Equal(t, e.Path, entity.Path)
			}
		}
	}
}

func TestDocumentRepository_Get(t *testing.T) {
	entities := useDocs(t, 10)

	for _, entity := range entities {
		got, err := tRepos.Docs.Get(context.Background(), entity.ID)

		assert.NoError(t, err)
		assert.Equal(t, entity.ID, got.ID)
		assert.Equal(t, entity.Title, got.Title)
		assert.Equal(t, entity.Path, got.Path)
	}
}

func TestDocumentRepository_Update(t *testing.T) {
	entities := useDocs(t, 10)

	for _, entity := range entities {
		got, err := tRepos.Docs.Get(context.Background(), entity.ID)

		assert.NoError(t, err)
		assert.Equal(t, entity.ID, got.ID)
		assert.Equal(t, entity.Title, got.Title)
		assert.Equal(t, entity.Path, got.Path)
	}

	for _, entity := range entities {
		updateData := types.DocumentUpdate{
			Title: fk.Str(10),
			Path:  fk.Path(),
		}

		updated, err := tRepos.Docs.Update(context.Background(), entity.ID, updateData)

		assert.NoError(t, err)
		assert.Equal(t, entity.ID, updated.ID)
		assert.Equal(t, updateData.Title, updated.Title)
		assert.Equal(t, updateData.Path, updated.Path)
	}
}

func TestDocumentRepository_Delete(t *testing.T) {
	entities := useDocs(t, 10)

	for _, entity := range entities {
		err := tRepos.Docs.Delete(context.Background(), entity.ID)
		assert.NoError(t, err)

		_, err = tRepos.Docs.Get(context.Background(), entity.ID)
		assert.Error(t, err)
	}
}
