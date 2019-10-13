// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "climatehero/models"
)

// CheckItemOKCode is the HTTP code returned for type CheckItemOK
const CheckItemOKCode int = 200

/*CheckItemOK success

swagger:response checkItemOK
*/
type CheckItemOK struct {

	/*
	  In: Body
	*/
	Payload *models.CheckReturn `json:"body,omitempty"`
}

// NewCheckItemOK creates CheckItemOK with default headers values
func NewCheckItemOK() *CheckItemOK {

	return &CheckItemOK{}
}

// WithPayload adds the payload to the check item o k response
func (o *CheckItemOK) WithPayload(payload *models.CheckReturn) *CheckItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check item o k response
func (o *CheckItemOK) SetPayload(payload *models.CheckReturn) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CheckItemDefault unexpected error

swagger:response checkItemDefault
*/
type CheckItemDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckItemDefault creates CheckItemDefault with default headers values
func NewCheckItemDefault(code int) *CheckItemDefault {
	if code <= 0 {
		code = 500
	}

	return &CheckItemDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the check item default response
func (o *CheckItemDefault) WithStatusCode(code int) *CheckItemDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the check item default response
func (o *CheckItemDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the check item default response
func (o *CheckItemDefault) WithPayload(payload *models.Error) *CheckItemDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check item default response
func (o *CheckItemDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckItemDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
