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

// Monitoring Monitoring
//
// swagger:model Monitoring
type Monitoring struct {

	// application family
	// Enum: ["CRM","ECOMMERCE","INTEGRATIONS","OPS"]
	ApplicationFamily string `json:"applicationFamily,omitempty"`

	// application Id
	ApplicationID string `json:"applicationId,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// new relic dashboard Url
	NewRelicDashboardURL string `json:"newRelicDashboardUrl,omitempty"`
}

// Validate validates this monitoring
func (m *Monitoring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationFamily(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var monitoringTypeApplicationFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CRM","ECOMMERCE","INTEGRATIONS","OPS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		monitoringTypeApplicationFamilyPropEnum = append(monitoringTypeApplicationFamilyPropEnum, v)
	}
}

const (

	// MonitoringApplicationFamilyCRM captures enum value "CRM"
	MonitoringApplicationFamilyCRM string = "CRM"

	// MonitoringApplicationFamilyECOMMERCE captures enum value "ECOMMERCE"
	MonitoringApplicationFamilyECOMMERCE string = "ECOMMERCE"

	// MonitoringApplicationFamilyINTEGRATIONS captures enum value "INTEGRATIONS"
	MonitoringApplicationFamilyINTEGRATIONS string = "INTEGRATIONS"

	// MonitoringApplicationFamilyOPS captures enum value "OPS"
	MonitoringApplicationFamilyOPS string = "OPS"
)

// prop value enum
func (m *Monitoring) validateApplicationFamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, monitoringTypeApplicationFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Monitoring) validateApplicationFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationFamily) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationFamilyEnum("applicationFamily", "body", m.ApplicationFamily); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this monitoring based on context it is used
func (m *Monitoring) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Monitoring) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Monitoring) UnmarshalBinary(b []byte) error {
	var res Monitoring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
