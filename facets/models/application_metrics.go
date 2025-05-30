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

// ApplicationMetrics ApplicationMetrics
//
// swagger:model ApplicationMetrics
type ApplicationMetrics struct {

	// alerts
	Alerts int32 `json:"alerts,omitempty"`

	// application Id
	ApplicationID string `json:"applicationId,omitempty"`

	// build failures
	BuildFailures int32 `json:"buildFailures,omitempty"`

	// critical code smells
	CriticalCodeSmells int32 `json:"criticalCodeSmells,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// errors
	Errors int32 `json:"errors,omitempty"`

	// fatal errors from logs
	FatalErrorsFromLogs int32 `json:"fatalErrorsFromLogs,omitempty"`

	// issues reported
	IssuesReported int32 `json:"issuesReported,omitempty"`

	// outages
	Outages int32 `json:"outages,omitempty"`

	// regression coverage
	RegressionCoverage int32 `json:"regressionCoverage,omitempty"`

	// regression failures
	RegressionFailures int32 `json:"regressionFailures,omitempty"`

	// response time
	ResponseTime float64 `json:"responseTime,omitempty"`

	// sonar Url
	SonarURL string `json:"sonarUrl,omitempty"`

	// unit test coverage
	UnitTestCoverage int32 `json:"unitTestCoverage,omitempty"`

	// unit tests
	UnitTests int32 `json:"unitTests,omitempty"`
}

// Validate validates this application metrics
func (m *ApplicationMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationMetrics) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application metrics based on context it is used
func (m *ApplicationMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationMetrics) UnmarshalBinary(b []byte) error {
	var res ApplicationMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
