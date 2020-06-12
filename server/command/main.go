package command

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

type HandlerFunc func(context *model.CommandArgs, args ...string) (*model.CommandResponse, *model.AppError)

type Handler struct {
	Command        *model.Command
	handlers       map[string]HandlerFunc
	defaultHandler HandlerFunc
}

func (ch Handler) Handle(context *model.CommandArgs, args ...string) (*model.CommandResponse, *model.AppError) {
	for n := len(args); n > 0; n-- {
		h := ch.handlers[strings.Join(args[:n], "/")]
		if h != nil {
			return h(context, args[n:]...)
		}
	}
	return ch.defaultHandler(context, args...)
}

var Handlers = map[string]Handler{
	CommandHandler.Command.Trigger: CommandHandler,
}
