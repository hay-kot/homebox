package repo

import "github.com/hay-kot/content/backend/ent"

// AllRepos is a container for all the repository interfaces
type AllRepos struct {
	Users      *UserRepository
	AuthTokens *TokenRepository
	Groups     *GroupRepository
	Locations  *LocationRepository
	Labels     *LabelRepository
	Items      *ItemsRepository
}

func EntAllRepos(db *ent.Client) *AllRepos {
	return &AllRepos{
		Users:      &UserRepository{db},
		AuthTokens: &TokenRepository{db},
		Groups:     &GroupRepository{db},
		Locations:  &LocationRepository{db},
		Labels:     &LabelRepository{db},
		Items:      &ItemsRepository{db},
	}
}
