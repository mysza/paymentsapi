package domain

import (
	"encoding/json"

	"github.com/google/uuid"
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

// CreateTestPayment creates a new instance of a payment
func CreateTestPayment() *Payment {
	return &Payment{
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Attributes: PaymentAttributes{
			Beneficiary: BeneficiaryPaymentParty{
				PaymentParty: PaymentParty{
					Account: Account{
						AccountNumber: "56781234",
						BankID:        "123123",
						BankIDCode:    "GBDSC",
					},
					AccountName:       "EJ Brown Black",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					Name:              "EJ Brown Black",
				},
				AccountType: 0,
			},
			Debtor: PaymentParty{
				Account: Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountName:       "EJ Brown Black",
				AccountNumberCode: "IBAN",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				Name:              "EJ Brown Black",
			},
			Sponsor: Account{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
			ChargesInformation: ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []Charge{
					Charge{Currency: "USD", Amount: "5.00"},
					Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			FX: FX{
				ContractReference: "FX123",
				ExchangeRate:      "2.00",
				OriginalAmount:    "100.12",
				OriginalCurrency:  "USD",
			},
			ProcessingDate:       "2017-01-18",
			Amount:               "100.12",
			Currency:             "USD",
			EndToEndReference:    "Some generic string",
			NumericReference:     "123456",
			PaymentID:            "123456789012345678",
			PaymentPurpose:       "Paying for goods/services",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			SchemePaymentType:    "InternetBanking",
			SchemePaymentSubType: "ImmediatePayment",
			Reference:            "Payment for Em's piano lessons",
		},
	}

}
