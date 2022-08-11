package covalent

import "context"

type Provider interface {
	GetTransactionHistory(ctx context.Context, address string, chainID int) (*GetTransactionHistoryResponse, error)
	GetTokenBalances(ctx context.Context, address string, chainID int) (*GetTokenBalancesResponse, error)
}
