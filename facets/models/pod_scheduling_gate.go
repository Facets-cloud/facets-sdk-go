// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodSchedulingGate PodSchedulingGate
//
// swagger:model PodSchedulingGate
type PodSchedulingGate struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this pod scheduling gate
func (m *PodSchedulingGate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod scheduling gate based on context it is used
func (m *PodSchedulingGate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodSchedulingGate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodSchedulingGate) UnmarshalBinary(b []byte) error {
	var res PodSchedulingGate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
