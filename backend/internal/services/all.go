package services

import "github.com/hay-kot/git-web-template/backend/internal/repo"

type AllServices struct {
	User  *UserService
	Admin *AdminService
}

func NewServices(repos *repo.AllRepos) *AllServices {
	return &AllServices{
		User:  &UserService{repos},
		Admin: &AdminService{repos},
	}
}
