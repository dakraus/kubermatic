// Code generated by go-swagger; DO NOT EDIT.

package packet

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

// NewListPacketSizesNoCredentialsV2Params creates a new ListPacketSizesNoCredentialsV2Params object
// with the default values initialized.
func NewListPacketSizesNoCredentialsV2Params() *ListPacketSizesNoCredentialsV2Params {
	var ()
	return &ListPacketSizesNoCredentialsV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPacketSizesNoCredentialsV2ParamsWithTimeout creates a new ListPacketSizesNoCredentialsV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPacketSizesNoCredentialsV2ParamsWithTimeout(timeout time.Duration) *ListPacketSizesNoCredentialsV2Params {
	var ()
	return &ListPacketSizesNoCredentialsV2Params{

		timeout: timeout,
	}
}

// NewListPacketSizesNoCredentialsV2ParamsWithContext creates a new ListPacketSizesNoCredentialsV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListPacketSizesNoCredentialsV2ParamsWithContext(ctx context.Context) *ListPacketSizesNoCredentialsV2Params {
	var ()
	return &ListPacketSizesNoCredentialsV2Params{

		Context: ctx,
	}
}

// NewListPacketSizesNoCredentialsV2ParamsWithHTTPClient creates a new ListPacketSizesNoCredentialsV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPacketSizesNoCredentialsV2ParamsWithHTTPClient(client *http.Client) *ListPacketSizesNoCredentialsV2Params {
	var ()
	return &ListPacketSizesNoCredentialsV2Params{
		HTTPClient: client,
	}
}

/*ListPacketSizesNoCredentialsV2Params contains all the parameters to send to the API endpoint
for the list packet sizes no credentials v2 operation typically these are written to a http.Request
*/
type ListPacketSizesNoCredentialsV2Params struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) WithTimeout(timeout time.Duration) *ListPacketSizesNoCredentialsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) WithContext(ctx context.Context) *ListPacketSizesNoCredentialsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) WithHTTPClient(client *http.Client) *ListPacketSizesNoCredentialsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) WithClusterID(clusterID string) *ListPacketSizesNoCredentialsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) WithProjectID(projectID string) *ListPacketSizesNoCredentialsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list packet sizes no credentials v2 params
func (o *ListPacketSizesNoCredentialsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPacketSizesNoCredentialsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}