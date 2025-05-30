// Code generated by go-swagger; DO NOT EDIT.

package ui_artifact_routing_rule_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteArtifactRoutingRuleUsingDELETEReader is a Reader for the DeleteArtifactRoutingRuleUsingDELETE structure.
type DeleteArtifactRoutingRuleUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtifactRoutingRuleUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteArtifactRoutingRuleUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteArtifactRoutingRuleUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteArtifactRoutingRuleUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArtifactRoutingRuleUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}] deleteArtifactRoutingRuleUsingDELETE", response, response.Code())
	}
}

// NewDeleteArtifactRoutingRuleUsingDELETEOK creates a DeleteArtifactRoutingRuleUsingDELETEOK with default headers values
func NewDeleteArtifactRoutingRuleUsingDELETEOK() *DeleteArtifactRoutingRuleUsingDELETEOK {
	return &DeleteArtifactRoutingRuleUsingDELETEOK{}
}

/*
DeleteArtifactRoutingRuleUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteArtifactRoutingRuleUsingDELETEOK struct {
}

// IsSuccess returns true when this delete artifact routing rule using d e l e t e o k response has a 2xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete artifact routing rule using d e l e t e o k response has a 3xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact routing rule using d e l e t e o k response has a 4xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact routing rule using d e l e t e o k response has a 5xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact routing rule using d e l e t e o k response a status code equal to that given
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete artifact routing rule using d e l e t e o k response
func (o *DeleteArtifactRoutingRuleUsingDELETEOK) Code() int {
	return 200
}

func (o *DeleteArtifactRoutingRuleUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEOK", 200)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEOK", 200)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactRoutingRuleUsingDELETENoContent creates a DeleteArtifactRoutingRuleUsingDELETENoContent with default headers values
func NewDeleteArtifactRoutingRuleUsingDELETENoContent() *DeleteArtifactRoutingRuleUsingDELETENoContent {
	return &DeleteArtifactRoutingRuleUsingDELETENoContent{}
}

/*
DeleteArtifactRoutingRuleUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteArtifactRoutingRuleUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete artifact routing rule using d e l e t e no content response has a 2xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete artifact routing rule using d e l e t e no content response has a 3xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact routing rule using d e l e t e no content response has a 4xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact routing rule using d e l e t e no content response has a 5xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact routing rule using d e l e t e no content response a status code equal to that given
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete artifact routing rule using d e l e t e no content response
func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETENoContent", 204)
}

func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETENoContent", 204)
}

func (o *DeleteArtifactRoutingRuleUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactRoutingRuleUsingDELETEUnauthorized creates a DeleteArtifactRoutingRuleUsingDELETEUnauthorized with default headers values
func NewDeleteArtifactRoutingRuleUsingDELETEUnauthorized() *DeleteArtifactRoutingRuleUsingDELETEUnauthorized {
	return &DeleteArtifactRoutingRuleUsingDELETEUnauthorized{}
}

/*
DeleteArtifactRoutingRuleUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteArtifactRoutingRuleUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete artifact routing rule using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact routing rule using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact routing rule using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact routing rule using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact routing rule using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete artifact routing rule using d e l e t e unauthorized response
func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) Code() int {
	return 401
}

func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEUnauthorized", 401)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEUnauthorized", 401)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactRoutingRuleUsingDELETEForbidden creates a DeleteArtifactRoutingRuleUsingDELETEForbidden with default headers values
func NewDeleteArtifactRoutingRuleUsingDELETEForbidden() *DeleteArtifactRoutingRuleUsingDELETEForbidden {
	return &DeleteArtifactRoutingRuleUsingDELETEForbidden{}
}

/*
DeleteArtifactRoutingRuleUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteArtifactRoutingRuleUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete artifact routing rule using d e l e t e forbidden response has a 2xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact routing rule using d e l e t e forbidden response has a 3xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact routing rule using d e l e t e forbidden response has a 4xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact routing rule using d e l e t e forbidden response has a 5xx status code
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact routing rule using d e l e t e forbidden response a status code equal to that given
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete artifact routing rule using d e l e t e forbidden response
func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) Code() int {
	return 403
}

func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEForbidden", 403)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifact-routing-rule/{ruleId}][%d] deleteArtifactRoutingRuleUsingDELETEForbidden", 403)
}

func (o *DeleteArtifactRoutingRuleUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
