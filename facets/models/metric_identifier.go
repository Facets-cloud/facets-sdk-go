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

// MetricIdentifier MetricIdentifier
//
// swagger:model MetricIdentifier
type MetricIdentifier struct {

	// name
	Name string `json:"name,omitempty"`

	// selector
	Selector *LabelSelector `json:"selector,omitempty"`
}

// Validate validates this metric identifier
func (m *MetricIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricIdentifier) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metric identifier based on the context it is used
func (m *MetricIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricIdentifier) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {

		if swag.IsZero(m.Selector) { // not required
			return nil
		}

		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricIdentifier) UnmarshalBinary(b []byte) error {
	var res MetricIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
