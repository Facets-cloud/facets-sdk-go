// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecurityReportSummary SecurityReportSummary
//
// swagger:model SecurityReportSummary
type SecurityReportSummary struct {

	// critical
	Critical int32 `json:"critical,omitempty"`

	// high
	High int32 `json:"high,omitempty"`

	// low
	Low int32 `json:"low,omitempty"`

	// medium
	Medium int32 `json:"medium,omitempty"`

	// unknown
	Unknown int32 `json:"unknown,omitempty"`
}

// Validate validates this security report summary
func (m *SecurityReportSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security report summary based on context it is used
func (m *SecurityReportSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityReportSummary) UnmarshalBinary(b []byte) error {
	var res SecurityReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
