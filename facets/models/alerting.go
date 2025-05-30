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

// Alerting Alerting
//
// swagger:model Alerting
type Alerting struct {

	// application family
	// Enum: ["CRM","ECOMMERCE","INTEGRATIONS","OPS"]
	ApplicationFamily string `json:"applicationFamily,omitempty"`

	// application Id
	ApplicationID string `json:"applicationId,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// new relic alerts Url
	NewRelicAlertsURL string `json:"newRelicAlertsUrl,omitempty"`
}

// Validate validates this alerting
func (m *Alerting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationFamily(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var alertingTypeApplicationFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CRM","ECOMMERCE","INTEGRATIONS","OPS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertingTypeApplicationFamilyPropEnum = append(alertingTypeApplicationFamilyPropEnum, v)
	}
}

const (

	// AlertingApplicationFamilyCRM captures enum value "CRM"
	AlertingApplicationFamilyCRM string = "CRM"

	// AlertingApplicationFamilyECOMMERCE captures enum value "ECOMMERCE"
	AlertingApplicationFamilyECOMMERCE string = "ECOMMERCE"

	// AlertingApplicationFamilyINTEGRATIONS captures enum value "INTEGRATIONS"
	AlertingApplicationFamilyINTEGRATIONS string = "INTEGRATIONS"

	// AlertingApplicationFamilyOPS captures enum value "OPS"
	AlertingApplicationFamilyOPS string = "OPS"
)

// prop value enum
func (m *Alerting) validateApplicationFamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertingTypeApplicationFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Alerting) validateApplicationFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationFamily) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationFamilyEnum("applicationFamily", "body", m.ApplicationFamily); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alerting based on context it is used
func (m *Alerting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Alerting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alerting) UnmarshalBinary(b []byte) error {
	var res Alerting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
