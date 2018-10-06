package repository

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
	"github.com/mysza/paymentsapi/domain"
	"github.com/stretchr/testify/assert"

	// blank import to initialize the SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func newPayment() *domain.Payment {
	return &domain.Payment{
		OrganisationID: uuid.New().String(),
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
			ProcessingDate:       time.Date(2018, time.October, 5, 12, 00, 00, 00, time.Local),
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

func createDB() (*badger.DB, string) {
	opts := badger.DefaultOptions
	dir, _ := ioutil.TempDir("", "badger")
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db, dir
}

func TestAdd(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := newPayment()
	assert := assert.New(t)
	t.Run("Add retuns ID of the created payment", func(t *testing.T) {
		id, err := repo.Add(payment)
		payment.ID = id
		assert.Nilf(err, "Error adding to repo: %v", err)
		assert.NotEmpty(id, "No ID was created")
	})
}

func TestGet(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := newPayment()
	id, _ := repo.Add(payment)
	payment.ID = id
	assert := assert.New(t)
	t.Run("Get returns previously added payment", func(t *testing.T) {
		retPayment, err := repo.Get(id)
		assert.Nilf(err, "Error getting from repo: %v", err)
		assert.NotEmpty(retPayment, "No payment returned")
		assert.Equalf(payment, retPayment, "Retruned payment differs from added")
	})
}

func TestGetAll(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := newPayment()
	// create 10 instances of payments
	for ix := 0; ix < 10; ix++ {
		repo.Add(payment)
	}
	assert := assert.New(t)
	t.Run("GetAll returns all payments", func(t *testing.T) {
		payments, err := repo.GetAll()
		assert.Nilf(err, "Error getting all from repo: %v", err)
		assert.NotEmpty(payments, "No payments returned")
		assert.Equalf(10, len(payments), "Number of payments is incorrect; expected: %v, actual: %v", 10, len(payments))
	})
}

func TestDelete(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := newPayment()
	id, _ := repo.Add(payment)
	assert := assert.New(t)
	t.Run("Delete deletes the payment", func(t *testing.T) {
		err := repo.Delete(id)
		assert.Nilf(err, "Error deleting payment: %v", err)
		assert.False(repo.Exists(id), "Repo reports payment still exists")
	})
}

func TestUpdate(t *testing.T) {
	db, dir := createDB()
	defer func() {
		// cleanup
		db.Close()
		os.RemoveAll(dir)
	}()
	repo := New(db)
	payment := newPayment()
	var id string
	// create 10 instances of payments
	for ix := 0; ix < 10; ix++ {
		id, _ = repo.Add(payment)
	}
	assert := assert.New(t)
	t.Run("Update updates the payment", func(t *testing.T) {
		payment, _ := repo.Get(id)
		orgID := "ACME Inc."
		payment.OrganisationID = orgID
		err := repo.Update(payment)
		updated, _ := repo.Get(id)
		assert.Nilf(err, "Error updating payment: %v", err)
		assert.Equalf(orgID, updated.OrganisationID, "Payment not updated; expected: %v, got: %v", orgID, updated.OrganisationID)
	})
}
