// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStocksListParams creates a new StocksListParams object
// with the default values initialized.
func NewStocksListParams() *StocksListParams {
	var ()
	return &StocksListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStocksListParamsWithTimeout creates a new StocksListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStocksListParamsWithTimeout(timeout time.Duration) *StocksListParams {
	var ()
	return &StocksListParams{

		timeout: timeout,
	}
}

// NewStocksListParamsWithContext creates a new StocksListParams object
// with the default values initialized, and the ability to set a context for a request
func NewStocksListParamsWithContext(ctx context.Context) *StocksListParams {
	var ()
	return &StocksListParams{

		Context: ctx,
	}
}

// NewStocksListParamsWithHTTPClient creates a new StocksListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStocksListParamsWithHTTPClient(client *http.Client) *StocksListParams {
	var ()
	return &StocksListParams{
		HTTPClient: client,
	}
}

/*StocksListParams contains all the parameters to send to the API endpoint
for the stocks list operation typically these are written to a http.Request
*/
type StocksListParams struct {

	/*ExchangeID
	  Exchange id.eg sz,sh...

	*/
	ExchangeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stocks list params
func (o *StocksListParams) WithTimeout(timeout time.Duration) *StocksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stocks list params
func (o *StocksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stocks list params
func (o *StocksListParams) WithContext(ctx context.Context) *StocksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stocks list params
func (o *StocksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stocks list params
func (o *StocksListParams) WithHTTPClient(client *http.Client) *StocksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stocks list params
func (o *StocksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExchangeID adds the exchangeID to the stocks list params
func (o *StocksListParams) WithExchangeID(exchangeID *string) *StocksListParams {
	o.SetExchangeID(exchangeID)
	return o
}

// SetExchangeID adds the exchangeId to the stocks list params
func (o *StocksListParams) SetExchangeID(exchangeID *string) {
	o.ExchangeID = exchangeID
}

// WriteToRequest writes these params to a swagger request
func (o *StocksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExchangeID != nil {

		// query param exchangeId
		var qrExchangeID string
		if o.ExchangeID != nil {
			qrExchangeID = *o.ExchangeID
		}
		qExchangeID := qrExchangeID
		if qExchangeID != "" {
			if err := r.SetQueryParam("exchangeId", qExchangeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}