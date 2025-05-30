// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MatchedResourceDTO MatchedResourceDTO
//
// swagger:model MatchedResourceDTO
type MatchedResourceDTO struct {

	// matched
	Matched bool `json:"matched,omitempty"`

	// module ref
	ModuleRef string `json:"moduleRef,omitempty"`
}

// Validate validates this matched resource d t o
func (m *MatchedResourceDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this matched resource d t o based on context it is used
func (m *MatchedResourceDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchedResourceDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchedResourceDTO) UnmarshalBinary(b []byte) error {
	var res MatchedResourceDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
