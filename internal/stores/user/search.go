package user

import "strings"

func (u *storeParams) SearchById(id int) []User {
	return filter(u.users, u.searchLimit, func(item User) bool {
		return item.Id == id
	})
}

func (u *storeParams) SearchByName(name string) []User {
	return filter(u.users, u.searchLimit, func(item User) bool {
		return strings.Contains(item.Name, name)
	})
}

func filter(source []User, maxItems int, filterFn func(User) bool) []User {
	filtered := make([]User, 0)

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
