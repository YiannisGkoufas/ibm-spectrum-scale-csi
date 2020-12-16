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

// NewDirectoryDeletev20Params creates a new DirectoryDeletev20Params object
// with the default values initialized.
func NewDirectoryDeletev20Params() *DirectoryDeletev20Params {
	var ()
	return &DirectoryDeletev20Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDirectoryDeletev20ParamsWithTimeout creates a new DirectoryDeletev20Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDirectoryDeletev20ParamsWithTimeout(timeout time.Duration) *DirectoryDeletev20Params {
	var ()
	return &DirectoryDeletev20Params{

		timeout: timeout,
	}
}

// NewDirectoryDeletev20ParamsWithContext creates a new DirectoryDeletev20Params object
// with the default values initialized, and the ability to set a context for a request
func NewDirectoryDeletev20ParamsWithContext(ctx context.Context) *DirectoryDeletev20Params {
	var ()
	return &DirectoryDeletev20Params{

		Context: ctx,
	}
}

// NewDirectoryDeletev20ParamsWithHTTPClient creates a new DirectoryDeletev20Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDirectoryDeletev20ParamsWithHTTPClient(client *http.Client) *DirectoryDeletev20Params {
	var ()
	return &DirectoryDeletev20Params{
		HTTPClient: client,
	}
}

/*DirectoryDeletev20Params contains all the parameters to send to the API endpoint
for the directory deletev2 0 operation typically these are written to a http.Request
*/
type DirectoryDeletev20Params struct {

	/*FilesetName
	  The fileset name

	*/
	FilesetName string
	/*FilesystemName
	  The name of the file system

	*/
	FilesystemName string
	/*Path
	  The directory path relative to the fileset path

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithTimeout(timeout time.Duration) *DirectoryDeletev20Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithContext(ctx context.Context) *DirectoryDeletev20Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithHTTPClient(client *http.Client) *DirectoryDeletev20Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilesetName adds the filesetName to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithFilesetName(filesetName string) *DirectoryDeletev20Params {
	o.SetFilesetName(filesetName)
	return o
}

// SetFilesetName adds the filesetName to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetFilesetName(filesetName string) {
	o.FilesetName = filesetName
}

// WithFilesystemName adds the filesystemName to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithFilesystemName(filesystemName string) *DirectoryDeletev20Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithPath adds the path to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) WithPath(path string) *DirectoryDeletev20Params {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the directory deletev2 0 params
func (o *DirectoryDeletev20Params) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DirectoryDeletev20Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filesetName
	if err := r.SetPathParam("filesetName", o.FilesetName); err != nil {
		return err
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