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

// VolumeProjection VolumeProjection
//
// swagger:model VolumeProjection
type VolumeProjection struct {

	// config map
	ConfigMap *ConfigMapProjection `json:"configMap,omitempty"`

	// downward API
	DownwardAPI *DownwardAPIProjection `json:"downwardAPI,omitempty"`

	// secret
	Secret *SecretProjection `json:"secret,omitempty"`

	// service account token
	ServiceAccountToken *ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}

// Validate validates this volume projection
func (m *VolumeProjection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownwardAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeProjection) validateConfigMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMap) { // not required
		return nil
	}

	if m.ConfigMap != nil {
		if err := m.ConfigMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) validateDownwardAPI(formats strfmt.Registry) error {
	if swag.IsZero(m.DownwardAPI) { // not required
		return nil
	}

	if m.DownwardAPI != nil {
		if err := m.DownwardAPI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downwardAPI")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downwardAPI")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) validateServiceAccountToken(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccountToken) { // not required
		return nil
	}

	if m.ServiceAccountToken != nil {
		if err := m.ServiceAccountToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceAccountToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceAccountToken")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume projection based on the context it is used
func (m *VolumeProjection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownwardAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceAccountToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeProjection) contextValidateConfigMap(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigMap != nil {

		if swag.IsZero(m.ConfigMap) { // not required
			return nil
		}

		if err := m.ConfigMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) contextValidateDownwardAPI(ctx context.Context, formats strfmt.Registry) error {

	if m.DownwardAPI != nil {

		if swag.IsZero(m.DownwardAPI) { // not required
			return nil
		}

		if err := m.DownwardAPI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downwardAPI")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downwardAPI")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {

		if swag.IsZero(m.Secret) { // not required
			return nil
		}

		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeProjection) contextValidateServiceAccountToken(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceAccountToken != nil {

		if swag.IsZero(m.ServiceAccountToken) { // not required
			return nil
		}

		if err := m.ServiceAccountToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceAccountToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceAccountToken")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeProjection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeProjection) UnmarshalBinary(b []byte) error {
	var res VolumeProjection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
