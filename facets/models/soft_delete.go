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

// SoftDelete SoftDelete
//
// swagger:model SoftDelete
type SoftDelete struct {

	// entity type
	// Enum: ["CLUSTER","BLUE_PRINT","TEMPLATE_INPUT","CONTROL_PLANE","IAC","ARTIFACT_CI","USER_GROUP","ACCOUNT","ARTIFACTORY"]
	EntityType string `json:"entityType,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this soft delete
func (m *SoftDelete) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var softDeleteTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLUSTER","BLUE_PRINT","TEMPLATE_INPUT","CONTROL_PLANE","IAC","ARTIFACT_CI","USER_GROUP","ACCOUNT","ARTIFACTORY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softDeleteTypeEntityTypePropEnum = append(softDeleteTypeEntityTypePropEnum, v)
	}
}

const (

	// SoftDeleteEntityTypeCLUSTER captures enum value "CLUSTER"
	SoftDeleteEntityTypeCLUSTER string = "CLUSTER"

	// SoftDeleteEntityTypeBLUEPRINT captures enum value "BLUE_PRINT"
	SoftDeleteEntityTypeBLUEPRINT string = "BLUE_PRINT"

	// SoftDeleteEntityTypeTEMPLATEINPUT captures enum value "TEMPLATE_INPUT"
	SoftDeleteEntityTypeTEMPLATEINPUT string = "TEMPLATE_INPUT"

	// SoftDeleteEntityTypeCONTROLPLANE captures enum value "CONTROL_PLANE"
	SoftDeleteEntityTypeCONTROLPLANE string = "CONTROL_PLANE"

	// SoftDeleteEntityTypeIAC captures enum value "IAC"
	SoftDeleteEntityTypeIAC string = "IAC"

	// SoftDeleteEntityTypeARTIFACTCI captures enum value "ARTIFACT_CI"
	SoftDeleteEntityTypeARTIFACTCI string = "ARTIFACT_CI"

	// SoftDeleteEntityTypeUSERGROUP captures enum value "USER_GROUP"
	SoftDeleteEntityTypeUSERGROUP string = "USER_GROUP"

	// SoftDeleteEntityTypeACCOUNT captures enum value "ACCOUNT"
	SoftDeleteEntityTypeACCOUNT string = "ACCOUNT"

	// SoftDeleteEntityTypeARTIFACTORY captures enum value "ARTIFACTORY"
	SoftDeleteEntityTypeARTIFACTORY string = "ARTIFACTORY"
)

// prop value enum
func (m *SoftDelete) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softDeleteTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftDelete) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this soft delete based on context it is used
func (m *SoftDelete) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoftDelete) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftDelete) UnmarshalBinary(b []byte) error {
	var res SoftDelete
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
