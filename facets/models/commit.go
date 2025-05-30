// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Commit Commit
//
// swagger:model Commit
type Commit struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this commit
func (m *Commit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this commit based on context it is used
func (m *Commit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Commit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commit) UnmarshalBinary(b []byte) error {
	var res Commit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
