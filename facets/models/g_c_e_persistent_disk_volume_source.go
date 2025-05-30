// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCEPersistentDiskVolumeSource GCEPersistentDiskVolumeSource
//
// swagger:model GCEPersistentDiskVolumeSource
type GCEPersistentDiskVolumeSource struct {

	// fs type
	FsType string `json:"fsType,omitempty"`

	// partition
	Partition int32 `json:"partition,omitempty"`

	// pd name
	PdName string `json:"pdName,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Validate validates this g c e persistent disk volume source
func (m *GCEPersistentDiskVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g c e persistent disk volume source based on context it is used
func (m *GCEPersistentDiskVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCEPersistentDiskVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCEPersistentDiskVolumeSource) UnmarshalBinary(b []byte) error {
	var res GCEPersistentDiskVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
