package services

import "github.com/hay-kot/homebox/backend/internal/repo"

type AllServices struct {
	User     *UserService
	Location *LocationService
	Labels   *LabelService
	Items    *ItemService
}

func NewServices(repos *repo.AllRepos) *AllServices {
	if repos == nil {
		panic("repos cannot be nil")
	}

	return &AllServices{
		User:     &UserService{repos},
		Location: &LocationService{repos},
		Labels:   &LabelService{repos},
		Items: &ItemService{
			repo: repos,
			at:   attachmentTokens{},
		},
	}
}
