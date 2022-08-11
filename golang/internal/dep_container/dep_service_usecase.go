package dep_container

import (
	"github.com/sarulabs/di"
	"silverspase/blockchain-api/golang/internal/service/blockchain"
	"silverspase/blockchain-api/golang/internal/service/usecase"
)

const usecaseDefName = "usecase-provider"

// RegisterUsecaseProvider registers Usecase Provider.
func RegisterUsecaseProvider(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: usecaseDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			blockchainProvider := ctn.Get(blockchainDefName).(blockchain.Provider)
			return usecase.NewProvider(blockchainProvider), nil
		},
	})
}
