package tests

import (
	userStore "github.com/drajk/mydesk-cli/internal/stores/user"
)

func getMockStore() userStore.IStore {
	return userStore.NewMockStore(
		[]userStore.User{
			{
				Id:         72,
				Name:       "Valentine Ashley",
				IsVerified: true,
			},
			{
				Id:         65,
				Name:       "Boone Cooke",
				IsVerified: true,
			},
		}, 10,
	)
}
