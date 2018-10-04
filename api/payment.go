package api

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mysza/paymentsapi/model"
)

// PaymentRepository defines database (or other storage medium) operations
// on a payment
type PaymentRepository interface {
	GetAll() (*[]model.Payment, error)
	Get(id string) (*model.Payment, error)
	Create( /* define input */ ) (string, error) // returns id
	Update(payment *model.Payment) error
	Delete(id string) error
}

// PaymentResource implements payments management handler
type PaymentResource struct {
	Store PaymentRepository
}

// NewPaymentResource creates and returns a payments resource.
func NewPaymentResource(repo PaymentRepository) *PaymentResource {
	return &PaymentResource{Store: repo}
}

func (rs *PaymentResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", rs.getAll)
	return r
}

func (rs *PaymentResource) getAll(w http.ResponseWriter, r *http.Request) {
	payments := []*model.Payment{
		&model.Payment{ID: uuid.New()},
		&model.Payment{ID: uuid.New()},
		&model.Payment{ID: uuid.New()},
		&model.Payment{ID: uuid.New()},
		&model.Payment{ID: uuid.New()},
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

func newPaymentListResponse(payments []*model.Payment) *PaymentListResponse {
	list := []*PaymentResponse{}
	for _, payment := range payments {
		list = append(list, &PaymentResponse{ID: payment.ID.String()})
	}
	return &PaymentListResponse{Data: list}
}
