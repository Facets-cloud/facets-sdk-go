// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PodDTO PodDTO
//
// swagger:model PodDTO
type PodDTO struct {

	// age in seconds
	AgeInSeconds int64 `json:"ageInSeconds,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ready
	Ready string `json:"ready,omitempty"`

	// restarts
	Restarts int32 `json:"restarts,omitempty"`

	// role
	// Enum: ["active","preview","na"]
	Role string `json:"role,omitempty"`

	// status
	// Enum: ["Pending","Running","Succeeded","Failed","Unknown","CrashLoopBackOff","ImagePullBackOff","ContainerCreating","Terminating","Evicted","Completed","Restarting","Error","InvalidImageName"]
	Status string `json:"status,omitempty"`
}

// Validate validates this pod d t o
func (m *PodDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var podDTOTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","preview","na"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		podDTOTypeRolePropEnum = append(podDTOTypeRolePropEnum, v)
	}
}

const (

	// PodDTORoleActive captures enum value "active"
	PodDTORoleActive string = "active"

	// PodDTORolePreview captures enum value "preview"
	PodDTORolePreview string = "preview"

	// PodDTORoleNa captures enum value "na"
	PodDTORoleNa string = "na"
)

// prop value enum
func (m *PodDTO) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, podDTOTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PodDTO) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var podDTOTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Running","Succeeded","Failed","Unknown","CrashLoopBackOff","ImagePullBackOff","ContainerCreating","Terminating","Evicted","Completed","Restarting","Error","InvalidImageName"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		podDTOTypeStatusPropEnum = append(podDTOTypeStatusPropEnum, v)
	}
}

const (

	// PodDTOStatusPending captures enum value "Pending"
	PodDTOStatusPending string = "Pending"

	// PodDTOStatusRunning captures enum value "Running"
	PodDTOStatusRunning string = "Running"

	// PodDTOStatusSucceeded captures enum value "Succeeded"
	PodDTOStatusSucceeded string = "Succeeded"

	// PodDTOStatusFailed captures enum value "Failed"
	PodDTOStatusFailed string = "Failed"

	// PodDTOStatusUnknown captures enum value "Unknown"
	PodDTOStatusUnknown string = "Unknown"

	// PodDTOStatusCrashLoopBackOff captures enum value "CrashLoopBackOff"
	PodDTOStatusCrashLoopBackOff string = "CrashLoopBackOff"

	// PodDTOStatusImagePullBackOff captures enum value "ImagePullBackOff"
	PodDTOStatusImagePullBackOff string = "ImagePullBackOff"

	// PodDTOStatusContainerCreating captures enum value "ContainerCreating"
	PodDTOStatusContainerCreating string = "ContainerCreating"

	// PodDTOStatusTerminating captures enum value "Terminating"
	PodDTOStatusTerminating string = "Terminating"

	// PodDTOStatusEvicted captures enum value "Evicted"
	PodDTOStatusEvicted string = "Evicted"

	// PodDTOStatusCompleted captures enum value "Completed"
	PodDTOStatusCompleted string = "Completed"

	// PodDTOStatusRestarting captures enum value "Restarting"
	PodDTOStatusRestarting string = "Restarting"

	// PodDTOStatusError captures enum value "Error"
	PodDTOStatusError string = "Error"

	// PodDTOStatusInvalidImageName captures enum value "InvalidImageName"
	PodDTOStatusInvalidImageName string = "InvalidImageName"
)

// prop value enum
func (m *PodDTO) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, podDTOTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PodDTO) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pod d t o based on context it is used
func (m *PodDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodDTO) UnmarshalBinary(b []byte) error {
	var res PodDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
