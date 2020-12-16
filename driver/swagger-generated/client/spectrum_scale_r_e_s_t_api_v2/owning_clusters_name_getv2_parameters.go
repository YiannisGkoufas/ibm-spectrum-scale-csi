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

// NewOwningClustersNameGetv2Params creates a new OwningClustersNameGetv2Params object
// with the default values initialized.
func NewOwningClustersNameGetv2Params() *OwningClustersNameGetv2Params {
	var ()
	return &OwningClustersNameGetv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewOwningClustersNameGetv2ParamsWithTimeout creates a new OwningClustersNameGetv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewOwningClustersNameGetv2ParamsWithTimeout(timeout time.Duration) *OwningClustersNameGetv2Params {
	var ()
	return &OwningClustersNameGetv2Params{

		timeout: timeout,
	}
}

// NewOwningClustersNameGetv2ParamsWithContext creates a new OwningClustersNameGetv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewOwningClustersNameGetv2ParamsWithContext(ctx context.Context) *OwningClustersNameGetv2Params {
	var ()
	return &OwningClustersNameGetv2Params{

		Context: ctx,
	}
}

// NewOwningClustersNameGetv2ParamsWithHTTPClient creates a new OwningClustersNameGetv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOwningClustersNameGetv2ParamsWithHTTPClient(client *http.Client) *OwningClustersNameGetv2Params {
	var ()
	return &OwningClustersNameGetv2Params{
		HTTPClient: client,
	}
}

/*OwningClustersNameGetv2Params contains all the parameters to send to the API endpoint
for the owning clusters name getv2 operation typically these are written to a http.Request
*/
type OwningClustersNameGetv2Params struct {

	/*OwningCluster
	  owning cluster name

	*/
	OwningCluster string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) WithTimeout(timeout time.Duration) *OwningClustersNameGetv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) WithContext(ctx context.Context) *OwningClustersNameGetv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) WithHTTPClient(client *http.Client) *OwningClustersNameGetv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwningCluster adds the owningCluster to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) WithOwningCluster(owningCluster string) *OwningClustersNameGetv2Params {
	o.SetOwningCluster(owningCluster)
	return o
}

// SetOwningCluster adds the owningCluster to the owning clusters name getv2 params
func (o *OwningClustersNameGetv2Params) SetOwningCluster(owningCluster string) {
	o.OwningCluster = owningCluster
}

// WriteToRequest writes these params to a swagger request
func (o *OwningClustersNameGetv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owningCluster
	if err := r.SetPathParam("owningCluster", o.OwningCluster); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}