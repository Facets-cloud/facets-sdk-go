// Code generated by go-swagger; DO NOT EDIT.

package ui_settings_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ui settings controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ui settings controller API client with basic auth credentials.
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

// New creates a new ui settings controller API client with a bearer token for authentication.
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
Client for ui settings controller API
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
	AddSettingValueUsingPUT(params *AddSettingValueUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddSettingValueUsingPUTOK, *AddSettingValueUsingPUTCreated, error)

	GetAllSettingsForEntityUsingGET(params *GetAllSettingsForEntityUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSettingsForEntityUsingGETOK, error)

	GetAllSettingsYamlUsingGET(params *GetAllSettingsYamlUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSettingsYamlUsingGETOK, error)

	GetSettingValueUsingGET(params *GetSettingValueUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingValueUsingGETOK, error)

	SetOnboardingDisplayUsingPUT(params *SetOnboardingDisplayUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetOnboardingDisplayUsingPUTOK, *SetOnboardingDisplayUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddSettingValueUsingPUT adds setting value
*/
func (a *Client) AddSettingValueUsingPUT(params *AddSettingValueUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddSettingValueUsingPUTOK, *AddSettingValueUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSettingValueUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addSettingValueUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cc-ui/v1/settings/value/{entityType}/{entityId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSettingValueUsingPUTReader{formats: a.formats},
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
	case *AddSettingValueUsingPUTOK:
		return value, nil, nil
	case *AddSettingValueUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_settings_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllSettingsForEntityUsingGET gets all settings for entity
*/
func (a *Client) GetAllSettingsForEntityUsingGET(params *GetAllSettingsForEntityUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSettingsForEntityUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllSettingsForEntityUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllSettingsForEntityUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/settings/entity/{entity}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllSettingsForEntityUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllSettingsForEntityUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllSettingsForEntityUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllSettingsYamlUsingGET gets all settings yaml
*/
func (a *Client) GetAllSettingsYamlUsingGET(params *GetAllSettingsYamlUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllSettingsYamlUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllSettingsYamlUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllSettingsYamlUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/settings/ui-yaml",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllSettingsYamlUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllSettingsYamlUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllSettingsYamlUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSettingValueUsingGET gets setting value
*/
func (a *Client) GetSettingValueUsingGET(params *GetSettingValueUsingGETParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingValueUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingValueUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSettingValueUsingGET",
		Method:             "GET",
		PathPattern:        "/cc-ui/v1/settings/value/{entityType}/{entityId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSettingValueUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetSettingValueUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSettingValueUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetOnboardingDisplayUsingPUT sets onboarding display
*/
func (a *Client) SetOnboardingDisplayUsingPUT(params *SetOnboardingDisplayUsingPUTParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetOnboardingDisplayUsingPUTOK, *SetOnboardingDisplayUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetOnboardingDisplayUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setOnboardingDisplayUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cc-ui/v1/settings/onboarding-display/{value}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetOnboardingDisplayUsingPUTReader{formats: a.formats},
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
	case *SetOnboardingDisplayUsingPUTOK:
		return value, nil, nil
	case *SetOnboardingDisplayUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ui_settings_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
