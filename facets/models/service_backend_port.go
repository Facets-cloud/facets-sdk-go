// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceBackendPort ServiceBackendPort
//
// swagger:model ServiceBackendPort
type ServiceBackendPort struct {

	// name
	Name string `json:"name,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`
}

// Validate validates this service backend port
func (m *ServiceBackendPort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service backend port based on context it is used
func (m *ServiceBackendPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBackendPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBackendPort) UnmarshalBinary(b []byte) error {
	var res ServiceBackendPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
