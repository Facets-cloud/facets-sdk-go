// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OutputReference OutputReference
//
// swagger:model OutputReference
type OutputReference struct {

	// output name
	OutputName string `json:"outputName,omitempty"`

	// output title
	OutputTitle string `json:"outputTitle,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}

// Validate validates this output reference
func (m *OutputReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this output reference based on context it is used
func (m *OutputReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutputReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutputReference) UnmarshalBinary(b []byte) error {
	var res OutputReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
