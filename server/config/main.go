package config

import (
	"github.com/mattermost/mattermost-server/plugin"
	"go.uber.org/atomic"
)

const (
	// PluginName ...
	PluginName                = "boilerplate"

	// CommandPrefix ...
	CommandPrefix             = PluginName

	// URLMappingKeyPrefix ...
	URLMappingKeyPrefix = "url_"

	// ServerExeToWebappRootPath ...
	ServerExeToWebappRootPath = "/../webapp"

	// URLPluginBase ...
	URLPluginBase = "/plugins/" + PluginName

	// URLStaticBase ...
	URLStaticBase = URLPluginBase + "/static"
)

var (
	config     atomic.Value

	// Mattermost ...
	Mattermost plugin.API
)

// Configuration ...
type Configuration struct {
	Foobar string
}

// GetConfig ...
func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

// SetConfig ...
func SetConfig(c *Configuration) {
	config.Store(c)
}

// ProcessConfiguration ...
// any post-processing on configurations goes here
func (c *Configuration) ProcessConfiguration() error {

	return nil
}

// IsValid ...
// Add config validations here.
// Check fot required fields, formats, etc.
func (c *Configuration) IsValid() error {

	return nil
}
