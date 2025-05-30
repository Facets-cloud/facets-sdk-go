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

// AppDeployment AppDeployment
//
// swagger:model AppDeployment
type AppDeployment struct {

	// app name
	AppName string `json:"appName,omitempty"`

	// artifact
	Artifact *Artifact `json:"artifact,omitempty"`
}

// Validate validates this app deployment
func (m *AppDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppDeployment) validateArtifact(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifact) { // not required
		return nil
	}

	if m.Artifact != nil {
		if err := m.Artifact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app deployment based on the context it is used
func (m *AppDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppDeployment) contextValidateArtifact(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifact != nil {

		if swag.IsZero(m.Artifact) { // not required
			return nil
		}

		if err := m.Artifact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppDeployment) UnmarshalBinary(b []byte) error {
	var res AppDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
