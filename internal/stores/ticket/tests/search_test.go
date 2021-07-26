package tests

import (
	"testing"

	ticketStore "github.com/drajk/mydesk-cli/internal/stores/ticket"
	"github.com/stretchr/testify/assert"
)

func getMockStore() ticketStore.IStore {
	return ticketStore.NewMockStore(
		[]ticketStore.Ticket{
			{
				Id:         "2614576f-98fb-4031-9e13-beca7a6a73ee",
				Type:       "question",
				Subject:    "A Nuisance in Aruba",
				Tags:       []string{"Mississippi", "Marshall Islands", "South Dakota", "Montana"},
				AssigneeId: 101,
			},
		}, 10,
	)
}

func Test_SearchByID_Success(t *testing.T) {
	cases := []struct {
		id             string
		expectedLength int
	}{
		{
			id:             "2614576f-98fb-4031-9e13-beca7a6a73ee",
			expectedLength: 1,
		},
		{
			id:             "some-random-keys",
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchById(c.id)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}

func Test_SearchByAssigneeId_Success(t *testing.T) {
	cases := []struct {
		assigneeId     int
		expectedLength int
	}{
		{
			assigneeId:     101,
			expectedLength: 1,
		},
		{
			assigneeId:     1,
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchByAssigneeId(c.assigneeId)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
