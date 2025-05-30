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

// StackGitLog StackGitLog
//
// swagger:model StackGitLog
type StackGitLog struct {

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`

	// commit log
	CommitLog *CommitLog `json:"commitLog,omitempty"`

	// directory
	Directory string `json:"directory,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// number of versions
	NumberOfVersions int32 `json:"numberOfVersions,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// versioning key
	VersioningKey string `json:"versioningKey,omitempty"`
}

// Validate validates this stack git log
func (m *StackGitLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackGitLog) validateCommitLog(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitLog) { // not required
		return nil
	}

	if m.CommitLog != nil {
		if err := m.CommitLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commitLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commitLog")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stack git log based on the context it is used
func (m *StackGitLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommitLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackGitLog) contextValidateCommitLog(ctx context.Context, formats strfmt.Registry) error {

	if m.CommitLog != nil {

		if swag.IsZero(m.CommitLog) { // not required
			return nil
		}

		if err := m.CommitLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commitLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commitLog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackGitLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackGitLog) UnmarshalBinary(b []byte) error {
	var res StackGitLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
