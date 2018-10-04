package model

// FX represents information about exchange rate in payment
type FX struct {
	ContractReference string   `json:"contract_reference"`
	ExchangeRate      string   `json:"exchange_rate"`
	OriginalAmount    Amount   `json:"original_amount"`
	OriginalCurrency  Currency `json:"original_currency"`
}
