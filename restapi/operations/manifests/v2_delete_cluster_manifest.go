// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2DeleteClusterManifestHandlerFunc turns a function with the right signature into a v2 delete cluster manifest handler
type V2DeleteClusterManifestHandlerFunc func(V2DeleteClusterManifestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2DeleteClusterManifestHandlerFunc) Handle(params V2DeleteClusterManifestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2DeleteClusterManifestHandler interface for that can handle valid v2 delete cluster manifest params
type V2DeleteClusterManifestHandler interface {
	Handle(V2DeleteClusterManifestParams, interface{}) middleware.Responder
}

// NewV2DeleteClusterManifest creates a new http.Handler for the v2 delete cluster manifest operation
func NewV2DeleteClusterManifest(ctx *middleware.Context, handler V2DeleteClusterManifestHandler) *V2DeleteClusterManifest {
	return &V2DeleteClusterManifest{Context: ctx, Handler: handler}
}

/*V2DeleteClusterManifest swagger:route DELETE /v2/clusters/{cluster_id}/manifests manifests v2DeleteClusterManifest

Deletes a manifest from the cluster.

*/
type V2DeleteClusterManifest struct {
	Context *middleware.Context
	Handler V2DeleteClusterManifestHandler
}

func (o *V2DeleteClusterManifest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV2DeleteClusterManifestParams()

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
