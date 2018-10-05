package service

import (
	"testing"
	"time"

	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/mysza/paymentsapi/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddReturnsErrorOnInvalidInput(t *testing.T) {
	t.Run("PaymentsService Add returns error if invalid input passed", func(t *testing.T) {
		ps := NewPaymentsService(nil)
		payment := &domain.Payment{
			OrganisationID: utils.NewUUID(),
		}
		if _, err := ps.Add(payment); err == nil {
			t.Error("Payments service didn't return error despite invalid input to Add")
		}
	})
}

func TestAddReturnsPaymentIDOnValidInput(t *testing.T) {
	payment := &domain.Payment{
		OrganisationID: utils.NewUUID(),
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
	assert := assert.New(t)
	t.Run("PaymentsService Add returns added payment ID if the input was valid", func(t *testing.T) {
		id, err := ps.Add(payment)
		assert.Equal("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", id, "Returned ID should equal expected value")
		assert.Nil(err)
		repo.AssertExpectations(t)
	})
}

func TestAddReturnsErrorIfIDSet(t *testing.T) {
	payment := &domain.Payment{
		ID:             utils.NewUUID(),
		OrganisationID: utils.NewUUID(),
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
	ps := NewPaymentsService(repo)
	assert := assert.New(t)
	t.Run("PaymentsService Add returns error if payment ID was set", func(t *testing.T) {
		id, err := ps.Add(payment)
		assert.Empty(id, "Returned ID should be empty")
		assert.Error(err, "Error should be set")
		repo.AssertExpectations(t)
	})
}

func TestGetAllReturnsAllPaymentsFromRepo(t *testing.T) {
	payments := []*domain.Payment{
		&domain.Payment{
			ID:             utils.NewUUID(),
			OrganisationID: utils.NewUUID(),
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
		},
		&domain.Payment{
			ID:             utils.NewUUID(),
			OrganisationID: utils.NewUUID(),
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
		},
	}
	repo := new(mocks.PaymentsRepository)
	repo.On("GetAll").Return(payments, nil)
	ps := NewPaymentsService(repo)
	assert := assert.New(t)
	t.Run("PaymentsService GetAll should return all payments from repository", func(t *testing.T) {
		retPayments, err := ps.GetAll()
		assert.Equal(payments, retPayments, "Returned payments should be all in repository")
		assert.Nil(err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateReturnsErrorOnInvalidInput(t *testing.T) {
	t.Run("PaymentsService Update returns error if invalid input passed", func(t *testing.T) {
		ps := NewPaymentsService(nil)
		payment := &domain.Payment{
			ID:             utils.NewUUID(),
			OrganisationID: utils.NewUUID(),
		}
		if err := ps.Update(payment); err == nil {
			t.Error("Payments service didn't return error despite invalid input to Update")
		}
	})
}

func TestUpdateReturnsErrorIfPaymentWithGiveIDDoesNotExist(t *testing.T) {
	paymentID := utils.NewUUID()
	payment := &domain.Payment{
		ID:             paymentID,
		OrganisationID: utils.NewUUID(),
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
	repo.On("Exists", paymentID).Return(false)
	ps := NewPaymentsService(repo)
	assert := assert.New(t)
	t.Run("PaymentsService Update returns error if invalid input passed", func(t *testing.T) {
		err := ps.Update(payment)
		assert.Error(err, "Update with non-existing payment shoud return error")
		repo.AssertExpectations(t)
	})
}

func TestUpdateReturnsNilOnValidInput(t *testing.T) {
	payment := &domain.Payment{
		ID:             utils.NewUUID(),
		OrganisationID: utils.NewUUID(),
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
	repo.On("Exists", payment.ID).Return(true)
	repo.On("Update", payment).Return(nil)
	ps := NewPaymentsService(repo)
	assert := assert.New(t)
	t.Run("PaymentsService Update returns nil if the input was valid", func(t *testing.T) {
		err := ps.Update(payment)
		assert.Nil(err)
		repo.AssertExpectations(t)
	})
}

func TestGetReturnsPaymentIfExistingIDPassed(t *testing.T) {
	id := utils.NewUUID()
	payment := &domain.Payment{
		ID:             id,
		OrganisationID: utils.NewUUID(),
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
	repo.On("Get", id).Return(payment, nil)
	ps := NewPaymentsService(repo)
	assert := assert.New(t)
	t.Run("PaymentsService Get returns payment if the input was its ID", func(t *testing.T) {
		retPayment, err := ps.Get(id)
		assert.Nil(err)
		assert.Equal(payment, retPayment, "Retrieved payment differs from expected payment after Get")
		repo.AssertExpectations(t)
	})
}

func TestGetReturnsErrorIfInvalidIDPassed(t *testing.T) {
	ps := NewPaymentsService(nil)
	assert := assert.New(t)
	t.Run("PaymentsService Get returns error if invalid ID passed", func(t *testing.T) {
		retPayment, err := ps.Get(nil)
		assert.Error(err)
		assert.Empty(retPayment)
	})
}
