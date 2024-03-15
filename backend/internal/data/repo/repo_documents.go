package repo

import (
	"context"
	"errors"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/blobstore"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
)

var ErrInvalidDocExtension = errors.New("invalid document extension")

type DocumentRepository struct {
	db *ent.Client
	bs blobstore.BlobStore
}

type (
	DocumentCreate struct {
		Title   string    `json:"title"`
		Content io.Reader `json:"content"`
	}

	DocumentOut struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
		Path  string    `json:"path"`
	}
)

func mapDocumentOut(doc *ent.Document) DocumentOut {
	return DocumentOut{
		ID:    doc.ID,
		Title: doc.Title,
		Path:  doc.Path,
	}
}

var (
	mapDocumentOutErr     = mapTErrFunc(mapDocumentOut)
	mapDocumentOutEachErr = mapTEachErrFunc(mapDocumentOut)
)

func (r *DocumentRepository) blobKey(gid uuid.UUID, ext string) string {
	return filepath.Join(gid.String(), "documents", uuid.NewString()+ext)
}

func (r *DocumentRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]DocumentOut, error) {
	return mapDocumentOutEachErr(r.db.Document.
		Query().
		Where(document.HasGroupWith(group.ID(gid))).
		All(ctx),
	)
}

func (r *DocumentRepository) Get(ctx context.Context, id uuid.UUID) (DocumentOut, error) {
	return mapDocumentOutErr(r.db.Document.Get(ctx, id))
}

func (r *DocumentRepository) Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error) {
	doc, err := r.db.Document.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	content, err := r.bs.Get(ctx, doc.Path)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (r *DocumentRepository) Create(ctx context.Context, gid uuid.UUID, doc DocumentCreate) (DocumentOut, error) {
	ext := filepath.Ext(doc.Title)
	if ext == "" {
		return DocumentOut{}, ErrInvalidDocExtension
	}

	key := r.blobKey(gid, ext)

	path, err := r.bs.Put(ctx, key, doc.Content)
	if err != nil {
		return DocumentOut{}, err
	}

	return mapDocumentOutErr(r.db.Document.Create().
		SetGroupID(gid).
		SetTitle(doc.Title).
		SetPath(path).
		Save(ctx),
	)
}

func (r *DocumentRepository) Rename(ctx context.Context, id uuid.UUID, title string) (DocumentOut, error) {
	return mapDocumentOutErr(r.db.Document.UpdateOneID(id).
		SetTitle(title).
		Save(ctx))
}

func (r *DocumentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	doc, err := r.db.Document.Get(ctx, id)
	if err != nil {
		return err
	}

	err = r.bs.Delete(ctx, doc.Path)
	if err != nil {
		return err
	}

	return r.db.Document.DeleteOneID(id).Exec(ctx)
}
