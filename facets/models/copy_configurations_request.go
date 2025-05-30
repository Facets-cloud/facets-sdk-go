// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CopyConfigurationsRequest CopyConfigurationsRequest
//
// swagger:model CopyConfigurationsRequest
type CopyConfigurationsRequest struct {

	// configuration types
	ConfigurationTypes []string `json:"configurationTypes"`

	// selection mode
	// Enum: ["INCLUDE","EXCLUDE"]
	SelectionMode string `json:"selectionMode,omitempty"`

	// source cluster Id
	SourceClusterID string `json:"sourceClusterId,omitempty"`
}

// Validate validates this copy configurations request
func (m *CopyConfigurationsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurationTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectionMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var copyConfigurationsRequestConfigurationTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VARIABLES_SECRETS","ARTIFACTS","ENVIRONMENT_SETTINGS","SCHEDULES","AVAILABILITY_SCHEDULES","OVERRIDES","TEMPLATE_INPUTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		copyConfigurationsRequestConfigurationTypesItemsEnum = append(copyConfigurationsRequestConfigurationTypesItemsEnum, v)
	}
}

func (m *CopyConfigurationsRequest) validateConfigurationTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, copyConfigurationsRequestConfigurationTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CopyConfigurationsRequest) validateConfigurationTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationTypes); i++ {

		// value enum
		if err := m.validateConfigurationTypesItemsEnum("configurationTypes"+"."+strconv.Itoa(i), "body", m.ConfigurationTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var copyConfigurationsRequestTypeSelectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCLUDE","EXCLUDE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		copyConfigurationsRequestTypeSelectionModePropEnum = append(copyConfigurationsRequestTypeSelectionModePropEnum, v)
	}
}

const (

	// CopyConfigurationsRequestSelectionModeINCLUDE captures enum value "INCLUDE"
	CopyConfigurationsRequestSelectionModeINCLUDE string = "INCLUDE"

	// CopyConfigurationsRequestSelectionModeEXCLUDE captures enum value "EXCLUDE"
	CopyConfigurationsRequestSelectionModeEXCLUDE string = "EXCLUDE"
)

// prop value enum
func (m *CopyConfigurationsRequest) validateSelectionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, copyConfigurationsRequestTypeSelectionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CopyConfigurationsRequest) validateSelectionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.SelectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateSelectionModeEnum("selectionMode", "body", m.SelectionMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this copy configurations request based on context it is used
func (m *CopyConfigurationsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyConfigurationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyConfigurationsRequest) UnmarshalBinary(b []byte) error {
	var res CopyConfigurationsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
