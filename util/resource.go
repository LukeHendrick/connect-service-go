package util

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tjarratt/babble"
)

// FormatARN receives the instanceID, resourceType, and GUID returning an ARN
func FormatARN(instanceID string, resourceType string, id string) string {
	return fmt.Sprintf("arn:aws:connect:us-east-1:000000000000:instance/%v/%v/%v", instanceID, resourceType, id)
}

// GenerateID generates a random UUIDv4
func GenerateID() string {
	return uuid.New().String()
}

func GenerateName() string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return babbler.Babble()
}
