package utils

import "github.com/google/uuid"

// NewUUID returns a pointer to uuid.UUID
func NewUUID() *uuid.UUID {
	uuid := uuid.New()
	return &uuid
}
