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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nzions/not-netbox/netbox/models"
)

// IpamIPAddressesUpdateReader is a Reader for the IpamIPAddressesUpdate structure.
type IpamIPAddressesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesUpdateOK creates a IpamIPAddressesUpdateOK with default headers values
func NewIpamIPAddressesUpdateOK() *IpamIPAddressesUpdateOK {
	return &IpamIPAddressesUpdateOK{}
}

/*IpamIPAddressesUpdateOK handles this case with default header values.

IpamIPAddressesUpdateOK ipam Ip addresses update o k
*/
type IpamIPAddressesUpdateOK struct {
	Payload *models.IPAddress
}

func (o *IpamIPAddressesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/{id}/][%d] ipamIpAddressesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesUpdateDefault creates a IpamIPAddressesUpdateDefault with default headers values
func NewIpamIPAddressesUpdateDefault(code int) *IpamIPAddressesUpdateDefault {
	return &IpamIPAddressesUpdateDefault{
		_statusCode: code,
	}
}

/*IpamIPAddressesUpdateDefault handles this case with default header values.

IpamIPAddressesUpdateDefault ipam ip addresses update default
*/
type IpamIPAddressesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses update default response
func (o *IpamIPAddressesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPAddressesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
