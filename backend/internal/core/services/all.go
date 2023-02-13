package services

import (
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/rs/zerolog/log"
)

type AllServices struct {
	User      *UserService
	Group     *GroupService
	Items     *ItemService
	Reporting *reporting.ReportingService
}

type OptionsFunc func(*options)

type options struct {
	autoIncrementAssetID bool
}

func WithAutoIncrementAssetID(v bool) func(*options) {
	return func(o *options) {
		o.autoIncrementAssetID = v
	}
}

func New(repos *repo.AllRepos, opts ...OptionsFunc) *AllServices {
	if repos == nil {
		panic("repos cannot be nil")
	}

	options := &options{
		autoIncrementAssetID: true,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &AllServices{
		User:  &UserService{repos},
		Group: &GroupService{repos},
		Items: &ItemService{
			repo:                 repos,
			autoIncrementAssetID: options.autoIncrementAssetID,
		},
		// TODO: don't use global logger
		Reporting: reporting.NewReportingService(repos, &log.Logger),
	}
}
