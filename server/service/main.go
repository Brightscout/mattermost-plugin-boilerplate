package service

import (
	"github.com/mattermost/mattermost-server/v5/plugin"

	"github.com/Brightscout/mattermost-plugin-boilerplate/server/store"
)

type Service struct {
	api     plugin.API
	helpers plugin.Helpers
	store   *store.Store
}

func NewService(api plugin.API, helpers plugin.Helpers, store *store.Store) *Service {
	return &Service{
		api,
		helpers,
		store,
	}
}
