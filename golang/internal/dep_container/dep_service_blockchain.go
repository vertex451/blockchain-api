package dep_container

import (
	"github.com/sarulabs/di"
	"silverspase/blockchain-api/golang/internal/config"
	"silverspase/blockchain-api/golang/internal/service/blockchain"
)

const blockchainDefName = "blockchain-provider"

// RegisterBlockchainProvider registers Blockchain Provider.
func RegisterBlockchainProvider(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: blockchainDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(configDefName).(*config.Config)
			return blockchain.NewProvider(cfg.BlockchainUrl, cfg.CovalentApiKey), nil
		},
	})
}
