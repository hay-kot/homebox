package repo

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/document"
	"github.com/hay-kot/homebox/backend/ent/group"
	"github.com/hay-kot/homebox/backend/pkgs/pathlib"
)

var (
	ErrInvalidDocExtension = errors.New("invalid document extension")
)

type DocumentRepository struct {
	db  *ent.Client
	dir string
}

type (
	DocumentCreate struct {
		Title   string
		Content io.Reader
	}
)

func (r *DocumentRepository) path(gid uuid.UUID, ext string) string {
	return pathlib.Safe(filepath.Join(r.dir, gid.String(), "documents", uuid.NewString()+ext))
}

func (r *DocumentRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]*ent.Document, error) {
	return r.db.Document.Query().
		Where(document.HasGroupWith(group.ID(gid))).
		All(ctx)
}

func (r *DocumentRepository) Get(ctx context.Context, id uuid.UUID) (*ent.Document, error) {
	return r.db.Document.Get(ctx, id)
}

func (r *DocumentRepository) Create(ctx context.Context, gid uuid.UUID, doc DocumentCreate) (*ent.Document, error) {
	ext := filepath.Ext(doc.Title)
	if ext == "" {
		return nil, ErrInvalidDocExtension
	}

	path := r.path(gid, ext)

	parent := filepath.Dir(path)
	err := os.MkdirAll(parent, 0755)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(f, doc.Content)
	if err != nil {
		return nil, err
	}

	return r.db.Document.Create().
		SetGroupID(gid).
		SetTitle(doc.Title).
		SetPath(path).
		Save(ctx)
}

func (r *DocumentRepository) Rename(ctx context.Context, id uuid.UUID, title string) (*ent.Document, error) {
	return r.db.Document.UpdateOneID(id).
		SetTitle(title).
		Save(ctx)
}

func (r *DocumentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	doc, err := r.db.Document.Get(ctx, id)
	if err != nil {
		return err
	}

	err = os.Remove(doc.Path)
	if err != nil {
		return err
	}

	return r.db.Document.DeleteOneID(id).Exec(ctx)
}
