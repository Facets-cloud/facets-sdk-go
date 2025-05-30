// Code generated by go-swagger; DO NOT EDIT.

package ui_chat_gpt_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ui chat gpt controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ui chat gpt controller API client with basic auth credentials.
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

// New creates a new ui chat gpt controller API client with a bearer token for authentication.
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
Client for ui chat gpt controller API
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
	AnalyzeKubernetesClusterUsingGET(params *AnalyzeKubernetesClusterUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AnalyzeKubernetesClusterUsingGETOK, error)

	ChatUsingPOST(params *ChatUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChatUsingPOSTOK, *ChatUsingPOSTCreated, error)

	CreateChatUsingPOST(params *CreateChatUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChatUsingPOSTOK, *CreateChatUsingPOSTCreated, error)

	GetAllChatsUsingGET(params *GetAllChatsUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllChatsUsingGETOK, error)

	GetAllStartersUsingGET(params *GetAllStartersUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllStartersUsingGETOK, error)

	GetChatByIDUsingGET(params *GetChatByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChatByIDUsingGETOK, error)

	GetK8sChatsUsingPOST(params *GetK8sChatsUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetK8sChatsUsingPOSTOK, *GetK8sChatsUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AnalyzeKubernetesClusterUsingGET analyzes kubernetes cluster
*/
func (a *Client) AnalyzeKubernetesClusterUsingGET(params *AnalyzeKubernetesClusterUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AnalyzeKubernetesClusterUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnalyzeKubernetesClusterUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "analyzeKubernetesClusterUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/clusters/{clusterId}/kubernetes/analyze",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AnalyzeKubernetesClusterUsingGETReader{formats: a.formats},
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
	success, ok := result.(*AnalyzeKubernetesClusterUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for analyzeKubernetesClusterUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ChatUsingPOST sends a message to a chat
*/
func (a *Client) ChatUsingPOST(params *ChatUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChatUsingPOSTOK, *ChatUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChatUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "chatUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/clusters/chat/{chatId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChatUsingPOSTReader{formats: a.formats},
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
	case *ChatUsingPOSTOK:
		return value, nil, nil
	case *ChatUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_chat_gpt_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateChatUsingPOST creates a new chat
*/
func (a *Client) CreateChatUsingPOST(params *CreateChatUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChatUsingPOSTOK, *CreateChatUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateChatUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createChatUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/clusters/{clusterId}/chat",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChatUsingPOSTReader{formats: a.formats},
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
	case *CreateChatUsingPOSTOK:
		return value, nil, nil
	case *CreateChatUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_chat_gpt_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllChatsUsingGET gets chats by cluster ID
*/
func (a *Client) GetAllChatsUsingGET(params *GetAllChatsUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllChatsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllChatsUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllChatsUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/clusters/{clusterId}/chat",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllChatsUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllChatsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllChatsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllStartersUsingGET gets chat starters metadata
*/
func (a *Client) GetAllStartersUsingGET(params *GetAllStartersUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllStartersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStartersUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllStartersUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/clusters/chat/metadata",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllStartersUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllStartersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllStartersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetChatByIDUsingGET gets chat by ID
*/
func (a *Client) GetChatByIDUsingGET(params *GetChatByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChatByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChatByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getChatByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/clusters/chat/{chatId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetChatByIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetChatByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getChatByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetK8sChatsUsingPOST asks questions about k8s operations
*/
func (a *Client) GetK8sChatsUsingPOST(params *GetK8sChatsUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetK8sChatsUsingPOSTOK, *GetK8sChatsUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetK8sChatsUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getK8sChatsUsingPOST",
		Method:             "POST",
		PathPattern:        "/cc-ui/v1/clusters/{clusterId}/k8s-chat",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetK8sChatsUsingPOSTReader{formats: a.formats},
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
	case *GetK8sChatsUsingPOSTOK:
		return value, nil, nil
	case *GetK8sChatsUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_chat_gpt_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
