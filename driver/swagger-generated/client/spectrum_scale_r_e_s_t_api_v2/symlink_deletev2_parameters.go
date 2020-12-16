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

// NewSymlinkDeletev2Params creates a new SymlinkDeletev2Params object
// with the default values initialized.
func NewSymlinkDeletev2Params() *SymlinkDeletev2Params {
	var ()
	return &SymlinkDeletev2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSymlinkDeletev2ParamsWithTimeout creates a new SymlinkDeletev2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSymlinkDeletev2ParamsWithTimeout(timeout time.Duration) *SymlinkDeletev2Params {
	var ()
	return &SymlinkDeletev2Params{

		timeout: timeout,
	}
}

// NewSymlinkDeletev2ParamsWithContext creates a new SymlinkDeletev2Params object
// with the default values initialized, and the ability to set a context for a request
func NewSymlinkDeletev2ParamsWithContext(ctx context.Context) *SymlinkDeletev2Params {
	var ()
	return &SymlinkDeletev2Params{

		Context: ctx,
	}
}

// NewSymlinkDeletev2ParamsWithHTTPClient creates a new SymlinkDeletev2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSymlinkDeletev2ParamsWithHTTPClient(client *http.Client) *SymlinkDeletev2Params {
	var ()
	return &SymlinkDeletev2Params{
		HTTPClient: client,
	}
}

/*SymlinkDeletev2Params contains all the parameters to send to the API endpoint
for the symlink deletev2 operation typically these are written to a http.Request
*/
type SymlinkDeletev2Params struct {

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

// WithTimeout adds the timeout to the symlink deletev2 params
func (o *SymlinkDeletev2Params) WithTimeout(timeout time.Duration) *SymlinkDeletev2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the symlink deletev2 params
func (o *SymlinkDeletev2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the symlink deletev2 params
func (o *SymlinkDeletev2Params) WithContext(ctx context.Context) *SymlinkDeletev2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the symlink deletev2 params
func (o *SymlinkDeletev2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the symlink deletev2 params
func (o *SymlinkDeletev2Params) WithHTTPClient(client *http.Client) *SymlinkDeletev2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the symlink deletev2 params
func (o *SymlinkDeletev2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilesystemName adds the filesystemName to the symlink deletev2 params
func (o *SymlinkDeletev2Params) WithFilesystemName(filesystemName string) *SymlinkDeletev2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the symlink deletev2 params
func (o *SymlinkDeletev2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithPath adds the path to the symlink deletev2 params
func (o *SymlinkDeletev2Params) WithPath(path string) *SymlinkDeletev2Params {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the symlink deletev2 params
func (o *SymlinkDeletev2Params) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *SymlinkDeletev2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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