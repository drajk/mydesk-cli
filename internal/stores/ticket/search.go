package ticket

import (
	"strings"

	"github.com/drajk/mydesk-cli/pkg/array"
)

func (u *storeParams) SearchById(id string) []Ticket {
	return filter(u.tickets, u.searchLimit, func(item Ticket) bool {
		return item.Id == id
	})
}

func (u *storeParams) SearchByAssigneeId(userId int) []Ticket {
	return filter(u.tickets, u.searchLimit, func(item Ticket) bool {
		return item.AssigneeId == userId
	})
}

func (u *storeParams) SearchByType(ticketType string) []Ticket {
	return filter(u.tickets, u.searchLimit, func(item Ticket) bool {
		return strings.Contains(item.Type, ticketType)
	})
}

func (u *storeParams) SearchBySubject(subject string) []Ticket {
	return filter(u.tickets, u.searchLimit, func(item Ticket) bool {
		return strings.Contains(item.Subject, subject)
	})
}

func (u *storeParams) SearchByTag(tag string) []Ticket {
	return filter(u.tickets, u.searchLimit, func(item Ticket) bool {
		return array.Contains(item.Tags, tag)
	})
}

func filter(source []Ticket, maxItems int, filterFn func(Ticket) bool) []Ticket {
	filtered := make([]Ticket, 0)

	for _, item := range source {
		if len(filtered) >= maxItems {
			break
		}

		if filterFn(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
