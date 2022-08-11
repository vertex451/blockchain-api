package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"silverspase/blockchain-api/golang/internal/service/blockchain/covalent"
)

const (
	CHAIN_ID = 941
)

type client struct {
	ethClient *ethclient.Client
	covalent  covalent.Provider
}

func NewProvider(url, covalentApiKey string) Provider {
	conn, err := ethclient.Dial(url)
	if err != nil {
		zap.L().Fatal("failed to connect to blockchain",
			zap.Error(err), zap.String("url", url))
	}

	zap.L().Info("eth client is connected", zap.String("url", url))

	return client{
		ethClient: conn,
		covalent:  covalent.NewService(covalentApiKey),
	}
}

func (c client) GetLatestBlockNumber(ctx context.Context) (int64, error) {
	header, err := c.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Int64(), nil
}

func (c client) GetTransactionHistory(ctx context.Context, address string, chainID int) (*covalent.GetTransactionHistoryResponse, error) {
	return c.covalent.GetTransactionHistory(ctx, address, chainID)
}

func (c client) GetTokenBalances(ctx context.Context, address string, chainID int) (*covalent.GetTokenBalancesResponse, error) {
	return c.covalent.GetTokenBalances(ctx, address, chainID)
}
