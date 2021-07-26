package ticket

import (
	"errors"
	"fmt"

	storeTickets "github.com/drajk/mydesk-cli/internal/stores/ticket"
	"github.com/urfave/cli/v2"
)

func handle(ticketStore storeTickets.IStore) func(c *cli.Context) error {
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
			fmt.Println()

			fmt.Printf("_id: \t %v \n", ticket.Id)
			fmt.Printf("assignee_id: \t %v \n", ticket.AssigneeId)
			fmt.Printf("type: \t %v \n", ticket.Type)
			fmt.Printf("subject: \t %v \n", ticket.Subject)
			fmt.Printf("tags: \t %v \n", ticket.Tags)
			fmt.Printf("created_at: \t %v \n", ticket.CreatedAt.Format("2006-01-02T15:04:05"))

			fmt.Println()
		}

		return
	}
}
