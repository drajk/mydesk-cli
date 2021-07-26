package ticket

import (
	"fmt"
	"time"

	"github.com/drajk/mydesk-cli/pkg/file"
)

type Ticket struct {
	Id         string    `json:"_id"`
	CreatedAt  time.Time `json:"created_at"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	AssigneeId int       `json:"assignee_id"`
	Tags       []string  `json:"tags"`
}

type IStore interface {
	SearchById(string) []Ticket
	SearchByAssigneeId(int) []Ticket
	SearchByType(string) []Ticket
	SearchBySubject(string) []Ticket
	SearchByTag(string) []Ticket
}

type storeParams struct {
	tickets     []Ticket
	searchLimit int
}

func NewStore(filePath string, searchLimit int) IStore {
	var tickets []Ticket
	err := file.ReadJSON(filePath, &tickets)

	if err != nil {
		panic(fmt.Errorf("failed to read %s. %v", filePath, err))
	}

	return &storeParams{
		tickets:     tickets,
		searchLimit: searchLimit,
	}
}
