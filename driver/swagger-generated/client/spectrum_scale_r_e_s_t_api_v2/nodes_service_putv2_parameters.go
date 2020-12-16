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

	"example.com/m/v2/models"
)

// NewNodesServicePutv2Params creates a new NodesServicePutv2Params object
// with the default values initialized.
func NewNodesServicePutv2Params() *NodesServicePutv2Params {
	var ()
	return &NodesServicePutv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesServicePutv2ParamsWithTimeout creates a new NodesServicePutv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesServicePutv2ParamsWithTimeout(timeout time.Duration) *NodesServicePutv2Params {
	var ()
	return &NodesServicePutv2Params{

		timeout: timeout,
	}
}

// NewNodesServicePutv2ParamsWithContext creates a new NodesServicePutv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewNodesServicePutv2ParamsWithContext(ctx context.Context) *NodesServicePutv2Params {
	var ()
	return &NodesServicePutv2Params{

		Context: ctx,
	}
}

// NewNodesServicePutv2ParamsWithHTTPClient creates a new NodesServicePutv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesServicePutv2ParamsWithHTTPClient(client *http.Client) *NodesServicePutv2Params {
	var ()
	return &NodesServicePutv2Params{
		HTTPClient: client,
	}
}

/*NodesServicePutv2Params contains all the parameters to send to the API endpoint
for the nodes service putv2 operation typically these are written to a http.Request
*/
type NodesServicePutv2Params struct {

	/*Body*/
	Body *models.ServiceAction
	/*Name
	  nodeclass or node name

	*/
	Name string
	/*ServiceName
	  name of the service to start/stop

	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithTimeout(timeout time.Duration) *NodesServicePutv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithContext(ctx context.Context) *NodesServicePutv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithHTTPClient(client *http.Client) *NodesServicePutv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithBody(body *models.ServiceAction) *NodesServicePutv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetBody(body *models.ServiceAction) {
	o.Body = body
}

// WithName adds the name to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithName(name string) *NodesServicePutv2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetName(name string) {
	o.Name = name
}

// WithServiceName adds the serviceName to the nodes service putv2 params
func (o *NodesServicePutv2Params) WithServiceName(serviceName string) *NodesServicePutv2Params {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the nodes service putv2 params
func (o *NodesServicePutv2Params) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *NodesServicePutv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}