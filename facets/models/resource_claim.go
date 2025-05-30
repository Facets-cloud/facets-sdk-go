// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceClaim ResourceClaim
//
// swagger:model ResourceClaim
type ResourceClaim struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this resource claim
func (m *ResourceClaim) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource claim based on context it is used
func (m *ResourceClaim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceClaim) UnmarshalBinary(b []byte) error {
	var res ResourceClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
