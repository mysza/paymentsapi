package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

// Payment is the base data structure provided by the service.
// It describes a single payment registered in the system.
type Payment struct {
	ID             string             `json:"id" validate:"-"`
	OrganisationID string             `json:"organisation_id" validate:"required"`
	Attributes     *PaymentAttributes `json:"attributes" validate:"required"`
	AttributesID   uint               `json:"-" gorm:""`
}

// Validate validates if a given Payment object is valid.
func (p *Payment) Validate(v *validator.Validate) error {
	return v.Struct(p)
}

// BeforeCreate assures the ID is set before the payment is created in DB.
func (p *Payment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
