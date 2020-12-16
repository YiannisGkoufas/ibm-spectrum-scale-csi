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

// NewSnapshotsFilesetsPostv2Params creates a new SnapshotsFilesetsPostv2Params object
// with the default values initialized.
func NewSnapshotsFilesetsPostv2Params() *SnapshotsFilesetsPostv2Params {
	var ()
	return &SnapshotsFilesetsPostv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotsFilesetsPostv2ParamsWithTimeout creates a new SnapshotsFilesetsPostv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSnapshotsFilesetsPostv2ParamsWithTimeout(timeout time.Duration) *SnapshotsFilesetsPostv2Params {
	var ()
	return &SnapshotsFilesetsPostv2Params{

		timeout: timeout,
	}
}

// NewSnapshotsFilesetsPostv2ParamsWithContext creates a new SnapshotsFilesetsPostv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewSnapshotsFilesetsPostv2ParamsWithContext(ctx context.Context) *SnapshotsFilesetsPostv2Params {
	var ()
	return &SnapshotsFilesetsPostv2Params{

		Context: ctx,
	}
}

// NewSnapshotsFilesetsPostv2ParamsWithHTTPClient creates a new SnapshotsFilesetsPostv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSnapshotsFilesetsPostv2ParamsWithHTTPClient(client *http.Client) *SnapshotsFilesetsPostv2Params {
	var ()
	return &SnapshotsFilesetsPostv2Params{
		HTTPClient: client,
	}
}

/*SnapshotsFilesetsPostv2Params contains all the parameters to send to the API endpoint
for the snapshots filesets postv2 operation typically these are written to a http.Request
*/
type SnapshotsFilesetsPostv2Params struct {

	/*Body*/
	Body *models.SnapshotCreate
	/*FilesetName
	  The fileset name

	*/
	FilesetName string
	/*FilesystemName
	  The filesystem name, :all:, :all\_local: or :all\_remote:

	*/
	FilesystemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithTimeout(timeout time.Duration) *SnapshotsFilesetsPostv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithContext(ctx context.Context) *SnapshotsFilesetsPostv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithHTTPClient(client *http.Client) *SnapshotsFilesetsPostv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithBody(body *models.SnapshotCreate) *SnapshotsFilesetsPostv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetBody(body *models.SnapshotCreate) {
	o.Body = body
}

// WithFilesetName adds the filesetName to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithFilesetName(filesetName string) *SnapshotsFilesetsPostv2Params {
	o.SetFilesetName(filesetName)
	return o
}

// SetFilesetName adds the filesetName to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetFilesetName(filesetName string) {
	o.FilesetName = filesetName
}

// WithFilesystemName adds the filesystemName to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) WithFilesystemName(filesystemName string) *SnapshotsFilesetsPostv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the snapshots filesets postv2 params
func (o *SnapshotsFilesetsPostv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotsFilesetsPostv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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