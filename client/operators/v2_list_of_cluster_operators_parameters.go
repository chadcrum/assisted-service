// Code generated by go-swagger; DO NOT EDIT.

package operators

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

// NewV2ListOfClusterOperatorsParams creates a new V2ListOfClusterOperatorsParams object
// with the default values initialized.
func NewV2ListOfClusterOperatorsParams() *V2ListOfClusterOperatorsParams {
	var ()
	return &V2ListOfClusterOperatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2ListOfClusterOperatorsParamsWithTimeout creates a new V2ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2ListOfClusterOperatorsParamsWithTimeout(timeout time.Duration) *V2ListOfClusterOperatorsParams {
	var ()
	return &V2ListOfClusterOperatorsParams{

		timeout: timeout,
	}
}

// NewV2ListOfClusterOperatorsParamsWithContext creates a new V2ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2ListOfClusterOperatorsParamsWithContext(ctx context.Context) *V2ListOfClusterOperatorsParams {
	var ()
	return &V2ListOfClusterOperatorsParams{

		Context: ctx,
	}
}

// NewV2ListOfClusterOperatorsParamsWithHTTPClient creates a new V2ListOfClusterOperatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2ListOfClusterOperatorsParamsWithHTTPClient(client *http.Client) *V2ListOfClusterOperatorsParams {
	var ()
	return &V2ListOfClusterOperatorsParams{
		HTTPClient: client,
	}
}

/*V2ListOfClusterOperatorsParams contains all the parameters to send to the API endpoint
for the v2 list of cluster operators operation typically these are written to a http.Request
*/
type V2ListOfClusterOperatorsParams struct {

	/*ClusterID
	  The cluster to return operators for.

	*/
	ClusterID strfmt.UUID
	/*OperatorName
	  An operator in the specified cluster to return its data.

	*/
	OperatorName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) WithTimeout(timeout time.Duration) *V2ListOfClusterOperatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) WithContext(ctx context.Context) *V2ListOfClusterOperatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) WithHTTPClient(client *http.Client) *V2ListOfClusterOperatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) WithClusterID(clusterID strfmt.UUID) *V2ListOfClusterOperatorsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithOperatorName adds the operatorName to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) WithOperatorName(operatorName *string) *V2ListOfClusterOperatorsParams {
	o.SetOperatorName(operatorName)
	return o
}

// SetOperatorName adds the operatorName to the v2 list of cluster operators params
func (o *V2ListOfClusterOperatorsParams) SetOperatorName(operatorName *string) {
	o.OperatorName = operatorName
}

// WriteToRequest writes these params to a swagger request
func (o *V2ListOfClusterOperatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.OperatorName != nil {

		// query param operator_name
		var qrOperatorName string
		if o.OperatorName != nil {
			qrOperatorName = *o.OperatorName
		}
		qOperatorName := qrOperatorName
		if qOperatorName != "" {
			if err := r.SetQueryParam("operator_name", qOperatorName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
