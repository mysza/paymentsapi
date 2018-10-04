package domain

import "github.com/google/uuid"

// Payment is the base data structure provided by the service.
// It describes a single payment registered in the system.
type Payment struct {
	// Type           string            `json:"type" validate:"required,"`
	// Version        int               `json:"version"`
	ID             uuid.UUID         `json:"id" validate:"required"`
	OrganisationID uuid.UUID         `json:"organisation_id" validate:"required"`
	Attributes     PaymentAttributes `json:"attributes" validate:"required"`
}
