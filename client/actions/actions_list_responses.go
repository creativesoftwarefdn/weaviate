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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsListReader is a Reader for the ActionsList structure.
type ActionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewActionsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActionsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewActionsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewActionsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsListOK creates a ActionsListOK with default headers values
func NewActionsListOK() *ActionsListOK {
	return &ActionsListOK{}
}

/*ActionsListOK handles this case with default header values.

Successful response.
*/
type ActionsListOK struct {
	Payload *models.ActionsListResponse
}

func (o *ActionsListOK) Error() string {
	return fmt.Sprintf("[GET /actions][%d] actionsListOK  %+v", 200, o.Payload)
}

func (o *ActionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsListUnauthorized creates a ActionsListUnauthorized with default headers values
func NewActionsListUnauthorized() *ActionsListUnauthorized {
	return &ActionsListUnauthorized{}
}

/*ActionsListUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsListUnauthorized struct {
}

func (o *ActionsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions][%d] actionsListUnauthorized ", 401)
}

func (o *ActionsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsListForbidden creates a ActionsListForbidden with default headers values
func NewActionsListForbidden() *ActionsListForbidden {
	return &ActionsListForbidden{}
}

/*ActionsListForbidden handles this case with default header values.

Forbidden
*/
type ActionsListForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsListForbidden) Error() string {
	return fmt.Sprintf("[GET /actions][%d] actionsListForbidden  %+v", 403, o.Payload)
}

func (o *ActionsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsListNotFound creates a ActionsListNotFound with default headers values
func NewActionsListNotFound() *ActionsListNotFound {
	return &ActionsListNotFound{}
}

/*ActionsListNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ActionsListNotFound struct {
}

func (o *ActionsListNotFound) Error() string {
	return fmt.Sprintf("[GET /actions][%d] actionsListNotFound ", 404)
}

func (o *ActionsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsListInternalServerError creates a ActionsListInternalServerError with default headers values
func NewActionsListInternalServerError() *ActionsListInternalServerError {
	return &ActionsListInternalServerError{}
}

/*ActionsListInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /actions][%d] actionsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}