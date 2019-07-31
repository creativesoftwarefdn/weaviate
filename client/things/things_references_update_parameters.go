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

package things

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

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewThingsReferencesUpdateParams creates a new ThingsReferencesUpdateParams object
// with the default values initialized.
func NewThingsReferencesUpdateParams() *ThingsReferencesUpdateParams {
	var ()
	return &ThingsReferencesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsReferencesUpdateParamsWithTimeout creates a new ThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsReferencesUpdateParamsWithTimeout(timeout time.Duration) *ThingsReferencesUpdateParams {
	var ()
	return &ThingsReferencesUpdateParams{

		timeout: timeout,
	}
}

// NewThingsReferencesUpdateParamsWithContext creates a new ThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsReferencesUpdateParamsWithContext(ctx context.Context) *ThingsReferencesUpdateParams {
	var ()
	return &ThingsReferencesUpdateParams{

		Context: ctx,
	}
}

// NewThingsReferencesUpdateParamsWithHTTPClient creates a new ThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsReferencesUpdateParamsWithHTTPClient(client *http.Client) *ThingsReferencesUpdateParams {
	var ()
	return &ThingsReferencesUpdateParams{
		HTTPClient: client,
	}
}

/*ThingsReferencesUpdateParams contains all the parameters to send to the API endpoint
for the things references update operation typically these are written to a http.Request
*/
type ThingsReferencesUpdateParams struct {

	/*Body*/
	Body models.MultipleRef
	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Thing.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things references update params
func (o *ThingsReferencesUpdateParams) WithTimeout(timeout time.Duration) *ThingsReferencesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things references update params
func (o *ThingsReferencesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things references update params
func (o *ThingsReferencesUpdateParams) WithContext(ctx context.Context) *ThingsReferencesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things references update params
func (o *ThingsReferencesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things references update params
func (o *ThingsReferencesUpdateParams) WithHTTPClient(client *http.Client) *ThingsReferencesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things references update params
func (o *ThingsReferencesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the things references update params
func (o *ThingsReferencesUpdateParams) WithBody(body models.MultipleRef) *ThingsReferencesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the things references update params
func (o *ThingsReferencesUpdateParams) SetBody(body models.MultipleRef) {
	o.Body = body
}

// WithID adds the id to the things references update params
func (o *ThingsReferencesUpdateParams) WithID(id strfmt.UUID) *ThingsReferencesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the things references update params
func (o *ThingsReferencesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the things references update params
func (o *ThingsReferencesUpdateParams) WithPropertyName(propertyName string) *ThingsReferencesUpdateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the things references update params
func (o *ThingsReferencesUpdateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsReferencesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}