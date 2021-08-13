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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nzions/not-netbox/netbox/models"
)

// DcimPlatformsUpdateReader is a Reader for the DcimPlatformsUpdate structure.
type DcimPlatformsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsUpdateOK creates a DcimPlatformsUpdateOK with default headers values
func NewDcimPlatformsUpdateOK() *DcimPlatformsUpdateOK {
	return &DcimPlatformsUpdateOK{}
}

/*DcimPlatformsUpdateOK handles this case with default header values.

DcimPlatformsUpdateOK dcim platforms update o k
*/
type DcimPlatformsUpdateOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/platforms/{id}/][%d] dcimPlatformsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsUpdateDefault creates a DcimPlatformsUpdateDefault with default headers values
func NewDcimPlatformsUpdateDefault(code int) *DcimPlatformsUpdateDefault {
	return &DcimPlatformsUpdateDefault{
		_statusCode: code,
	}
}

/*DcimPlatformsUpdateDefault handles this case with default header values.

DcimPlatformsUpdateDefault dcim platforms update default
*/
type DcimPlatformsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms update default response
func (o *DcimPlatformsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPlatformsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/platforms/{id}/][%d] dcim_platforms_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
