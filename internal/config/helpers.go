package config

import (
	"github.com/urfave/cli/v2"
)

func (app *App) WithCommands(commands ...*cli.Command) *App {
	app.commands = append(app.commands, commands...)
	return app
}
