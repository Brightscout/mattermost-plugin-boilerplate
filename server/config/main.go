package config

import (
	"github.com/mattermost/mattermost-server/plugin"
	"go.uber.org/atomic"
)

const (
	CommandPrefix = PluginName

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"

	BotUserName    = ""
	BotDisplayName = ""
	BotDescription = ""
)

var (
	config     atomic.Value
	Mattermost plugin.API
	BotUserID  string
)

type Configuration struct {
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

// ProcessConfiguration is used for post-processing on the configuration.
func (c *Configuration) ProcessConfiguration() error {
	return nil
}

// IsValid is used for config validations.
func (c *Configuration) IsValid() error {
	return nil
}
