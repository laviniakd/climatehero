// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// UpdateListHandlerFunc turns a function with the right signature into a update list handler
type UpdateListHandlerFunc func(UpdateListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateListHandlerFunc) Handle(params UpdateListParams) middleware.Responder {
	return fn(params)
}

// UpdateListHandler interface for that can handle valid update list params
type UpdateListHandler interface {
	Handle(UpdateListParams) middleware.Responder
}

// NewUpdateList creates a new http.Handler for the update list operation
func NewUpdateList(ctx *middleware.Context, handler UpdateListHandler) *UpdateList {
	return &UpdateList{Context: ctx, Handler: handler}
}

/*UpdateList swagger:route POST /updateList updateList

updates a user list

*/
type UpdateList struct {
	Context *middleware.Context
	Handler UpdateListHandler
}

func (o *UpdateList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateListBody update list body
// swagger:model UpdateListBody
type UpdateListBody struct {

	// email
	Email string `json:"email,omitempty"`

	// list items
	ListItems []string `json:"listItems"`
}

// Validate validates this update list body
func (o *UpdateListBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateListBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateListBody) UnmarshalBinary(b []byte) error {
	var res UpdateListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
