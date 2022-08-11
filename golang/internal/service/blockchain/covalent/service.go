package covalent

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://api.covalenthq.com/v1"

type service struct {
	apiKey string
}

func NewService(apiKey string) Provider {
	// TODO add ping

	return service{
		apiKey: apiKey,
	}
}

func (s service) GetTransactionHistory( // nolint: dupl
	ctx context.Context,
	address string,
	chainID int,
) (*GetTransactionHistoryResponse, error) {
	// curl -X GET https://api.covalenthq.com/v1/1/address/0xa79E63e78Eec28741e711f89A672A4C40876Ebf3/transactions_v2/?key=KEY
	requestUrl := fmt.Sprintf("%s/%v/address/%s/transactions_v2/?key=%s", BASE_URL, chainID, address, s.apiKey)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		return nil, err
	}

	var res GetTransactionHistoryResponse
	if err = json.Unmarshal(body, &res); err != nil { // Parse []byte to go struct pointer
		return nil, err
	}

	if res.Error {
		return nil, fmt.Errorf("covalent returned an error: %v", res.ErrorMessage)
	}

	return &res, nil
}

func (s service) GetTokenBalances(ctx context.Context, address string, chainID int) (*GetTokenBalancesResponse, error) { // nolint: dupl
	// curl -X GET https://api.covalenthq.com/v1/1/address/0xa79E63e78Eec28741e711f89A672A4C40876Ebf3/balances_v2/?key=KEY
	requestUrl := fmt.Sprintf("%s/%v/address/%s/balances_v2/?key=%s", BASE_URL, chainID, address, s.apiKey)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res GetTokenBalancesResponse
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if res.Error {
		return nil, fmt.Errorf("covalent returned an error: %v", res.ErrorMessage)
	}

	return &res, nil
}
