package usecase

import (
	"context"

	"silverspase/blockchain-api/golang/internal/service/blockchain/covalent"
)

type Provider interface {
	GetLatestBlockNumber(ctx context.Context) (int64, error)
	GetTransactionHistory(ctx context.Context, address string, chainID int) (*covalent.GetTransactionHistoryResponse, error)
	GetTokenBalances(ctx context.Context, address string, chainID int) (*covalent.GetTokenBalancesResponse, error)
}
