package ticket

import (
	ticketStore "github.com/drajk/mydesk-cli/internal/stores/ticket"
	"github.com/urfave/cli/v2"
)

type container interface {
	TicketStore() ticketStore.IStore
}

const (
	Name  = "ticket"
	Usage = "Search tickets with id, type, user, tags and other parameters"
)

const (
	ParamID         = "id"
	ParamAssigneeID = "assigneeid"
	ParamType       = "type"
	ParamSubject    = "subject"
	ParamTag        = "tag"
)

func Command(c container) *cli.Command {
	return &cli.Command{
		Name:   Name,
		Usage:  Usage,
		Flags:  Flags(),
		Action: handle(c.TicketStore()),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  ParamID,
			Usage: "--id \"2614576f-98fb-4031-9e13-beca7a6a73ee\"",
		},
		&cli.IntFlag{
			Name:  ParamAssigneeID,
			Usage: "--assigneeid 21",
		},
		&cli.StringFlag{
			Name:  ParamType,
			Usage: "--type \"question\"",
		},
		&cli.StringFlag{
			Name:  ParamSubject,
			Usage: "--subject \"A Nuisance in Kiribati\" or  --subject \"Nuisance\" or --subject \"Kiribati\"",
		},
		&cli.StringFlag{
			Name:  ParamTag,
			Usage: "--tag \"Arizona\" or --tag \"Delaware\" or --tag \"New Hampshire\"",
		},
	}
}
