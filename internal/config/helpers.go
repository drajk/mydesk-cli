package config

import (
	"github.com/urfave/cli/v2"
)

func (app *App) WithCommand(commands ...*cli.Command) *App {
	app.commands = append(app.commands, commands...)
	return app
}
