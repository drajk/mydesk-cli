package config

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Route define structure that handle request for route
type App struct {
	commands []*cli.Command
}

func NewApp() *App {
	return &App{
		commands: make([]*cli.Command, 0),
	}
}

func (a *App) Start() {
	app := &cli.App{
		Name:                 ServiceName,
		UsageText:            InfoUsage,
		Commands:             a.commands,
		EnableBashCompletion: true,
	}

	err := app.Run(os.Args)

	if len(os.Args) < 2 {
		fmt.Println(ErrorUseHelp)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(ErrorUseHelp)
	}
}
