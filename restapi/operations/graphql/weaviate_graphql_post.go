/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateGraphqlPostHandlerFunc turns a function with the right signature into a weaviate graphql post handler
type WeaviateGraphqlPostHandlerFunc func(WeaviateGraphqlPostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGraphqlPostHandlerFunc) Handle(params WeaviateGraphqlPostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGraphqlPostHandler interface for that can handle valid weaviate graphql post params
type WeaviateGraphqlPostHandler interface {
	Handle(WeaviateGraphqlPostParams, interface{}) middleware.Responder
}

// NewWeaviateGraphqlPost creates a new http.Handler for the weaviate graphql post operation
func NewWeaviateGraphqlPost(ctx *middleware.Context, handler WeaviateGraphqlPostHandler) *WeaviateGraphqlPost {
	return &WeaviateGraphqlPost{Context: ctx, Handler: handler}
}

/*WeaviateGraphqlPost swagger:route POST /graphql graphql weaviateGraphqlPost

Get a response based on GraphQL

Query Weaviate with GraphQL

*/
type WeaviateGraphqlPost struct {
	Context *middleware.Context
	Handler WeaviateGraphqlPostHandler
}

func (o *WeaviateGraphqlPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateGraphqlPostParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
