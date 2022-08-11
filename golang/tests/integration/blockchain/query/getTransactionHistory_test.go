package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"silverspase/blockchain-api/golang/tests/clients/graphql"
)

type GetTransactionHistoryTestSuite struct {
	suite.Suite
	client *graphql.Client
}

func TestAddLotToUserFavoritesTestSuite(t *testing.T) {
	suite.Run(t, new(GetTransactionHistoryTestSuite))
}

// SetupSuite prepare everything for tests.
func (s *GetTransactionHistoryTestSuite) SetupSuite() {
	s.client = graphql.NewClient(fmt.Sprintf("%s/query", "http://localhost:8080"))
}

func (s *GetTransactionHistoryTestSuite) TestGetTransactionHistory_Ok() {
	var getTransactionHistoryQuery struct {
		Response struct {
			Address string `graphql:"address"`
			ChainID int    `graphql:"chainID"`
			Items   []struct {
				From string `graphql:"From"`
			} `graphql:"items"`
		} `graphql:"getTransactionHistory(address: $address, chain_id: $chain_id)"`
	}

	err := s.client.Query(&getTransactionHistoryQuery, map[string]interface{}{
		"address":  graphql.String("0xa79E63e78Eec28741e711f89A672A4C40876Ebf3"),
		"chain_id": graphql.Int(1),
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), getTransactionHistoryQuery.Response.Items[0].From, "0xa79e63e78eec28741e711f89a672a4c40876ebf3")
}
