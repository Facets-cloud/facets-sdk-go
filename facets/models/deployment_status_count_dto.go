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

// DeploymentStatusCountDto DeploymentStatusCountDto
//
// swagger:model DeploymentStatusCountDto
type DeploymentStatusCountDto struct {

	// count
	Count int64 `json:"count,omitempty"`

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// status
	// Enum: ["SUCCEEDED","FAILED","FAULT","TIMED_OUT","IN_PROGRESS","STOPPED","INVALID","STARTED","UNKNOWN","QUEUED","PENDING_APPROVAL","APPROVED","REJECTED"]
	Status string `json:"status,omitempty"`
}

// Validate validates this deployment status count dto
func (m *DeploymentStatusCountDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
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

func (m *DeploymentStatusCountDto) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var deploymentStatusCountDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCEEDED","FAILED","FAULT","TIMED_OUT","IN_PROGRESS","STOPPED","INVALID","STARTED","UNKNOWN","QUEUED","PENDING_APPROVAL","APPROVED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentStatusCountDtoTypeStatusPropEnum = append(deploymentStatusCountDtoTypeStatusPropEnum, v)
	}
}

const (

	// DeploymentStatusCountDtoStatusSUCCEEDED captures enum value "SUCCEEDED"
	DeploymentStatusCountDtoStatusSUCCEEDED string = "SUCCEEDED"

	// DeploymentStatusCountDtoStatusFAILED captures enum value "FAILED"
	DeploymentStatusCountDtoStatusFAILED string = "FAILED"

	// DeploymentStatusCountDtoStatusFAULT captures enum value "FAULT"
	DeploymentStatusCountDtoStatusFAULT string = "FAULT"

	// DeploymentStatusCountDtoStatusTIMEDOUT captures enum value "TIMED_OUT"
	DeploymentStatusCountDtoStatusTIMEDOUT string = "TIMED_OUT"

	// DeploymentStatusCountDtoStatusINPROGRESS captures enum value "IN_PROGRESS"
	DeploymentStatusCountDtoStatusINPROGRESS string = "IN_PROGRESS"

	// DeploymentStatusCountDtoStatusSTOPPED captures enum value "STOPPED"
	DeploymentStatusCountDtoStatusSTOPPED string = "STOPPED"

	// DeploymentStatusCountDtoStatusINVALID captures enum value "INVALID"
	DeploymentStatusCountDtoStatusINVALID string = "INVALID"

	// DeploymentStatusCountDtoStatusSTARTED captures enum value "STARTED"
	DeploymentStatusCountDtoStatusSTARTED string = "STARTED"

	// DeploymentStatusCountDtoStatusUNKNOWN captures enum value "UNKNOWN"
	DeploymentStatusCountDtoStatusUNKNOWN string = "UNKNOWN"

	// DeploymentStatusCountDtoStatusQUEUED captures enum value "QUEUED"
	DeploymentStatusCountDtoStatusQUEUED string = "QUEUED"

	// DeploymentStatusCountDtoStatusPENDINGAPPROVAL captures enum value "PENDING_APPROVAL"
	DeploymentStatusCountDtoStatusPENDINGAPPROVAL string = "PENDING_APPROVAL"

	// DeploymentStatusCountDtoStatusAPPROVED captures enum value "APPROVED"
	DeploymentStatusCountDtoStatusAPPROVED string = "APPROVED"

	// DeploymentStatusCountDtoStatusREJECTED captures enum value "REJECTED"
	DeploymentStatusCountDtoStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *DeploymentStatusCountDto) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentStatusCountDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentStatusCountDto) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deployment status count dto based on context it is used
func (m *DeploymentStatusCountDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentStatusCountDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentStatusCountDto) UnmarshalBinary(b []byte) error {
	var res DeploymentStatusCountDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
