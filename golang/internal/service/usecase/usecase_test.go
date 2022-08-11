package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"silverspase/blockchain-api/golang/internal/service/blockchain"
)

func TestGetLatestBlockNumber(t *testing.T) {
	var expectedBlockNumber int64 = 123456

	// create an instance which will replace actual blockchain.GetLatestBlockNumber() call
	blockchainMockProvider := &blockchain.MockProvider{}

	// here we say that we expect blockchainMockProvider to receive a call of GetLatestBlockNumber method
	// and define the return values
	blockchainMockProvider.On("GetLatestBlockNumber", mock.Anything).
		Return(expectedBlockNumber, nil)

	// to call usecase.GetLatestBlockNumber we need an usecase instance.
	uc := NewProvider(blockchainMockProvider)
	blockNumber, err := uc.GetLatestBlockNumber(context.Background())

	// usual tests assertions
	assert.Nil(t, err)
	assert.Equal(t, blockNumber, expectedBlockNumber)
}
