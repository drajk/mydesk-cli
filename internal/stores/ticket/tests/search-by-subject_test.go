package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchBySubject_Success(t *testing.T) {
	cases := []struct {
		subject        string
		expectedLength int
	}{
		{
			subject:        "A store in morroco",
			expectedLength: 0,
		},
		{
			subject:        "A Problem in Denmark",
			expectedLength: 1,
		},
		{
			subject:        "Problem",
			expectedLength: 1,
		},
		{
			subject:        "Morroco",
			expectedLength: 0,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchBySubject(c.subject)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
