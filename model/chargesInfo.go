package model

// ChargesInformation holds information about charges applied to a payment
type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code"`
	SenderCharges           []Charge `json:"sender_charges"`
	ReceiverChargesAmount   Amount   `json:"receiver_charges_amount"`
	ReceiverChargesCurrency Currency `json:"receiver_charges_currency"`
}
