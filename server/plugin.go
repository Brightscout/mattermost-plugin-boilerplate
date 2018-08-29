package main

import (
	"net/http"

	"fmt"
	"os"
	"path/filepath"

	"github.com/Brightscout/mattermost-plugin-boilerplate/server/config"
	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin ...
type Plugin struct {
	plugin.MattermostPlugin

	handler http.Handler
}

// OnActivate ...
func (p *Plugin) OnActivate() error {
	fmt.Println("activated")

	config.Mattermost = p.API

	if err := p.setupStaticFileServer(); err != nil {
		return err
	}

	//if err := p.OnConfigurationChange(); err != nil {
	//	return err
	//}

	return nil
}

func (p *Plugin) setupStaticFileServer() error {
	var exe, err = os.Executable()
	if err != nil {
		return err
	}
	p.handler = http.FileServer(http.Dir(filepath.Dir(exe) + config.ServerExeToWebappRootPath))
	return nil
}

// OnConfigurationChange ...
func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost != nil {
		var configuration config.Configuration

		if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
			return err
		}

		if err := configuration.ProcessConfiguration(); err != nil {
			return err
		}

		if err := configuration.IsValid(); err != nil {
			return err
		}

		config.SetConfig(&configuration)

		fmt.Println(config.GetConfig().Foobar)
	}
	return nil
}

//
//func (p *Plugin) RegisterCommands() error {
//	for _, c := range command.Commands {
//		if err := config.Mattermost.RegisterCommand(c.Command); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

//func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
//	// cant use strings.split as it includes empty string if deliminator
//	// is the last character in input string
//	var split, argErr = util.SplitArgs(args.Command)
//	if argErr != nil {
//		return &model.CommandResponse{
//			Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
//			Text: argErr.Error(),
//		}, nil
//	}
//
//	var function = split[0]
//	var params []string
//
//	if len(split) > 2 {
//		params = split[2:]
//	}
//
//	var commandConfig = command.Commands[function]
//	if commandConfig == nil {
//		return nil, &model.AppError{Message: "Unknown command: [" + function + "] encountered"}
//	}
//
//	var context = p.prepareContext(args)
//	if err := commandConfig.Validate(params, context); err != nil {
//		return err, nil
//	}
//
//	return commandConfig.Execute(params, context)
//}
//
//func (p *Plugin) prepareContext(args *model.CommandArgs) command.CommandContext {
//	return command.CommandContext{
//		CommandArgs: args,
//	}
//}
//

// ServeHTTP ...
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

	//var path = r.URL.Path
	//var endpoint = controller.Endpoints[path]
	//
	//if endpoint != nil {
	//	if endpoint.RequiresAuth {
	//		if controller.Authenticated(w, r) {
	//			endpoint.Execute(w, r)
	//		}
	//	} else {
	//		endpoint.Execute(w, r)
	//	}
	//} else {
	//	p.handler.ServeHTTP(w, r)
	//}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
