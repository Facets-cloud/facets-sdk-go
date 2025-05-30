// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodCondition PodCondition
//
// swagger:model PodCondition
type PodCondition struct {

	// last probe time
	LastProbeTime string `json:"lastProbeTime,omitempty"`

	// last transition time
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this pod condition
func (m *PodCondition) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod condition based on context it is used
func (m *PodCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodCondition) UnmarshalBinary(b []byte) error {
	var res PodCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
