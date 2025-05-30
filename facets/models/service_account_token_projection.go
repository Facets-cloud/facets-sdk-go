// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceAccountTokenProjection ServiceAccountTokenProjection
//
// swagger:model ServiceAccountTokenProjection
type ServiceAccountTokenProjection struct {

	// audience
	Audience string `json:"audience,omitempty"`

	// expiration seconds
	ExpirationSeconds int64 `json:"expirationSeconds,omitempty"`

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this service account token projection
func (m *ServiceAccountTokenProjection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service account token projection based on context it is used
func (m *ServiceAccountTokenProjection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAccountTokenProjection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAccountTokenProjection) UnmarshalBinary(b []byte) error {
	var res ServiceAccountTokenProjection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
