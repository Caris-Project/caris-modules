package utils

import "github.com/google/uuid"

// GenerateUUID ...
func GenerateUUID() string {
	value, _ := uuid.NewRandom()
	return value.String()
}

// IsValidUUID ...
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// IsEmptyUUID ...
func IsEmptyUUID(id string) bool {
	return id == ""
}
