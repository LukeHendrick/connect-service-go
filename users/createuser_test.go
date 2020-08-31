package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Run("Create User with Good Data", func(t *testing.T) {
		username := "lhendrick"
		password := "abc123"
		routing := "abc"
		securityProfileIds := []string{"abc"}
		phoneConfig := PhoneConfig{}
		phoneType := "SOFT_PHONE"
		phoneConfig.PhoneType = &phoneType
		goodCreateUser := CreateUserPayload{
			Username:           &username,
			Password:           &password,
			RoutingProfileID:   &routing,
			SecurityProfileIds: &securityProfileIds,
			PhoneConfig:        phoneConfig,
		}
		request := newCreateUserRequest(goodCreateUser)
		response := httptest.NewRecorder()
		CreateUser(response, request)
		assertStatus(t, response.Code, 200)
	})
}

func newCreateUserRequest(c CreateUserPayload) *http.Request {
	cByte, _ := json.Marshal(c)
	cReader := bytes.NewReader(cByte)
	req, _ := http.NewRequest(http.MethodPut, "/users/abc", cReader)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
