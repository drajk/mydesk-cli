package config

// IContainer Interface
type IContainer interface {
	ServiceName() string
}

type container struct {
	serviceName string
}

// NewContainer - creates new app container required parameters
func NewContainer(serviceName string) IContainer {
	return &container{
		serviceName: serviceName,
	}
}

//ServiceName returns the name of service
func (c *container) ServiceName() string {
	return c.serviceName
}
