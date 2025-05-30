// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Toleration Toleration
//
// swagger:model Toleration
type Toleration struct {

	// effect
	Effect string `json:"effect,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// operator
	Operator string `json:"operator,omitempty"`

	// toleration seconds
	TolerationSeconds int64 `json:"tolerationSeconds,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this toleration
func (m *Toleration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this toleration based on context it is used
func (m *Toleration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Toleration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Toleration) UnmarshalBinary(b []byte) error {
	var res Toleration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
