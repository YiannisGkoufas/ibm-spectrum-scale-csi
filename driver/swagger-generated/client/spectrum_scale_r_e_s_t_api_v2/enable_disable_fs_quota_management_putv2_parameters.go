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

// NewEnableDisableFsQuotaManagementPutv2Params creates a new EnableDisableFsQuotaManagementPutv2Params object
// with the default values initialized.
func NewEnableDisableFsQuotaManagementPutv2Params() *EnableDisableFsQuotaManagementPutv2Params {
	var ()
	return &EnableDisableFsQuotaManagementPutv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableDisableFsQuotaManagementPutv2ParamsWithTimeout creates a new EnableDisableFsQuotaManagementPutv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableDisableFsQuotaManagementPutv2ParamsWithTimeout(timeout time.Duration) *EnableDisableFsQuotaManagementPutv2Params {
	var ()
	return &EnableDisableFsQuotaManagementPutv2Params{

		timeout: timeout,
	}
}

// NewEnableDisableFsQuotaManagementPutv2ParamsWithContext creates a new EnableDisableFsQuotaManagementPutv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewEnableDisableFsQuotaManagementPutv2ParamsWithContext(ctx context.Context) *EnableDisableFsQuotaManagementPutv2Params {
	var ()
	return &EnableDisableFsQuotaManagementPutv2Params{

		Context: ctx,
	}
}

// NewEnableDisableFsQuotaManagementPutv2ParamsWithHTTPClient creates a new EnableDisableFsQuotaManagementPutv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableDisableFsQuotaManagementPutv2ParamsWithHTTPClient(client *http.Client) *EnableDisableFsQuotaManagementPutv2Params {
	var ()
	return &EnableDisableFsQuotaManagementPutv2Params{
		HTTPClient: client,
	}
}

/*EnableDisableFsQuotaManagementPutv2Params contains all the parameters to send to the API endpoint
for the enable disable fs quota management putv2 operation typically these are written to a http.Request
*/
type EnableDisableFsQuotaManagementPutv2Params struct {

	/*Body*/
	Body *models.FilesystemQuotaManagement
	/*FilesystemName
	  The filesystem name, :all:, :all\_local: or :all\_remote:

	*/
	FilesystemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) WithTimeout(timeout time.Duration) *EnableDisableFsQuotaManagementPutv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) WithContext(ctx context.Context) *EnableDisableFsQuotaManagementPutv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) WithHTTPClient(client *http.Client) *EnableDisableFsQuotaManagementPutv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) WithBody(body *models.FilesystemQuotaManagement) *EnableDisableFsQuotaManagementPutv2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) SetBody(body *models.FilesystemQuotaManagement) {
	o.Body = body
}

// WithFilesystemName adds the filesystemName to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) WithFilesystemName(filesystemName string) *EnableDisableFsQuotaManagementPutv2Params {
	o.SetFilesystemName(filesystemName)
	return o
}

// SetFilesystemName adds the filesystemName to the enable disable fs quota management putv2 params
func (o *EnableDisableFsQuotaManagementPutv2Params) SetFilesystemName(filesystemName string) {
	o.FilesystemName = filesystemName
}

// WriteToRequest writes these params to a swagger request
func (o *EnableDisableFsQuotaManagementPutv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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