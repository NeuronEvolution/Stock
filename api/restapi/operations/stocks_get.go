// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StocksGetHandlerFunc turns a function with the right signature into a stocks get handler
type StocksGetHandlerFunc func(StocksGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StocksGetHandlerFunc) Handle(params StocksGetParams) middleware.Responder {
	return fn(params)
}

// StocksGetHandler interface for that can handle valid stocks get params
type StocksGetHandler interface {
	Handle(StocksGetParams) middleware.Responder
}

// NewStocksGet creates a new http.Handler for the stocks get operation
func NewStocksGet(ctx *middleware.Context, handler StocksGetHandler) *StocksGet {
	return &StocksGet{Context: ctx, Handler: handler}
}

/*StocksGet swagger:route GET /stocks/{stockId} stocksGet

Get stock,sh_{code},sz_{code}

*/
type StocksGet struct {
	Context *middleware.Context
	Handler StocksGetHandler
}

func (o *StocksGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStocksGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}