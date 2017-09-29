// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NeuronEvolution/Stock/api/models"
)

// StocksGetOKCode is the HTTP code returned for type StocksGetOK
const StocksGetOKCode int = 200

/*StocksGetOK Stock

swagger:response stocksGetOK
*/
type StocksGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Stock `json:"body,omitempty"`
}

// NewStocksGetOK creates StocksGetOK with default headers values
func NewStocksGetOK() *StocksGetOK {
	return &StocksGetOK{}
}

// WithPayload adds the payload to the stocks get o k response
func (o *StocksGetOK) WithPayload(payload *models.Stock) *StocksGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stocks get o k response
func (o *StocksGetOK) SetPayload(payload *models.Stock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StocksGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StocksGetBadRequestCode is the HTTP code returned for type StocksGetBadRequest
const StocksGetBadRequestCode int = 400

/*StocksGetBadRequest Bad request

swagger:response stocksGetBadRequest
*/
type StocksGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStocksGetBadRequest creates StocksGetBadRequest with default headers values
func NewStocksGetBadRequest() *StocksGetBadRequest {
	return &StocksGetBadRequest{}
}

// WithPayload adds the payload to the stocks get bad request response
func (o *StocksGetBadRequest) WithPayload(payload string) *StocksGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stocks get bad request response
func (o *StocksGetBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StocksGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// StocksGetNotFoundCode is the HTTP code returned for type StocksGetNotFound
const StocksGetNotFoundCode int = 404

/*StocksGetNotFound Not found

swagger:response stocksGetNotFound
*/
type StocksGetNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStocksGetNotFound creates StocksGetNotFound with default headers values
func NewStocksGetNotFound() *StocksGetNotFound {
	return &StocksGetNotFound{}
}

// WithPayload adds the payload to the stocks get not found response
func (o *StocksGetNotFound) WithPayload(payload string) *StocksGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stocks get not found response
func (o *StocksGetNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StocksGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// StocksGetInternalServerErrorCode is the HTTP code returned for type StocksGetInternalServerError
const StocksGetInternalServerErrorCode int = 500

/*StocksGetInternalServerError Internal error

swagger:response stocksGetInternalServerError
*/
type StocksGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStocksGetInternalServerError creates StocksGetInternalServerError with default headers values
func NewStocksGetInternalServerError() *StocksGetInternalServerError {
	return &StocksGetInternalServerError{}
}

// WithPayload adds the payload to the stocks get internal server error response
func (o *StocksGetInternalServerError) WithPayload(payload string) *StocksGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stocks get internal server error response
func (o *StocksGetInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StocksGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
