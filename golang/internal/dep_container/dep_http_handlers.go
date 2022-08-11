package dep_container

import (
	"github.com/sarulabs/di"
	"silverspase/blockchain-api/golang/internal/service/transport/http_handlers"
)

const httpHandlersDefName = "http-handlers-handlers"

// RegisterHttpHandlers registers service which provides http-handlers handlers
func RegisterHttpHandlers(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: httpHandlersDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			return http_handlers.NewController(), nil
		},
	})
}
