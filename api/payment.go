package api

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service"

	"github.com/mysza/paymentsapi/utils"
)

// PaymentResource implements payments management handler
type PaymentResource struct {
	service *service.PaymentsService
}

// NewPaymentResource creates and returns a payments resource.
func NewPaymentResource(repo service.PaymentsRepository) *PaymentResource {
	service := service.NewPaymentsService(repo)
	return &PaymentResource{service}
}

func (rs *PaymentResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", rs.getAll)
	return r
}

func (rs *PaymentResource) getAll(w http.ResponseWriter, r *http.Request) {
	payments := []*domain.Payment{
		&domain.Payment{ID: utils.NewUUID()},
		&domain.Payment{ID: utils.NewUUID()},
		&domain.Payment{ID: utils.NewUUID()},
		&domain.Payment{ID: utils.NewUUID()},
		&domain.Payment{ID: utils.NewUUID()},
	}
	response := newPaymentListResponse(payments)
	render.Render(w, r, response)
}

// PaymentResponse is the response payload for Payment data model.
type PaymentResponse struct {
	ID string `json:"id"`
}

// Render implements the chi Renderer interface
func (pr *PaymentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newPaymentResponse(id uuid.UUID) *PaymentResponse {
	return &PaymentResponse{ID: id.String()}
}

// PaymentListResponse is the response payload for a list of Payments.
type PaymentListResponse struct {
	Data []*PaymentResponse `json:"data"`
}

// Render implements the chi Renderer interface
func (pl *PaymentListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newPaymentListResponse(payments []*domain.Payment) *PaymentListResponse {
	list := []*PaymentResponse{}
	for _, payment := range payments {
		list = append(list, &PaymentResponse{ID: payment.ID.String()})
	}
	return &PaymentListResponse{Data: list}
}
