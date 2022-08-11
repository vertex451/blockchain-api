package covalent

type GetTransactionHistoryResponse struct {
	Data struct {
		Address string
		ChainID int `json:"chain_id"`
		Items   []Transaction
	}
	Error        bool
	ErrorMessage string `json:"error_message"`
	ErrorCode    int
}

type Transaction struct {
	TxHash        string
	Successful    bool
	From          string `json:"from_address"`
	To            string `json:"to_address"`
	Value         string
	BlockSignedAt string `json:"block_signed_at"`
}

type GetTokenBalancesResponse struct {
	Data struct {
		Address string
		ChainID int `json:"chain_id"`
		Items   []TokenBalance
	}
	Error        bool
	ErrorMessage string `json:"error_message"`
	ErrorCode    int
}

type TokenBalance struct {
	LogoURL              string  `json:"logo_url"`
	Balance              string  `json:"balance"`
	ContractTickerSymbol string  `json:"contract_ticker_symbol"`
	ContractDecimals     int     `json:"contract_decimals"`
	QuoteRate            float64 `json:"quote_rate"`
}
