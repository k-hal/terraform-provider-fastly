package fastly

import (
	"fmt"
	"time"

	"github.com/google/jsonapi"
)

type SAUser struct {
	ID string `jsonapi:"primary,user"`
}

type SAService struct {
	ID string `jsonapi:"primary,service"`
}

type ServiceAuthorization struct {
	ID         string     `jsonapi:"primary,service_authorization"`
	Permission string     `jsonapi:"attr,permission,omitempty"`
	CreatedAt  *time.Time `jsonapi:"attr,created_at,iso8601"`
	UpdatedAt  *time.Time `jsonapi:"attr,updated_at,iso8601"`
	DeltedAt   *time.Time `jsonapi:"attr,deleted_at,iso8601"`
	User       *SAUser    `jsonapi:"relation,user,omitempty"`
	Service    *SAService `jsonapi:"relation,service,omitempty"`
}

// GetServiceAuthorizationInput is used as input to the GetServiceAuthorization function.
type GetServiceAuthorizationInput struct {
	// ID of the service authorization to retrieve.
	ID string
}

// GetServiceAuthorization retrieves an existing service authorization using its ID.
func (c *Client) GetServiceAuthorization(i *GetServiceAuthorizationInput) (*ServiceAuthorization, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service-authorizations/%s", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var sa ServiceAuthorization
	if err := jsonapi.UnmarshalPayload(resp.Body, &sa); err != nil {
		return nil, err
	}

	return &sa, nil
}

// CreateServiceAuthorizationInput is used as input to the CreateServiceAuthorization function.
type CreateServiceAuthorizationInput struct {
	// ID value is ignored and should not be set, needed to make JSONAPI work correctly.
	ID string `jsonapi:"primary,service_authorization"`

	// Permission is the level of permissions to grant the user to the service. Valid values are "full", "read_only", "purge_select" or "purge_all".
	Permission string `jsonapi:"attr,permission,omitempty"`

	// ServiceID is the ID of the service to grant permissions for.
	Service *SAService `jsonapi:"relation,service,omitempty"`

	// UserID is the ID of the user which should have its permissions set.
	User *SAUser `jsonapi:"relation,user,omitempty"`
}

// CreateServiceAuthorization creates a new service authorization granting granular service and user permissions.
func (c *Client) CreateServiceAuthorization(i *CreateServiceAuthorizationInput) (*ServiceAuthorization, error) {
	if i.Service == nil || i.Service.ID == "" {
		return nil, ErrMissingServiceAuthorizationsService
	}
	if i.User == nil || i.User.ID == "" {
		return nil, ErrMissingServiceAuthorizationsUser
	}

	resp, err := c.PostJSONAPI("/service-authorizations", i, nil)
	if err != nil {
		return nil, err
	}

	var sa ServiceAuthorization
	if err := jsonapi.UnmarshalPayload(resp.Body, &sa); err != nil {
		return nil, err
	}

	return &sa, nil
}

// UpdateServiceAuthorizationInput is used as input to the UpdateServiceAuthorization function.
type UpdateServiceAuthorizationInput struct {
	// ID uniquely identifies the service authorization (service and user pair) to be updated.
	ID string `jsonapi:"primary,service_authorization"`

	// The permission to grant the user to the service referenced by this service authorization.
	Permissions string `jsonapi:"attr,permission,omitempty"`
}

// UpdateServiceAuthorization updates an exisitng service authorization. The ID must be known.
func (c *Client) UpdateServiceAuthorization(i *UpdateServiceAuthorizationInput) (*ServiceAuthorization, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	if i.Permissions == "" {
		return nil, ErrMissingPermissions
	}

	path := fmt.Sprintf("/service-authorizations/%s", i.ID)
	resp, err := c.PatchJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var sa ServiceAuthorization
	if err := jsonapi.UnmarshalPayload(resp.Body, &sa); err != nil {
		return nil, err
	}

	return &sa, nil
}

// DeleteServiceAuthorizationInput is used as input to the DeleteServiceAuthorization function.
type DeleteServiceAuthorizationInput struct {
	// ID of the service authorization to delete.
	ID string
}

// DeleteServiceAuthorization deletes an existing service authorization using the ID.
func (c *Client) DeleteServiceAuthorization(i *DeleteServiceAuthorizationInput) error {
	if i.ID == "" {
		return ErrMissingID
	}

	path := fmt.Sprintf("/service-authorizations/%s", i.ID)
	_, err := c.Delete(path, nil)

	return err
}