package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchByType_Success(t *testing.T) {
	cases := []struct {
		ticketType     string
		expectedLength int
	}{
		{
			ticketType:     "question",
			expectedLength: 1,
		},
		{
			ticketType:     "incident",
			expectedLength: 1,
		},
		{
			ticketType:     "ques",
			expectedLength: 1,
		},
		{
			ticketType:     "inci",
			expectedLength: 1,
		},
		{
			ticketType:     "emergency",
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchByType(c.ticketType)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
