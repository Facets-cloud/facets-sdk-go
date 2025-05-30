// Code generated by go-swagger; DO NOT EDIT.

package ui_artifact_hub_controller

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

// SearchPackagesUsingGETReader is a Reader for the SearchPackagesUsingGET structure.
type SearchPackagesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPackagesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPackagesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchPackagesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchPackagesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchPackagesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/artifactHub/search-packages] searchPackagesUsingGET", response, response.Code())
	}
}

// NewSearchPackagesUsingGETOK creates a SearchPackagesUsingGETOK with default headers values
func NewSearchPackagesUsingGETOK() *SearchPackagesUsingGETOK {
	return &SearchPackagesUsingGETOK{}
}

/*
SearchPackagesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type SearchPackagesUsingGETOK struct {
	Payload *models.PackageResponse
}

// IsSuccess returns true when this search packages using g e t o k response has a 2xx status code
func (o *SearchPackagesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search packages using g e t o k response has a 3xx status code
func (o *SearchPackagesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search packages using g e t o k response has a 4xx status code
func (o *SearchPackagesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search packages using g e t o k response has a 5xx status code
func (o *SearchPackagesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search packages using g e t o k response a status code equal to that given
func (o *SearchPackagesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search packages using g e t o k response
func (o *SearchPackagesUsingGETOK) Code() int {
	return 200
}

func (o *SearchPackagesUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETOK %s", 200, payload)
}

func (o *SearchPackagesUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETOK %s", 200, payload)
}

func (o *SearchPackagesUsingGETOK) GetPayload() *models.PackageResponse {
	return o.Payload
}

func (o *SearchPackagesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPackagesUsingGETUnauthorized creates a SearchPackagesUsingGETUnauthorized with default headers values
func NewSearchPackagesUsingGETUnauthorized() *SearchPackagesUsingGETUnauthorized {
	return &SearchPackagesUsingGETUnauthorized{}
}

/*
SearchPackagesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchPackagesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this search packages using g e t unauthorized response has a 2xx status code
func (o *SearchPackagesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search packages using g e t unauthorized response has a 3xx status code
func (o *SearchPackagesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search packages using g e t unauthorized response has a 4xx status code
func (o *SearchPackagesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search packages using g e t unauthorized response has a 5xx status code
func (o *SearchPackagesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search packages using g e t unauthorized response a status code equal to that given
func (o *SearchPackagesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the search packages using g e t unauthorized response
func (o *SearchPackagesUsingGETUnauthorized) Code() int {
	return 401
}

func (o *SearchPackagesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETUnauthorized", 401)
}

func (o *SearchPackagesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETUnauthorized", 401)
}

func (o *SearchPackagesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchPackagesUsingGETForbidden creates a SearchPackagesUsingGETForbidden with default headers values
func NewSearchPackagesUsingGETForbidden() *SearchPackagesUsingGETForbidden {
	return &SearchPackagesUsingGETForbidden{}
}

/*
SearchPackagesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchPackagesUsingGETForbidden struct {
}

// IsSuccess returns true when this search packages using g e t forbidden response has a 2xx status code
func (o *SearchPackagesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search packages using g e t forbidden response has a 3xx status code
func (o *SearchPackagesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search packages using g e t forbidden response has a 4xx status code
func (o *SearchPackagesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search packages using g e t forbidden response has a 5xx status code
func (o *SearchPackagesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search packages using g e t forbidden response a status code equal to that given
func (o *SearchPackagesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search packages using g e t forbidden response
func (o *SearchPackagesUsingGETForbidden) Code() int {
	return 403
}

func (o *SearchPackagesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETForbidden", 403)
}

func (o *SearchPackagesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETForbidden", 403)
}

func (o *SearchPackagesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchPackagesUsingGETNotFound creates a SearchPackagesUsingGETNotFound with default headers values
func NewSearchPackagesUsingGETNotFound() *SearchPackagesUsingGETNotFound {
	return &SearchPackagesUsingGETNotFound{}
}

/*
SearchPackagesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchPackagesUsingGETNotFound struct {
}

// IsSuccess returns true when this search packages using g e t not found response has a 2xx status code
func (o *SearchPackagesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search packages using g e t not found response has a 3xx status code
func (o *SearchPackagesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search packages using g e t not found response has a 4xx status code
func (o *SearchPackagesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search packages using g e t not found response has a 5xx status code
func (o *SearchPackagesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search packages using g e t not found response a status code equal to that given
func (o *SearchPackagesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search packages using g e t not found response
func (o *SearchPackagesUsingGETNotFound) Code() int {
	return 404
}

func (o *SearchPackagesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETNotFound", 404)
}

func (o *SearchPackagesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/artifactHub/search-packages][%d] searchPackagesUsingGETNotFound", 404)
}

func (o *SearchPackagesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
