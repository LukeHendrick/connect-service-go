package util

import (
	"fmt"

	"github.com/google/uuid"
)

// FormatARN receives the instanceID, resourceType, and GUID returning an ARN
func FormatARN(instanceID string, resourceType string, id string) string {
	return fmt.Sprintf("arn:aws:connect:us-east-1:000000000000:instance/%v/agent/%v", instanceID, id)
}

// GenerateID generates a random UUIDv4
func GenerateID() string {
	return uuid.New().String()
}
