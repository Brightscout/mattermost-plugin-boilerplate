package command

import "github.com/mattermost/mattermost-server/model"

// Context ...
type Context struct {
	CommandArgs *model.CommandArgs
}

// Config ...
type Config struct {
	Command  *model.Command
	Syntax   string
	Execute  func([]string, Context) (*model.CommandResponse, *model.AppError)
	Validate func([]string, Context) *model.CommandResponse
}

// Commands ...
var Commands = map[string]*Config{
	// Add command mappings here.
	// Map trigger to corresponding Config object. Example -
	// SomeCommand().Trigger: SomeCommand()s
}
