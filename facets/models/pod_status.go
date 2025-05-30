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

// PodStatus PodStatus
//
// swagger:model PodStatus
type PodStatus struct {

	// conditions
	Conditions []*PodCondition `json:"conditions"`

	// container statuses
	ContainerStatuses []*ContainerStatus `json:"containerStatuses"`

	// ephemeral container statuses
	EphemeralContainerStatuses []*ContainerStatus `json:"ephemeralContainerStatuses"`

	// host IP
	HostIP string `json:"hostIP,omitempty"`

	// init container statuses
	InitContainerStatuses []*ContainerStatus `json:"initContainerStatuses"`

	// message
	Message string `json:"message,omitempty"`

	// nominated node name
	NominatedNodeName string `json:"nominatedNodeName,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// pod IP
	PodIP string `json:"podIP,omitempty"`

	// pod i ps
	PodIPs []*PodIP `json:"podIPs"`

	// qos class
	QosClass string `json:"qosClass,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this pod status
func (m *PodStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodIPs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) validateContainerStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.ContainerStatuses); i++ {
		if swag.IsZero(m.ContainerStatuses[i]) { // not required
			continue
		}

		if m.ContainerStatuses[i] != nil {
			if err := m.ContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) validateEphemeralContainerStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.EphemeralContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.EphemeralContainerStatuses); i++ {
		if swag.IsZero(m.EphemeralContainerStatuses[i]) { // not required
			continue
		}

		if m.EphemeralContainerStatuses[i] != nil {
			if err := m.EphemeralContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeralContainerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ephemeralContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) validateInitContainerStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.InitContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.InitContainerStatuses); i++ {
		if swag.IsZero(m.InitContainerStatuses[i]) { // not required
			continue
		}

		if m.InitContainerStatuses[i] != nil {
			if err := m.InitContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initContainerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) validatePodIPs(formats strfmt.Registry) error {
	if swag.IsZero(m.PodIPs) { // not required
		return nil
	}

	for i := 0; i < len(m.PodIPs); i++ {
		if swag.IsZero(m.PodIPs[i]) { // not required
			continue
		}

		if m.PodIPs[i] != nil {
			if err := m.PodIPs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podIPs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pod status based on the context it is used
func (m *PodStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainerStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEphemeralContainerStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitContainerStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodIPs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {

			if swag.IsZero(m.Conditions[i]) { // not required
				return nil
			}

			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) contextValidateContainerStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContainerStatuses); i++ {

		if m.ContainerStatuses[i] != nil {

			if swag.IsZero(m.ContainerStatuses[i]) { // not required
				return nil
			}

			if err := m.ContainerStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) contextValidateEphemeralContainerStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EphemeralContainerStatuses); i++ {

		if m.EphemeralContainerStatuses[i] != nil {

			if swag.IsZero(m.EphemeralContainerStatuses[i]) { // not required
				return nil
			}

			if err := m.EphemeralContainerStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeralContainerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ephemeralContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) contextValidateInitContainerStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InitContainerStatuses); i++ {

		if m.InitContainerStatuses[i] != nil {

			if swag.IsZero(m.InitContainerStatuses[i]) { // not required
				return nil
			}

			if err := m.InitContainerStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initContainerStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodStatus) contextValidatePodIPs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PodIPs); i++ {

		if m.PodIPs[i] != nil {

			if swag.IsZero(m.PodIPs[i]) { // not required
				return nil
			}

			if err := m.PodIPs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podIPs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodStatus) UnmarshalBinary(b []byte) error {
	var res PodStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
