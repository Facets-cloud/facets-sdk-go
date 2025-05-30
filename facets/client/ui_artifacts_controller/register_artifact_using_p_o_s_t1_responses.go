// Code generated by go-swagger; DO NOT EDIT.

package ui_artifacts_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// RegisterArtifactUsingPOST1Reader is a Reader for the RegisterArtifactUsingPOST1 structure.
type RegisterArtifactUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterArtifactUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterArtifactUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRegisterArtifactUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegisterArtifactUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegisterArtifactUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegisterArtifactUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/artifacts/register] registerArtifactUsingPOST_1", response, response.Code())
	}
}

// NewRegisterArtifactUsingPOST1OK creates a RegisterArtifactUsingPOST1OK with default headers values
func NewRegisterArtifactUsingPOST1OK() *RegisterArtifactUsingPOST1OK {
	return &RegisterArtifactUsingPOST1OK{}
}

/*
RegisterArtifactUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type RegisterArtifactUsingPOST1OK struct {
	Payload []*models.Artifact
}

// IsSuccess returns true when this register artifact using p o s t1 o k response has a 2xx status code
func (o *RegisterArtifactUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register artifact using p o s t1 o k response has a 3xx status code
func (o *RegisterArtifactUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register artifact using p o s t1 o k response has a 4xx status code
func (o *RegisterArtifactUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this register artifact using p o s t1 o k response has a 5xx status code
func (o *RegisterArtifactUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this register artifact using p o s t1 o k response a status code equal to that given
func (o *RegisterArtifactUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the register artifact using p o s t1 o k response
func (o *RegisterArtifactUsingPOST1OK) Code() int {
	return 200
}

func (o *RegisterArtifactUsingPOST1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1OK %s", 200, payload)
}

func (o *RegisterArtifactUsingPOST1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1OK %s", 200, payload)
}

func (o *RegisterArtifactUsingPOST1OK) GetPayload() []*models.Artifact {
	return o.Payload
}

func (o *RegisterArtifactUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterArtifactUsingPOST1Created creates a RegisterArtifactUsingPOST1Created with default headers values
func NewRegisterArtifactUsingPOST1Created() *RegisterArtifactUsingPOST1Created {
	return &RegisterArtifactUsingPOST1Created{}
}

/*
RegisterArtifactUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type RegisterArtifactUsingPOST1Created struct {
}

// IsSuccess returns true when this register artifact using p o s t1 created response has a 2xx status code
func (o *RegisterArtifactUsingPOST1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register artifact using p o s t1 created response has a 3xx status code
func (o *RegisterArtifactUsingPOST1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register artifact using p o s t1 created response has a 4xx status code
func (o *RegisterArtifactUsingPOST1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this register artifact using p o s t1 created response has a 5xx status code
func (o *RegisterArtifactUsingPOST1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this register artifact using p o s t1 created response a status code equal to that given
func (o *RegisterArtifactUsingPOST1Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the register artifact using p o s t1 created response
func (o *RegisterArtifactUsingPOST1Created) Code() int {
	return 201
}

func (o *RegisterArtifactUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Created", 201)
}

func (o *RegisterArtifactUsingPOST1Created) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Created", 201)
}

func (o *RegisterArtifactUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterArtifactUsingPOST1Unauthorized creates a RegisterArtifactUsingPOST1Unauthorized with default headers values
func NewRegisterArtifactUsingPOST1Unauthorized() *RegisterArtifactUsingPOST1Unauthorized {
	return &RegisterArtifactUsingPOST1Unauthorized{}
}

/*
RegisterArtifactUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegisterArtifactUsingPOST1Unauthorized struct {
}

// IsSuccess returns true when this register artifact using p o s t1 unauthorized response has a 2xx status code
func (o *RegisterArtifactUsingPOST1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register artifact using p o s t1 unauthorized response has a 3xx status code
func (o *RegisterArtifactUsingPOST1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register artifact using p o s t1 unauthorized response has a 4xx status code
func (o *RegisterArtifactUsingPOST1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this register artifact using p o s t1 unauthorized response has a 5xx status code
func (o *RegisterArtifactUsingPOST1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this register artifact using p o s t1 unauthorized response a status code equal to that given
func (o *RegisterArtifactUsingPOST1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the register artifact using p o s t1 unauthorized response
func (o *RegisterArtifactUsingPOST1Unauthorized) Code() int {
	return 401
}

func (o *RegisterArtifactUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Unauthorized", 401)
}

func (o *RegisterArtifactUsingPOST1Unauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Unauthorized", 401)
}

func (o *RegisterArtifactUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterArtifactUsingPOST1Forbidden creates a RegisterArtifactUsingPOST1Forbidden with default headers values
func NewRegisterArtifactUsingPOST1Forbidden() *RegisterArtifactUsingPOST1Forbidden {
	return &RegisterArtifactUsingPOST1Forbidden{}
}

/*
RegisterArtifactUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RegisterArtifactUsingPOST1Forbidden struct {
}

// IsSuccess returns true when this register artifact using p o s t1 forbidden response has a 2xx status code
func (o *RegisterArtifactUsingPOST1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register artifact using p o s t1 forbidden response has a 3xx status code
func (o *RegisterArtifactUsingPOST1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register artifact using p o s t1 forbidden response has a 4xx status code
func (o *RegisterArtifactUsingPOST1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this register artifact using p o s t1 forbidden response has a 5xx status code
func (o *RegisterArtifactUsingPOST1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this register artifact using p o s t1 forbidden response a status code equal to that given
func (o *RegisterArtifactUsingPOST1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the register artifact using p o s t1 forbidden response
func (o *RegisterArtifactUsingPOST1Forbidden) Code() int {
	return 403
}

func (o *RegisterArtifactUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Forbidden", 403)
}

func (o *RegisterArtifactUsingPOST1Forbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1Forbidden", 403)
}

func (o *RegisterArtifactUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterArtifactUsingPOST1NotFound creates a RegisterArtifactUsingPOST1NotFound with default headers values
func NewRegisterArtifactUsingPOST1NotFound() *RegisterArtifactUsingPOST1NotFound {
	return &RegisterArtifactUsingPOST1NotFound{}
}

/*
RegisterArtifactUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type RegisterArtifactUsingPOST1NotFound struct {
}

// IsSuccess returns true when this register artifact using p o s t1 not found response has a 2xx status code
func (o *RegisterArtifactUsingPOST1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register artifact using p o s t1 not found response has a 3xx status code
func (o *RegisterArtifactUsingPOST1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register artifact using p o s t1 not found response has a 4xx status code
func (o *RegisterArtifactUsingPOST1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this register artifact using p o s t1 not found response has a 5xx status code
func (o *RegisterArtifactUsingPOST1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this register artifact using p o s t1 not found response a status code equal to that given
func (o *RegisterArtifactUsingPOST1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the register artifact using p o s t1 not found response
func (o *RegisterArtifactUsingPOST1NotFound) Code() int {
	return 404
}

func (o *RegisterArtifactUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1NotFound", 404)
}

func (o *RegisterArtifactUsingPOST1NotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifacts/register][%d] registerArtifactUsingPOST1NotFound", 404)
}

func (o *RegisterArtifactUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
