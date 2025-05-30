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

// EnvSpecificAccess EnvSpecificAccess
//
// swagger:model EnvSpecificAccess
type EnvSpecificAccess struct {

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`

	// env name
	EnvName string `json:"envName,omitempty"`

	// project
	Project string `json:"project,omitempty"`

	// role info
	RoleInfo *RoleInfo `json:"roleInfo,omitempty"`
}

// Validate validates this env specific access
func (m *EnvSpecificAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvSpecificAccess) validateRoleInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleInfo) { // not required
		return nil
	}

	if m.RoleInfo != nil {
		if err := m.RoleInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this env specific access based on the context it is used
func (m *EnvSpecificAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvSpecificAccess) contextValidateRoleInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleInfo != nil {

		if swag.IsZero(m.RoleInfo) { // not required
			return nil
		}

		if err := m.RoleInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvSpecificAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvSpecificAccess) UnmarshalBinary(b []byte) error {
	var res EnvSpecificAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
