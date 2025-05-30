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

// ActionExecution ActionExecution
//
// swagger:model ActionExecution
type ActionExecution struct {

	// action
	Action *ApplicationAction `json:"action,omitempty"`

	// application Id
	ApplicationID string `json:"applicationId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// output
	Output string `json:"output,omitempty"`

	// trigger exception
	TriggerException string `json:"triggerException,omitempty"`

	// trigger status
	// Enum: ["SUCCESS","FAILURE"]
	TriggerStatus string `json:"triggerStatus,omitempty"`

	// trigger time
	TriggerTime int64 `json:"triggerTime,omitempty"`
}

// Validate validates this action execution
func (m *ActionExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionExecution) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

var actionExecutionTypeTriggerStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionExecutionTypeTriggerStatusPropEnum = append(actionExecutionTypeTriggerStatusPropEnum, v)
	}
}

const (

	// ActionExecutionTriggerStatusSUCCESS captures enum value "SUCCESS"
	ActionExecutionTriggerStatusSUCCESS string = "SUCCESS"

	// ActionExecutionTriggerStatusFAILURE captures enum value "FAILURE"
	ActionExecutionTriggerStatusFAILURE string = "FAILURE"
)

// prop value enum
func (m *ActionExecution) validateTriggerStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actionExecutionTypeTriggerStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActionExecution) validateTriggerStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.TriggerStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerStatusEnum("triggerStatus", "body", m.TriggerStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this action execution based on the context it is used
func (m *ActionExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionExecution) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {

		if swag.IsZero(m.Action) { // not required
			return nil
		}

		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionExecution) UnmarshalBinary(b []byte) error {
	var res ActionExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
