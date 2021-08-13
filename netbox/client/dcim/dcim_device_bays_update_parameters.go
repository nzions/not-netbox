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

	"github.com/nzions/not-netbox/netbox/models"
)

// NewDcimDeviceBaysUpdateParams creates a new DcimDeviceBaysUpdateParams object
// with the default values initialized.
func NewDcimDeviceBaysUpdateParams() *DcimDeviceBaysUpdateParams {
	var ()
	return &DcimDeviceBaysUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysUpdateParamsWithTimeout creates a new DcimDeviceBaysUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBaysUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysUpdateParams {
	var ()
	return &DcimDeviceBaysUpdateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBaysUpdateParamsWithContext creates a new DcimDeviceBaysUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBaysUpdateParamsWithContext(ctx context.Context) *DcimDeviceBaysUpdateParams {
	var ()
	return &DcimDeviceBaysUpdateParams{

		Context: ctx,
	}
}

// NewDcimDeviceBaysUpdateParamsWithHTTPClient creates a new DcimDeviceBaysUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBaysUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysUpdateParams {
	var ()
	return &DcimDeviceBaysUpdateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBaysUpdateParams contains all the parameters to send to the API endpoint
for the dcim device bays update operation typically these are written to a http.Request
*/
type DcimDeviceBaysUpdateParams struct {

	/*Data*/
	Data *models.WritableDeviceBay
	/*ID
	  A unique integer value identifying this device bay.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithContext(ctx context.Context) *DcimDeviceBaysUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithData(data *models.WritableDeviceBay) *DcimDeviceBaysUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetData(data *models.WritableDeviceBay) {
	o.Data = data
}

// WithID adds the id to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithID(id int64) *DcimDeviceBaysUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
