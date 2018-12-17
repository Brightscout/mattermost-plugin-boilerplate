package command

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
)

type Context struct {
	CommandArgs *model.CommandArgs
	Props       map[string]interface{}
}

type Config struct {
	Command  *model.Command
	HelpText string
	Execute  func([]string, Context) (*model.CommandResponse, *model.AppError)
	Validate func([]string, Context) (*model.CommandResponse, *model.AppError)
}

func (c *Config) Syntax() string {
	return fmt.Sprintf("/%s %s", c.Command.Trigger, c.Command.AutoCompleteHint)
}

var Commands map[string]*Config

func init() {
	Commands = map[string]*Config{
		// Add command mappings here.
		// Map trigger to corresponding Config object. Example -
		// SomeCommand().Trigger: SomeCommand()s
	}
}
