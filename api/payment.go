package api

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

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
	r.Get("/{paymentID}", rs.get)
	r.Delete("/{paymentID}", rs.delete)
	return r
}

func (rs *PaymentResource) getAll(w http.ResponseWriter, r *http.Request) {
	payments, err := rs.service.GetAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/getAll",
			"details":  "service.GetAll",
			"error":    err,
		}).Warn("Error getting payments from repository")
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
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/add",
			"details":  "render.Bind",
			"error":    err,
		}).Warn("Error binding to the input")
		render.Render(w, r, ErrBadRequest)
		return
	}
	id, err := rs.service.Add(input.Payment)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/add",
			"details":  "service.Add",
			"error":    err,
		}).Warn("Error adding by service")
		render.Render(w, r, ErrBadRequest)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/payments/%v", id))
	logrus.WithField("location", "api/payment/add").Infof("Added new Payment with ID: %s", id)
	w.WriteHeader(http.StatusCreated)
}

func (rs *PaymentResource) update(w http.ResponseWriter, r *http.Request) {
	input := &paymentRequest{}
	if err := render.Bind(r, input); err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/update",
			"details":  "render.Bind",
			"error":    err,
		}).Warn("Error binding to the input")
		render.Render(w, r, ErrBadRequest)
	}
	err := rs.service.Update(input.Payment)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/update",
			"details":  "service.Update",
			"error":    err,
		}).Warn("Error updating by service")
		switch err.(type) {
		case *service.InputError:
			render.Render(w, r, ErrBadRequest)
		case *service.NotFoundError:
			render.Render(w, r, ErrNotFound)
		default:
			render.Render(w, r, ErrInternalServerError)
		}
	}
	logrus.WithField("location", "api/payment/update").Infof("Updated Payment with ID: %s", input.ID)
	render.NoContent(w, r)
}

func (rs *PaymentResource) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "paymentID")
	payment, err := rs.service.Get(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/get",
			"details":  "service.Get",
			"id":       id,
			"error":    err,
		}).Warn("Error getting by service")
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
	id := chi.URLParam(r, "paymentID")
	err := rs.service.Delete(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/payment/delete",
			"details":  "service.Delete",
			"id":       id,
			"error":    err,
		}).Warn("Error deleting by service")
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
