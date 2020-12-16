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

// NewCosDirPostv2Params creates a new CosDirPostv2Params object
// with the default values initialized.
func NewCosDirPostv2Params() *CosDirPostv2Params {
	var ()
	return &CosDirPostv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCosDirPostv2ParamsWithTimeout creates a new CosDirPostv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCosDirPostv2ParamsWithTimeout(timeout time.Duration) *CosDirPostv2Params {
	var ()
	return &CosDirPostv2Params{

		timeout: timeout,
	}
}

// NewCosDirPostv2ParamsWithContext creates a new CosDirPostv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCosDirPostv2ParamsWithContext(ctx context.Context) *CosDirPostv2Params {
	var ()
	return &CosDirPostv2Params{

		Context: ctx,
	}
}

// NewCosDirPostv2ParamsWithHTTPClient creates a new CosDirPostv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCosDirPostv2ParamsWithHTTPClient(client *http.Client) *CosDirPostv2Params {
	var ()
	return &CosDirPostv2Params{
		HTTPClient: client,
	}
}

/*CosDirPostv2Params contains all the parameters to send to the API endpoint
for the cos dir postv2 operation typically these are written to a http.Request
*/
type CosDirPostv2Params struct {

	/*Body*/
	Body *models.CosAccessParams
	/*FilesetName
	  The fileset name

	*/
	FilesetName string
	/*FilesystemName
	  The filesystem name

	*/
	FilesystemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cos dir postv2 params
func (o *CosDirPostv2Params) WithTimeout(timeout time.Duration) *CosDirPostv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cos dir postv2 params
func (o *CosDirPostv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cos dir postv2 params
func (o *CosDirPostv2Params) WithContext(ctx context.Context) *CosDirPostv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cos dir postv2 params
func (o *CosDirPostv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cos dir postv2 params
func (o *CosDirPostv2Params) WithHTTPClient(client *http.Client) *CosDirPostv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cos dir postv2 params
func (o *CosDirPostv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cos dir postv2 params
func (o *CosDirPostv2Params) WithBody(body *models.CosAccessParams) *CosDirPostv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cos dir postv2 params
func (o *CosDirPostv2Params) SetBody(body *models.CosAccessParams) {
	o.Body = body
}

// WithFilesetName adds the filesetName to the cos dir postv2 params
func (o *CosDirPostv2Params) WithFilesetName(filesetName string) *CosDirPostv2Params {
	o.SetFilesetName(filesetName)
	return o
}

// SetFilesetName adds the filesetName to the cos dir postv2 params
func (o *CosDirPostv2Params) SetFilesetName(filesetName string) {
	o.FilesetName = filesetName
}

// WithFilesystemName adds the filesystemName to the cos dir postv2 params
func (o *CosDirPostv2Params) WithFilesystemName(filesystemName string) *CosDirPostv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the cos dir postv2 params
func (o *CosDirPostv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WriteToRequest writes these params to a swagger request
func (o *CosDirPostv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param filesetName
	if err := r.SetPathParam("filesetName", o.FilesetName); err != nil {
		return err
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