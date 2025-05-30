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

// ResourceHistory ResourceHistory
//
// swagger:model ResourceHistory
type ResourceHistory struct {

	// alpha
	Alpha bool `json:"alpha,omitempty"`

	// artifact version
	ArtifactVersion *HistoryMetadata `json:"artifactVersion,omitempty"`

	// blueprint version
	BlueprintVersion *HistoryMetadata `json:"blueprintVersion,omitempty"`

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// override version
	OverrideVersion *HistoryMetadata `json:"overrideVersion,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource status
	// Enum: ["RUNNING","STOPPED"]
	ResourceStatus string `json:"resourceStatus,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}

// Validate validates this resource history
func (m *ResourceHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlueprintVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrideVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceHistory) validateArtifactVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.ArtifactVersion) { // not required
		return nil
	}

	if m.ArtifactVersion != nil {
		if err := m.ArtifactVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifactVersion")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceHistory) validateBlueprintVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.BlueprintVersion) { // not required
		return nil
	}

	if m.BlueprintVersion != nil {
		if err := m.BlueprintVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blueprintVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blueprintVersion")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceHistory) validateOverrideVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.OverrideVersion) { // not required
		return nil
	}

	if m.OverrideVersion != nil {
		if err := m.OverrideVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrideVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overrideVersion")
			}
			return err
		}
	}

	return nil
}

var resourceHistoryTypeResourceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceHistoryTypeResourceStatusPropEnum = append(resourceHistoryTypeResourceStatusPropEnum, v)
	}
}

const (

	// ResourceHistoryResourceStatusRUNNING captures enum value "RUNNING"
	ResourceHistoryResourceStatusRUNNING string = "RUNNING"

	// ResourceHistoryResourceStatusSTOPPED captures enum value "STOPPED"
	ResourceHistoryResourceStatusSTOPPED string = "STOPPED"
)

// prop value enum
func (m *ResourceHistory) validateResourceStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceHistoryTypeResourceStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResourceHistory) validateResourceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceStatusEnum("resourceStatus", "body", m.ResourceStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resource history based on the context it is used
func (m *ResourceHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifactVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlueprintVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverrideVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceHistory) contextValidateArtifactVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.ArtifactVersion != nil {

		if swag.IsZero(m.ArtifactVersion) { // not required
			return nil
		}

		if err := m.ArtifactVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifactVersion")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceHistory) contextValidateBlueprintVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.BlueprintVersion != nil {

		if swag.IsZero(m.BlueprintVersion) { // not required
			return nil
		}

		if err := m.BlueprintVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blueprintVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blueprintVersion")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceHistory) contextValidateOverrideVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.OverrideVersion != nil {

		if swag.IsZero(m.OverrideVersion) { // not required
			return nil
		}

		if err := m.OverrideVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrideVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overrideVersion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceHistory) UnmarshalBinary(b []byte) error {
	var res ResourceHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
