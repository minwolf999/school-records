package utility

import "github.com/gofrs/uuid"

// These function return a new uuid withot the error
func NewUUID() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}