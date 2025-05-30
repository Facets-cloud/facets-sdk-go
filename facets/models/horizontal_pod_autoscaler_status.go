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

// HorizontalPodAutoscalerStatus HorizontalPodAutoscalerStatus
//
// swagger:model HorizontalPodAutoscalerStatus
type HorizontalPodAutoscalerStatus struct {

	// conditions
	Conditions []*HorizontalPodAutoscalerCondition `json:"conditions"`

	// current metrics
	CurrentMetrics []*MetricStatus `json:"currentMetrics"`

	// current replicas
	CurrentReplicas int32 `json:"currentReplicas,omitempty"`

	// desired replicas
	DesiredReplicas int32 `json:"desiredReplicas,omitempty"`

	// last scale time
	LastScaleTime string `json:"lastScaleTime,omitempty"`

	// observed generation
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// Validate validates this horizontal pod autoscaler status
func (m *HorizontalPodAutoscalerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HorizontalPodAutoscalerStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *HorizontalPodAutoscalerStatus) validateCurrentMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentMetrics) { // not required
		return nil
	}

	for i := 0; i < len(m.CurrentMetrics); i++ {
		if swag.IsZero(m.CurrentMetrics[i]) { // not required
			continue
		}

		if m.CurrentMetrics[i] != nil {
			if err := m.CurrentMetrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("currentMetrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("currentMetrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this horizontal pod autoscaler status based on the context it is used
func (m *HorizontalPodAutoscalerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HorizontalPodAutoscalerStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HorizontalPodAutoscalerStatus) contextValidateCurrentMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CurrentMetrics); i++ {

		if m.CurrentMetrics[i] != nil {

			if swag.IsZero(m.CurrentMetrics[i]) { // not required
				return nil
			}

			if err := m.CurrentMetrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("currentMetrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("currentMetrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HorizontalPodAutoscalerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HorizontalPodAutoscalerStatus) UnmarshalBinary(b []byte) error {
	var res HorizontalPodAutoscalerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
