package command

import "github.com/mattermost/mattermost-server/model"

type CommandContext struct {
	CommandArgs *model.CommandArgs
}

type Config struct {
	Command  *model.Command
	Syntax   string
	Execute  func([]string, CommandContext) (*model.CommandResponse, *model.AppError)
	Validate func([]string, CommandContext) *model.CommandResponse
}

var Commands = map[string]*Config{
	// Add command mappings here.
	// Map trigger to corresponding Config object. Example -
	// SomeCommand().Trigger: SomeCommand()s
}
