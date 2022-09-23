package services

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
	"github.com/hay-kot/homebox/backend/pkgs/pathlib"
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

func (svc *ItemService) AttachmentToken(ctx ServiceContext, itemId, attachmentId uuid.UUID) (string, error) {
	item, err := svc.repo.Items.GetOne(ctx, itemId)
	if err != nil {
		return "", err
	}
	if item.Edges.Group.ID != ctx.GID {
		return "", ErrNotOwner
	}

	token := hasher.GenerateToken()

	svc.at.Add(token.Raw, attachmentId)

	return token.Raw, nil
}

func (svc *ItemService) attachmentPath(gid, itemId uuid.UUID, filename string) string {
	path := filepath.Join(svc.filepath, gid.String(), itemId.String(), filename)
	path = pathlib.Safe(path)
	log.Debug().Str("path", path).Msg("attachment path")
	return path
}

func (svc *ItemService) AttachmentPath(ctx context.Context, token string) (string, error) {
	attachmentId, ok := svc.at.Get(token)
	if !ok {
		return "", ErrNotFound
	}

	attachment, err := svc.repo.Attachments.Get(ctx, attachmentId)
	if err != nil {
		return "", err
	}

	return attachment.Edges.Document.Path, nil
}

func (svc *ItemService) AttachmentUpdate(ctx ServiceContext, itemId uuid.UUID, data *types.ItemAttachmentUpdate) (*types.ItemOut, error) {
	// Update Properties
	attachment, err := svc.repo.Attachments.Update(ctx, data.ID, attachment.Type(data.Type))
	if err != nil {
		return nil, err
	}

	attDoc := attachment.Edges.Document

	if data.Title != attachment.Edges.Document.Title {
		newPath := pathlib.Safe(svc.attachmentPath(ctx.GID, itemId, data.Title))

		// Move File
		err = os.Rename(attachment.Edges.Document.Path, newPath)
		if err != nil {
			return nil, err
		}

		_, err = svc.repo.Docs.Update(ctx, attDoc.ID, types.DocumentUpdate{
			Title: data.Title,
			Path:  newPath,
		})
		if err != nil {
			return nil, err
		}
	}

	return svc.GetOne(ctx, ctx.GID, itemId)
}

// AttachmentAdd adds an attachment to an item by creating an entry in the Documents table and linking it to the Attachment
// Table and Items table. The file provided via the reader is stored on the file system based on the provided
// relative path during construction of the service.
func (svc *ItemService) AttachmentAdd(ctx ServiceContext, itemId uuid.UUID, filename string, attachmentType attachment.Type, file io.Reader) (*types.ItemOut, error) {
	// Get the Item
	item, err := svc.repo.Items.GetOne(ctx, itemId)
	if err != nil {
		return nil, err
	}

	if item.Edges.Group.ID != ctx.GID {
		return nil, ErrNotOwner
	}

	fp := svc.attachmentPath(ctx.GID, itemId, filename)
	filename = filepath.Base(fp)

	// Create the document
	doc, err := svc.repo.Docs.Create(ctx, ctx.GID, types.DocumentCreate{
		Title: filename,
		Path:  fp,
	})
	if err != nil {
		return nil, err
	}

	// Create the attachment
	_, err = svc.repo.Attachments.Create(ctx, itemId, doc.ID, attachmentType)
	if err != nil {
		return nil, err
	}

	// Read the contents and write them to a file on the file system
	err = os.MkdirAll(filepath.Dir(doc.Path), os.ModePerm)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(doc.Path)
	if err != nil {
		log.Err(err).Msg("failed to create file")
		return nil, err
	}

	_, err = io.Copy(f, file)
	if err != nil {
		return nil, err
	}

	return svc.GetOne(ctx, ctx.GID, itemId)
}

func (svc *ItemService) AttachmentDelete(ctx context.Context, gid, itemId, attachmentId uuid.UUID) error {
	// Get the Item
	item, err := svc.repo.Items.GetOne(ctx, itemId)
	if err != nil {
		return err
	}

	if item.Edges.Group.ID != gid {
		return ErrNotOwner
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
