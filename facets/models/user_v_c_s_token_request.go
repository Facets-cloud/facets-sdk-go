// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserVCSTokenRequest UserVCSTokenRequest
//
// swagger:model UserVCSTokenRequest
type UserVCSTokenRequest struct {

	// host name
	HostName string `json:"hostName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// personal access token
	PersonalAccessToken string `json:"personalAccessToken,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this user v c s token request
func (m *UserVCSTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user v c s token request based on context it is used
func (m *UserVCSTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserVCSTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserVCSTokenRequest) UnmarshalBinary(b []byte) error {
	var res UserVCSTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
