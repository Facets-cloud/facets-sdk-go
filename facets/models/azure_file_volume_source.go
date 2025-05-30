// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureFileVolumeSource AzureFileVolumeSource
//
// swagger:model AzureFileVolumeSource
type AzureFileVolumeSource struct {

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// secret name
	SecretName string `json:"secretName,omitempty"`

	// share name
	ShareName string `json:"shareName,omitempty"`
}

// Validate validates this azure file volume source
func (m *AzureFileVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure file volume source based on context it is used
func (m *AzureFileVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureFileVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureFileVolumeSource) UnmarshalBinary(b []byte) error {
	var res AzureFileVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
