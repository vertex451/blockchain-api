package usecase

import (
	"context"
	"errors"

	"silverspase/blockchain-api/golang/internal/service/blockchain"
	"silverspase/blockchain-api/golang/internal/service/blockchain/covalent"
)

type useCase struct {
	blockchainProvider blockchain.Provider
	//repoProvider goes here
}

func NewProvider(provider blockchain.Provider) Provider {
	return useCase{
		blockchainProvider: provider,
	}
}

func (u useCase) GetLatestBlockNumber(ctx context.Context) (int64, error) {
	return u.blockchainProvider.GetLatestBlockNumber(ctx)
}

func (u useCase) GetTransactionHistory(ctx context.Context, address string, chainID int) (*covalent.GetTransactionHistoryResponse, error) {
	switch chainID {
	case blockchain.CHAIN_ID:
		// here we should go to our indexed db for example
		return nil, errors.New("not implemented")
	default:
		return u.blockchainProvider.GetTransactionHistory(ctx, address, chainID)
	}
}

func (u useCase) GetTokenBalances(ctx context.Context, address string, chainID int) (*covalent.GetTokenBalancesResponse, error) {
	switch chainID {
	case blockchain.CHAIN_ID:
		// Im not sure if we would index in the db something like token balances...so we might remove this switch statement
		return nil, errors.New("not implemented")
	default:
		return u.blockchainProvider.GetTokenBalances(ctx, address, chainID)
	}
}
