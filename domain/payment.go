package domain

import (
	"encoding/json"
)

// Payment is the base data structure provided by the service.
// It describes a single payment registered in the system.
type Payment struct {
	ID             string            `json:"id" validate:"-"`
	OrganisationID string            `json:"organisation_id" validate:"required"`
	Attributes     PaymentAttributes `json:"attributes" validate:"required"`
}

// PaymentToByteSlice encodes the Payment to byte slice.
func PaymentToByteSlice(p *Payment) ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// PaymentFromByteSlice decodes the Payment from a byte slice.
func PaymentFromByteSlice(data []byte) (*Payment, error) {
	var p Payment
	err := json.Unmarshal(data, &p)
	return &p, err
}
