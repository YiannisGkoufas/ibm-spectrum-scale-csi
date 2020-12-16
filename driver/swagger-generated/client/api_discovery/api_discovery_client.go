// Code generated by go-swagger; DO NOT EDIT.

package api_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api discovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api discovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPICDocs(params *GetAPICDocsParams) (*GetAPICDocsOK, error)

	GetDocs(params *GetDocsParams) (*GetDocsOK, error)

	PostApimProduct(params *PostApimProductParams) (*PostApimProductOK, error)

	PostSubscription(params *PostSubscriptionParams) (*PostSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPICDocs produces a swagger document that is ready to be pulled into API connect

  Returns a valid Swagger 2.0 document with all application REST APIs available from a Liberty instance merged into a single document and with the appropriate gateway attributes and assembly setup to integrate with API Connect. By default, the document is returned in JSON format. Including an optional 'Accept' header with an 'application/yaml' value provides the document in YAML format. This endpoint has a multiple-cardinality, optional query parameter called 'root' that can filter the found context roots. The returned document is always in a non-compact format.
*/
func (a *Client) GetAPICDocs(params *GetAPICDocsParams) (*GetAPICDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPICDocsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAPICDocs",
		Method:             "GET",
		PathPattern:        "/ibm/api/docs/apiconnect",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPICDocsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPICDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAPICDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDocs gets a valid swagger 2 0 document with all r e s t a p is available from a liberty instance

  Returns a valid Swagger 2.0 document with all REST APIs available from a Liberty instance merged into a single document. This is useful for consumer applications that want to programmatically navigate the set of available APIs, such as an API Management solution. By default, the document is returned in JSON format. Including an optional 'Accept' header with an 'application/yaml' value provides the document in YAML format. This endpoint has two multiple-cardinality, optional query parameters called 'root' that can filter the found context roots and 'compact' that can control whether the output should be formatted or not.
*/
func (a *Client) GetDocs(params *GetDocsParams) (*GetDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDocsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDocs",
		Method:             "GET",
		PathPattern:        "/ibm/api/docs/",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDocsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostApimProduct pushes liberty r e s t endpoints into an API connect catalog

  Push Liberty REST endpoints into an API Connect catalog.
*/
func (a *Client) PostApimProduct(params *PostApimProductParams) (*PostApimProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostApimProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postApimProduct",
		Method:             "POST",
		PathPattern:        "/ibm/api/docs/apiconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostApimProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApimProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postApimProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostSubscription subscribes to receive r e s t API updates

  Allows users to subscribe to any REST API update, such as new APIs being available or old APIs being removed. This is useful when a user wants to be notified right away of any changes in the endpoints that are provided by a particular Liberty instance. Note that in addition to the base 'apiDiscovery-1.0' configuration, it is required to also configure either 'websocket-1.0' or 'websocket-1.1' in the 'server.xml'
*/
func (a *Client) PostSubscription(params *PostSubscriptionParams) (*PostSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSubscription",
		Method:             "POST",
		PathPattern:        "/ibm/api/docs/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSubscriptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}