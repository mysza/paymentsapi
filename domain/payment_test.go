package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestPayment() *Payment {
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

func TestPaymentMarshalling(t *testing.T) {
	payment := createTestPayment()
	assert := assert.New(t)
	var (
		data []byte
		err  error
	)
	t.Run("Should marshal without error", func(t *testing.T) {
		data, err = PaymentToByteSlice(payment)
		assert.Nil(err)
		assert.NotEmpty(data)
	})
	t.Run("Should unmarshal without error", func(t *testing.T) {
		retPayment, err := PaymentFromByteSlice(data)
		assert.Nil(err)
		assert.NotEmpty(retPayment)
	})
}
