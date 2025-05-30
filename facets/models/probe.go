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

// Probe Probe
//
// swagger:model Probe
type Probe struct {

	// exec
	Exec *ExecAction `json:"exec,omitempty"`

	// failure threshold
	FailureThreshold int32 `json:"failureThreshold,omitempty"`

	// grpc
	Grpc *GRPCAction `json:"grpc,omitempty"`

	// http get
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`

	// initial delay seconds
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`

	// period seconds
	PeriodSeconds int32 `json:"periodSeconds,omitempty"`

	// success threshold
	SuccessThreshold int32 `json:"successThreshold,omitempty"`

	// tcp socket
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty"`

	// termination grace period seconds
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// timeout seconds
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

// Validate validates this probe
func (m *Probe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPGet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPSocket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Probe) validateExec(formats strfmt.Registry) error {
	if swag.IsZero(m.Exec) { // not required
		return nil
	}

	if m.Exec != nil {
		if err := m.Exec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) validateGrpc(formats strfmt.Registry) error {
	if swag.IsZero(m.Grpc) { // not required
		return nil
	}

	if m.Grpc != nil {
		if err := m.Grpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grpc")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) validateHTTPGet(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPGet) { // not required
		return nil
	}

	if m.HTTPGet != nil {
		if err := m.HTTPGet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) validateTCPSocket(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPSocket) { // not required
		return nil
	}

	if m.TCPSocket != nil {
		if err := m.TCPSocket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this probe based on the context it is used
func (m *Probe) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPGet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPSocket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Probe) contextValidateExec(ctx context.Context, formats strfmt.Registry) error {

	if m.Exec != nil {

		if swag.IsZero(m.Exec) { // not required
			return nil
		}

		if err := m.Exec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) contextValidateGrpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Grpc != nil {

		if swag.IsZero(m.Grpc) { // not required
			return nil
		}

		if err := m.Grpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grpc")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) contextValidateHTTPGet(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPGet != nil {

		if swag.IsZero(m.HTTPGet) { // not required
			return nil
		}

		if err := m.HTTPGet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *Probe) contextValidateTCPSocket(ctx context.Context, formats strfmt.Registry) error {

	if m.TCPSocket != nil {

		if swag.IsZero(m.TCPSocket) { // not required
			return nil
		}

		if err := m.TCPSocket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Probe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Probe) UnmarshalBinary(b []byte) error {
	var res Probe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
