package controller

import (
	"net/http"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/config"
)

type Endpoint struct {
	Path         string
	Execute      func(w http.ResponseWriter, r *http.Request)
	RequiresAuth bool
}

var Endpoints = map[string]*Endpoint{
	// add endpoints here.
	// Map endpoint URL to Endpoint object. Example -
	// GetMetadata.Path: GetMetadata
}

// verifies if provided request is performed by a logged in Mattermost user.
func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	var userId = r.Header.Get(config.HeaderMattermostUserId)

	if userId == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}
