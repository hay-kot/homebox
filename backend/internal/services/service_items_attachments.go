package services

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
	"github.com/rs/zerolog/log"
)

// TODO: this isn't a scalable solution, tokens should be stored in the database
type attachmentTokens map[string]uuid.UUID

func (at attachmentTokens) Add(token string, id uuid.UUID) {
	at[token] = id

	log.Debug().Str("token", token).Str("uuid", id.String()).Msg("added token")

	go func() {
		ch := time.After(1 * time.Minute)
		<-ch
		at.Delete(token)
		log.Debug().Str("token", token).Msg("deleted token")
	}()
}

func (at attachmentTokens) Get(token string) (uuid.UUID, bool) {
	id, ok := at[token]
	return id, ok
}

func (at attachmentTokens) Delete(token string) {
	delete(at, token)
}

func (svc *ItemService) AttachmentToken(ctx Context, itemId, attachmentId uuid.UUID) (string, error) {
	_, err := svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemId)
	if err != nil {
		return "", err
	}

	token := hasher.GenerateToken()

	// Ensure that the file exists
	attachment, err := svc.repo.Attachments.Get(ctx, attachmentId)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(attachment.Edges.Document.Path); os.IsNotExist(err) {
		_ = svc.AttachmentDelete(ctx, ctx.GID, itemId, attachmentId)
		return "", ErrNotFound
	}

	svc.at.Add(token.Raw, attachmentId)

	return token.Raw, nil
}

func (svc *ItemService) AttachmentPath(ctx context.Context, token string) (*ent.Document, error) {
	attachmentId, ok := svc.at.Get(token)
	if !ok {
		return nil, ErrNotFound
	}

	attachment, err := svc.repo.Attachments.Get(ctx, attachmentId)
	if err != nil {
		return nil, err
	}

	return attachment.Edges.Document, nil
}

func (svc *ItemService) AttachmentUpdate(ctx Context, itemId uuid.UUID, data *repo.ItemAttachmentUpdate) (repo.ItemOut, error) {
	// Update Attachment
	attachment, err := svc.repo.Attachments.Update(ctx, data.ID, attachment.Type(data.Type))
	if err != nil {
		return repo.ItemOut{}, err
	}

	// Update Document
	attDoc := attachment.Edges.Document
	_, err = svc.repo.Docs.Rename(ctx, attDoc.ID, data.Title)
	if err != nil {
		return repo.ItemOut{}, err
	}

	return svc.GetOne(ctx, ctx.GID, itemId)
}

// AttachmentAdd adds an attachment to an item by creating an entry in the Documents table and linking it to the Attachment
// Table and Items table. The file provided via the reader is stored on the file system based on the provided
// relative path during construction of the service.
func (svc *ItemService) AttachmentAdd(ctx Context, itemId uuid.UUID, filename string, attachmentType attachment.Type, file io.Reader) (repo.ItemOut, error) {
	// Get the Item
	_, err := svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemId)
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
	_, err = svc.repo.Attachments.Create(ctx, itemId, doc.ID, attachmentType)
	if err != nil {
		log.Err(err).Msg("failed to create attachment")
		return repo.ItemOut{}, err
	}

	return svc.GetOne(ctx, ctx.GID, itemId)
}

func (svc *ItemService) AttachmentDelete(ctx context.Context, gid, itemId, attachmentId uuid.UUID) error {
	// Get the Item
	_, err := svc.repo.Items.GetOneByGroup(ctx, gid, itemId)
	if err != nil {
		return err
	}

	attachment, err := svc.repo.Attachments.Get(ctx, attachmentId)
	if err != nil {
		return err
	}

	// Delete the attachment
	err = svc.repo.Attachments.Delete(ctx, attachmentId)
	if err != nil {
		return err
	}

	// Remove File
	err = os.Remove(attachment.Edges.Document.Path)

	return err
}
