package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
