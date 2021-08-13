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

package ipam

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

// NewIpamVlanGroupsPartialUpdateParams creates a new IpamVlanGroupsPartialUpdateParams object
// with the default values initialized.
func NewIpamVlanGroupsPartialUpdateParams() *IpamVlanGroupsPartialUpdateParams {
	var ()
	return &IpamVlanGroupsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsPartialUpdateParamsWithTimeout creates a new IpamVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamVlanGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsPartialUpdateParams {
	var ()
	return &IpamVlanGroupsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIpamVlanGroupsPartialUpdateParamsWithContext creates a new IpamVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamVlanGroupsPartialUpdateParamsWithContext(ctx context.Context) *IpamVlanGroupsPartialUpdateParams {
	var ()
	return &IpamVlanGroupsPartialUpdateParams{

		Context: ctx,
	}
}

// NewIpamVlanGroupsPartialUpdateParamsWithHTTPClient creates a new IpamVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamVlanGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsPartialUpdateParams {
	var ()
	return &IpamVlanGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IpamVlanGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam vlan groups partial update operation typically these are written to a http.Request
*/
type IpamVlanGroupsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableVLANGroup
	/*ID
	  A unique integer value identifying this VLAN group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) WithContext(ctx context.Context) *IpamVlanGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) WithData(data *models.WritableVLANGroup) *IpamVlanGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WithID adds the id to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) WithID(id int64) *IpamVlanGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups partial update params
func (o *IpamVlanGroupsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
