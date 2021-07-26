package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	assert.True(t, Contains([]string{"one", "two"}, "one"))
	assert.False(t, Contains([]string{"one", "two"}, "three"))
	assert.False(t, Contains([]string{"one", "two"}, ""))
}
