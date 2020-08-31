package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukehendrick/connectService/util"
)

// CreateUserPayload is the Body received from the request
type CreateUserPayload struct {
	DirectoryUserID    *string      `json:"DirectoryUserId,omitempty"`
	HierarchyGroupID   *string      `json:"HierarchyGroupId,omitempty"`
	IdentityInfo       IdentityInfo `json:"IdentityInfo,omitempty"`
	Password           *string      `json:"Password,omitempty"`
	PhoneConfig        PhoneConfig  `json:"PhoneConfig"`
	RoutingProfileID   *string      `json:"RoutingProfileId"`
	SecurityProfileIds *[]string    `json:"SecurityProfileIds"`
	Tags               Tags         `json:"Tags,omitempty"`
	Username           *string      `json:"Username"`
}

// PhoneConfig is the users...Phone configuration
type PhoneConfig struct {
	AfterContactWorkTimeLimit *int    `json:"AfterContactWorkTimeLimit"`
	AutoAccept                *bool   `json:"AutoAccept"`
	DeskPhoneNumber           *string `json:"DeskPhoneNumber,omitempty"`
	PhoneType                 *string `json:"PhoneType"`
}

// IdentityInfo is the users...identity information
type IdentityInfo struct {
	Email     *string `json:"Email"`
	FirstName *string `json:"FirstName"`
	LastName  *string `json:"LastName"`
}

// Tags are tags to be assigned to the user
type Tags struct {
	String *string `json:"string"`
}

// CreateUserResponse contains the generated ARN and ID for the user
type CreateUserResponse struct {
	UserArn string `json:"UserArn"`
	UserID  string `json:"UserId"`
}

// CreateUser receives and verifies a payload, then returns some user info
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userPayload := CreateUserPayload{}
	instanceID := chi.URLParam(r, "instanceID")
	fmt.Printf("InstanceID %v", instanceID)
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		util.ReturnError(w, r, err, http.StatusBadRequest)
		return
	}
	err = validateCreateUserPayload(userPayload)
	if err != nil {
		util.ReturnError(w, r, err, http.StatusBadRequest)
		return
	}
	id := util.GenerateID()
	arn := util.FormatARN(instanceID, "agent", id)
	userResponse := CreateUserResponse{
		UserArn: arn,
		UserID:  id,
	}
	util.MarshalAndWriteResponse(w, r, userResponse)
}

func validateCreateUserPayload(payload CreateUserPayload) error {
	username := payload.Username
	routingProfile := payload.RoutingProfileID
	securityProfile := payload.SecurityProfileIds
	phoneType := payload.PhoneConfig.PhoneType
	if username == nil {
		return fmt.Errorf("Missing Username")
	}
	if routingProfile == nil {
		return fmt.Errorf("Missing Routing Profile")
	}
	if securityProfile == nil {
		return fmt.Errorf("Missing Security Profile")
	}
	if phoneType == nil {
		return fmt.Errorf("Missing Phone Type from Phone Config")
	}
	return nil
}
