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

// NewFilesetsPostv2Params creates a new FilesetsPostv2Params object
// with the default values initialized.
func NewFilesetsPostv2Params() *FilesetsPostv2Params {
	var ()
	return &FilesetsPostv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilesetsPostv2ParamsWithTimeout creates a new FilesetsPostv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilesetsPostv2ParamsWithTimeout(timeout time.Duration) *FilesetsPostv2Params {
	var ()
	return &FilesetsPostv2Params{

		timeout: timeout,
	}
}

// NewFilesetsPostv2ParamsWithContext creates a new FilesetsPostv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewFilesetsPostv2ParamsWithContext(ctx context.Context) *FilesetsPostv2Params {
	var ()
	return &FilesetsPostv2Params{

		Context: ctx,
	}
}

// NewFilesetsPostv2ParamsWithHTTPClient creates a new FilesetsPostv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFilesetsPostv2ParamsWithHTTPClient(client *http.Client) *FilesetsPostv2Params {
	var ()
	return &FilesetsPostv2Params{
		HTTPClient: client,
	}
}

/*FilesetsPostv2Params contains all the parameters to send to the API endpoint
for the filesets postv2 operation typically these are written to a http.Request
*/
type FilesetsPostv2Params struct {

	/*Body*/
	Body *models.FilesetCreate
	/*FilesystemName
	  The filesystem name, :all:, :all\_local: or :all\_remote:

	*/
	FilesystemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the filesets postv2 params
func (o *FilesetsPostv2Params) WithTimeout(timeout time.Duration) *FilesetsPostv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the filesets postv2 params
func (o *FilesetsPostv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the filesets postv2 params
func (o *FilesetsPostv2Params) WithContext(ctx context.Context) *FilesetsPostv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the filesets postv2 params
func (o *FilesetsPostv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the filesets postv2 params
func (o *FilesetsPostv2Params) WithHTTPClient(client *http.Client) *FilesetsPostv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the filesets postv2 params
func (o *FilesetsPostv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the filesets postv2 params
func (o *FilesetsPostv2Params) WithBody(body *models.FilesetCreate) *FilesetsPostv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the filesets postv2 params
func (o *FilesetsPostv2Params) SetBody(body *models.FilesetCreate) {
	o.Body = body
}

// WithFilesystemName adds the filesystemName to the filesets postv2 params
func (o *FilesetsPostv2Params) WithFilesystemName(filesystemName string) *FilesetsPostv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the filesets postv2 params
func (o *FilesetsPostv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WriteToRequest writes these params to a swagger request
func (o *FilesetsPostv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}