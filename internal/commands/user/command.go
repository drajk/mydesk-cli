package user

import (
	ticketStore "github.com/drajk/mydesk-cli/internal/stores/ticket"
	userStore "github.com/drajk/mydesk-cli/internal/stores/user"
	"github.com/urfave/cli/v2"
)

type container interface {
	UserStore() userStore.IStore
	TicketStore() ticketStore.IStore
}

const (
	Name  = "user"
	Usage = "Search user with id, name and other parameters"
)

const (
	ParamID   = "id"
	ParamName = "name"
)

func Command(c container) *cli.Command {
	return &cli.Command{
		Name:   Name,
		Usage:  Usage,
		Flags:  Flags(),
		Action: handle(c.UserStore(), c.TicketStore()),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:  ParamID,
			Usage: "--id 1",
		},
		&cli.StringFlag{
			Name:  ParamName,
			Usage: "--name \"Francisca Morgan\"",
		},
	}
}
