package internal

import (
	"github.com/drajk/mydesk-cli/internal/commands/ticket"
	"github.com/drajk/mydesk-cli/internal/commands/user"
	"github.com/drajk/mydesk-cli/internal/config"
)

//Start - creates app container and starts the app
func Start() {
	container := config.NewContainer(config.ServiceName)

	app := config.
		NewApp().
		WithCommand(user.Command(container)).
		WithCommand(ticket.Command(container))

	app.Start()
}
