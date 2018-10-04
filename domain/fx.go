package domain

// FX represents information about exchange rate in payment
type FX struct {
	ContractReference string   `json:"contract_reference" validate:"required,alphanum"`
	ExchangeRate      string   `json:"exchange_rate" validate:"required,numeric"`
	OriginalAmount    Amount   `json:"original_amount" validate:"required,numeric"`
	OriginalCurrency  Currency `json:"original_currency" validate:"required,len=3,alpha"`
}
