package internal

import (
	"fmt"

	"github.com/drajk/mydesk-cli/internal/config"
)

//Start starts the cli with given configuration
func Start(container config.IContainer) {
	fmt.Println("Hello, " + container.ServiceName() + "!")
}
