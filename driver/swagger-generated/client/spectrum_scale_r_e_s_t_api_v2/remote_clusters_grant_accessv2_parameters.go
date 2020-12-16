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

// NewRemoteClustersGrantAccessv2Params creates a new RemoteClustersGrantAccessv2Params object
// with the default values initialized.
func NewRemoteClustersGrantAccessv2Params() *RemoteClustersGrantAccessv2Params {
	var ()
	return &RemoteClustersGrantAccessv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoteClustersGrantAccessv2ParamsWithTimeout creates a new RemoteClustersGrantAccessv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoteClustersGrantAccessv2ParamsWithTimeout(timeout time.Duration) *RemoteClustersGrantAccessv2Params {
	var ()
	return &RemoteClustersGrantAccessv2Params{

		timeout: timeout,
	}
}

// NewRemoteClustersGrantAccessv2ParamsWithContext creates a new RemoteClustersGrantAccessv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewRemoteClustersGrantAccessv2ParamsWithContext(ctx context.Context) *RemoteClustersGrantAccessv2Params {
	var ()
	return &RemoteClustersGrantAccessv2Params{

		Context: ctx,
	}
}

// NewRemoteClustersGrantAccessv2ParamsWithHTTPClient creates a new RemoteClustersGrantAccessv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoteClustersGrantAccessv2ParamsWithHTTPClient(client *http.Client) *RemoteClustersGrantAccessv2Params {
	var ()
	return &RemoteClustersGrantAccessv2Params{
		HTTPClient: client,
	}
}

/*RemoteClustersGrantAccessv2Params contains all the parameters to send to the API endpoint
for the remote clusters grant accessv2 operation typically these are written to a http.Request
*/
type RemoteClustersGrantAccessv2Params struct {

	/*Body*/
	Body *models.OwningFilesystemAccess
	/*OwningClusterFilesystem
	  filesystem name or 'all'

	*/
	OwningClusterFilesystem string
	/*RemoteCluster
	  remote cluster name

	*/
	RemoteCluster string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithTimeout(timeout time.Duration) *RemoteClustersGrantAccessv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithContext(ctx context.Context) *RemoteClustersGrantAccessv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithHTTPClient(client *http.Client) *RemoteClustersGrantAccessv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithBody(body *models.OwningFilesystemAccess) *RemoteClustersGrantAccessv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetBody(body *models.OwningFilesystemAccess) {
	o.Body = body
}

// WithOwningClusterFilesystem adds the owningClusterFilesystem to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithOwningClusterFilesystem(owningClusterFilesystem string) *RemoteClustersGrantAccessv2Params {
	o.SetOwningClusterFilesystem(owningClusterFilesystem)
	return o
}

// SetOwningClusterFilesystem adds the owningClusterFilesystem to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetOwningClusterFilesystem(owningClusterFilesystem string) {
	o.OwningClusterFilesystem = owningClusterFilesystem
}

// WithRemoteCluster adds the remoteCluster to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) WithRemoteCluster(remoteCluster string) *RemoteClustersGrantAccessv2Params {
	o.SetRemoteCluster(remoteCluster)
	return o
}

// SetRemoteCluster adds the remoteCluster to the remote clusters grant accessv2 params
func (o *RemoteClustersGrantAccessv2Params) SetRemoteCluster(remoteCluster string) {
	o.RemoteCluster = remoteCluster
}

// WriteToRequest writes these params to a swagger request
func (o *RemoteClustersGrantAccessv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owningClusterFilesystem
	if err := r.SetPathParam("owningClusterFilesystem", o.OwningClusterFilesystem); err != nil {
		return err
	}

	// path param remoteCluster
	if err := r.SetPathParam("remoteCluster", o.RemoteCluster); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}