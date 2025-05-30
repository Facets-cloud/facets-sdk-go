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

// ServiceDTO ServiceDTO
//
// swagger:model ServiceDTO
type ServiceDTO struct {

	// age
	Age int64 `json:"age,omitempty"`

	// cluster Ip
	ClusterIP string `json:"clusterIp,omitempty"`

	// external Ip
	ExternalIP string `json:"externalIp,omitempty"`

	// ip families
	IPFamilies []string `json:"ipFamilies"`

	// name
	Name string `json:"name,omitempty"`

	// port
	Port []*ServicePortDTO `json:"port"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this service d t o
func (m *ServiceDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTO) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	for i := 0; i < len(m.Port); i++ {
		if swag.IsZero(m.Port[i]) { // not required
			continue
		}

		if m.Port[i] != nil {
			if err := m.Port[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service d t o based on the context it is used
func (m *ServiceDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTO) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Port); i++ {

		if m.Port[i] != nil {

			if swag.IsZero(m.Port[i]) { // not required
				return nil
			}

			if err := m.Port[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTO) UnmarshalBinary(b []byte) error {
	var res ServiceDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
