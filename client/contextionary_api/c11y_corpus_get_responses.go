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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// C11yCorpusGetReader is a Reader for the C11yCorpusGet structure.
type C11yCorpusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *C11yCorpusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 501:
		result := NewC11yCorpusGetNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewC11yCorpusGetNotImplemented creates a C11yCorpusGetNotImplemented with default headers values
func NewC11yCorpusGetNotImplemented() *C11yCorpusGetNotImplemented {
	return &C11yCorpusGetNotImplemented{}
}

/*C11yCorpusGetNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type C11yCorpusGetNotImplemented struct {
}

func (o *C11yCorpusGetNotImplemented) Error() string {
	return fmt.Sprintf("[POST /c11y/corpus][%d] c11yCorpusGetNotImplemented ", 501)
}

func (o *C11yCorpusGetNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*C11yCorpusGetBody The text corpus.
swagger:model C11yCorpusGetBody
*/
type C11yCorpusGetBody struct {

	// corpus
	Corpus string `json:"corpus,omitempty"`
}

// Validate validates this c11y corpus get body
func (o *C11yCorpusGetBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *C11yCorpusGetBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *C11yCorpusGetBody) UnmarshalBinary(b []byte) error {
	var res C11yCorpusGetBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}