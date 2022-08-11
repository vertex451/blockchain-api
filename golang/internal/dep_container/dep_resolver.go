package dep_container

import (
	"github.com/sarulabs/di"
	"silverspase/blockchain-api/golang/internal/service/transport/graphql"
	"silverspase/blockchain-api/golang/internal/service/usecase"
)

const resolverServiceDefName = "resolver"

// RegisterResolver registers Resolver dependency.
func RegisterResolver(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: resolverServiceDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			usecaseProvider := ctn.Get(usecaseDefName).(usecase.Provider)

			return graphql.NewResolver(usecaseProvider), nil
		},
	})
}
