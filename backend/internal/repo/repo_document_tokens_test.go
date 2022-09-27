package repo

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/documenttoken"
	"github.com/stretchr/testify/assert"
)

func TestDocumentTokensRepository_Create(t *testing.T) {
	entities := useDocs(t, 1)
	doc := entities[0]
	expires := fk.Time()

	type args struct {
		ctx  context.Context
		data DocumentTokenCreate
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.DocumentToken
		wantErr bool
	}{
		{
			name: "create document token",
			args: args{
				ctx: context.Background(),
				data: DocumentTokenCreate{
					DocumentID: doc.ID,
					TokenHash:  []byte("token"),
					ExpiresAt:  expires,
				},
			},
			want: &ent.DocumentToken{
				Edges: ent.DocumentTokenEdges{
					Document: &ent.Document{
						ID: doc.ID,
					},
				},
				Token:     []byte("token"),
				ExpiresAt: expires,
			},
			wantErr: false,
		},
		{
			name: "create document token with empty token",
			args: args{
				ctx: context.Background(),
				data: DocumentTokenCreate{
					DocumentID: doc.ID,
					TokenHash:  []byte(""),
					ExpiresAt:  expires,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create document token with empty document id",
			args: args{
				ctx: context.Background(),
				data: DocumentTokenCreate{
					DocumentID: uuid.Nil,
					TokenHash:  []byte("token"),
					ExpiresAt:  expires,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	ids := make([]uuid.UUID, 0, len(tests))

	t.Cleanup(func() {
		for _, id := range ids {
			_ = tRepos.DocTokens.Delete(context.Background(), id)
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tRepos.DocTokens.Create(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentTokensRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			assert.Equal(t, tt.want.Token, got.TokenHash)
			assert.WithinDuration(t, tt.want.ExpiresAt, got.ExpiresAt, time.Duration(1)*time.Second)
			assert.Equal(t, tt.want.Edges.Document.ID, got.DocumentID)
		})

	}
}

func useDocTokens(t *testing.T, num int) []DocumentToken {
	entity := useDocs(t, 1)[0]

	results := make([]DocumentToken, 0, num)

	ids := make([]uuid.UUID, 0, num)
	t.Cleanup(func() {
		for _, id := range ids {
			_ = tRepos.DocTokens.Delete(context.Background(), id)
		}
	})

	for i := 0; i < num; i++ {
		e, err := tRepos.DocTokens.Create(context.Background(), DocumentTokenCreate{
			DocumentID: entity.ID,
			TokenHash:  []byte(fk.Str(10)),
			ExpiresAt:  fk.Time(),
		})

		assert.NoError(t, err)
		results = append(results, e)
		ids = append(ids, e.ID)
	}

	return results
}

func TestDocumentTokensRepository_PurgeExpiredTokens(t *testing.T) {
	entities := useDocTokens(t, 2)

	// set expired token
	tRepos.DocTokens.db.DocumentToken.Update().
		Where(documenttoken.ID(entities[0].ID)).
		SetExpiresAt(time.Now().Add(-time.Hour)).
		ExecX(context.Background())

	count, err := tRepos.DocTokens.PurgeExpiredTokens(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	all, err := tRepos.DocTokens.db.DocumentToken.Query().All(context.Background())
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, entities[1].ID, all[0].ID)
}
