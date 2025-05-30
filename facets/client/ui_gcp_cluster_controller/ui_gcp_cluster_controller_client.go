// Code generated by go-swagger; DO NOT EDIT.

package ui_gcp_cluster_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ui gcp cluster controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ui gcp cluster controller API client with basic auth credentials.
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

// New creates a new ui gcp cluster controller API client with a bearer token for authentication.
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
Client for ui gcp cluster controller API
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
	ConfigureDraftClusterUsingPOST2(params *ConfigureDraftClusterUsingPOST2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConfigureDraftClusterUsingPOST2OK, *ConfigureDraftClusterUsingPOST2Created, error)

	CreateGCPClusterUsingPOST(params *CreateGCPClusterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGCPClusterUsingPOSTOK, *CreateGCPClusterUsingPOSTCreated, error)

	GetGCPClusterUsingGET(params *GetGCPClusterUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGCPClusterUsingGETOK, error)

	UpdateGCPClusterUsingPUT(params *UpdateGCPClusterUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGCPClusterUsingPUTOK, *UpdateGCPClusterUsingPUTCreated, error)

	ValidateVpcIDUsingGET1(params *ValidateVpcIDUsingGET1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateVpcIDUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConfigureDraftClusterUsingPOST2 configures draft cluster
*/
func (a *Client) ConfigureDraftClusterUsingPOST2(params *ConfigureDraftClusterUsingPOST2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConfigureDraftClusterUsingPOST2OK, *ConfigureDraftClusterUsingPOST2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigureDraftClusterUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "configureDraftClusterUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/gcp/clusters/configure/{clusterId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConfigureDraftClusterUsingPOST2Reader{formats: a.formats},
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
	case *ConfigureDraftClusterUsingPOST2OK:
		return value, nil, nil
	case *ConfigureDraftClusterUsingPOST2Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_gcp_cluster_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateGCPClusterUsingPOST creates g c p cluster
*/
func (a *Client) CreateGCPClusterUsingPOST(params *CreateGCPClusterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGCPClusterUsingPOSTOK, *CreateGCPClusterUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGCPClusterUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createGCPClusterUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/gcp/clusters",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGCPClusterUsingPOSTReader{formats: a.formats},
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
	case *CreateGCPClusterUsingPOSTOK:
		return value, nil, nil
	case *CreateGCPClusterUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_gcp_cluster_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGCPClusterUsingGET gets g c p cluster
*/
func (a *Client) GetGCPClusterUsingGET(params *GetGCPClusterUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGCPClusterUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGCPClusterUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGCPClusterUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/gcp/clusters/{clusterId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGCPClusterUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetGCPClusterUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGCPClusterUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateGCPClusterUsingPUT updates g c p cluster
*/
func (a *Client) UpdateGCPClusterUsingPUT(params *UpdateGCPClusterUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGCPClusterUsingPUTOK, *UpdateGCPClusterUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGCPClusterUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateGCPClusterUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cc-ui/v1/gcp/clusters/{clusterId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGCPClusterUsingPUTReader{formats: a.formats},
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
	case *UpdateGCPClusterUsingPUTOK:
		return value, nil, nil
	case *UpdateGCPClusterUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_gcp_cluster_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateVpcIDUsingGET1 validates vpc Id
*/
func (a *Client) ValidateVpcIDUsingGET1(params *ValidateVpcIDUsingGET1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateVpcIDUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateVpcIDUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateVpcIdUsingGET_1",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/gcp/clusters/validate-vpcId",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateVpcIDUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*ValidateVpcIDUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateVpcIdUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
