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

// LocalClusterRequest LocalClusterRequest
//
// swagger:model LocalClusterRequest
type LocalClusterRequest struct {

	// auto sign off schedule
	AutoSignOffSchedule string `json:"autoSignOffSchedule,omitempty"`

	// base cluster Id
	BaseClusterID string `json:"baseClusterId,omitempty"`

	// cd pipeline parent
	CdPipelineParent string `json:"cdPipelineParent,omitempty"`

	// cloud
	// Enum: ["AWS","AZURE","LOCAL","GCP","KUBERNETES"]
	Cloud string `json:"cloud,omitempty"`

	// cloud account Id
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// component versions
	ComponentVersions map[string]string `json:"componentVersions,omitempty"`

	// enable auto sign off
	EnableAutoSignOff bool `json:"enableAutoSignOff,omitempty"`

	// is ephemeral
	IsEphemeral bool `json:"isEphemeral,omitempty"`

	// k8s requests to limits ratio
	K8sRequestsToLimitsRatio float64 `json:"k8sRequestsToLimitsRatio,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// release stream
	ReleaseStream string `json:"releaseStream,omitempty"`

	// require sign off
	RequireSignOff bool `json:"requireSignOff,omitempty"`

	// schedules
	Schedules map[string]string `json:"schedules,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// tz
	Tz *TimeZone `json:"tz,omitempty"`
}

// Validate validates this local cluster request
func (m *LocalClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTz(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var localClusterRequestTypeCloudPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","AZURE","LOCAL","GCP","KUBERNETES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		localClusterRequestTypeCloudPropEnum = append(localClusterRequestTypeCloudPropEnum, v)
	}
}

const (

	// LocalClusterRequestCloudAWS captures enum value "AWS"
	LocalClusterRequestCloudAWS string = "AWS"

	// LocalClusterRequestCloudAZURE captures enum value "AZURE"
	LocalClusterRequestCloudAZURE string = "AZURE"

	// LocalClusterRequestCloudLOCAL captures enum value "LOCAL"
	LocalClusterRequestCloudLOCAL string = "LOCAL"

	// LocalClusterRequestCloudGCP captures enum value "GCP"
	LocalClusterRequestCloudGCP string = "GCP"

	// LocalClusterRequestCloudKUBERNETES captures enum value "KUBERNETES"
	LocalClusterRequestCloudKUBERNETES string = "KUBERNETES"
)

// prop value enum
func (m *LocalClusterRequest) validateCloudEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, localClusterRequestTypeCloudPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocalClusterRequest) validateCloud(formats strfmt.Registry) error {
	if swag.IsZero(m.Cloud) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudEnum("cloud", "body", m.Cloud); err != nil {
		return err
	}

	return nil
}

func (m *LocalClusterRequest) validateTz(formats strfmt.Registry) error {
	if swag.IsZero(m.Tz) { // not required
		return nil
	}

	if m.Tz != nil {
		if err := m.Tz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tz")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local cluster request based on the context it is used
func (m *LocalClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTz(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalClusterRequest) contextValidateTz(ctx context.Context, formats strfmt.Registry) error {

	if m.Tz != nil {

		if swag.IsZero(m.Tz) { // not required
			return nil
		}

		if err := m.Tz.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tz")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalClusterRequest) UnmarshalBinary(b []byte) error {
	var res LocalClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
