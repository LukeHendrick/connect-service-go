package util

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// FormatARN receives the instanceID, resourceType, and GUID returning an ARN
func FormatARN(instanceID string, resourceType string, id string) string {
	return fmt.Sprintf("arn:aws:connect:us-east-1:000000000000:instance/%v/%v/%v", instanceID, resourceType, id)
}

// GenerateID generates a random UUIDv4
func GenerateID() string {
	return uuid.New().String()
}

// GenerateName generates a random string for resource names
func GenerateName(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
