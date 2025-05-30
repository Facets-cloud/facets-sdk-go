// Code generated by go-swagger; DO NOT EDIT.

package ui_deployment_controller

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

// GetReleaseChangesUsingGETReader is a Reader for the GetReleaseChangesUsingGET structure.
type GetReleaseChangesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseChangesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseChangesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReleaseChangesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReleaseChangesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleaseChangesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes] getReleaseChangesUsingGET", response, response.Code())
	}
}

// NewGetReleaseChangesUsingGETOK creates a GetReleaseChangesUsingGETOK with default headers values
func NewGetReleaseChangesUsingGETOK() *GetReleaseChangesUsingGETOK {
	return &GetReleaseChangesUsingGETOK{}
}

/*
GetReleaseChangesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetReleaseChangesUsingGETOK struct {
	Payload *models.ReleaseChanges
}

// IsSuccess returns true when this get release changes using g e t o k response has a 2xx status code
func (o *GetReleaseChangesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get release changes using g e t o k response has a 3xx status code
func (o *GetReleaseChangesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release changes using g e t o k response has a 4xx status code
func (o *GetReleaseChangesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get release changes using g e t o k response has a 5xx status code
func (o *GetReleaseChangesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get release changes using g e t o k response a status code equal to that given
func (o *GetReleaseChangesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get release changes using g e t o k response
func (o *GetReleaseChangesUsingGETOK) Code() int {
	return 200
}

func (o *GetReleaseChangesUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETOK %s", 200, payload)
}

func (o *GetReleaseChangesUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETOK %s", 200, payload)
}

func (o *GetReleaseChangesUsingGETOK) GetPayload() *models.ReleaseChanges {
	return o.Payload
}

func (o *GetReleaseChangesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseChanges)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseChangesUsingGETUnauthorized creates a GetReleaseChangesUsingGETUnauthorized with default headers values
func NewGetReleaseChangesUsingGETUnauthorized() *GetReleaseChangesUsingGETUnauthorized {
	return &GetReleaseChangesUsingGETUnauthorized{}
}

/*
GetReleaseChangesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReleaseChangesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get release changes using g e t unauthorized response has a 2xx status code
func (o *GetReleaseChangesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release changes using g e t unauthorized response has a 3xx status code
func (o *GetReleaseChangesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release changes using g e t unauthorized response has a 4xx status code
func (o *GetReleaseChangesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release changes using g e t unauthorized response has a 5xx status code
func (o *GetReleaseChangesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get release changes using g e t unauthorized response a status code equal to that given
func (o *GetReleaseChangesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get release changes using g e t unauthorized response
func (o *GetReleaseChangesUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetReleaseChangesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETUnauthorized", 401)
}

func (o *GetReleaseChangesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETUnauthorized", 401)
}

func (o *GetReleaseChangesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleaseChangesUsingGETForbidden creates a GetReleaseChangesUsingGETForbidden with default headers values
func NewGetReleaseChangesUsingGETForbidden() *GetReleaseChangesUsingGETForbidden {
	return &GetReleaseChangesUsingGETForbidden{}
}

/*
GetReleaseChangesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReleaseChangesUsingGETForbidden struct {
}

// IsSuccess returns true when this get release changes using g e t forbidden response has a 2xx status code
func (o *GetReleaseChangesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release changes using g e t forbidden response has a 3xx status code
func (o *GetReleaseChangesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release changes using g e t forbidden response has a 4xx status code
func (o *GetReleaseChangesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release changes using g e t forbidden response has a 5xx status code
func (o *GetReleaseChangesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get release changes using g e t forbidden response a status code equal to that given
func (o *GetReleaseChangesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get release changes using g e t forbidden response
func (o *GetReleaseChangesUsingGETForbidden) Code() int {
	return 403
}

func (o *GetReleaseChangesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETForbidden", 403)
}

func (o *GetReleaseChangesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETForbidden", 403)
}

func (o *GetReleaseChangesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleaseChangesUsingGETNotFound creates a GetReleaseChangesUsingGETNotFound with default headers values
func NewGetReleaseChangesUsingGETNotFound() *GetReleaseChangesUsingGETNotFound {
	return &GetReleaseChangesUsingGETNotFound{}
}

/*
GetReleaseChangesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReleaseChangesUsingGETNotFound struct {
}

// IsSuccess returns true when this get release changes using g e t not found response has a 2xx status code
func (o *GetReleaseChangesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release changes using g e t not found response has a 3xx status code
func (o *GetReleaseChangesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release changes using g e t not found response has a 4xx status code
func (o *GetReleaseChangesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release changes using g e t not found response has a 5xx status code
func (o *GetReleaseChangesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get release changes using g e t not found response a status code equal to that given
func (o *GetReleaseChangesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get release changes using g e t not found response
func (o *GetReleaseChangesUsingGETNotFound) Code() int {
	return 404
}

func (o *GetReleaseChangesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETNotFound", 404)
}

func (o *GetReleaseChangesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/release-changes][%d] getReleaseChangesUsingGETNotFound", 404)
}

func (o *GetReleaseChangesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
