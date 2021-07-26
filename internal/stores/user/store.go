package user

import (
	"fmt"
	"time"

	"github.com/drajk/mydesk-cli/pkg/file"
)

type User struct {
	Id         int       `json:"_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	IsVerified bool      `json:"verified"`
}

type IStore interface {
	SearchById(int) []User
	SearchByName(string) []User
}

type storeParams struct {
	users       []User
	searchLimit int
}

func NewStore(filePath string, searchLimit int) IStore {
	var usersDB []User
	err := file.ReadJSON(filePath, &usersDB)

	if err != nil {
		panic(fmt.Errorf("failed to read %s. %v", filePath, err))
	}

	return &storeParams{
		users:       usersDB,
		searchLimit: searchLimit,
	}
}
