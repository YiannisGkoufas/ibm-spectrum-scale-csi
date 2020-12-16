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

// NewRemoteFilesystemsDeletev2Params creates a new RemoteFilesystemsDeletev2Params object
// with the default values initialized.
func NewRemoteFilesystemsDeletev2Params() *RemoteFilesystemsDeletev2Params {
	var ()
	return &RemoteFilesystemsDeletev2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoteFilesystemsDeletev2ParamsWithTimeout creates a new RemoteFilesystemsDeletev2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoteFilesystemsDeletev2ParamsWithTimeout(timeout time.Duration) *RemoteFilesystemsDeletev2Params {
	var ()
	return &RemoteFilesystemsDeletev2Params{

		timeout: timeout,
	}
}

// NewRemoteFilesystemsDeletev2ParamsWithContext creates a new RemoteFilesystemsDeletev2Params object
// with the default values initialized, and the ability to set a context for a request
func NewRemoteFilesystemsDeletev2ParamsWithContext(ctx context.Context) *RemoteFilesystemsDeletev2Params {
	var ()
	return &RemoteFilesystemsDeletev2Params{

		Context: ctx,
	}
}

// NewRemoteFilesystemsDeletev2ParamsWithHTTPClient creates a new RemoteFilesystemsDeletev2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoteFilesystemsDeletev2ParamsWithHTTPClient(client *http.Client) *RemoteFilesystemsDeletev2Params {
	var ()
	return &RemoteFilesystemsDeletev2Params{
		HTTPClient: client,
	}
}

/*RemoteFilesystemsDeletev2Params contains all the parameters to send to the API endpoint
for the remote filesystems deletev2 operation typically these are written to a http.Request
*/
type RemoteFilesystemsDeletev2Params struct {

	/*Force
	  Force deletion of the remote filesystem

	*/
	Force *string
	/*RemoteFilesystem
	  remote filesystem name

	*/
	RemoteFilesystem string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) WithTimeout(timeout time.Duration) *RemoteFilesystemsDeletev2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) WithContext(ctx context.Context) *RemoteFilesystemsDeletev2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) WithHTTPClient(client *http.Client) *RemoteFilesystemsDeletev2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) WithForce(force *string) *RemoteFilesystemsDeletev2Params {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) SetForce(force *string) {
	o.Force = force
}

// WithRemoteFilesystem adds the remoteFilesystem to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) WithRemoteFilesystem(remoteFilesystem string) *RemoteFilesystemsDeletev2Params {
	o.SetRemoteFilesystem(remoteFilesystem)
	return o
}

// SetRemoteFilesystem adds the remoteFilesystem to the remote filesystems deletev2 params
func (o *RemoteFilesystemsDeletev2Params) SetRemoteFilesystem(remoteFilesystem string) {
	o.RemoteFilesystem = remoteFilesystem
}

// WriteToRequest writes these params to a swagger request
func (o *RemoteFilesystemsDeletev2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce string
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := qrForce
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	// path param remoteFilesystem
	if err := r.SetPathParam("remoteFilesystem", o.RemoteFilesystem); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}