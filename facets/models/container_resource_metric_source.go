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

// ContainerResourceMetricSource ContainerResourceMetricSource
//
// swagger:model ContainerResourceMetricSource
type ContainerResourceMetricSource struct {

	// container
	Container string `json:"container,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// target
	Target *MetricTarget `json:"target,omitempty"`
}

// Validate validates this container resource metric source
func (m *ContainerResourceMetricSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerResourceMetricSource) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container resource metric source based on the context it is used
func (m *ContainerResourceMetricSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerResourceMetricSource) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerResourceMetricSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerResourceMetricSource) UnmarshalBinary(b []byte) error {
	var res ContainerResourceMetricSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
