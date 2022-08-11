package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	gen_models "silverspase/blockchain-api/golang/internal/service/transport/graphql/models"
)

// GetLatestBlockNumber is the resolver for the getLatestBlockNumber field.
func (r *queryResolver) GetLatestBlockNumber(ctx context.Context) (int, error) {
	latestBlock, err := r.usecase.GetLatestBlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return int(latestBlock), nil
}

// GetTransactionHistory is the resolver for the getTransactionHistory field.
func (r *queryResolver) GetTransactionHistory(ctx context.Context, address string, chainID int) (*gen_models.GetTransactionHistoryResponse, error) {
	res, err := r.usecase.GetTransactionHistory(ctx, address, chainID)
	if err != nil {
		return nil, err
	}

	var items []*gen_models.Transaction
	for _, item := range res.Data.Items {
		items = append(items, &gen_models.Transaction{
			From:          item.From,
			To:            item.To,
			Value:         item.Value,
			BlockSignedAt: item.BlockSignedAt,
		})
	}

	return &gen_models.GetTransactionHistoryResponse{
		Address: res.Data.Address,
		ChainID: res.Data.ChainID,
		Items:   items,
	}, nil
}

// GetTokenBalances is the resolver for the getTokenBalances field.
func (r *queryResolver) GetTokenBalances(ctx context.Context, address string, chainID int) (*gen_models.GetTokenBalancesResponse, error) {
	res, err := r.usecase.GetTokenBalances(ctx, address, chainID)
	if err != nil {
		return nil, err
	}

	var items []*gen_models.TokenBalance
	for _, item := range res.Data.Items {
		items = append(items, &gen_models.TokenBalance{
			LogoURL:              item.LogoURL,
			Balance:              item.Balance,
			ContractTickerSymbol: item.ContractTickerSymbol,
			ContractDecimals:     item.ContractDecimals,
			QuoteRate:            item.QuoteRate,
		})
	}

	return &gen_models.GetTokenBalancesResponse{
		Address: res.Data.Address,
		ChainID: res.Data.ChainID,
		Items:   items,
	}, nil
}
