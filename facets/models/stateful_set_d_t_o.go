// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatefulSetDTO StatefulSetDTO
//
// swagger:model StatefulSetDTO
type StatefulSetDTO struct {

	// age in seconds
	AgeInSeconds int64 `json:"ageInSeconds,omitempty"`

	// desired replicas
	DesiredReplicas int32 `json:"desiredReplicas,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pod count
	PodCount int32 `json:"podCount,omitempty"`

	// ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// replicas
	Replicas int32 `json:"replicas,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this stateful set d t o
func (m *StatefulSetDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stateful set d t o based on context it is used
func (m *StatefulSetDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatefulSetDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatefulSetDTO) UnmarshalBinary(b []byte) error {
	var res StatefulSetDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
