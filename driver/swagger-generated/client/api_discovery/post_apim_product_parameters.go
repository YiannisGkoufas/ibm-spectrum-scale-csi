// Code generated by go-swagger; DO NOT EDIT.

package api_discovery

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
	"github.com/go-openapi/swag"

	"example.com/m/v2/models"
)

// NewPostApimProductParams creates a new PostApimProductParams object
// with the default values initialized.
func NewPostApimProductParams() *PostApimProductParams {
	var (
		contentTypeDefault = string("application/json")
		stageOnlyDefault   = bool(false)
	)
	return &PostApimProductParams{
		ContentType: &contentTypeDefault,
		StageOnly:   &stageOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApimProductParamsWithTimeout creates a new PostApimProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApimProductParamsWithTimeout(timeout time.Duration) *PostApimProductParams {
	var (
		contentTypeDefault = string("application/json")
		stageOnlyDefault   = bool(false)
	)
	return &PostApimProductParams{
		ContentType: &contentTypeDefault,
		StageOnly:   &stageOnlyDefault,

		timeout: timeout,
	}
}

// NewPostApimProductParamsWithContext creates a new PostApimProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApimProductParamsWithContext(ctx context.Context) *PostApimProductParams {
	var (
		contentTypeDefault = string("application/json")
		stageOnlyDefault   = bool(false)
	)
	return &PostApimProductParams{
		ContentType: &contentTypeDefault,
		StageOnly:   &stageOnlyDefault,

		Context: ctx,
	}
}

// NewPostApimProductParamsWithHTTPClient creates a new PostApimProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApimProductParamsWithHTTPClient(client *http.Client) *PostApimProductParams {
	var (
		contentTypeDefault = string("application/json")
		stageOnlyDefault   = bool(false)
	)
	return &PostApimProductParams{
		ContentType: &contentTypeDefault,
		StageOnly:   &stageOnlyDefault,
		HTTPClient:  client,
	}
}

/*PostApimProductParams contains all the parameters to send to the API endpoint
for the post apim product operation typically these are written to a http.Request
*/
type PostApimProductParams struct {

	/*ContentType
	  Format of the input payload, which could be JSON (default) or YAML

	*/
	ContentType *string
	/*XAPIMAuthorization
	  Credentials to connect to API Connect. It can have two forms, (1) 'username:password', or (2) 'Basic xyz' where 'xyz' is the base64 encoded version of username:password. Standard API Management login requires the username to start with apimanager/

	*/
	XAPIMAuthorization strfmt.Password
	/*Body
	  A JSON or YAML representation of the API Connect Product

	*/
	Body *models.Product
	/*Catalog
	  Name of the catalog you want to push your APIs into.

	*/
	Catalog string
	/*Organization
	  Name of the organization.  In an on-premises environment the name of the organization is displayed on the API Connect Manager's left side menu.  In Bluemix, the organization is a combination of the Bluemix username (with all @ and . removed), followed by a dash and the space name.  For example, if the Bluemix login username is joe@ibm.com and the space is dev, then the API Connect organization is joeibmcom-dev

	*/
	Organization string
	/*Root
	  Specify the APIs (context roots) that you want to push into API Connect, or leave empty to push all deployed APIs.

	*/
	Root []string
	/*Server
	  API Connect's server address.  An on-premises example is https://apimdev0056.ibm.com/apim, and a Bluemix example is https://us.apiconnect.ibmcloud.com (please note the URL will change depending on the region you're using).

	*/
	Server string
	/*StageOnly
	  Whether or not the product should be only staged, not published into the catalog.

	*/
	StageOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post apim product params
func (o *PostApimProductParams) WithTimeout(timeout time.Duration) *PostApimProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post apim product params
func (o *PostApimProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post apim product params
func (o *PostApimProductParams) WithContext(ctx context.Context) *PostApimProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post apim product params
func (o *PostApimProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post apim product params
func (o *PostApimProductParams) WithHTTPClient(client *http.Client) *PostApimProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post apim product params
func (o *PostApimProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the post apim product params
func (o *PostApimProductParams) WithContentType(contentType *string) *PostApimProductParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the post apim product params
func (o *PostApimProductParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithXAPIMAuthorization adds the xAPIMAuthorization to the post apim product params
func (o *PostApimProductParams) WithXAPIMAuthorization(xAPIMAuthorization strfmt.Password) *PostApimProductParams {
	o.SetXAPIMAuthorization(xAPIMAuthorization)
	return o
}

// SetXAPIMAuthorization adds the xApiMAuthorization to the post apim product params
func (o *PostApimProductParams) SetXAPIMAuthorization(xAPIMAuthorization strfmt.Password) {
	o.XAPIMAuthorization = xAPIMAuthorization
}

// WithBody adds the body to the post apim product params
func (o *PostApimProductParams) WithBody(body *models.Product) *PostApimProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post apim product params
func (o *PostApimProductParams) SetBody(body *models.Product) {
	o.Body = body
}

// WithCatalog adds the catalog to the post apim product params
func (o *PostApimProductParams) WithCatalog(catalog string) *PostApimProductParams {
	o.SetCatalog(catalog)
	return o
}

// SetCatalog adds the catalog to the post apim product params
func (o *PostApimProductParams) SetCatalog(catalog string) {
	o.Catalog = catalog
}

// WithOrganization adds the organization to the post apim product params
func (o *PostApimProductParams) WithOrganization(organization string) *PostApimProductParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the post apim product params
func (o *PostApimProductParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithRoot adds the root to the post apim product params
func (o *PostApimProductParams) WithRoot(root []string) *PostApimProductParams {
	o.SetRoot(root)
	return o
}

// SetRoot adds the root to the post apim product params
func (o *PostApimProductParams) SetRoot(root []string) {
	o.Root = root
}

// WithServer adds the server to the post apim product params
func (o *PostApimProductParams) WithServer(server string) *PostApimProductParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the post apim product params
func (o *PostApimProductParams) SetServer(server string) {
	o.Server = server
}

// WithStageOnly adds the stageOnly to the post apim product params
func (o *PostApimProductParams) WithStageOnly(stageOnly *bool) *PostApimProductParams {
	o.SetStageOnly(stageOnly)
	return o
}

// SetStageOnly adds the stageOnly to the post apim product params
func (o *PostApimProductParams) SetStageOnly(stageOnly *bool) {
	o.StageOnly = stageOnly
}

// WriteToRequest writes these params to a swagger request
func (o *PostApimProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// header param Content-Type
		if err := r.SetHeaderParam("Content-Type", *o.ContentType); err != nil {
			return err
		}

	}

	// header param X-APIM-Authorization
	if err := r.SetHeaderParam("X-APIM-Authorization", o.XAPIMAuthorization.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param catalog
	qrCatalog := o.Catalog
	qCatalog := qrCatalog
	if qCatalog != "" {
		if err := r.SetQueryParam("catalog", qCatalog); err != nil {
			return err
		}
	}

	// query param organization
	qrOrganization := o.Organization
	qOrganization := qrOrganization
	if qOrganization != "" {
		if err := r.SetQueryParam("organization", qOrganization); err != nil {
			return err
		}
	}

	valuesRoot := o.Root

	joinedRoot := swag.JoinByFormat(valuesRoot, "multi")
	// query array param root
	if err := r.SetQueryParam("root", joinedRoot...); err != nil {
		return err
	}

	// query param server
	qrServer := o.Server
	qServer := qrServer
	if qServer != "" {
		if err := r.SetQueryParam("server", qServer); err != nil {
			return err
		}
	}

	if o.StageOnly != nil {

		// query param stageOnly
		var qrStageOnly bool
		if o.StageOnly != nil {
			qrStageOnly = *o.StageOnly
		}
		qStageOnly := swag.FormatBool(qrStageOnly)
		if qStageOnly != "" {
			if err := r.SetQueryParam("stageOnly", qStageOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}