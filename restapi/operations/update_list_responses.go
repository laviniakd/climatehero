// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "climatehero/models"
)

// UpdateListOKCode is the HTTP code returned for type UpdateListOK
const UpdateListOKCode int = 200

/*UpdateListOK success

swagger:response updateListOK
*/
type UpdateListOK struct {
}

// NewUpdateListOK creates UpdateListOK with default headers values
func NewUpdateListOK() *UpdateListOK {

	return &UpdateListOK{}
}

// WriteResponse to the client
func (o *UpdateListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*UpdateListDefault unexpected error

swagger:response updateListDefault
*/
type UpdateListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateListDefault creates UpdateListDefault with default headers values
func NewUpdateListDefault(code int) *UpdateListDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update list default response
func (o *UpdateListDefault) WithStatusCode(code int) *UpdateListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update list default response
func (o *UpdateListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update list default response
func (o *UpdateListDefault) WithPayload(payload *models.Error) *UpdateListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update list default response
func (o *UpdateListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
