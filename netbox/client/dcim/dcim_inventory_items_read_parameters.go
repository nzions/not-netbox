// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimInventoryItemsReadParams creates a new DcimInventoryItemsReadParams object
// with the default values initialized.
func NewDcimInventoryItemsReadParams() *DcimInventoryItemsReadParams {
	var ()
	return &DcimInventoryItemsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsReadParamsWithTimeout creates a new DcimInventoryItemsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInventoryItemsReadParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsReadParams {
	var ()
	return &DcimInventoryItemsReadParams{

		timeout: timeout,
	}
}

// NewDcimInventoryItemsReadParamsWithContext creates a new DcimInventoryItemsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInventoryItemsReadParamsWithContext(ctx context.Context) *DcimInventoryItemsReadParams {
	var ()
	return &DcimInventoryItemsReadParams{

		Context: ctx,
	}
}

// NewDcimInventoryItemsReadParamsWithHTTPClient creates a new DcimInventoryItemsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInventoryItemsReadParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsReadParams {
	var ()
	return &DcimInventoryItemsReadParams{
		HTTPClient: client,
	}
}

/*DcimInventoryItemsReadParams contains all the parameters to send to the API endpoint
for the dcim inventory items read operation typically these are written to a http.Request
*/
type DcimInventoryItemsReadParams struct {

	/*ID
	  A unique integer value identifying this inventory item.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithContext(ctx context.Context) *DcimInventoryItemsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) WithID(id int64) *DcimInventoryItemsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items read params
func (o *DcimInventoryItemsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
