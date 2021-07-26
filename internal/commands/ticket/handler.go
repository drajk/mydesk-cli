package ticket

import (
	"errors"
	"fmt"
	"strings"

	storeTickets "github.com/drajk/mydesk-cli/internal/stores/ticket"
	storeUser "github.com/drajk/mydesk-cli/internal/stores/user"

	"github.com/drajk/mydesk-cli/pkg/format"
	"github.com/urfave/cli/v2"
)

func handle(userStore storeUser.IStore, ticketStore storeTickets.IStore) func(c *cli.Context) error {
	return func(c *cli.Context) (err error) {
		var searchResult []storeTickets.Ticket

		id := c.String(ParamID)
		assigneeId := c.Int(ParamAssigneeID)
		ticketType := c.String(ParamType)
		subject := c.String(ParamSubject)
		tag := c.String(ParamTag)

		switch {
		case id != "":
			fmt.Printf("Searching tickets for %s with a value of %v \n", ParamID, id)
			searchResult = ticketStore.SearchById(id)

		case assigneeId > 0:
			fmt.Printf("Searching tickets for %s with a value of %v \n", ParamAssigneeID, assigneeId)
			searchResult = ticketStore.SearchByAssigneeId(assigneeId)

		case ticketType != "":
			fmt.Printf("Searching tickets for %s with a value of %v \n", ParamType, ticketType)
			searchResult = ticketStore.SearchByType(ticketType)

		case subject != "":
			fmt.Printf("Searching tickets for %s with a value of %v \n", ParamSubject, subject)
			searchResult = ticketStore.SearchBySubject(subject)

		case tag != "":
			fmt.Printf("Searching tickets for %s with a value of %v \n", ParamTag, tag)
			searchResult = ticketStore.SearchByTag(tag)

		default:
			err = errors.New("invalid search param")
			return
		}

		if len(searchResult) < 1 {
			err = errors.New("no results found")
			return
		}

		fmt.Println("Displaying first 10 results")

		for _, ticket := range searchResult {
			matchingUsers := userStore.SearchById(ticket.AssigneeId)

			assigneeName := ""
			if len(matchingUsers) > 0 {
				assigneeName = matchingUsers[0].Name
			}

			formatted, err := format.ToIndentedKeyValue(
				"Id", ticket.Id,
				"Type", ticket.Type,
				"Subject", ticket.Subject,
				"Tags", strings.Join(ticket.Tags, ", "),
				"Assignee Id", fmt.Sprint(ticket.AssigneeId),
				"Assignee Name", assigneeName,
			)

			if err != nil {
				return err
			}

			fmt.Print(formatted)
		}

		return
	}
}
