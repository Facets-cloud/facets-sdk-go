// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricTarget MetricTarget
//
// swagger:model MetricTarget
type MetricTarget struct {

	// average utilization
	AverageUtilization int32 `json:"averageUtilization,omitempty"`

	// average value
	AverageValue *Quantity `json:"averageValue,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value *Quantity `json:"value,omitempty"`
}

// Validate validates this metric target
func (m *MetricTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricTarget) validateAverageValue(formats strfmt.Registry) error {
	if swag.IsZero(m.AverageValue) { // not required
		return nil
	}

	if m.AverageValue != nil {
		if err := m.AverageValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("averageValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("averageValue")
			}
			return err
		}
	}

	return nil
}

func (m *MetricTarget) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metric target based on the context it is used
func (m *MetricTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAverageValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricTarget) contextValidateAverageValue(ctx context.Context, formats strfmt.Registry) error {

	if m.AverageValue != nil {

		if swag.IsZero(m.AverageValue) { // not required
			return nil
		}

		if err := m.AverageValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("averageValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("averageValue")
			}
			return err
		}
	}

	return nil
}

func (m *MetricTarget) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricTarget) UnmarshalBinary(b []byte) error {
	var res MetricTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
