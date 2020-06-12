package command

import (
	"github.com/mattermost/mattermost-server/model"

	"github.com/Brightscout/mattermost-plugin-boilerplate/server/config"
)

const (
	invalidCommand = "Invalid command parameters. Please use `" + config.CommandPrefix + " help` for more information."
)

var CommandHandler = Handler{
	Command: &model.Command{
		Trigger:     config.CommandPrefix,
		Description: "",
		DisplayName: "",
	},
	handlers: map[string]HandlerFunc{},
	defaultHandler: func(context *model.CommandArgs, args ...string) (*model.CommandResponse, *model.AppError) {
		return nil, nil
	},
}
