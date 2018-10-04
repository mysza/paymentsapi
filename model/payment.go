package model

import "github.com/google/uuid"

// Payment is the base data structure provided by the service.
// It describes a single payment registered in the system.
type Payment struct {
	Type           string            `json:"type"`
	ID             uuid.UUID         `json:"id"`
	Version        int               `json:"version"`
	OrganisationID uuid.UUID         `json:"organisation_id"`
	Attributes     PaymentAttributes `json:"attributes"`
}
