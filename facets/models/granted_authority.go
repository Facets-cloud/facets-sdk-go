// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GrantedAuthority GrantedAuthority
//
// swagger:model GrantedAuthority
type GrantedAuthority struct {

	// authority
	Authority string `json:"authority,omitempty"`
}

// Validate validates this granted authority
func (m *GrantedAuthority) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this granted authority based on context it is used
func (m *GrantedAuthority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrantedAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrantedAuthority) UnmarshalBinary(b []byte) error {
	var res GrantedAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
