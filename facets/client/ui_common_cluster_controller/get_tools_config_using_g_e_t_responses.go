// Code generated by go-swagger; DO NOT EDIT.

package ui_common_cluster_controller

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

// GetToolsConfigUsingGETReader is a Reader for the GetToolsConfigUsingGET structure.
type GetToolsConfigUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetToolsConfigUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetToolsConfigUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetToolsConfigUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetToolsConfigUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetToolsConfigUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/tools-config] getToolsConfigUsingGET", response, response.Code())
	}
}

// NewGetToolsConfigUsingGETOK creates a GetToolsConfigUsingGETOK with default headers values
func NewGetToolsConfigUsingGETOK() *GetToolsConfigUsingGETOK {
	return &GetToolsConfigUsingGETOK{}
}

/*
GetToolsConfigUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetToolsConfigUsingGETOK struct {
	Payload *models.ToolsConfig
}

// IsSuccess returns true when this get tools config using g e t o k response has a 2xx status code
func (o *GetToolsConfigUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tools config using g e t o k response has a 3xx status code
func (o *GetToolsConfigUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tools config using g e t o k response has a 4xx status code
func (o *GetToolsConfigUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tools config using g e t o k response has a 5xx status code
func (o *GetToolsConfigUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tools config using g e t o k response a status code equal to that given
func (o *GetToolsConfigUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tools config using g e t o k response
func (o *GetToolsConfigUsingGETOK) Code() int {
	return 200
}

func (o *GetToolsConfigUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETOK %s", 200, payload)
}

func (o *GetToolsConfigUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETOK %s", 200, payload)
}

func (o *GetToolsConfigUsingGETOK) GetPayload() *models.ToolsConfig {
	return o.Payload
}

func (o *GetToolsConfigUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ToolsConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetToolsConfigUsingGETUnauthorized creates a GetToolsConfigUsingGETUnauthorized with default headers values
func NewGetToolsConfigUsingGETUnauthorized() *GetToolsConfigUsingGETUnauthorized {
	return &GetToolsConfigUsingGETUnauthorized{}
}

/*
GetToolsConfigUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetToolsConfigUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get tools config using g e t unauthorized response has a 2xx status code
func (o *GetToolsConfigUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tools config using g e t unauthorized response has a 3xx status code
func (o *GetToolsConfigUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tools config using g e t unauthorized response has a 4xx status code
func (o *GetToolsConfigUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tools config using g e t unauthorized response has a 5xx status code
func (o *GetToolsConfigUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tools config using g e t unauthorized response a status code equal to that given
func (o *GetToolsConfigUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tools config using g e t unauthorized response
func (o *GetToolsConfigUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetToolsConfigUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETUnauthorized", 401)
}

func (o *GetToolsConfigUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETUnauthorized", 401)
}

func (o *GetToolsConfigUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetToolsConfigUsingGETForbidden creates a GetToolsConfigUsingGETForbidden with default headers values
func NewGetToolsConfigUsingGETForbidden() *GetToolsConfigUsingGETForbidden {
	return &GetToolsConfigUsingGETForbidden{}
}

/*
GetToolsConfigUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetToolsConfigUsingGETForbidden struct {
}

// IsSuccess returns true when this get tools config using g e t forbidden response has a 2xx status code
func (o *GetToolsConfigUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tools config using g e t forbidden response has a 3xx status code
func (o *GetToolsConfigUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tools config using g e t forbidden response has a 4xx status code
func (o *GetToolsConfigUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tools config using g e t forbidden response has a 5xx status code
func (o *GetToolsConfigUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get tools config using g e t forbidden response a status code equal to that given
func (o *GetToolsConfigUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get tools config using g e t forbidden response
func (o *GetToolsConfigUsingGETForbidden) Code() int {
	return 403
}

func (o *GetToolsConfigUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETForbidden", 403)
}

func (o *GetToolsConfigUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETForbidden", 403)
}

func (o *GetToolsConfigUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetToolsConfigUsingGETNotFound creates a GetToolsConfigUsingGETNotFound with default headers values
func NewGetToolsConfigUsingGETNotFound() *GetToolsConfigUsingGETNotFound {
	return &GetToolsConfigUsingGETNotFound{}
}

/*
GetToolsConfigUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetToolsConfigUsingGETNotFound struct {
}

// IsSuccess returns true when this get tools config using g e t not found response has a 2xx status code
func (o *GetToolsConfigUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tools config using g e t not found response has a 3xx status code
func (o *GetToolsConfigUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tools config using g e t not found response has a 4xx status code
func (o *GetToolsConfigUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tools config using g e t not found response has a 5xx status code
func (o *GetToolsConfigUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tools config using g e t not found response a status code equal to that given
func (o *GetToolsConfigUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tools config using g e t not found response
func (o *GetToolsConfigUsingGETNotFound) Code() int {
	return 404
}

func (o *GetToolsConfigUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETNotFound", 404)
}

func (o *GetToolsConfigUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/tools-config][%d] getToolsConfigUsingGETNotFound", 404)
}

func (o *GetToolsConfigUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
