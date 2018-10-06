package domain

// ChargesInformation holds information about charges applied to a payment
type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code" validate:"required"`
	SenderCharges           []Charge `json:"sender_charges" validate:"required,dive" gorm:"many2many:chargesinfo_sendercharges"`
	ReceiverChargesAmount   string   `json:"receiver_charges_amount" validate:"required,numeric"`
	ReceiverChargesCurrency string   `json:"receiver_charges_currency" validate:"required,len=3,alpha"`
}
