package main

import (
	"github.com/drajk/mydesk-cli/internal"
	"github.com/drajk/mydesk-cli/internal/config"
)

const ServiceName = "MyDesk CLI"

func main() {
	container := config.NewContainer(ServiceName)
	internal.Start(container)
}
