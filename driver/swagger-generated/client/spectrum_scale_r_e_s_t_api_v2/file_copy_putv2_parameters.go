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

// NewFileCopyPutv2Params creates a new FileCopyPutv2Params object
// with the default values initialized.
func NewFileCopyPutv2Params() *FileCopyPutv2Params {
	var ()
	return &FileCopyPutv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewFileCopyPutv2ParamsWithTimeout creates a new FileCopyPutv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewFileCopyPutv2ParamsWithTimeout(timeout time.Duration) *FileCopyPutv2Params {
	var ()
	return &FileCopyPutv2Params{

		timeout: timeout,
	}
}

// NewFileCopyPutv2ParamsWithContext creates a new FileCopyPutv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewFileCopyPutv2ParamsWithContext(ctx context.Context) *FileCopyPutv2Params {
	var ()
	return &FileCopyPutv2Params{

		Context: ctx,
	}
}

// NewFileCopyPutv2ParamsWithHTTPClient creates a new FileCopyPutv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFileCopyPutv2ParamsWithHTTPClient(client *http.Client) *FileCopyPutv2Params {
	var ()
	return &FileCopyPutv2Params{
		HTTPClient: client,
	}
}

/*FileCopyPutv2Params contains all the parameters to send to the API endpoint
for the file copy putv2 operation typically these are written to a http.Request
*/
type FileCopyPutv2Params struct {

	/*Body*/
	Body *models.CopyConfig
	/*FilesetName
	  The fileset name

	*/
	FilesetName string
	/*FilesystemName
	  The filesystem name

	*/
	FilesystemName string
	/*SourcePath
	  The source path relative to the fileset path

	*/
	SourcePath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the file copy putv2 params
func (o *FileCopyPutv2Params) WithTimeout(timeout time.Duration) *FileCopyPutv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file copy putv2 params
func (o *FileCopyPutv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file copy putv2 params
func (o *FileCopyPutv2Params) WithContext(ctx context.Context) *FileCopyPutv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file copy putv2 params
func (o *FileCopyPutv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file copy putv2 params
func (o *FileCopyPutv2Params) WithHTTPClient(client *http.Client) *FileCopyPutv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file copy putv2 params
func (o *FileCopyPutv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the file copy putv2 params
func (o *FileCopyPutv2Params) WithBody(body *models.CopyConfig) *FileCopyPutv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the file copy putv2 params
func (o *FileCopyPutv2Params) SetBody(body *models.CopyConfig) {
	o.Body = body
}

// WithFilesetName adds the filesetName to the file copy putv2 params
func (o *FileCopyPutv2Params) WithFilesetName(filesetName string) *FileCopyPutv2Params {
	o.SetFilesetName(filesetName)
	return o
}

// SetFilesetName adds the filesetName to the file copy putv2 params
func (o *FileCopyPutv2Params) SetFilesetName(filesetName string) {
	o.FilesetName = filesetName
}

// WithFilesystemName adds the filesystemName to the file copy putv2 params
func (o *FileCopyPutv2Params) WithFilesystemName(filesystemName string) *FileCopyPutv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the file copy putv2 params
func (o *FileCopyPutv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithSourcePath adds the sourcePath to the file copy putv2 params
func (o *FileCopyPutv2Params) WithSourcePath(sourcePath string) *FileCopyPutv2Params {
	o.SetSourcePath(sourcePath)
	return o
}

// SetSourcePath adds the sourcePath to the file copy putv2 params
func (o *FileCopyPutv2Params) SetSourcePath(sourcePath string) {
	o.SourcePath = sourcePath
}

// WriteToRequest writes these params to a swagger request
func (o *FileCopyPutv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sourcePath
	if err := r.SetPathParam("sourcePath", o.SourcePath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}