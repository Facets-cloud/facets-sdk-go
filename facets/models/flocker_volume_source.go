// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlockerVolumeSource FlockerVolumeSource
//
// swagger:model FlockerVolumeSource
type FlockerVolumeSource struct {

	// dataset name
	DatasetName string `json:"datasetName,omitempty"`

	// dataset UUID
	DatasetUUID string `json:"datasetUUID,omitempty"`
}

// Validate validates this flocker volume source
func (m *FlockerVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flocker volume source based on context it is used
func (m *FlockerVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlockerVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlockerVolumeSource) UnmarshalBinary(b []byte) error {
	var res FlockerVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
