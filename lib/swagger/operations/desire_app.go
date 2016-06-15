package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DesireAppHandlerFunc turns a function with the right signature into a desire app handler
type DesireAppHandlerFunc func(DesireAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DesireAppHandlerFunc) Handle(params DesireAppParams) middleware.Responder {
	return fn(params)
}

// DesireAppHandler interface for that can handle valid desire app params
type DesireAppHandler interface {
	Handle(DesireAppParams) middleware.Responder
}

// NewDesireApp creates a new http.Handler for the desire app operation
func NewDesireApp(ctx *middleware.Context, handler DesireAppHandler) *DesireApp {
	return &DesireApp{Context: ctx, Handler: handler}
}

/*DesireApp swagger:route PUT /apps/{process_guid} desireApp

DesireApp desire app API

*/
type DesireApp struct {
	Context *middleware.Context
	Handler DesireAppHandler
}

func (o *DesireApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDesireAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}