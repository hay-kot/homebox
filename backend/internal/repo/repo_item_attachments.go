package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/attachment"
)

// AttachmentRepo is a repository for Attachments table that links Items to Documents
// While also specifying the type of the attachment. This _ONLY_ provides basic Create Update
// And Delete operations. For accessing the actual documents, use the Items repository since it
// provides the attachments with the documents.
type AttachmentRepo struct {
	db *ent.Client
}

type (
	ItemAttachment struct {
		ID        uuid.UUID   `json:"id"`
		CreatedAt time.Time   `json:"createdAt"`
		UpdatedAt time.Time   `json:"updatedAt"`
		Type      string      `json:"type"`
		Document  DocumentOut `json:"document"`
	}

	ItemAttachmentUpdate struct {
		ID    uuid.UUID `json:"-"`
		Type  string    `json:"type"`
		Title string    `json:"title"`
	}
)

func ToItemAttachment(attachment *ent.Attachment) ItemAttachment {
	return ItemAttachment{
		ID:        attachment.ID,
		CreatedAt: attachment.CreatedAt,
		UpdatedAt: attachment.UpdatedAt,
		Type:      attachment.Type.String(),
		Document: DocumentOut{
			ID:    attachment.Edges.Document.ID,
			Title: attachment.Edges.Document.Title,
			Path:  attachment.Edges.Document.Path,
		},
	}
}

func (r *AttachmentRepo) Create(ctx context.Context, itemId, docId uuid.UUID, typ attachment.Type) (*ent.Attachment, error) {
	return r.db.Attachment.Create().
		SetType(typ).
		SetDocumentID(docId).
		SetItemID(itemId).
		Save(ctx)
}

func (r *AttachmentRepo) Get(ctx context.Context, id uuid.UUID) (*ent.Attachment, error) {
	return r.db.Attachment.
		Query().
		Where(attachment.ID(id)).
		WithItem().
		WithDocument().
		Only(ctx)
}

func (r *AttachmentRepo) Update(ctx context.Context, itemId uuid.UUID, typ attachment.Type) (*ent.Attachment, error) {
	itm, err := r.db.Attachment.UpdateOneID(itemId).
		SetType(typ).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.Get(ctx, itm.ID)
}

func (r *AttachmentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Attachment.DeleteOneID(id).Exec(ctx)
}
