package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"

	"github.com/Brightscout/mattermost-plugin-boilerplate/server/command"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/config"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/controller"
	"github.com/Brightscout/mattermost-plugin-boilerplate/server/util"
)

type Plugin struct {
	plugin.MattermostPlugin

	handler http.Handler
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.setupStaticFileServer(); err != nil {
		p.API.LogError(err.Error())
		return err
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

func (p *Plugin) setupStaticFileServer() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	p.handler = http.FileServer(http.Dir(filepath.Dir(exe) + config.ServerExeToWebappRootPath))
	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost != nil {
		var configuration config.Configuration

		if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
			config.Mattermost.LogError("Error in LoadPluginConfiguration: " + err.Error())
			return err
		}

		if err := configuration.ProcessConfiguration(); err != nil {
			config.Mattermost.LogError("Error in ProcessConfiguration: " + err.Error())
			return err
		}

		if err := configuration.IsValid(); err != nil {
			config.Mattermost.LogError("Error in Validating Configuration: " + err.Error())
			return err
		}

		config.SetConfig(&configuration)
	}
	return nil
}

func (p *Plugin) registerCommands() error {
	for _, c := range command.Commands {
		if err := config.Mattermost.RegisterCommand(c.Command); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split, argErr := util.SplitArgs(args.Command)
	if argErr != nil {
		return util.CommandError(argErr.Error())
	}

	cmdName := split[0]
	var params []string

	if len(split) > 1 {
		params = split[1:]
	}

	commandConfig := command.Commands[cmdName]
	if commandConfig == nil {
		return nil, &model.AppError{Message: "Unknown command: [" + cmdName + "] encountered"}
	}

	context := p.prepareContext(args)
	if response, err := commandConfig.Validate(params, context); response != nil {
		return response, err
	}

	config.Mattermost.LogInfo("Executing command: " + cmdName + " with params: [" + strings.Join(params, ", ") + "]")
	return commandConfig.Execute(params, context)
}

func (p *Plugin) prepareContext(args *model.CommandArgs) command.Context {
	return command.Context{
		CommandArgs: args,
		Props:       make(map[string]interface{}),
	}
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured: " + err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	path := r.URL.Path
	endpoint := controller.Endpoints[path]

	if endpoint == nil {
		p.handler.ServeHTTP(w, r)
	} else if !endpoint.RequiresAuth || controller.Authenticated(w, r) {
		endpoint.Execute(w, r)
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
