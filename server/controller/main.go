package controller

import (
	"net/http"

	"github.com/Brightscout/mattermost-plugin-googledrive/server/config"
)

// Endpoint ...
type Endpoint struct {
	Path         string
	Execute      func(w http.ResponseWriter, r *http.Request)
	RequiresAuth bool
}

// Endpoints ...
// add endpoints here.
// Map endpoint URL to Endpoint object. Example -
// GetMetadata.Path: GetMetadata
var Endpoints = map[string]*Endpoint{

}

// Authenticated ...
// verifies if provided request is performed by a logged in Mattermost user.
func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	var userID = r.Header.Get(config.HeaderMattermostUserId)

	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}
