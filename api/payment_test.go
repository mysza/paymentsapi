package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	validPaymentJSONAdd = `
{
	"version": 0,
	"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	"attributes": {
		"amount": "100.21",
		"beneficiary_party": {
			"account_name": "W Owens",
			"account_number": "31926819",
			"account_number_code": "BBAN",
			"account_type": 0,
			"address": "1 The Beneficiary Localtown SE2",
			"bank_id": "403000",
			"bank_id_code": "GBDSC",
			"name": "Wilfred Jeremiah Owens"
		},
		"charges_information": {
			"bearer_code": "SHAR",
			"sender_charges": [
				{ "amount": "5.00", "currency": "GBP" },
				{ "amount": "10.00", "currency": "USD" }
			],
			"receiver_charges_amount": "1.00",
			"receiver_charges_currency": "USD"
		},
		"currency": "GBP",
		"debtor_party": {
			"account_name": "EJ Brown Black",
			"account_number": "GB29XABC10161234567801",
			"account_number_code": "IBAN",
			"address": "10 Debtor Crescent Sourcetown NE1",
			"bank_id": "203301",
			"bank_id_code": "GBDSC",
			"name": "Emelia Jane Brown"
		},
		"end_to_end_reference": "Wil piano Jan",
		"fx": {
			"contract_reference": "FX123",
			"exchange_rate": "2.00000",
			"original_amount": "200.42",
			"original_currency": "USD"
		},
		"numeric_reference": "1002001",
		"payment_id": "123456789012345678",
		"payment_purpose": "Paying for goods/services",
		"payment_scheme": "FPS",
		"payment_type": "Credit",
		"processing_date": "2017-01-18",
		"reference": "Payment for Em's piano lessons",
		"scheme_payment_sub_type": "InternetBanking",
		"scheme_payment_type": "ImmediatePayment",
		"sponsor_party": {
			"account_number": "56781234",
			"bank_id": "123123",
			"bank_id_code": "GBDSC"
		}
	}
}`

	invalidPaymentJSON = `
{
	"version": 0,
	"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	"attributes": {}
}`

	validPaymentJSONUpdate = `
{
	"version": 0,
	"id": "502758ff-505f-4d81-b9d2-83aa9c01ebe2",
	"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	"attributes": {
		"amount": "100.21",
		"beneficiary_party": {
			"account_name": "W Owens",
			"account_number": "31926819",
			"account_number_code": "BBAN",
			"account_type": 0,
			"address": "1 The Beneficiary Localtown SE2",
			"bank_id": "403000",
			"bank_id_code": "GBDSC",
			"name": "Wilfred Jeremiah Owens"
		},
		"charges_information": {
			"bearer_code": "SHAR",
			"sender_charges": [
				{ "amount": "5.00", "currency": "GBP" },
				{ "amount": "10.00", "currency": "USD" }
			],
			"receiver_charges_amount": "1.00",
			"receiver_charges_currency": "USD"
		},
		"currency": "GBP",
		"debtor_party": {
			"account_name": "EJ Brown Black",
			"account_number": "GB29XABC10161234567801",
			"account_number_code": "IBAN",
			"address": "10 Debtor Crescent Sourcetown NE1",
			"bank_id": "203301",
			"bank_id_code": "GBDSC",
			"name": "Emelia Jane Brown"
		},
		"end_to_end_reference": "Wil piano Jan",
		"fx": {
			"contract_reference": "FX123",
			"exchange_rate": "2.00000",
			"original_amount": "200.42",
			"original_currency": "USD"
		},
		"numeric_reference": "1002001",
		"payment_id": "123456789012345678",
		"payment_purpose": "Paying for goods/services",
		"payment_scheme": "FPS",
		"payment_type": "Credit",
		"processing_date": "2017-01-18",
		"reference": "Payment for Em's piano lessons",
		"scheme_payment_sub_type": "InternetBanking",
		"scheme_payment_type": "ImmediatePayment",
		"sponsor_party": {
			"account_number": "56781234",
			"bank_id": "123123",
			"bank_id_code": "GBDSC"
		}
	}
}`
)

func createTestPayment() *domain.Payment {
	return &domain.Payment{
		ID:             uuid.New().String(),
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

func TestAddHandler(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		body := []byte(validPaymentJSONAdd)
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
		repo := new(mocks.PaymentsRepository)

		repo.On("Add", mock.Anything).Return(id, nil)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.add)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusCreated, recorder.Code)
		assert.Equal(fmt.Sprintf("/payments/%v", id), recorder.Header().Get("Location"))
	})
	t.Run("Invalid input", func(t *testing.T) {
		body := []byte(invalidPaymentJSON)
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		repo := new(mocks.PaymentsRepository)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.add)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusBadRequest, recorder.Code)
	})
}

func TestUpdateHandler(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		body := []byte(validPaymentJSONUpdate)
		req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		repo := new(mocks.PaymentsRepository)
		repo.On("Exists", mock.Anything).Return(true)
		repo.On("Update", mock.Anything).Return(nil)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.update)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusNoContent, recorder.Code)
	})
	t.Run("Invalid input", func(t *testing.T) {
		body := []byte(invalidPaymentJSON)
		req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		repo := new(mocks.PaymentsRepository)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.update)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusBadRequest, recorder.Code)
	})
	t.Run("Non-existing payment", func(t *testing.T) {
		body := []byte(validPaymentJSONUpdate)
		req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		repo := new(mocks.PaymentsRepository)
		repo.On("Exists", mock.Anything).Return(false)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.update)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusNotFound, recorder.Code)
	})
}

func TestGetAllHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	repo := new(mocks.PaymentsRepository)
	payments := []*domain.Payment{
		createTestPayment(),
		createTestPayment(),
		createTestPayment(),
	}
	repo.On("GetAll").Return(payments, nil)

	paymentResource := NewPaymentResource(repo)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(paymentResource.getAll)
	assert := assert.New(t)

	handler.ServeHTTP(recorder, req)

	repo.AssertExpectations(t)
	assert.Equal(http.StatusOK, recorder.Code)
}

func TestGetHandler(t *testing.T) {
	t.Run("Payment exists", func(t *testing.T) {
		id := "502758ff-505f-4d81-b9d2-83aa9c01ebe2"
		req, err := http.NewRequest("GET", fmt.Sprintf("/%v", id), nil)
		if err != nil {
			t.Fatal(err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("paymentID", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		repo := new(mocks.PaymentsRepository)
		payment := createTestPayment()
		payment.ID = id
		repo.On("Get", id).Return(payment, nil)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.get)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.NotNil(recorder.Body)
		assert.Equal(http.StatusOK, recorder.Code)

	})
	t.Run("Payment does not exist", func(t *testing.T) {
		id := "502758ff-505f-4d81-b9d2-83aa9c01ebe2"
		req, err := http.NewRequest("GET", fmt.Sprintf("/%v", id), nil)
		if err != nil {
			t.Fatal(err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("paymentID", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		repo := new(mocks.PaymentsRepository)
		payment := createTestPayment()
		payment.ID = id
		repo.On("Get", id).Return(nil, errors.New("not found"))

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.get)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusNotFound, recorder.Code)

	})
}

func TestDeleteHandler(t *testing.T) {
	t.Run("Payment exists", func(t *testing.T) {
		id := "502758ff-505f-4d81-b9d2-83aa9c01ebe2"
		req, err := http.NewRequest("DELETE", "/502758ff-505f-4d81-b9d2-83aa9c01ebe2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("paymentID", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		repo := new(mocks.PaymentsRepository)
		payment := createTestPayment()
		payment.ID = id
		repo.On("Exists", mock.Anything).Return(true)
		repo.On("Delete", mock.Anything).Return(nil)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.delete)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusNoContent, recorder.Code)
	})
	t.Run("Payment does not exist", func(t *testing.T) {
		id := "502758ff-505f-4d81-b9d2-83aa9c01ebe2"
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/%v", id), nil)
		if err != nil {
			t.Fatal(err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("paymentID", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		repo := new(mocks.PaymentsRepository)
		payment := createTestPayment()
		payment.ID = id
		repo.On("Exists", id).Return(false)

		paymentResource := NewPaymentResource(repo)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(paymentResource.delete)
		assert := assert.New(t)

		handler.ServeHTTP(recorder, req)

		repo.AssertExpectations(t)
		assert.Equal(http.StatusNotFound, recorder.Code)

	})
}
