// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodOS PodOS
//
// swagger:model PodOS
type PodOS struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this pod o s
func (m *PodOS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod o s based on context it is used
func (m *PodOS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodOS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodOS) UnmarshalBinary(b []byte) error {
	var res PodOS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
