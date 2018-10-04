package domain

// ChargesInformation holds information about charges applied to a payment
type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code" validate:"required"`
	SenderCharges           []Charge `json:"sender_charges" validate:"required,dive"`
	ReceiverChargesAmount   Amount   `json:"receiver_charges_amount" validate:"required,numeric"`
	ReceiverChargesCurrency Currency `json:"receiver_charges_currency" validate:"required,len=3,alpha"`
}
