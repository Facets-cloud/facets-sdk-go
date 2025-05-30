// Code generated by go-swagger; DO NOT EDIT.

package ui_application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ClusterSyncWithGitUsingPOSTReader is a Reader for the ClusterSyncWithGitUsingPOST structure.
type ClusterSyncWithGitUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterSyncWithGitUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterSyncWithGitUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewClusterSyncWithGitUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewClusterSyncWithGitUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewClusterSyncWithGitUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClusterSyncWithGitUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git] clusterSyncWithGitUsingPOST", response, response.Code())
	}
}

// NewClusterSyncWithGitUsingPOSTOK creates a ClusterSyncWithGitUsingPOSTOK with default headers values
func NewClusterSyncWithGitUsingPOSTOK() *ClusterSyncWithGitUsingPOSTOK {
	return &ClusterSyncWithGitUsingPOSTOK{}
}

/*
ClusterSyncWithGitUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type ClusterSyncWithGitUsingPOSTOK struct {
}

// IsSuccess returns true when this cluster sync with git using p o s t o k response has a 2xx status code
func (o *ClusterSyncWithGitUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster sync with git using p o s t o k response has a 3xx status code
func (o *ClusterSyncWithGitUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster sync with git using p o s t o k response has a 4xx status code
func (o *ClusterSyncWithGitUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster sync with git using p o s t o k response has a 5xx status code
func (o *ClusterSyncWithGitUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster sync with git using p o s t o k response a status code equal to that given
func (o *ClusterSyncWithGitUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster sync with git using p o s t o k response
func (o *ClusterSyncWithGitUsingPOSTOK) Code() int {
	return 200
}

func (o *ClusterSyncWithGitUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTOK", 200)
}

func (o *ClusterSyncWithGitUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTOK", 200)
}

func (o *ClusterSyncWithGitUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterSyncWithGitUsingPOSTCreated creates a ClusterSyncWithGitUsingPOSTCreated with default headers values
func NewClusterSyncWithGitUsingPOSTCreated() *ClusterSyncWithGitUsingPOSTCreated {
	return &ClusterSyncWithGitUsingPOSTCreated{}
}

/*
ClusterSyncWithGitUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type ClusterSyncWithGitUsingPOSTCreated struct {
}

// IsSuccess returns true when this cluster sync with git using p o s t created response has a 2xx status code
func (o *ClusterSyncWithGitUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster sync with git using p o s t created response has a 3xx status code
func (o *ClusterSyncWithGitUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster sync with git using p o s t created response has a 4xx status code
func (o *ClusterSyncWithGitUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster sync with git using p o s t created response has a 5xx status code
func (o *ClusterSyncWithGitUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster sync with git using p o s t created response a status code equal to that given
func (o *ClusterSyncWithGitUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the cluster sync with git using p o s t created response
func (o *ClusterSyncWithGitUsingPOSTCreated) Code() int {
	return 201
}

func (o *ClusterSyncWithGitUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTCreated", 201)
}

func (o *ClusterSyncWithGitUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTCreated", 201)
}

func (o *ClusterSyncWithGitUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterSyncWithGitUsingPOSTUnauthorized creates a ClusterSyncWithGitUsingPOSTUnauthorized with default headers values
func NewClusterSyncWithGitUsingPOSTUnauthorized() *ClusterSyncWithGitUsingPOSTUnauthorized {
	return &ClusterSyncWithGitUsingPOSTUnauthorized{}
}

/*
ClusterSyncWithGitUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ClusterSyncWithGitUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this cluster sync with git using p o s t unauthorized response has a 2xx status code
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cluster sync with git using p o s t unauthorized response has a 3xx status code
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster sync with git using p o s t unauthorized response has a 4xx status code
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cluster sync with git using p o s t unauthorized response has a 5xx status code
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster sync with git using p o s t unauthorized response a status code equal to that given
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cluster sync with git using p o s t unauthorized response
func (o *ClusterSyncWithGitUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *ClusterSyncWithGitUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTUnauthorized", 401)
}

func (o *ClusterSyncWithGitUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTUnauthorized", 401)
}

func (o *ClusterSyncWithGitUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterSyncWithGitUsingPOSTForbidden creates a ClusterSyncWithGitUsingPOSTForbidden with default headers values
func NewClusterSyncWithGitUsingPOSTForbidden() *ClusterSyncWithGitUsingPOSTForbidden {
	return &ClusterSyncWithGitUsingPOSTForbidden{}
}

/*
ClusterSyncWithGitUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ClusterSyncWithGitUsingPOSTForbidden struct {
}

// IsSuccess returns true when this cluster sync with git using p o s t forbidden response has a 2xx status code
func (o *ClusterSyncWithGitUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cluster sync with git using p o s t forbidden response has a 3xx status code
func (o *ClusterSyncWithGitUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster sync with git using p o s t forbidden response has a 4xx status code
func (o *ClusterSyncWithGitUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cluster sync with git using p o s t forbidden response has a 5xx status code
func (o *ClusterSyncWithGitUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster sync with git using p o s t forbidden response a status code equal to that given
func (o *ClusterSyncWithGitUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cluster sync with git using p o s t forbidden response
func (o *ClusterSyncWithGitUsingPOSTForbidden) Code() int {
	return 403
}

func (o *ClusterSyncWithGitUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTForbidden", 403)
}

func (o *ClusterSyncWithGitUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTForbidden", 403)
}

func (o *ClusterSyncWithGitUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterSyncWithGitUsingPOSTNotFound creates a ClusterSyncWithGitUsingPOSTNotFound with default headers values
func NewClusterSyncWithGitUsingPOSTNotFound() *ClusterSyncWithGitUsingPOSTNotFound {
	return &ClusterSyncWithGitUsingPOSTNotFound{}
}

/*
ClusterSyncWithGitUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ClusterSyncWithGitUsingPOSTNotFound struct {
}

// IsSuccess returns true when this cluster sync with git using p o s t not found response has a 2xx status code
func (o *ClusterSyncWithGitUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cluster sync with git using p o s t not found response has a 3xx status code
func (o *ClusterSyncWithGitUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster sync with git using p o s t not found response has a 4xx status code
func (o *ClusterSyncWithGitUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cluster sync with git using p o s t not found response has a 5xx status code
func (o *ClusterSyncWithGitUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster sync with git using p o s t not found response a status code equal to that given
func (o *ClusterSyncWithGitUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cluster sync with git using p o s t not found response
func (o *ClusterSyncWithGitUsingPOSTNotFound) Code() int {
	return 404
}

func (o *ClusterSyncWithGitUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTNotFound", 404)
}

func (o *ClusterSyncWithGitUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/sync-with-git][%d] clusterSyncWithGitUsingPOSTNotFound", 404)
}

func (o *ClusterSyncWithGitUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
