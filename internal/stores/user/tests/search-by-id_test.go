package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchByID_Success(t *testing.T) {
	cases := []struct {
		id             int
		expectedLength int
	}{
		{
			id:             72,
			expectedLength: 1,
		},
		{
			id:             65,
			expectedLength: 1,
		},
		{
			id:             44,
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchById(c.id)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
