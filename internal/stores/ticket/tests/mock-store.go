package tests

import (
	ticketStore "github.com/drajk/mydesk-cli/internal/stores/ticket"
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
			{
				Id:         "530bc434-9984-4a54-8a74-83433d3da340",
				Type:       "incident",
				Subject:    "A Problem in Denmark",
				Tags:       []string{"Connecticut", "Arkansas", "Missouri", "Alabama"},
				AssigneeId: 22,
			},
		}, 10,
	)
}
