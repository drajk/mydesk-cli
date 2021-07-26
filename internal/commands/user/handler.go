package user

import (
	"errors"
	"fmt"

	storeUser "github.com/drajk/mydesk-cli/internal/stores/user"
	"github.com/urfave/cli/v2"
)

func handle(userStore storeUser.IStore) func(c *cli.Context) error {
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
			fmt.Println()

			fmt.Printf("_id: \t %v \n", user.Id)
			fmt.Printf("name: \t %v \n", user.Name)
			fmt.Printf("verified: \t %v \n", user.IsVerified)
			fmt.Printf("created_at: \t %v \n", user.CreatedAt.Format("2006-01-02T15:04:05"))

			fmt.Println()
		}

		return
	}
}
