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

// Lifecycle Lifecycle
//
// swagger:model Lifecycle
type Lifecycle struct {

	// post start
	PostStart *LifecycleHandler `json:"postStart,omitempty"`

	// pre stop
	PreStop *LifecycleHandler `json:"preStop,omitempty"`
}

// Validate validates this lifecycle
func (m *Lifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreStop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Lifecycle) validatePostStart(formats strfmt.Registry) error {
	if swag.IsZero(m.PostStart) { // not required
		return nil
	}

	if m.PostStart != nil {
		if err := m.PostStart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *Lifecycle) validatePreStop(formats strfmt.Registry) error {
	if swag.IsZero(m.PreStop) { // not required
		return nil
	}

	if m.PreStop != nil {
		if err := m.PreStop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lifecycle based on the context it is used
func (m *Lifecycle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePostStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreStop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Lifecycle) contextValidatePostStart(ctx context.Context, formats strfmt.Registry) error {

	if m.PostStart != nil {

		if swag.IsZero(m.PostStart) { // not required
			return nil
		}

		if err := m.PostStart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *Lifecycle) contextValidatePreStop(ctx context.Context, formats strfmt.Registry) error {

	if m.PreStop != nil {

		if swag.IsZero(m.PreStop) { // not required
			return nil
		}

		if err := m.PreStop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Lifecycle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Lifecycle) UnmarshalBinary(b []byte) error {
	var res Lifecycle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
