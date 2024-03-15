package services

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/rs/zerolog/log"
)

func (svc *ItemService) AttachmentData(ctx context.Context, attachmentID uuid.UUID) (*ent.Document, io.ReadCloser, error) {
	attachment, err := svc.repo.Attachments.Get(ctx, attachmentID)
	if err != nil {
		return nil, nil, err
	}

	content, err := svc.repo.Docs.Read(ctx, attachment.Edges.Document.ID)
	if err != nil {
		return nil, nil, err
	}

	return attachment.Edges.Document, content, nil
}

func (svc *ItemService) AttachmentUpdate(ctx Context, itemID uuid.UUID, data *repo.ItemAttachmentUpdate) (repo.ItemOut, error) {
	// Update Attachment
	attachment, err := svc.repo.Attachments.Update(ctx, data.ID, data)
	if err != nil {
		return repo.ItemOut{}, err
	}

	// Update Document
	attDoc := attachment.Edges.Document
	_, err = svc.repo.Docs.Rename(ctx, attDoc.ID, data.Title)
	if err != nil {
		return repo.ItemOut{}, err
	}

	return svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
}

// AttachmentAdd adds an attachment to an item by creating an entry in the Documents table and linking it to the Attachment
// Table and Items table. The file provided via the reader is stored on the file system based on the provided
// relative path during construction of the service.
func (svc *ItemService) AttachmentAdd(ctx Context, itemID uuid.UUID, filename string, attachmentType attachment.Type, file io.Reader) (repo.ItemOut, error) {
	// Get the Item
	_, err := svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
	if err != nil {
		return repo.ItemOut{}, err
	}

	// Create the document
	doc, err := svc.repo.Docs.Create(ctx, ctx.GID, repo.DocumentCreate{Title: filename, Content: file})
	if err != nil {
		log.Err(err).Msg("failed to create document")
		return repo.ItemOut{}, err
	}

	// Create the attachment
	_, err = svc.repo.Attachments.Create(ctx, itemID, doc.ID, attachmentType)
	if err != nil {
		log.Err(err).Msg("failed to create attachment")
		return repo.ItemOut{}, err
	}

	return svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
}

func (svc *ItemService) AttachmentDelete(ctx context.Context, gid, itemID, attachmentID uuid.UUID) error {
	// Get the Item
	_, err := svc.repo.Items.GetOneByGroup(ctx, gid, itemID)
	if err != nil {
		return err
	}

	attachment, err := svc.repo.Attachments.Get(ctx, attachmentID)
	if err != nil {
		return err
	}

	// Delete the attachment
	err = svc.repo.Attachments.Delete(ctx, attachmentID)
	if err != nil {
		return err
	}

	// Delete the document
	err = svc.repo.Docs.Delete(ctx, attachment.Edges.Document.ID)

	return err
}
