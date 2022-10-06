package repo

import "github.com/hay-kot/homebox/backend/ent"

// AllRepos is a container for all the repository interfaces
type AllRepos struct {
	Users       *UserRepository
	AuthTokens  *TokenRepository
	Groups      *GroupRepository
	Locations   *LocationRepository
	Labels      *LabelRepository
	Items       *ItemsRepository
	Docs        *DocumentRepository
	DocTokens   *DocumentTokensRepository
	Attachments *AttachmentRepo
}

func New(db *ent.Client, root string) *AllRepos {
	return &AllRepos{
		Users:       &UserRepository{db},
		AuthTokens:  &TokenRepository{db},
		Groups:      &GroupRepository{db},
		Locations:   &LocationRepository{db},
		Labels:      &LabelRepository{db},
		Items:       &ItemsRepository{db},
		Docs:        &DocumentRepository{db, root},
		DocTokens:   &DocumentTokensRepository{db},
		Attachments: &AttachmentRepo{db},
	}
}
