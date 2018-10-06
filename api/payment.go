package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mysza/paymentsapi/domain"
	"github.com/mysza/paymentsapi/service"
)

type ctxKey int

const (
	ctxPaymentID ctxKey = iota
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
	r.Post("/", rs.add)
	r.Put("/", rs.update)
	r.Route("/{paymentID}", func(r chi.Router) {
		r.Use(rs.paymentCtx)
		r.Get("/", rs.get)
		r.Delete("/", rs.delete)
	})
	return r
}

func (rs *PaymentResource) getAll(w http.ResponseWriter, r *http.Request) {
	payments, err := rs.service.GetAll()
	if err != nil {
		render.Render(w, r, &ErrResponse{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			StatusText: http.StatusText(http.StatusInternalServerError),
		})
	}
	render.Respond(w, r, newPaymentListResponse(payments))
}

func (rs *PaymentResource) add(w http.ResponseWriter, r *http.Request) {
	input := &paymentRequest{}
	if err := render.Bind(r, input); err != nil {
		render.Render(w, r, ErrBadRequest)
	}
	id, err := rs.service.Add(input.Payment)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
	}
	w.Header().Set("Location", fmt.Sprintf("/payments/%v", id))
	render.NoContent(w, r)
}

func (rs *PaymentResource) update(w http.ResponseWriter, r *http.Request) {
	input := &paymentRequest{}
	if err := render.Bind(r, input); err != nil {
		render.Render(w, r, ErrBadRequest)
	}
	err := rs.service.Update(input.Payment)
	if err != nil {
		switch err.(type) {
		case *service.InputError:
			render.Render(w, r, ErrBadRequest)
		case *service.NotFoundError:
			render.Render(w, r, ErrNotFound)
		default:
			render.Render(w, r, ErrInternalServerError)
		}
	}
	render.NoContent(w, r)
}

func (rs *PaymentResource) paymentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "paymentID"))
		if err != nil {
			render.Render(w, r, ErrBadRequest)
		}
		ctx := context.WithValue(r.Context(), ctxPaymentID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs *PaymentResource) get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxPaymentID).(string)
	payment, err := rs.service.Get(id)
	if err != nil {
		switch err.(type) {
		case *service.InputError:
			render.Render(w, r, ErrBadRequest)
		case *service.NotFoundError:
			render.Render(w, r, ErrNotFound)
		default:
			render.Render(w, r, ErrInternalServerError)
		}
	}
	render.Respond(w, r, newPaymentResponse(payment))
}

func (rs *PaymentResource) delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxPaymentID).(string)
	err := rs.service.Delete(id)
	if err != nil {
		switch err.(type) {
		case *service.InputError:
			render.Render(w, r, ErrBadRequest)
		case *service.NotFoundError:
			render.Render(w, r, ErrNotFound)
		default:
			render.Render(w, r, ErrInternalServerError)
		}
	}
	render.NoContent(w, r)
}

type paymentRequest struct {
	*domain.Payment
}

func (c *paymentRequest) Bind(r *http.Request) error {
	return nil
}

type paymentResponse struct {
	*domain.Payment
	Type    string `json:"type"`
	Version int    `json:"version"`
}

func newPaymentResponse(payment *domain.Payment) *paymentResponse {
	return &paymentResponse{Payment: payment, Type: "Payment", Version: 0}
}

// PaymentListResponse is the response payload for a list of Payments.
type paymentListResponse struct {
	Data []*paymentResponse `json:"data"`
}

func newPaymentListResponse(payments []*domain.Payment) *paymentListResponse {
	list := []*paymentResponse{}
	for _, payment := range payments {
		list = append(list, newPaymentResponse(payment))
	}
	return &paymentListResponse{Data: list}
}
