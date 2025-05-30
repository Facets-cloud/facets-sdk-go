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

// TopologySpreadConstraint TopologySpreadConstraint
//
// swagger:model TopologySpreadConstraint
type TopologySpreadConstraint struct {

	// label selector
	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`

	// match label keys
	MatchLabelKeys []string `json:"matchLabelKeys"`

	// max skew
	MaxSkew int32 `json:"maxSkew,omitempty"`

	// min domains
	MinDomains int32 `json:"minDomains,omitempty"`

	// node affinity policy
	NodeAffinityPolicy string `json:"nodeAffinityPolicy,omitempty"`

	// node taints policy
	NodeTaintsPolicy string `json:"nodeTaintsPolicy,omitempty"`

	// topology key
	TopologyKey string `json:"topologyKey,omitempty"`

	// when unsatisfiable
	WhenUnsatisfiable string `json:"whenUnsatisfiable,omitempty"`
}

// Validate validates this topology spread constraint
func (m *TopologySpreadConstraint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologySpreadConstraint) validateLabelSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelSelector) { // not required
		return nil
	}

	if m.LabelSelector != nil {
		if err := m.LabelSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelSelector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this topology spread constraint based on the context it is used
func (m *TopologySpreadConstraint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologySpreadConstraint) contextValidateLabelSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelSelector != nil {

		if swag.IsZero(m.LabelSelector) { // not required
			return nil
		}

		if err := m.LabelSelector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelSelector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologySpreadConstraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologySpreadConstraint) UnmarshalBinary(b []byte) error {
	var res TopologySpreadConstraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
