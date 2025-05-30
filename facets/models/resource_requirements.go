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

// ResourceRequirements ResourceRequirements
//
// swagger:model ResourceRequirements
type ResourceRequirements struct {

	// claims
	Claims []*ResourceClaim `json:"claims"`

	// limits
	Limits map[string]Quantity `json:"limits,omitempty"`

	// requests
	Requests map[string]Quantity `json:"requests,omitempty"`
}

// Validate validates this resource requirements
func (m *ResourceRequirements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaims(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceRequirements) validateClaims(formats strfmt.Registry) error {
	if swag.IsZero(m.Claims) { // not required
		return nil
	}

	for i := 0; i < len(m.Claims); i++ {
		if swag.IsZero(m.Claims[i]) { // not required
			continue
		}

		if m.Claims[i] != nil {
			if err := m.Claims[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceRequirements) validateLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	for k := range m.Limits {

		if err := validate.Required("limits"+"."+k, "body", m.Limits[k]); err != nil {
			return err
		}
		if val, ok := m.Limits[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("limits" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceRequirements) validateRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	for k := range m.Requests {

		if err := validate.Required("requests"+"."+k, "body", m.Requests[k]); err != nil {
			return err
		}
		if val, ok := m.Requests[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requests" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requests" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this resource requirements based on the context it is used
func (m *ResourceRequirements) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClaims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceRequirements) contextValidateClaims(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Claims); i++ {

		if m.Claims[i] != nil {

			if swag.IsZero(m.Claims[i]) { // not required
				return nil
			}

			if err := m.Claims[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceRequirements) contextValidateLimits(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Limits {

		if val, ok := m.Limits[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ResourceRequirements) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Requests {

		if val, ok := m.Requests[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceRequirements) UnmarshalBinary(b []byte) error {
	var res ResourceRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
