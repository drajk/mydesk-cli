package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchByName_Success(t *testing.T) {
	cases := []struct {
		userName       string
		expectedLength int
	}{
		{
			userName:       "Valentine",
			expectedLength: 1,
		},
		{
			userName:       "Ashley",
			expectedLength: 1,
		},
		{
			userName:       "ashley",
			expectedLength: 0,
		},
		{
			userName:       "Boone",
			expectedLength: 1,
		},
		{
			userName:       "Carter",
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchByName(c.userName)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
