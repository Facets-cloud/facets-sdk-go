// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangedAttribute ChangedAttribute
//
// swagger:model ChangedAttribute
type ChangedAttribute struct {

	// attribute
	Attribute string `json:"attribute,omitempty"`

	// new value
	NewValue string `json:"newValue,omitempty"`

	// old value
	OldValue string `json:"oldValue,omitempty"`
}

// Validate validates this changed attribute
func (m *ChangedAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this changed attribute based on context it is used
func (m *ChangedAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangedAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangedAttribute) UnmarshalBinary(b []byte) error {
	var res ChangedAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
