package config

import (
	"github.com/mattermost/mattermost-server/plugin"
	"go.uber.org/atomic"
)

const (
	PluginName                = "boilerplate"
	CommandPrefix             = PluginName
	UrlMappingKeyPrefix       = "url_"
	ServerExeToWebappRootPath = "/../webapp"

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserId = "Mattermost-User-Id"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
	Foobar string
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	// any post-processing on configurations goes here

	return nil
}

func (c *Configuration) IsValid() error {
	// Add config validations here.
	// Check fot required fields, formats, etc.

	return nil
}
