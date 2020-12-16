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

// NewRemoteClustersChangeAccessv2Params creates a new RemoteClustersChangeAccessv2Params object
// with the default values initialized.
func NewRemoteClustersChangeAccessv2Params() *RemoteClustersChangeAccessv2Params {
	var ()
	return &RemoteClustersChangeAccessv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoteClustersChangeAccessv2ParamsWithTimeout creates a new RemoteClustersChangeAccessv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoteClustersChangeAccessv2ParamsWithTimeout(timeout time.Duration) *RemoteClustersChangeAccessv2Params {
	var ()
	return &RemoteClustersChangeAccessv2Params{

		timeout: timeout,
	}
}

// NewRemoteClustersChangeAccessv2ParamsWithContext creates a new RemoteClustersChangeAccessv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewRemoteClustersChangeAccessv2ParamsWithContext(ctx context.Context) *RemoteClustersChangeAccessv2Params {
	var ()
	return &RemoteClustersChangeAccessv2Params{

		Context: ctx,
	}
}

// NewRemoteClustersChangeAccessv2ParamsWithHTTPClient creates a new RemoteClustersChangeAccessv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoteClustersChangeAccessv2ParamsWithHTTPClient(client *http.Client) *RemoteClustersChangeAccessv2Params {
	var ()
	return &RemoteClustersChangeAccessv2Params{
		HTTPClient: client,
	}
}

/*RemoteClustersChangeAccessv2Params contains all the parameters to send to the API endpoint
for the remote clusters change accessv2 operation typically these are written to a http.Request
*/
type RemoteClustersChangeAccessv2Params struct {

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

// WithTimeout adds the timeout to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithTimeout(timeout time.Duration) *RemoteClustersChangeAccessv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithContext(ctx context.Context) *RemoteClustersChangeAccessv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithHTTPClient(client *http.Client) *RemoteClustersChangeAccessv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithBody(body *models.OwningFilesystemAccess) *RemoteClustersChangeAccessv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetBody(body *models.OwningFilesystemAccess) {
	o.Body = body
}

// WithOwningClusterFilesystem adds the owningClusterFilesystem to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithOwningClusterFilesystem(owningClusterFilesystem string) *RemoteClustersChangeAccessv2Params {
	o.SetOwningClusterFilesystem(owningClusterFilesystem)
	return o
}

// SetOwningClusterFilesystem adds the owningClusterFilesystem to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetOwningClusterFilesystem(owningClusterFilesystem string) {
	o.OwningClusterFilesystem = owningClusterFilesystem
}

// WithRemoteCluster adds the remoteCluster to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) WithRemoteCluster(remoteCluster string) *RemoteClustersChangeAccessv2Params {
	o.SetRemoteCluster(remoteCluster)
	return o
}

// SetRemoteCluster adds the remoteCluster to the remote clusters change accessv2 params
func (o *RemoteClustersChangeAccessv2Params) SetRemoteCluster(remoteCluster string) {
	o.RemoteCluster = remoteCluster
}

// WriteToRequest writes these params to a swagger request
func (o *RemoteClustersChangeAccessv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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