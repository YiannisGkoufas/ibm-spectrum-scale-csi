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

// NewCesaddressesCesAddressGetv2Params creates a new CesaddressesCesAddressGetv2Params object
// with the default values initialized.
func NewCesaddressesCesAddressGetv2Params() *CesaddressesCesAddressGetv2Params {
	var ()
	return &CesaddressesCesAddressGetv2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCesaddressesCesAddressGetv2ParamsWithTimeout creates a new CesaddressesCesAddressGetv2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCesaddressesCesAddressGetv2ParamsWithTimeout(timeout time.Duration) *CesaddressesCesAddressGetv2Params {
	var ()
	return &CesaddressesCesAddressGetv2Params{

		timeout: timeout,
	}
}

// NewCesaddressesCesAddressGetv2ParamsWithContext creates a new CesaddressesCesAddressGetv2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCesaddressesCesAddressGetv2ParamsWithContext(ctx context.Context) *CesaddressesCesAddressGetv2Params {
	var ()
	return &CesaddressesCesAddressGetv2Params{

		Context: ctx,
	}
}

// NewCesaddressesCesAddressGetv2ParamsWithHTTPClient creates a new CesaddressesCesAddressGetv2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCesaddressesCesAddressGetv2ParamsWithHTTPClient(client *http.Client) *CesaddressesCesAddressGetv2Params {
	var ()
	return &CesaddressesCesAddressGetv2Params{
		HTTPClient: client,
	}
}

/*CesaddressesCesAddressGetv2Params contains all the parameters to send to the API endpoint
for the cesaddresses ces address getv2 operation typically these are written to a http.Request
*/
type CesaddressesCesAddressGetv2Params struct {

	/*CesAddress
	  IP address to query

	*/
	CesAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) WithTimeout(timeout time.Duration) *CesaddressesCesAddressGetv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) WithContext(ctx context.Context) *CesaddressesCesAddressGetv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) WithHTTPClient(client *http.Client) *CesaddressesCesAddressGetv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCesAddress adds the cesAddress to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) WithCesAddress(cesAddress string) *CesaddressesCesAddressGetv2Params {
	o.SetCesAddress(cesAddress)
	return o
}

// SetCesAddress adds the cesAddress to the cesaddresses ces address getv2 params
func (o *CesaddressesCesAddressGetv2Params) SetCesAddress(cesAddress string) {
	o.CesAddress = cesAddress
}

// WriteToRequest writes these params to a swagger request
func (o *CesaddressesCesAddressGetv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cesAddress
	if err := r.SetPathParam("cesAddress", o.CesAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}