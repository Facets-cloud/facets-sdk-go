// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_module_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ui tf module controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ui tf module controller API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new ui tf module controller API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for ui tf module controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptStarStar sets the Accept header to "*/*".
func WithAcceptStarStar(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"*/*"}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	BootstrapModulesUsingPOST(params *BootstrapModulesUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapModulesUsingPOSTOK, *BootstrapModulesUsingPOSTCreated, error)

	DeleteTfModuleUsingDELETE(params *DeleteTfModuleUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTfModuleUsingDELETEOK, *DeleteTfModuleUsingDELETENoContent, error)

	DownloadModuleByIDUsingGET(params *DownloadModuleByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadModuleByIDUsingGETOK, error)

	DownloadModuleByVersionIDUsingGET(params *DownloadModuleByVersionIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadModuleByVersionIDUsingGETOK, error)

	GetAllModulesLiteUsingGET(params *GetAllModulesLiteUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllModulesLiteUsingGETOK, error)

	GetAllModulesUsingGET(params *GetAllModulesUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllModulesUsingGETOK, error)

	GetAllUsingGET4(params *GetAllUsingGET4Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllUsingGET4OK, error)

	GetByIDUsingGET(params *GetByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDUsingGETOK, error)

	GetFacetsYamlByStackUsingGET(params *GetFacetsYamlByStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFacetsYamlByStackUsingGETOK, error)

	GetGroupedModulesForStackUsingGET(params *GetGroupedModulesForStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupedModulesForStackUsingGETOK, error)

	GetIntentAddOnModulesUsingGET(params *GetIntentAddOnModulesUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIntentAddOnModulesUsingGETOK, error)

	GetModuleForIFVAndStackUsingGET(params *GetModuleForIFVAndStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetModuleForIFVAndStackUsingGETOK, error)

	MarkAsPublishedByIDUsingPOST(params *MarkAsPublishedByIDUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkAsPublishedByIDUsingPOSTOK, *MarkAsPublishedByIDUsingPOSTCreated, error)

	MarkAsPublishedUsingPOST(params *MarkAsPublishedUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkAsPublishedUsingPOSTOK, *MarkAsPublishedUsingPOSTCreated, error)

	UploadModuleUsingPOST(params *UploadModuleUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadModuleUsingPOSTOK, *UploadModuleUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	BootstrapModulesUsingPOST bootstraps modules

	- **Description:** Bootstraps modules for initialization.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) BootstrapModulesUsingPOST(params *BootstrapModulesUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapModulesUsingPOSTOK, *BootstrapModulesUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapModulesUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "bootstrapModulesUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/modules/bootstrap",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BootstrapModulesUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *BootstrapModulesUsingPOSTOK:
		return value, nil, nil
	case *BootstrapModulesUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_tf_module_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DeleteTfModuleUsingDELETE deletes a module

	- **Description:** Deletes a module by ID.

- **Permissions:** Requires MODULE_DELETE permission.
- **Audit Logging:** Yes
*/
func (a *Client) DeleteTfModuleUsingDELETE(params *DeleteTfModuleUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTfModuleUsingDELETEOK, *DeleteTfModuleUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTfModuleUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTfModuleUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/cc-ui/v1/modules/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTfModuleUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteTfModuleUsingDELETEOK:
		return value, nil, nil
	case *DeleteTfModuleUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_tf_module_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadModuleByIDUsingGET downloads module by Id
*/
func (a *Client) DownloadModuleByIDUsingGET(params *DownloadModuleByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadModuleByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadModuleByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadModuleByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/{id}/download",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadModuleByIDUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadModuleByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadModuleByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadModuleByVersionIDUsingGET downloads module by version Id
*/
func (a *Client) DownloadModuleByVersionIDUsingGET(params *DownloadModuleByVersionIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadModuleByVersionIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadModuleByVersionIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadModuleByVersionIdUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/version/{versionId}/download",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadModuleByVersionIDUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadModuleByVersionIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadModuleByVersionIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllModulesLiteUsingGET gets all modules lite
*/
func (a *Client) GetAllModulesLiteUsingGET(params *GetAllModulesLiteUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllModulesLiteUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllModulesLiteUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllModulesLiteUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/modules-lite",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllModulesLiteUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllModulesLiteUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllModulesLiteUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllModulesUsingGET gets all modules
*/
func (a *Client) GetAllModulesUsingGET(params *GetAllModulesUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllModulesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllModulesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllModulesUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllModulesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllModulesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllModulesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetAllUsingGET4 gets all modules

	- **Description:** Retrieves all modules with optional filtering.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) GetAllUsingGET4(params *GetAllUsingGET4Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllUsingGET_4",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/all",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllUsingGET4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetByIDUsingGET gets module by ID

	- **Description:** Retrieves a module by its ID.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) GetByIDUsingGET(params *GetByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetByIDUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetFacetsYamlByStackUsingGET gets facets yaml by stack

	- **Description:** Retrieves facets.yaml by stack information.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) GetFacetsYamlByStackUsingGET(params *GetFacetsYamlByStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFacetsYamlByStackUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFacetsYamlByStackUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFacetsYamlByStackUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/stack/{stackName}/{intent}/{flavor}/{version}/fields",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFacetsYamlByStackUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFacetsYamlByStackUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFacetsYamlByStackUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetGroupedModulesForStackUsingGET gets grouped modules for stack

	- **Description:** Retrieves grouped modules specifically for a stack name.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) GetGroupedModulesForStackUsingGET(params *GetGroupedModulesForStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupedModulesForStackUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupedModulesForStackUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupedModulesForStackUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/stack/{stackName}/grouped",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupedModulesForStackUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupedModulesForStackUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupedModulesForStackUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetIntentAddOnModulesUsingGET gets add on modules

	- **Description:** Retrieves all add-on modules based on intent and flavor, optionally by cloud.

- **Audit Logging:** No specific audit logging at the moment.
*/
func (a *Client) GetIntentAddOnModulesUsingGET(params *GetIntentAddOnModulesUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIntentAddOnModulesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntentAddOnModulesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIntentAddOnModulesUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/intent/{intent}/flavor/{flavor}/add-ons",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIntentAddOnModulesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntentAddOnModulesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIntentAddOnModulesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetModuleForIFVAndStackUsingGET gets module for i f v and stack
*/
func (a *Client) GetModuleForIFVAndStackUsingGET(params *GetModuleForIFVAndStackUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetModuleForIFVAndStackUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetModuleForIFVAndStackUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getModuleForIFVAndStackUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/modules/stack/{stackName}/{intent}/{flavor}/{version}/module",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetModuleForIFVAndStackUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetModuleForIFVAndStackUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getModuleForIFVAndStackUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	MarkAsPublishedByIDUsingPOST marks module as published

	- **Description:** Marks a specific module version as published.

- **Permissions:** Requires MODULE_WRITE permission.
- **Audit Logging:** Yes
*/
func (a *Client) MarkAsPublishedByIDUsingPOST(params *MarkAsPublishedByIDUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkAsPublishedByIDUsingPOSTOK, *MarkAsPublishedByIDUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMarkAsPublishedByIDUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "markAsPublishedByIdUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/modules/{id}/mark-published",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MarkAsPublishedByIDUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *MarkAsPublishedByIDUsingPOSTOK:
		return value, nil, nil
	case *MarkAsPublishedByIDUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_tf_module_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	MarkAsPublishedUsingPOST marks module as published

	- **Description:** Marks a specific module version as published.

- **Permissions:** Requires MODULE_WRITE permission.
- **Audit Logging:** Yes
*/
func (a *Client) MarkAsPublishedUsingPOST(params *MarkAsPublishedUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MarkAsPublishedUsingPOSTOK, *MarkAsPublishedUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMarkAsPublishedUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "markAsPublishedUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/modules/intent/{intent}/flavor/{flavor}/version/{version}/mark-published",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MarkAsPublishedUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *MarkAsPublishedUsingPOSTOK:
		return value, nil, nil
	case *MarkAsPublishedUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_tf_module_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UploadModuleUsingPOST uploads a module

	- **Description:** Uploads a module using a file.

- **Permissions:** Requires MODULE_WRITE permission.
- **Parameters:**
  - `file`: The module file to upload
  - `metadata` (optional): Additional module metadata including:
  - `gitUrl`: Web URL of the git repository (Expected to embed the commit id)
  - `gitRef`: Git reference (branch, tag, or commit)
  - `isFeatureBranch`: If this is true, this preview module cannot be directly marked as published until we register the module again with this as false.

- **Audit Logging:** Yes
*/
func (a *Client) UploadModuleUsingPOST(params *UploadModuleUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadModuleUsingPOSTOK, *UploadModuleUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadModuleUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadModuleUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/modules/upload",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadModuleUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UploadModuleUsingPOSTOK:
		return value, nil, nil
	case *UploadModuleUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_tf_module_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
