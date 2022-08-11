package dep_container

import (
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

type dep func(builder *di.Builder) error

var dependencyList = []dep{
	RegisterBlockchainProvider,
	RegisterConfig,
	RegisterGraphqlServer,
	RegisterHTTPServer,
	RegisterHttpHandlers,
	RegisterLogger,
	RegisterResolver,
	RegisterUsecaseProvider,
}

// Container represents DI container.
type Container struct {
	container di.Container
}

// New creates new Application Dependency Container.
func New() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	c := Container{}
	for _, dep := range dependencyList {
		if err := dep(builder); err != nil {
			return nil, err
		}
	}

	c.container = builder.Build()

	c.InitLogger()

	return &c, nil
}

// Delete deletes current container.
func (c Container) Delete() error {
	if err := c.container.Delete(); err != nil {
		return errors.Wrap(err, "error deleting container")
	}
	return nil
}
