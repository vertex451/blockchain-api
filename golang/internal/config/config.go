package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"silverspase/blockchain-api/golang/pkg/config"
)

// Support application environments.
const (
	configPrefix = "APP"
)

// Config represents main application Config object to hold all ENV parameters.
type Config struct {
	Env            string `split_words:"true" default:"local"`
	HttpPort       string `split_words:"true" default:"8080"`
	BlockchainUrl  string `split_words:"true" default:"https://mainnet.infura.io/v3"`
	CovalentApiKey string `split_words:"true" required:"true"`
}

// LoadConfig loads config from ENV variables.
func LoadConfig() *Config {
	if err := config.PopulateFromDotEnv(); err != nil {
		zap.L().Error("failed to process config", zap.Error(err))
		panic(fmt.Sprintf("failed to process config: %v", err))
	}

	var cfg Config
	if err := envconfig.Process(configPrefix, &cfg); err != nil {
		zap.L().Error("failed to process config", zap.Error(err))
		panic(fmt.Sprintf("failed to process config: %v", err))
	}

	return &cfg
}
