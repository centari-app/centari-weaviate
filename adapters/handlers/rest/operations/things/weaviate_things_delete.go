/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateThingsDeleteHandlerFunc turns a function with the right signature into a weaviate things delete handler
type WeaviateThingsDeleteHandlerFunc func(WeaviateThingsDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsDeleteHandlerFunc) Handle(params WeaviateThingsDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsDeleteHandler interface for that can handle valid weaviate things delete params
type WeaviateThingsDeleteHandler interface {
	Handle(WeaviateThingsDeleteParams, *models.Principal) middleware.Responder
}

// NewWeaviateThingsDelete creates a new http.Handler for the weaviate things delete operation
func NewWeaviateThingsDelete(ctx *middleware.Context, handler WeaviateThingsDeleteHandler) *WeaviateThingsDelete {
	return &WeaviateThingsDelete{Context: ctx, Handler: handler}
}

/*WeaviateThingsDelete swagger:route DELETE /things/{id} things weaviateThingsDelete

Delete a Thing based on its UUID.

Deletes a Thing from the system. All Actions pointing to this Thing, where the Thing is the object of the Action, are also being deleted.

*/
type WeaviateThingsDelete struct {
	Context *middleware.Context
	Handler WeaviateThingsDeleteHandler
}

func (o *WeaviateThingsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}