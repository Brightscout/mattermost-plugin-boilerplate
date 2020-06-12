package main

import (
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"

	"github.com/Brightscout/mattermost-plugin-boilerplate/server/command"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/config"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/controller"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/util"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.initBotUser(); err != nil {
		config.Mattermost.LogError("Failed to create a bot user", "Error", err.Error())
	}

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	if err := p.registerCommands(); err != nil {
		config.Mattermost.LogError(err.Error())
		return err
	}

	return nil
}

func (p *Plugin) initBotUser() error {
	botUserID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    config.BotUserName,
		DisplayName: config.BotDisplayName,
		Description: config.BotDescription,
	})
	if err != nil {
		return errors.Wrap(err, "failed to ensure bot")
	}

	config.BotUserID = botUserID
	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	// If OnActivate has not been run yet.
	if config.Mattermost == nil {
		return nil
	}
	var configuration config.Configuration

	if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
		config.Mattermost.LogError("Error in LoadPluginConfiguration.", "Error", err.Error())
		return err
	}

	if err := configuration.ProcessConfiguration(); err != nil {
		config.Mattermost.LogError("Error in ProcessConfiguration.", "Error", err.Error())
		return err
	}

	if err := configuration.IsValid(); err != nil {
		config.Mattermost.LogError("Error in Validating Configuration.", "Error", err.Error())
		return err
	}

	config.SetConfig(&configuration)
	return nil
}

func (p *Plugin) registerCommands() error {
	for trigger, handler := range command.Handlers {
		if err := config.Mattermost.RegisterCommand(handler.Command); err != nil {
			return errors.Wrap(err, "failed to register slash command: "+trigger)
		}
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split, argErr := util.SplitArgs(args.Command)
	if argErr != nil {
		return util.SendEphemeralCommandResponse(argErr.Error())
	}

	cmdName := split[0][1:]
	var params []string

	if len(split) > 1 {
		params = split[1:]
	}

	handler, ok := command.Handlers[cmdName]
	if !ok {
		return util.SendEphemeralCommandResponse("Unknown command: [" + cmdName + "] encountered")
	}

	config.Mattermost.LogDebug("Executing command: " + cmdName + " with params: [" + strings.Join(params, ", ") + "]")
	return handler.Handle(args, params...)
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("New request:", "Host", r.Host, "RequestURI", r.RequestURI, "Method", r.Method)

	conf := config.GetConfig()
	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured.", "Error", err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	controller.InitAPI().ServeHTTP(w, r)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
