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

// NewACLDeletev2Params creates a new ACLDeletev2Params object
// with the default values initialized.
func NewACLDeletev2Params() *ACLDeletev2Params {
	var ()
	return &ACLDeletev2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewACLDeletev2ParamsWithTimeout creates a new ACLDeletev2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewACLDeletev2ParamsWithTimeout(timeout time.Duration) *ACLDeletev2Params {
	var ()
	return &ACLDeletev2Params{

		timeout: timeout,
	}
}

// NewACLDeletev2ParamsWithContext creates a new ACLDeletev2Params object
// with the default values initialized, and the ability to set a context for a request
func NewACLDeletev2ParamsWithContext(ctx context.Context) *ACLDeletev2Params {
	var ()
	return &ACLDeletev2Params{

		Context: ctx,
	}
}

// NewACLDeletev2ParamsWithHTTPClient creates a new ACLDeletev2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewACLDeletev2ParamsWithHTTPClient(client *http.Client) *ACLDeletev2Params {
	var ()
	return &ACLDeletev2Params{
		HTTPClient: client,
	}
}

/*ACLDeletev2Params contains all the parameters to send to the API endpoint
for the acl deletev2 operation typically these are written to a http.Request
*/
type ACLDeletev2Params struct {

	/*ShareName
	  The name of the smb share

	*/
	ShareName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the acl deletev2 params
func (o *ACLDeletev2Params) WithTimeout(timeout time.Duration) *ACLDeletev2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the acl deletev2 params
func (o *ACLDeletev2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the acl deletev2 params
func (o *ACLDeletev2Params) WithContext(ctx context.Context) *ACLDeletev2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the acl deletev2 params
func (o *ACLDeletev2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the acl deletev2 params
func (o *ACLDeletev2Params) WithHTTPClient(client *http.Client) *ACLDeletev2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the acl deletev2 params
func (o *ACLDeletev2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShareName adds the shareName to the acl deletev2 params
func (o *ACLDeletev2Params) WithShareName(shareName string) *ACLDeletev2Params {
	o.SetShareName(shareName)
	return o
}

// SetShareName adds the shareName to the acl deletev2 params
func (o *ACLDeletev2Params) SetShareName(shareName string) {
	o.ShareName = shareName
}

// WriteToRequest writes these params to a swagger request
func (o *ACLDeletev2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param shareName
	if err := r.SetPathParam("shareName", o.ShareName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}