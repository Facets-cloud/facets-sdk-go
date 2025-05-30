// Code generated by go-swagger; DO NOT EDIT.

package ui_artifactory_controller

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

// GetAllArtifactoriesForStackUsingGETReader is a Reader for the GetAllArtifactoriesForStackUsingGET structure.
type GetAllArtifactoriesForStackUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllArtifactoriesForStackUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllArtifactoriesForStackUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllArtifactoriesForStackUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllArtifactoriesForStackUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllArtifactoriesForStackUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/artifactories/stack/{stackName}] getAllArtifactoriesForStackUsingGET", response, response.Code())
	}
}

// NewGetAllArtifactoriesForStackUsingGETOK creates a GetAllArtifactoriesForStackUsingGETOK with default headers values
func NewGetAllArtifactoriesForStackUsingGETOK() *GetAllArtifactoriesForStackUsingGETOK {
	return &GetAllArtifactoriesForStackUsingGETOK{}
}

/*
GetAllArtifactoriesForStackUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllArtifactoriesForStackUsingGETOK struct {
	Payload []*models.Artifactory
}

// IsSuccess returns true when this get all artifactories for stack using g e t o k response has a 2xx status code
func (o *GetAllArtifactoriesForStackUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all artifactories for stack using g e t o k response has a 3xx status code
func (o *GetAllArtifactoriesForStackUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all artifactories for stack using g e t o k response has a 4xx status code
func (o *GetAllArtifactoriesForStackUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all artifactories for stack using g e t o k response has a 5xx status code
func (o *GetAllArtifactoriesForStackUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all artifactories for stack using g e t o k response a status code equal to that given
func (o *GetAllArtifactoriesForStackUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all artifactories for stack using g e t o k response
func (o *GetAllArtifactoriesForStackUsingGETOK) Code() int {
	return 200
}

func (o *GetAllArtifactoriesForStackUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETOK %s", 200, payload)
}

func (o *GetAllArtifactoriesForStackUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETOK %s", 200, payload)
}

func (o *GetAllArtifactoriesForStackUsingGETOK) GetPayload() []*models.Artifactory {
	return o.Payload
}

func (o *GetAllArtifactoriesForStackUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllArtifactoriesForStackUsingGETUnauthorized creates a GetAllArtifactoriesForStackUsingGETUnauthorized with default headers values
func NewGetAllArtifactoriesForStackUsingGETUnauthorized() *GetAllArtifactoriesForStackUsingGETUnauthorized {
	return &GetAllArtifactoriesForStackUsingGETUnauthorized{}
}

/*
GetAllArtifactoriesForStackUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllArtifactoriesForStackUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all artifactories for stack using g e t unauthorized response has a 2xx status code
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all artifactories for stack using g e t unauthorized response has a 3xx status code
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all artifactories for stack using g e t unauthorized response has a 4xx status code
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all artifactories for stack using g e t unauthorized response has a 5xx status code
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all artifactories for stack using g e t unauthorized response a status code equal to that given
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all artifactories for stack using g e t unauthorized response
func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETUnauthorized", 401)
}

func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETUnauthorized", 401)
}

func (o *GetAllArtifactoriesForStackUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllArtifactoriesForStackUsingGETForbidden creates a GetAllArtifactoriesForStackUsingGETForbidden with default headers values
func NewGetAllArtifactoriesForStackUsingGETForbidden() *GetAllArtifactoriesForStackUsingGETForbidden {
	return &GetAllArtifactoriesForStackUsingGETForbidden{}
}

/*
GetAllArtifactoriesForStackUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllArtifactoriesForStackUsingGETForbidden struct {
}

// IsSuccess returns true when this get all artifactories for stack using g e t forbidden response has a 2xx status code
func (o *GetAllArtifactoriesForStackUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all artifactories for stack using g e t forbidden response has a 3xx status code
func (o *GetAllArtifactoriesForStackUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all artifactories for stack using g e t forbidden response has a 4xx status code
func (o *GetAllArtifactoriesForStackUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all artifactories for stack using g e t forbidden response has a 5xx status code
func (o *GetAllArtifactoriesForStackUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all artifactories for stack using g e t forbidden response a status code equal to that given
func (o *GetAllArtifactoriesForStackUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all artifactories for stack using g e t forbidden response
func (o *GetAllArtifactoriesForStackUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAllArtifactoriesForStackUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETForbidden", 403)
}

func (o *GetAllArtifactoriesForStackUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETForbidden", 403)
}

func (o *GetAllArtifactoriesForStackUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllArtifactoriesForStackUsingGETNotFound creates a GetAllArtifactoriesForStackUsingGETNotFound with default headers values
func NewGetAllArtifactoriesForStackUsingGETNotFound() *GetAllArtifactoriesForStackUsingGETNotFound {
	return &GetAllArtifactoriesForStackUsingGETNotFound{}
}

/*
GetAllArtifactoriesForStackUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllArtifactoriesForStackUsingGETNotFound struct {
}

// IsSuccess returns true when this get all artifactories for stack using g e t not found response has a 2xx status code
func (o *GetAllArtifactoriesForStackUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all artifactories for stack using g e t not found response has a 3xx status code
func (o *GetAllArtifactoriesForStackUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all artifactories for stack using g e t not found response has a 4xx status code
func (o *GetAllArtifactoriesForStackUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all artifactories for stack using g e t not found response has a 5xx status code
func (o *GetAllArtifactoriesForStackUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all artifactories for stack using g e t not found response a status code equal to that given
func (o *GetAllArtifactoriesForStackUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all artifactories for stack using g e t not found response
func (o *GetAllArtifactoriesForStackUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAllArtifactoriesForStackUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETNotFound", 404)
}

func (o *GetAllArtifactoriesForStackUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactories/stack/{stackName}][%d] getAllArtifactoriesForStackUsingGETNotFound", 404)
}

func (o *GetAllArtifactoriesForStackUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
