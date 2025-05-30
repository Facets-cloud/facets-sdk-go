// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DailyCost DailyCost
//
// swagger:model DailyCost
type DailyCost struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// per service cost
	PerServiceCost []*ServiceCost `json:"perServiceCost"`

	// unit
	Unit string `json:"unit,omitempty"`
}

// Validate validates this daily cost
func (m *DailyCost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerServiceCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DailyCost) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DailyCost) validatePerServiceCost(formats strfmt.Registry) error {
	if swag.IsZero(m.PerServiceCost) { // not required
		return nil
	}

	for i := 0; i < len(m.PerServiceCost); i++ {
		if swag.IsZero(m.PerServiceCost[i]) { // not required
			continue
		}

		if m.PerServiceCost[i] != nil {
			if err := m.PerServiceCost[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("perServiceCost" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("perServiceCost" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this daily cost based on the context it is used
func (m *DailyCost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerServiceCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DailyCost) contextValidatePerServiceCost(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerServiceCost); i++ {

		if m.PerServiceCost[i] != nil {

			if swag.IsZero(m.PerServiceCost[i]) { // not required
				return nil
			}

			if err := m.PerServiceCost[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("perServiceCost" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("perServiceCost" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DailyCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DailyCost) UnmarshalBinary(b []byte) error {
	var res DailyCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
