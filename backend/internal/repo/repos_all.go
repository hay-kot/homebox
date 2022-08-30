package repo

import "github.com/hay-kot/git-web-template/backend/ent"

// AllRepos is a container for all the repository interfaces
type AllRepos struct {
	Users      UserRepository
	AuthTokens TokenRepository
}

func EntAllRepos(db *ent.Client) *AllRepos {
	return &AllRepos{
		Users:      &EntUserRepository{db},
		AuthTokens: &EntTokenRepository{db},
	}
}
