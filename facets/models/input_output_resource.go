// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InputOutputResource InputOutputResource
//
// swagger:model InputOutputResource
type InputOutputResource struct {

	// output name
	OutputName string `json:"outputName,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}

// Validate validates this input output resource
func (m *InputOutputResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this input output resource based on context it is used
func (m *InputOutputResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InputOutputResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputOutputResource) UnmarshalBinary(b []byte) error {
	var res InputOutputResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
