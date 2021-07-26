package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToIndentedKeyValue_Should_Work_With_Zero_Arguments(t *testing.T) {
	assert.Equal(t, "", ToIndentedKeyValue())
}

func Test_ToIndentedKeyValue_Should_Work_With_Even_Number_Of_Arguments(t *testing.T) {
	assert.Equal(t, "key: \t value \n\n", ToIndentedKeyValue("key", "value"))
}

func Test_ToIndentedKeyValue_Should_Return_Empty_When_Even_Number_Of_Arguments(t *testing.T) {
	assert.Equal(t, "", ToIndentedKeyValue("key", "value", "random"))
}
