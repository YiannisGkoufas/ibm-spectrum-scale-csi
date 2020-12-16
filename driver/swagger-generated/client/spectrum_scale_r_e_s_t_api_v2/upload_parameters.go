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

// NewUploadParams creates a new UploadParams object
// with the default values initialized.
func NewUploadParams() *UploadParams {
	var ()
	return &UploadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadParamsWithTimeout creates a new UploadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadParamsWithTimeout(timeout time.Duration) *UploadParams {
	var ()
	return &UploadParams{

		timeout: timeout,
	}
}

// NewUploadParamsWithContext creates a new UploadParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadParamsWithContext(ctx context.Context) *UploadParams {
	var ()
	return &UploadParams{

		Context: ctx,
	}
}

// NewUploadParamsWithHTTPClient creates a new UploadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadParamsWithHTTPClient(client *http.Client) *UploadParams {
	var ()
	return &UploadParams{
		HTTPClient: client,
	}
}

/*UploadParams contains all the parameters to send to the API endpoint
for the upload operation typically these are written to a http.Request
*/
type UploadParams struct {

	/*Body*/
	Body *models.CosUploadParams
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

// WithTimeout adds the timeout to the upload params
func (o *UploadParams) WithTimeout(timeout time.Duration) *UploadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload params
func (o *UploadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload params
func (o *UploadParams) WithContext(ctx context.Context) *UploadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload params
func (o *UploadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload params
func (o *UploadParams) WithHTTPClient(client *http.Client) *UploadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload params
func (o *UploadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upload params
func (o *UploadParams) WithBody(body *models.CosUploadParams) *UploadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload params
func (o *UploadParams) SetBody(body *models.CosUploadParams) {
	o.Body = body
}

// WithFilesetName adds the filesetName to the upload params
func (o *UploadParams) WithFilesetName(filesetName string) *UploadParams {
	o.SetFilesetName(filesetName)
	return o
}

// SetFilesetName adds the filesetName to the upload params
func (o *UploadParams) SetFilesetName(filesetName string) {
	o.FilesetName = filesetName
}

// WithFilesystemName adds the filesystemName to the upload params
func (o *UploadParams) WithFilesystemName(filesystemName string) *UploadParams {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the upload params
func (o *UploadParams) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WriteToRequest writes these params to a swagger request
func (o *UploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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