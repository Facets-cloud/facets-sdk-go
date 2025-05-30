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
)

// ModuleInputDTO ModuleInputDTO
//
// swagger:model ModuleInputDTO
type ModuleInputDTO struct {

	// compatible resources
	CompatibleResources []*InputOutputResource `json:"compatibleResources"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// optional
	Optional bool `json:"optional,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this module input d t o
func (m *ModuleInputDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatibleResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleInputDTO) validateCompatibleResources(formats strfmt.Registry) error {
	if swag.IsZero(m.CompatibleResources) { // not required
		return nil
	}

	for i := 0; i < len(m.CompatibleResources); i++ {
		if swag.IsZero(m.CompatibleResources[i]) { // not required
			continue
		}

		if m.CompatibleResources[i] != nil {
			if err := m.CompatibleResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compatibleResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compatibleResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this module input d t o based on the context it is used
func (m *ModuleInputDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompatibleResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleInputDTO) contextValidateCompatibleResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CompatibleResources); i++ {

		if m.CompatibleResources[i] != nil {

			if swag.IsZero(m.CompatibleResources[i]) { // not required
				return nil
			}

			if err := m.CompatibleResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compatibleResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compatibleResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModuleInputDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleInputDTO) UnmarshalBinary(b []byte) error {
	var res ModuleInputDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
