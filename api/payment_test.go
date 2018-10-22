package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/mock"

	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service"
	"github.com/mysza/paymentsapi/service/mocks"
	"github.com/mysza/paymentsapi/test"
	"github.com/stretchr/testify/assert"
)

type httpRequestContext struct {
	name  string
	value string
}

type header struct {
	name  string
	value string
}

type apiTest struct {
	name           string
	request        *http.Request
	handler        http.HandlerFunc
	expectedCode   int
	expectedHeader *header
}

func createHTTPRequest(method, path string, input *domain.Payment, ctx *httpRequestContext) *http.Request {
	var body io.Reader
	if input != nil {
		bytesInput, _ := domain.PaymentToByteSlice(input)
		body = bytes.NewBuffer(bytesInput)
	} else {
		body = nil
	}
	req, _ := http.NewRequest(method, path, body)
	if ctx != nil {
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add(ctx.name, ctx.value)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	}
	req.Header.Add("content-type", "application/json")
	return req
}

func prepareRepository(newID string, existing, notExisting *domain.Payment) service.PaymentsRepository {
	repo := new(mocks.PaymentsRepository)
	repo.On("Add", mock.Anything).Return(newID, nil)
	repo.On("Exists", existing.ID).Return(true)
	repo.On("Exists", notExisting.ID).Return(false)
	repo.On("Update", existing).Return(nil)
	repo.On("GetAll").Return([]*domain.Payment{existing, existing, existing}, nil)
	repo.On("Get", existing.ID).Return(existing, nil)
	repo.On("Get", notExisting.ID).Return(nil, errors.New("not found"))
	repo.On("Delete", existing.ID).Return(nil)
	return repo
}

func TestAPI(t *testing.T) {
	testDataDir := filepath.Join("..", "testdata")

	validPayment := test.PaymentFromFile(t, filepath.Join(testDataDir, "validPayment.json"))
	invalidPayment := test.PaymentFromFile(t, filepath.Join(testDataDir, "invalidPayment.json"))

	var validPaymentNoID = &domain.Payment{}
	copier.Copy(&validPaymentNoID, validPayment)
	validPaymentNoID.ID = ""

	var validPaymentNotExisting = &domain.Payment{}
	copier.Copy(&validPaymentNotExisting, validPayment)
	validPaymentNotExisting.ID = "30c85da3-244f-4fc4-86bb-312ce8ffa52a"

	newID := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"

	repo := prepareRepository(newID, validPayment, validPaymentNotExisting)

	paymentResource := NewPaymentResource(repo)
	addHandler := http.HandlerFunc(paymentResource.add)
	updateHandler := http.HandlerFunc(paymentResource.update)
	getAllHandler := http.HandlerFunc(paymentResource.getAll)
	getHandler := http.HandlerFunc(paymentResource.get)
	deleteHandler := http.HandlerFunc(paymentResource.delete)

	cases := []apiTest{
		{
			name:           "POST: valid input",
			request:        createHTTPRequest("POST", "/", validPaymentNoID, nil),
			handler:        addHandler,
			expectedCode:   http.StatusCreated,
			expectedHeader: &header{"Location", fmt.Sprintf("/payments/%v", newID)},
		},
		{
			name:         "POST: invalid input",
			request:      createHTTPRequest("POST", "/", invalidPayment, nil),
			handler:      addHandler,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "PUT: valid input",
			request:      createHTTPRequest("PUT", "/", validPayment, nil),
			handler:      updateHandler,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "PUT: invalid input",
			request:      createHTTPRequest("PUT", "/", invalidPayment, nil),
			handler:      updateHandler,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "PUT: non-existing payment",
			request:      createHTTPRequest("PUT", "/", validPaymentNotExisting, nil),
			handler:      updateHandler,
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "GET: get all",
			request:      createHTTPRequest("GET", "/", nil, nil),
			handler:      getAllHandler,
			expectedCode: http.StatusOK,
		},
		{
			name:         "GET: existing payment",
			request:      createHTTPRequest("GET", fmt.Sprintf("/%v", validPayment.ID), nil, &httpRequestContext{"paymentID", validPayment.ID}),
			handler:      getHandler,
			expectedCode: http.StatusOK,
		},
		{
			name:         "GET: non-existing payment",
			request:      createHTTPRequest("GET", fmt.Sprintf("/%v", validPaymentNotExisting.ID), nil, &httpRequestContext{"paymentID", validPaymentNotExisting.ID}),
			handler:      getHandler,
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "DELETE: existing payment",
			request:      createHTTPRequest("DELETE", fmt.Sprintf("/%v", validPayment.ID), nil, &httpRequestContext{"paymentID", validPayment.ID}),
			handler:      deleteHandler,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "DELETE: non-existing payment",
			request:      createHTTPRequest("GET", fmt.Sprintf("/%v", validPaymentNotExisting.ID), nil, &httpRequestContext{"paymentID", validPaymentNotExisting.ID}),
			handler:      deleteHandler,
			expectedCode: http.StatusNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			assert := assert.New(t)
			recorder := httptest.NewRecorder()

			testCase.handler.ServeHTTP(recorder, testCase.request)

			assert.Equal(testCase.expectedCode, recorder.Code)
			if testCase.expectedHeader != nil {
				assert.Equal(testCase.expectedHeader.value, recorder.Header().Get(testCase.expectedHeader.name))
			}
		})
	}
}
