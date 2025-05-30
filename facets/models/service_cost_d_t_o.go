// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceCostDTO ServiceCostDTO
//
// swagger:model ServiceCostDTO
type ServiceCostDTO struct {

	// cost
	Cost float64 `json:"cost,omitempty"`

	// previous cost
	PreviousCost float64 `json:"previousCost,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this service cost d t o
func (m *ServiceCostDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service cost d t o based on context it is used
func (m *ServiceCostDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCostDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCostDTO) UnmarshalBinary(b []byte) error {
	var res ServiceCostDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
