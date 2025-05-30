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

// ArtifactRequest ArtifactRequest
//
// swagger:model ArtifactRequest
type ArtifactRequest struct {

	// application name
	ApplicationName string `json:"applicationName,omitempty"`

	// artifact Uri
	ArtifactURI string `json:"artifactUri,omitempty"`

	// artifactory
	Artifactory string `json:"artifactory,omitempty"`

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// promoted
	Promoted bool `json:"promoted,omitempty"`

	// registered for
	// Enum: ["CLUSTER","RELEASE_STREAM"]
	RegisteredFor string `json:"registeredFor,omitempty"`

	// registration type
	// Enum: ["ENVIRONMENT","RELEASE_STREAM","HYBRID"]
	RegistrationType string `json:"registrationType,omitempty"`

	// release stream
	ReleaseStream string `json:"releaseStream,omitempty"`

	// rule Id
	RuleID string `json:"ruleId,omitempty"`

	// rule name
	RuleName string `json:"ruleName,omitempty"`

	// rule version Id
	RuleVersionID string `json:"ruleVersionId,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this artifact request
func (m *ArtifactRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegisteredFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var artifactRequestTypeRegisteredForPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLUSTER","RELEASE_STREAM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactRequestTypeRegisteredForPropEnum = append(artifactRequestTypeRegisteredForPropEnum, v)
	}
}

const (

	// ArtifactRequestRegisteredForCLUSTER captures enum value "CLUSTER"
	ArtifactRequestRegisteredForCLUSTER string = "CLUSTER"

	// ArtifactRequestRegisteredForRELEASESTREAM captures enum value "RELEASE_STREAM"
	ArtifactRequestRegisteredForRELEASESTREAM string = "RELEASE_STREAM"
)

// prop value enum
func (m *ArtifactRequest) validateRegisteredForEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactRequestTypeRegisteredForPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArtifactRequest) validateRegisteredFor(formats strfmt.Registry) error {
	if swag.IsZero(m.RegisteredFor) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegisteredForEnum("registeredFor", "body", m.RegisteredFor); err != nil {
		return err
	}

	return nil
}

var artifactRequestTypeRegistrationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENVIRONMENT","RELEASE_STREAM","HYBRID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactRequestTypeRegistrationTypePropEnum = append(artifactRequestTypeRegistrationTypePropEnum, v)
	}
}

const (

	// ArtifactRequestRegistrationTypeENVIRONMENT captures enum value "ENVIRONMENT"
	ArtifactRequestRegistrationTypeENVIRONMENT string = "ENVIRONMENT"

	// ArtifactRequestRegistrationTypeRELEASESTREAM captures enum value "RELEASE_STREAM"
	ArtifactRequestRegistrationTypeRELEASESTREAM string = "RELEASE_STREAM"

	// ArtifactRequestRegistrationTypeHYBRID captures enum value "HYBRID"
	ArtifactRequestRegistrationTypeHYBRID string = "HYBRID"
)

// prop value enum
func (m *ArtifactRequest) validateRegistrationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactRequestTypeRegistrationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArtifactRequest) validateRegistrationType(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistrationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegistrationTypeEnum("registrationType", "body", m.RegistrationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this artifact request based on context it is used
func (m *ArtifactRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactRequest) UnmarshalBinary(b []byte) error {
	var res ArtifactRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
