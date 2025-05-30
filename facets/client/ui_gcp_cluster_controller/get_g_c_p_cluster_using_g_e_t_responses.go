// Code generated by go-swagger; DO NOT EDIT.

package ui_gcp_cluster_controller

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

// GetGCPClusterUsingGETReader is a Reader for the GetGCPClusterUsingGET structure.
type GetGCPClusterUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCPClusterUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCPClusterUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCPClusterUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCPClusterUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGCPClusterUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/gcp/clusters/{clusterId}] getGCPClusterUsingGET", response, response.Code())
	}
}

// NewGetGCPClusterUsingGETOK creates a GetGCPClusterUsingGETOK with default headers values
func NewGetGCPClusterUsingGETOK() *GetGCPClusterUsingGETOK {
	return &GetGCPClusterUsingGETOK{}
}

/*
GetGCPClusterUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetGCPClusterUsingGETOK struct {
	Payload *models.GCPCluster
}

// IsSuccess returns true when this get g c p cluster using g e t o k response has a 2xx status code
func (o *GetGCPClusterUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get g c p cluster using g e t o k response has a 3xx status code
func (o *GetGCPClusterUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get g c p cluster using g e t o k response has a 4xx status code
func (o *GetGCPClusterUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get g c p cluster using g e t o k response has a 5xx status code
func (o *GetGCPClusterUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get g c p cluster using g e t o k response a status code equal to that given
func (o *GetGCPClusterUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get g c p cluster using g e t o k response
func (o *GetGCPClusterUsingGETOK) Code() int {
	return 200
}

func (o *GetGCPClusterUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETOK %s", 200, payload)
}

func (o *GetGCPClusterUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETOK %s", 200, payload)
}

func (o *GetGCPClusterUsingGETOK) GetPayload() *models.GCPCluster {
	return o.Payload
}

func (o *GetGCPClusterUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GCPCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCPClusterUsingGETUnauthorized creates a GetGCPClusterUsingGETUnauthorized with default headers values
func NewGetGCPClusterUsingGETUnauthorized() *GetGCPClusterUsingGETUnauthorized {
	return &GetGCPClusterUsingGETUnauthorized{}
}

/*
GetGCPClusterUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGCPClusterUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get g c p cluster using g e t unauthorized response has a 2xx status code
func (o *GetGCPClusterUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get g c p cluster using g e t unauthorized response has a 3xx status code
func (o *GetGCPClusterUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get g c p cluster using g e t unauthorized response has a 4xx status code
func (o *GetGCPClusterUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get g c p cluster using g e t unauthorized response has a 5xx status code
func (o *GetGCPClusterUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get g c p cluster using g e t unauthorized response a status code equal to that given
func (o *GetGCPClusterUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get g c p cluster using g e t unauthorized response
func (o *GetGCPClusterUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetGCPClusterUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETUnauthorized", 401)
}

func (o *GetGCPClusterUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETUnauthorized", 401)
}

func (o *GetGCPClusterUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGCPClusterUsingGETForbidden creates a GetGCPClusterUsingGETForbidden with default headers values
func NewGetGCPClusterUsingGETForbidden() *GetGCPClusterUsingGETForbidden {
	return &GetGCPClusterUsingGETForbidden{}
}

/*
GetGCPClusterUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGCPClusterUsingGETForbidden struct {
}

// IsSuccess returns true when this get g c p cluster using g e t forbidden response has a 2xx status code
func (o *GetGCPClusterUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get g c p cluster using g e t forbidden response has a 3xx status code
func (o *GetGCPClusterUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get g c p cluster using g e t forbidden response has a 4xx status code
func (o *GetGCPClusterUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get g c p cluster using g e t forbidden response has a 5xx status code
func (o *GetGCPClusterUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get g c p cluster using g e t forbidden response a status code equal to that given
func (o *GetGCPClusterUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get g c p cluster using g e t forbidden response
func (o *GetGCPClusterUsingGETForbidden) Code() int {
	return 403
}

func (o *GetGCPClusterUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETForbidden", 403)
}

func (o *GetGCPClusterUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETForbidden", 403)
}

func (o *GetGCPClusterUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGCPClusterUsingGETNotFound creates a GetGCPClusterUsingGETNotFound with default headers values
func NewGetGCPClusterUsingGETNotFound() *GetGCPClusterUsingGETNotFound {
	return &GetGCPClusterUsingGETNotFound{}
}

/*
GetGCPClusterUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGCPClusterUsingGETNotFound struct {
}

// IsSuccess returns true when this get g c p cluster using g e t not found response has a 2xx status code
func (o *GetGCPClusterUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get g c p cluster using g e t not found response has a 3xx status code
func (o *GetGCPClusterUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get g c p cluster using g e t not found response has a 4xx status code
func (o *GetGCPClusterUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get g c p cluster using g e t not found response has a 5xx status code
func (o *GetGCPClusterUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get g c p cluster using g e t not found response a status code equal to that given
func (o *GetGCPClusterUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get g c p cluster using g e t not found response
func (o *GetGCPClusterUsingGETNotFound) Code() int {
	return 404
}

func (o *GetGCPClusterUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETNotFound", 404)
}

func (o *GetGCPClusterUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/gcp/clusters/{clusterId}][%d] getGCPClusterUsingGETNotFound", 404)
}

func (o *GetGCPClusterUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
