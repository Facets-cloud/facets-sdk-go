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

// WebComponentDTO WebComponentDTO
//
// # Model representing a Web Component for embedding custom applications
//
// swagger:model WebComponentDTO
type WebComponentDTO struct {

	// Additional contextual attributes for component configuration
	// Example: {"theme":"dark","width":"100%"}
	ContextualAttributes interface{} `json:"contextualAttributes,omitempty"`

	// Flag to enable/disable component visibility
	// Example: true
	// Required: true
	Enabled *bool `json:"enabled"`

	// URL of the icon to be displayed in the UI
	// Example: https://example.com/icon.png
	IconURL string `json:"iconURL,omitempty"`

	// Unique identifier for the web component
	// Example: abc123
	ID string `json:"id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// last modified date
	// Format: date-time
	LastModifiedDate strfmt.DateTime `json:"lastModifiedDate,omitempty"`

	// Unique identifier for the web component
	// Example: custom-dashboard
	// Required: true
	Name *string `json:"name"`

	// Position priority of the component
	// Example: 1
	Order int32 `json:"order,omitempty"`

	// Hosted location of the component script
	// Example: https://example.com/component.js
	// Required: true
	RemoteURL *string `json:"remoteURL"`

	// Tooltip text to display on hover
	// Example: Custom Dashboard Component
	Tooltip string `json:"tooltip,omitempty"`

	// Type of web component (NAV_TYPE or TAB_TYPE)
	// Example: NAV_APP
	// Required: true
	// Enum: ["NAV_APP","TAB_APP"]
	Type *string `json:"type"`
}

// Validate validates this web component d t o
func (m *WebComponentDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebComponentDTO) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *WebComponentDTO) validateLastModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedDate", "body", "date-time", m.LastModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebComponentDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WebComponentDTO) validateRemoteURL(formats strfmt.Registry) error {

	if err := validate.Required("remoteURL", "body", m.RemoteURL); err != nil {
		return err
	}

	return nil
}

var webComponentDTOTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NAV_APP","TAB_APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webComponentDTOTypeTypePropEnum = append(webComponentDTOTypeTypePropEnum, v)
	}
}

const (

	// WebComponentDTOTypeNAVAPP captures enum value "NAV_APP"
	WebComponentDTOTypeNAVAPP string = "NAV_APP"

	// WebComponentDTOTypeTABAPP captures enum value "TAB_APP"
	WebComponentDTOTypeTABAPP string = "TAB_APP"
)

// prop value enum
func (m *WebComponentDTO) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webComponentDTOTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebComponentDTO) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web component d t o based on context it is used
func (m *WebComponentDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebComponentDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebComponentDTO) UnmarshalBinary(b []byte) error {
	var res WebComponentDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
