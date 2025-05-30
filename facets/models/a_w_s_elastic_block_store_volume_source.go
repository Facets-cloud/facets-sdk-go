// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AWSElasticBlockStoreVolumeSource AWSElasticBlockStoreVolumeSource
//
// swagger:model AWSElasticBlockStoreVolumeSource
type AWSElasticBlockStoreVolumeSource struct {

	// fs type
	FsType string `json:"fsType,omitempty"`

	// partition
	Partition int32 `json:"partition,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// volume ID
	VolumeID string `json:"volumeID,omitempty"`
}

// Validate validates this a w s elastic block store volume source
func (m *AWSElasticBlockStoreVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a w s elastic block store volume source based on context it is used
func (m *AWSElasticBlockStoreVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AWSElasticBlockStoreVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSElasticBlockStoreVolumeSource) UnmarshalBinary(b []byte) error {
	var res AWSElasticBlockStoreVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
