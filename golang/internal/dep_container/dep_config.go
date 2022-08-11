package dep_container

import (
	"github.com/sarulabs/di"
	"silverspase/blockchain-api/golang/internal/config"
)

const configDefName = "config"

// RegisterConfig registers Config dependency.
func RegisterConfig(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: configDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.LoadConfig(), nil
		},
	})
}

// GetConfig returns COnfig object.
func (c Container) GetConfig() *config.Config {
	return c.container.Get(configDefName).(*config.Config)
}
