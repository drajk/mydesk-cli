package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchByTag_Success(t *testing.T) {
	cases := []struct {
		tag            string
		expectedLength int
	}{
		{
			tag:            "Morroco",
			expectedLength: 0,
		},
		{
			tag:            "Canada",
			expectedLength: 0,
		},
		{
			tag:            "Mississippi",
			expectedLength: 1,
		},
		{
			tag:            "Alabama",
			expectedLength: 1,
		},
		{
			tag:            "Missouri",
			expectedLength: 1,
		},
	}

	for _, c := range cases {
		matchingTickets := getMockStore().SearchByTag(c.tag)
		assert.Equal(t, c.expectedLength, len(matchingTickets))
	}
}
