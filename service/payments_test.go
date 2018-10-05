package service

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddReturnsErrorOnInvalidInput(t *testing.T) {
	t.Run("PaymentsService Add returns error if invalid input passed", func(t *testing.T) {
		ps := NewPaymentsService(nil)
		payment := &domain.Payment{
			ID:             uuid.New(),
			OrganisationID: uuid.New(),
		}
		if _, err := ps.Add(payment); err == nil {
			t.Error("Payments service didn't return error despite invalid input")
		}
	})
}

func TestAddReturnsPaymentIDOnValidInput(t *testing.T) {
	t.Run("PaymentsService Add returns added payment ID if the input was valid", func(t *testing.T) {
		payment := &domain.Payment{
			ID:             uuid.New(),
			OrganisationID: uuid.New(),
			Attributes: domain.PaymentAttributes{
				Beneficiary: domain.BeneficiaryPaymentParty{
					PaymentParty: domain.PaymentParty{
						Account: domain.Account{
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
				Debtor: domain.PaymentParty{
					Account: domain.Account{
						AccountNumber: "56781234",
						BankID:        "123123",
						BankIDCode:    "GBDSC",
					},
					AccountName:       "EJ Brown Black",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					Name:              "EJ Brown Black",
				},
				Sponsor: domain.Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				ChargesInformation: domain.ChargesInformation{
					BearerCode:              "SHAR",
					ReceiverChargesAmount:   "100.12",
					ReceiverChargesCurrency: "USD",
					SenderCharges: []domain.Charge{
						domain.Charge{Currency: "USD", Amount: "5.00"},
						domain.Charge{Currency: "GBP", Amount: "15.00"},
					},
				},
				FX: domain.FX{
					ContractReference: "FX123",
					ExchangeRate:      "2.00",
					OriginalAmount:    "100.12",
					OriginalCurrency:  "USD",
				},
				ProcessingDate:       time.Now(),
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
		repo := new(mocks.PaymentsRepository)
		repo.On("Add", payment).Return("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
		ps := NewPaymentsService(repo)
		id, err := ps.Add(payment)
		assert := assert.New(t)
		assert.Equal("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", id, "Returned ID should equal expected value")
		assert.Nil(err)
		repo.AssertExpectations(t)
	})
}
