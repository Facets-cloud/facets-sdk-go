// Code generated by go-swagger; DO NOT EDIT.

package ui_cloud_cost_explorer_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ui cloud cost explorer controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ui cloud cost explorer controller API client with basic auth credentials.
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

// New creates a new ui cloud cost explorer controller API client with a bearer token for authentication.
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
Client for ui cloud cost explorer controller API
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
	GetDailyCloudCostUsingGET(params *GetDailyCloudCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDailyCloudCostUsingGETOK, error)

	GetServiceWiseCostUsingGET(params *GetServiceWiseCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceWiseCostUsingGETOK, error)

	IsAwsCostExplorerEnabledUsingGET(params *IsAwsCostExplorerEnabledUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IsAwsCostExplorerEnabledUsingGETOK, error)

	SyncCloudCostUsingGET(params *SyncCloudCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SyncCloudCostUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDailyCloudCostUsingGET gets daily cloud cost
*/
func (a *Client) GetDailyCloudCostUsingGET(params *GetDailyCloudCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDailyCloudCostUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDailyCloudCostUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDailyCloudCostUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/cost-explorer/stack/{stackName}/daily-cost",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDailyCloudCostUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetDailyCloudCostUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDailyCloudCostUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServiceWiseCostUsingGET gets service wise cost
*/
func (a *Client) GetServiceWiseCostUsingGET(params *GetServiceWiseCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceWiseCostUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceWiseCostUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServiceWiseCostUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/cost-explorer/service-wise-cost/{clusterId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceWiseCostUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetServiceWiseCostUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceWiseCostUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IsAwsCostExplorerEnabledUsingGET is aws cost explorer enabled
*/
func (a *Client) IsAwsCostExplorerEnabledUsingGET(params *IsAwsCostExplorerEnabledUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IsAwsCostExplorerEnabledUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsAwsCostExplorerEnabledUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "isAwsCostExplorerEnabledUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/cost-explorer/aws/enabled",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IsAwsCostExplorerEnabledUsingGETReader{formats: a.formats},
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
	success, ok := result.(*IsAwsCostExplorerEnabledUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for isAwsCostExplorerEnabledUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SyncCloudCostUsingGET syncs cloud cost
*/
func (a *Client) SyncCloudCostUsingGET(params *SyncCloudCostUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SyncCloudCostUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncCloudCostUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "syncCloudCostUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/cost-explorer/sync-cost",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SyncCloudCostUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SyncCloudCostUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for syncCloudCostUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
