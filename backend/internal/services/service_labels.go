package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
)

type LabelService struct {
	repos *repo.AllRepos
}

func (svc *LabelService) Create(ctx context.Context, groupId uuid.UUID, data repo.LabelCreate) (repo.LabelOut, error) {
	return svc.repos.Labels.Create(ctx, groupId, data)
}

func (svc *LabelService) Update(ctx context.Context, groupId uuid.UUID, data repo.LabelUpdate) (repo.LabelOut, error) {
	return svc.repos.Labels.Update(ctx, data)
}

func (svc *LabelService) Delete(ctx context.Context, gid uuid.UUID, id uuid.UUID) error {
	_, err := svc.repos.Labels.GetOneByGroup(ctx, gid, id)
	if err != nil {
		return err
	}
	return svc.repos.Labels.Delete(ctx, id)
}

func (svc *LabelService) Get(ctx context.Context, gid uuid.UUID, id uuid.UUID) (repo.LabelOut, error) {
	return svc.repos.Labels.GetOneByGroup(ctx, gid, id)

}

func (svc *LabelService) GetAll(ctx context.Context, groupId uuid.UUID) ([]repo.LabelSummary, error) {
	return svc.repos.Labels.GetAll(ctx, groupId)
}
