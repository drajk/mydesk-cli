package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
