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

// NewDirectoryPostv2Params creates a new DirectoryPostv2Params object
// with the default values initialized.
func NewDirectoryPostv2Params() *DirectoryPostv2Params {
	var ()
	return &DirectoryPostv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDirectoryPostv2ParamsWithTimeout creates a new DirectoryPostv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDirectoryPostv2ParamsWithTimeout(timeout time.Duration) *DirectoryPostv2Params {
	var ()
	return &DirectoryPostv2Params{

		timeout: timeout,
	}
}

// NewDirectoryPostv2ParamsWithContext creates a new DirectoryPostv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDirectoryPostv2ParamsWithContext(ctx context.Context) *DirectoryPostv2Params {
	var ()
	return &DirectoryPostv2Params{

		Context: ctx,
	}
}

// NewDirectoryPostv2ParamsWithHTTPClient creates a new DirectoryPostv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDirectoryPostv2ParamsWithHTTPClient(client *http.Client) *DirectoryPostv2Params {
	var ()
	return &DirectoryPostv2Params{
		HTTPClient: client,
	}
}

/*DirectoryPostv2Params contains all the parameters to send to the API endpoint
for the directory postv2 operation typically these are written to a http.Request
*/
type DirectoryPostv2Params struct {

	/*Body*/
	Body *models.FileOwner
	/*FilesystemName
	  The name of the file system

	*/
	FilesystemName string
	/*Path
	  The directory path relative to the filesystem's mount point

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the directory postv2 params
func (o *DirectoryPostv2Params) WithTimeout(timeout time.Duration) *DirectoryPostv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the directory postv2 params
func (o *DirectoryPostv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the directory postv2 params
func (o *DirectoryPostv2Params) WithContext(ctx context.Context) *DirectoryPostv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the directory postv2 params
func (o *DirectoryPostv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the directory postv2 params
func (o *DirectoryPostv2Params) WithHTTPClient(client *http.Client) *DirectoryPostv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the directory postv2 params
func (o *DirectoryPostv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the directory postv2 params
func (o *DirectoryPostv2Params) WithBody(body *models.FileOwner) *DirectoryPostv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the directory postv2 params
func (o *DirectoryPostv2Params) SetBody(body *models.FileOwner) {
	o.Body = body
}

// WithFilesystemName adds the filesystemName to the directory postv2 params
func (o *DirectoryPostv2Params) WithFilesystemName(filesystemName string) *DirectoryPostv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the directory postv2 params
func (o *DirectoryPostv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithPath adds the path to the directory postv2 params
func (o *DirectoryPostv2Params) WithPath(path string) *DirectoryPostv2Params {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the directory postv2 params
func (o *DirectoryPostv2Params) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DirectoryPostv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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