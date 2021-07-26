package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_WithCommand_Should_Add_Commmand(t *testing.T) {
	// arrange
	app := NewApp()

	// act and assert
	assert.Equal(t, 0, len(app.commands))

	app.WithCommand(&cli.Command{})
	assert.Equal(t, 1, len(app.commands))
}
