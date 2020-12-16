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

// NewOwnerPostv2Params creates a new OwnerPostv2Params object
// with the default values initialized.
func NewOwnerPostv2Params() *OwnerPostv2Params {
	var ()
	return &OwnerPostv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewOwnerPostv2ParamsWithTimeout creates a new OwnerPostv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewOwnerPostv2ParamsWithTimeout(timeout time.Duration) *OwnerPostv2Params {
	var ()
	return &OwnerPostv2Params{

		timeout: timeout,
	}
}

// NewOwnerPostv2ParamsWithContext creates a new OwnerPostv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewOwnerPostv2ParamsWithContext(ctx context.Context) *OwnerPostv2Params {
	var ()
	return &OwnerPostv2Params{

		Context: ctx,
	}
}

// NewOwnerPostv2ParamsWithHTTPClient creates a new OwnerPostv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOwnerPostv2ParamsWithHTTPClient(client *http.Client) *OwnerPostv2Params {
	var ()
	return &OwnerPostv2Params{
		HTTPClient: client,
	}
}

/*OwnerPostv2Params contains all the parameters to send to the API endpoint
for the owner postv2 operation typically these are written to a http.Request
*/
type OwnerPostv2Params struct {

	/*Body*/
	Body *models.FileOwner
	/*FilesystemName
	  name of the filesystem

	*/
	FilesystemName string
	/*Path
	  The file path relative to the filesystem's mount point

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the owner postv2 params
func (o *OwnerPostv2Params) WithTimeout(timeout time.Duration) *OwnerPostv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the owner postv2 params
func (o *OwnerPostv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the owner postv2 params
func (o *OwnerPostv2Params) WithContext(ctx context.Context) *OwnerPostv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the owner postv2 params
func (o *OwnerPostv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the owner postv2 params
func (o *OwnerPostv2Params) WithHTTPClient(client *http.Client) *OwnerPostv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the owner postv2 params
func (o *OwnerPostv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the owner postv2 params
func (o *OwnerPostv2Params) WithBody(body *models.FileOwner) *OwnerPostv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the owner postv2 params
func (o *OwnerPostv2Params) SetBody(body *models.FileOwner) {
	o.Body = body
}

// WithFilesystemName adds the filesystemName to the owner postv2 params
func (o *OwnerPostv2Params) WithFilesystemName(filesystemName string) *OwnerPostv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the owner postv2 params
func (o *OwnerPostv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithPath adds the path to the owner postv2 params
func (o *OwnerPostv2Params) WithPath(path string) *OwnerPostv2Params {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the owner postv2 params
func (o *OwnerPostv2Params) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *OwnerPostv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param filesystemName
	if err := r.SetPathParam("filesystemName", o.FilesystemName); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}