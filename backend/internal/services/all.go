package services

import "github.com/hay-kot/homebox/backend/internal/repo"

type AllServices struct {
	User  *UserService
	Group *GroupService
	Items *ItemService
}

func New(repos *repo.AllRepos) *AllServices {
	if repos == nil {
		panic("repos cannot be nil")
	}

	return &AllServices{
		User:  &UserService{repos},
		Group: &GroupService{repos},
		Items: &ItemService{
			repo: repos,
			at:   attachmentTokens{},
		},
	}
}
