// Code generated by go-swagger; DO NOT EDIT.

package spectrum_scale_r_e_s_t_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewNodesDeleteMappingv2Params creates a new NodesDeleteMappingv2Params object
// with the default values initialized.
func NewNodesDeleteMappingv2Params() *NodesDeleteMappingv2Params {
	var ()
	return &NodesDeleteMappingv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesDeleteMappingv2ParamsWithTimeout creates a new NodesDeleteMappingv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesDeleteMappingv2ParamsWithTimeout(timeout time.Duration) *NodesDeleteMappingv2Params {
	var ()
	return &NodesDeleteMappingv2Params{

		timeout: timeout,
	}
}

// NewNodesDeleteMappingv2ParamsWithContext creates a new NodesDeleteMappingv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewNodesDeleteMappingv2ParamsWithContext(ctx context.Context) *NodesDeleteMappingv2Params {
	var ()
	return &NodesDeleteMappingv2Params{

		Context: ctx,
	}
}

// NewNodesDeleteMappingv2ParamsWithHTTPClient creates a new NodesDeleteMappingv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesDeleteMappingv2ParamsWithHTTPClient(client *http.Client) *NodesDeleteMappingv2Params {
	var ()
	return &NodesDeleteMappingv2Params{
		HTTPClient: client,
	}
}

/*NodesDeleteMappingv2Params contains all the parameters to send to the API endpoint
for the nodes delete mappingv2 operation typically these are written to a http.Request
*/
type NodesDeleteMappingv2Params struct {

	/*MappingName
	  node mapping name

	*/
	MappingName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) WithTimeout(timeout time.Duration) *NodesDeleteMappingv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) WithContext(ctx context.Context) *NodesDeleteMappingv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) WithHTTPClient(client *http.Client) *NodesDeleteMappingv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMappingName adds the mappingName to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) WithMappingName(mappingName string) *NodesDeleteMappingv2Params {
	o.SetMappingName(mappingName)
	return o
}

// SetMappingName adds the mappingName to the nodes delete mappingv2 params
func (o *NodesDeleteMappingv2Params) SetMappingName(mappingName string) {
	o.MappingName = mappingName
}

// WriteToRequest writes these params to a swagger request
func (o *NodesDeleteMappingv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mappingName
	if err := r.SetPathParam("mappingName", o.MappingName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}