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

// DownwardAPIVolumeFile DownwardAPIVolumeFile
//
// swagger:model DownwardAPIVolumeFile
type DownwardAPIVolumeFile struct {

	// field ref
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty"`

	// mode
	Mode int32 `json:"mode,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// resource field ref
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

// Validate validates this downward API volume file
func (m *DownwardAPIVolumeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DownwardAPIVolumeFile) validateFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldRef) { // not required
		return nil
	}

	if m.FieldRef != nil {
		if err := m.FieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *DownwardAPIVolumeFile) validateResourceFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceFieldRef) { // not required
		return nil
	}

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this downward API volume file based on the context it is used
func (m *DownwardAPIVolumeFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DownwardAPIVolumeFile) contextValidateFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.FieldRef != nil {

		if swag.IsZero(m.FieldRef) { // not required
			return nil
		}

		if err := m.FieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *DownwardAPIVolumeFile) contextValidateResourceFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceFieldRef != nil {

		if swag.IsZero(m.ResourceFieldRef) { // not required
			return nil
		}

		if err := m.ResourceFieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DownwardAPIVolumeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownwardAPIVolumeFile) UnmarshalBinary(b []byte) error {
	var res DownwardAPIVolumeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
