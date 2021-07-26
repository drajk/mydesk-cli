package config

import (
	"github.com/drajk/mydesk-cli/internal/stores/ticket"
	"github.com/drajk/mydesk-cli/internal/stores/user"
)

// IContainer Interface
type IContainer interface {
	ServiceName() string
	UserStore() user.IStore
	TicketStore() ticket.IStore
}

type container struct {
	serviceName string
	userStore   user.IStore
	ticketStore ticket.IStore
}

// NewContainer - creates new app container required parameters
func NewContainer(serviceName string) IContainer {
	return &container{
		serviceName: serviceName,
		userStore:   user.NewStore("mocks/users.json", 10),
		ticketStore: ticket.NewStore("mocks/tickets.json", 10),
	}
}

//ServiceName returns the name of service
func (c *container) ServiceName() string {
	return c.serviceName
}

//UserStore returns user store
func (c *container) UserStore() user.IStore {
	return c.userStore
}

//TicketStore returns ticket store
func (c *container) TicketStore() ticket.IStore {
	return c.ticketStore
}
