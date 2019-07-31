//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBatchingActionsCreateParams creates a new BatchingActionsCreateParams object
// with the default values initialized.
func NewBatchingActionsCreateParams() *BatchingActionsCreateParams {
	var ()
	return &BatchingActionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchingActionsCreateParamsWithTimeout creates a new BatchingActionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchingActionsCreateParamsWithTimeout(timeout time.Duration) *BatchingActionsCreateParams {
	var ()
	return &BatchingActionsCreateParams{

		timeout: timeout,
	}
}

// NewBatchingActionsCreateParamsWithContext creates a new BatchingActionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchingActionsCreateParamsWithContext(ctx context.Context) *BatchingActionsCreateParams {
	var ()
	return &BatchingActionsCreateParams{

		Context: ctx,
	}
}

// NewBatchingActionsCreateParamsWithHTTPClient creates a new BatchingActionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchingActionsCreateParamsWithHTTPClient(client *http.Client) *BatchingActionsCreateParams {
	var ()
	return &BatchingActionsCreateParams{
		HTTPClient: client,
	}
}

/*BatchingActionsCreateParams contains all the parameters to send to the API endpoint
for the batching actions create operation typically these are written to a http.Request
*/
type BatchingActionsCreateParams struct {

	/*Body*/
	Body BatchingActionsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batching actions create params
func (o *BatchingActionsCreateParams) WithTimeout(timeout time.Duration) *BatchingActionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batching actions create params
func (o *BatchingActionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batching actions create params
func (o *BatchingActionsCreateParams) WithContext(ctx context.Context) *BatchingActionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batching actions create params
func (o *BatchingActionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batching actions create params
func (o *BatchingActionsCreateParams) WithHTTPClient(client *http.Client) *BatchingActionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batching actions create params
func (o *BatchingActionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batching actions create params
func (o *BatchingActionsCreateParams) WithBody(body BatchingActionsCreateBody) *BatchingActionsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batching actions create params
func (o *BatchingActionsCreateParams) SetBody(body BatchingActionsCreateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchingActionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}