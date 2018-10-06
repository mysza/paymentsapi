package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/mysza/paymentsapi/domain"

	// blank import to initialize the SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestAddReturnsNewID(t *testing.T) {
	payment := &domain.Payment{
		OrganisationID: uuid.New().String(),
		Attributes: &domain.PaymentAttributes{
			Beneficiary: &domain.BeneficiaryPaymentParty{
				PaymentParty: &domain.PaymentParty{
					Account: &domain.Account{
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
			Debtor: &domain.PaymentParty{
				Account: &domain.Account{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
				AccountName:       "EJ Brown Black",
				AccountNumberCode: "IBAN",
				Address:           "10 Debtor Crescent Sourcetown NE1",
				Name:              "EJ Brown Black",
			},
			Sponsor: &domain.Account{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
			ChargesInformation: &domain.ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   "100.12",
				ReceiverChargesCurrency: "USD",
				SenderCharges: []*domain.Charge{
					&domain.Charge{Currency: "USD", Amount: "5.00"},
					&domain.Charge{Currency: "GBP", Amount: "15.00"},
				},
			},
			FX: &domain.FX{
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
	db, _ := gorm.Open("sqlite3", ":memory:")
	repo := New(db)
	id, err := repo.Add(payment)
	if err != nil {
		t.Errorf("Error adding to repo: %v", err)
	}
	if id == "" {
		t.Errorf("No ID was created")
	}
}
