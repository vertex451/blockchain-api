package dep_container

import (
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"silverspase/blockchain-api/golang/pkg/logger"
)

const loggerDefName = "logger"

// RegisterLogger registers Logger dependency.
func RegisterLogger(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: loggerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger.Init(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*zap.Logger).Sync()
			return nil
		},
	})
}

// InitLogger init Logger.
func (c Container) InitLogger() {
	c.container.Get(loggerDefName)
}
