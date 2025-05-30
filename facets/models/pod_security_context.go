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

// PodSecurityContext PodSecurityContext
//
// swagger:model PodSecurityContext
type PodSecurityContext struct {

	// fs group
	FsGroup int64 `json:"fsGroup,omitempty"`

	// fs group change policy
	FsGroupChangePolicy string `json:"fsGroupChangePolicy,omitempty"`

	// run as group
	RunAsGroup int64 `json:"runAsGroup,omitempty"`

	// run as non root
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`

	// run as user
	RunAsUser int64 `json:"runAsUser,omitempty"`

	// se linux options
	SeLinuxOptions *SELinuxOptions `json:"seLinuxOptions,omitempty"`

	// seccomp profile
	SeccompProfile *SeccompProfile `json:"seccompProfile,omitempty"`

	// supplemental groups
	SupplementalGroups []int64 `json:"supplementalGroups"`

	// sysctls
	Sysctls []*Sysctl `json:"sysctls"`

	// windows options
	WindowsOptions *WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

// Validate validates this pod security context
func (m *PodSecurityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeLinuxOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeccompProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSysctls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindowsOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodSecurityContext) validateSeLinuxOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SeLinuxOptions) { // not required
		return nil
	}

	if m.SeLinuxOptions != nil {
		if err := m.SeLinuxOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *PodSecurityContext) validateSeccompProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.SeccompProfile) { // not required
		return nil
	}

	if m.SeccompProfile != nil {
		if err := m.SeccompProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seccompProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seccompProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PodSecurityContext) validateSysctls(formats strfmt.Registry) error {
	if swag.IsZero(m.Sysctls) { // not required
		return nil
	}

	for i := 0; i < len(m.Sysctls); i++ {
		if swag.IsZero(m.Sysctls[i]) { // not required
			continue
		}

		if m.Sysctls[i] != nil {
			if err := m.Sysctls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sysctls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sysctls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSecurityContext) validateWindowsOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.WindowsOptions) { // not required
		return nil
	}

	if m.WindowsOptions != nil {
		if err := m.WindowsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pod security context based on the context it is used
func (m *PodSecurityContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSeLinuxOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeccompProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSysctls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWindowsOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodSecurityContext) contextValidateSeLinuxOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.SeLinuxOptions != nil {

		if swag.IsZero(m.SeLinuxOptions) { // not required
			return nil
		}

		if err := m.SeLinuxOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *PodSecurityContext) contextValidateSeccompProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.SeccompProfile != nil {

		if swag.IsZero(m.SeccompProfile) { // not required
			return nil
		}

		if err := m.SeccompProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seccompProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seccompProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PodSecurityContext) contextValidateSysctls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sysctls); i++ {

		if m.Sysctls[i] != nil {

			if swag.IsZero(m.Sysctls[i]) { // not required
				return nil
			}

			if err := m.Sysctls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sysctls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sysctls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSecurityContext) contextValidateWindowsOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.WindowsOptions != nil {

		if swag.IsZero(m.WindowsOptions) { // not required
			return nil
		}

		if err := m.WindowsOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodSecurityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodSecurityContext) UnmarshalBinary(b []byte) error {
	var res PodSecurityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
