// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OverrideObject OverrideObject
//
// swagger:model OverrideObject
type OverrideObject struct {

	// change log
	ChangeLog string `json:"changeLog,omitempty"`

	// commit Id
	CommitID string `json:"commitId,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// last modified date
	// Format: date-time
	LastModifiedDate strfmt.DateTime `json:"lastModifiedDate,omitempty"`

	// number of versions
	NumberOfVersions int32 `json:"numberOfVersions,omitempty"`

	// overrides
	Overrides interface{} `json:"overrides,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// versioning key
	VersioningKey string `json:"versioningKey,omitempty"`
}

// Validate validates this override object
func (m *OverrideObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OverrideObject) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OverrideObject) validateLastModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedDate", "body", "date-time", m.LastModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this override object based on context it is used
func (m *OverrideObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OverrideObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OverrideObject) UnmarshalBinary(b []byte) error {
	var res OverrideObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
