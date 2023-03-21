package repo

import "github.com/hay-kot/homebox/backend/internal/data/ent"

// AllRepos is a container for all the repository interfaces
type AllRepos struct {
	Users       *UserRepository
	AuthTokens  *TokenRepository
	Groups      *GroupRepository
	Locations   *LocationRepository
	Labels      *LabelRepository
	Items       *ItemsRepository
	Docs        *DocumentRepository
	Attachments *AttachmentRepo
	MaintEntry  *MaintenanceEntryRepository
	Notifiers   *NotifierRepository
}

func New(db *ent.Client, root string) *AllRepos {
	return &AllRepos{
		Users:       &UserRepository{db},
		AuthTokens:  &TokenRepository{db},
		Groups:      NewGroupRepository(db),
		Locations:   &LocationRepository{db},
		Labels:      &LabelRepository{db},
		Items:       &ItemsRepository{db},
		Docs:        &DocumentRepository{db, root},
		Attachments: &AttachmentRepo{db},
		MaintEntry:  &MaintenanceEntryRepository{db},
		Notifiers:   NewNotifierRepository(db),
	}
}
