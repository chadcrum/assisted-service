// Code generated by go-swagger; DO NOT EDIT.

package installer

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

	"github.com/openshift/assisted-service/models"
)

// NewV2RegisterClusterParams creates a new V2RegisterClusterParams object
// with the default values initialized.
func NewV2RegisterClusterParams() *V2RegisterClusterParams {
	var ()
	return &V2RegisterClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2RegisterClusterParamsWithTimeout creates a new V2RegisterClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2RegisterClusterParamsWithTimeout(timeout time.Duration) *V2RegisterClusterParams {
	var ()
	return &V2RegisterClusterParams{

		timeout: timeout,
	}
}

// NewV2RegisterClusterParamsWithContext creates a new V2RegisterClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2RegisterClusterParamsWithContext(ctx context.Context) *V2RegisterClusterParams {
	var ()
	return &V2RegisterClusterParams{

		Context: ctx,
	}
}

// NewV2RegisterClusterParamsWithHTTPClient creates a new V2RegisterClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2RegisterClusterParamsWithHTTPClient(client *http.Client) *V2RegisterClusterParams {
	var ()
	return &V2RegisterClusterParams{
		HTTPClient: client,
	}
}

/*V2RegisterClusterParams contains all the parameters to send to the API endpoint
for the v2 register cluster operation typically these are written to a http.Request
*/
type V2RegisterClusterParams struct {

	/*NewClusterParams
	  The properties describing the new cluster.

	*/
	NewClusterParams *models.ClusterCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 register cluster params
func (o *V2RegisterClusterParams) WithTimeout(timeout time.Duration) *V2RegisterClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 register cluster params
func (o *V2RegisterClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 register cluster params
func (o *V2RegisterClusterParams) WithContext(ctx context.Context) *V2RegisterClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 register cluster params
func (o *V2RegisterClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 register cluster params
func (o *V2RegisterClusterParams) WithHTTPClient(client *http.Client) *V2RegisterClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 register cluster params
func (o *V2RegisterClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewClusterParams adds the newClusterParams to the v2 register cluster params
func (o *V2RegisterClusterParams) WithNewClusterParams(newClusterParams *models.ClusterCreateParams) *V2RegisterClusterParams {
	o.SetNewClusterParams(newClusterParams)
	return o
}

// SetNewClusterParams adds the newClusterParams to the v2 register cluster params
func (o *V2RegisterClusterParams) SetNewClusterParams(newClusterParams *models.ClusterCreateParams) {
	o.NewClusterParams = newClusterParams
}

// WriteToRequest writes these params to a swagger request
func (o *V2RegisterClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NewClusterParams != nil {
		if err := r.SetBodyParam(o.NewClusterParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
