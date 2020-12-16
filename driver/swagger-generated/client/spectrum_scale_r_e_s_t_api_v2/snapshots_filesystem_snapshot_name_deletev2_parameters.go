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

// NewSnapshotsFilesystemSnapshotNameDeletev2Params creates a new SnapshotsFilesystemSnapshotNameDeletev2Params object
// with the default values initialized.
func NewSnapshotsFilesystemSnapshotNameDeletev2Params() *SnapshotsFilesystemSnapshotNameDeletev2Params {
	var ()
	return &SnapshotsFilesystemSnapshotNameDeletev2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithTimeout creates a new SnapshotsFilesystemSnapshotNameDeletev2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithTimeout(timeout time.Duration) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	var ()
	return &SnapshotsFilesystemSnapshotNameDeletev2Params{

		timeout: timeout,
	}
}

// NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithContext creates a new SnapshotsFilesystemSnapshotNameDeletev2Params object
// with the default values initialized, and the ability to set a context for a request
func NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithContext(ctx context.Context) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	var ()
	return &SnapshotsFilesystemSnapshotNameDeletev2Params{

		Context: ctx,
	}
}

// NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithHTTPClient creates a new SnapshotsFilesystemSnapshotNameDeletev2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSnapshotsFilesystemSnapshotNameDeletev2ParamsWithHTTPClient(client *http.Client) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	var ()
	return &SnapshotsFilesystemSnapshotNameDeletev2Params{
		HTTPClient: client,
	}
}

/*SnapshotsFilesystemSnapshotNameDeletev2Params contains all the parameters to send to the API endpoint
for the snapshots filesystem snapshot name deletev2 operation typically these are written to a http.Request
*/
type SnapshotsFilesystemSnapshotNameDeletev2Params struct {

	/*FilesystemName
	  The filesystem name, :all:, :all\_local: or :all\_remote:

	*/
	FilesystemName string
	/*SnapshotName
	  The snapshot name

	*/
	SnapshotName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WithTimeout(timeout time.Duration) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WithContext(ctx context.Context) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WithHTTPClient(client *http.Client) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilesystemName adds the filesystemName to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WithFilesystemName(filesystemName string) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WithSnapshotName adds the snapshotName to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WithSnapshotName(snapshotName string) *SnapshotsFilesystemSnapshotNameDeletev2Params {
	o.SetSnapshotName(snapshotName)
	return o
}

// SetSnapshotName adds the snapshotName to the snapshots filesystem snapshot name deletev2 params
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) SetSnapshotName(snapshotName string) {
	o.SnapshotName = snapshotName
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotsFilesystemSnapshotNameDeletev2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filesystemName
	if err := r.SetPathParam("filesystemName", o.FilesystemName); err != nil {
		return err
	}

	// path param snapshotName
	if err := r.SetPathParam("snapshotName", o.SnapshotName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}