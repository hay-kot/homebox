package repo

import "github.com/hay-kot/content/backend/ent"

// AllRepos is a container for all the repository interfaces
type AllRepos struct {
	Users      *EntUserRepository
	AuthTokens *EntTokenRepository
	Groups     *EntGroupRepository
	Locations  *EntLocationRepository
	Labels     *EntLabelRepository
	Items      *ItemsRepository
}

func EntAllRepos(db *ent.Client) *AllRepos {
	return &AllRepos{
		Users:      &EntUserRepository{db},
		AuthTokens: &EntTokenRepository{db},
		Groups:     &EntGroupRepository{db},
		Locations:  &EntLocationRepository{db},
		Labels:     &EntLabelRepository{db},
		Items:      &ItemsRepository{db},
	}
}
