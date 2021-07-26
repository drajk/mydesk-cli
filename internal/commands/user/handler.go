package user

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
		var searchResult []storeUser.User

		id := c.Int(ParamID)
		name := c.String(ParamName)

		switch {
		case id > 0:
			fmt.Printf("Searching users for %s with a value of %v \n", ParamID, id)
			searchResult = userStore.SearchById(id)

		case name != "":
			fmt.Printf("Searching users for %s with a value of %v \n", ParamName, name)
			searchResult = userStore.SearchByName(name)

		default:
			err = errors.New("invalid search param")
			return
		}

		if len(searchResult) < 1 {
			err = errors.New("no results found")
			return
		}

		fmt.Println("Displaying first 10 results")

		for _, user := range searchResult {
			matchingTickets := ticketStore.SearchByAssigneeId(user.Id)

			var assignedTickets []string
			for _, ticket := range matchingTickets {
				assignedTickets = append(assignedTickets, ticket.Subject)
			}

			formatted, err := format.ToIndentedKeyValue(
				"Id", fmt.Sprint(user.Id),
				"Name", user.Name,
				"Verified", fmt.Sprint(user.IsVerified),
				"CreatedAt", fmt.Sprint(user.CreatedAt),
				"Tickets", strings.Join(assignedTickets, ", "),
			)

			if err != nil {
				return err
			}

			fmt.Print(formatted)
		}

		return
	}
}
