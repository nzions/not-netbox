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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nzions/not-netbox/netbox/models"
)

// VirtualizationClusterTypesPartialUpdateReader is a Reader for the VirtualizationClusterTypesPartialUpdate structure.
type VirtualizationClusterTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesPartialUpdateOK creates a VirtualizationClusterTypesPartialUpdateOK with default headers values
func NewVirtualizationClusterTypesPartialUpdateOK() *VirtualizationClusterTypesPartialUpdateOK {
	return &VirtualizationClusterTypesPartialUpdateOK{}
}

/*VirtualizationClusterTypesPartialUpdateOK handles this case with default header values.

VirtualizationClusterTypesPartialUpdateOK virtualization cluster types partial update o k
*/
type VirtualizationClusterTypesPartialUpdateOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesPartialUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesPartialUpdateDefault creates a VirtualizationClusterTypesPartialUpdateDefault with default headers values
func NewVirtualizationClusterTypesPartialUpdateDefault(code int) *VirtualizationClusterTypesPartialUpdateDefault {
	return &VirtualizationClusterTypesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationClusterTypesPartialUpdateDefault handles this case with default header values.

VirtualizationClusterTypesPartialUpdateDefault virtualization cluster types partial update default
*/
type VirtualizationClusterTypesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types partial update default response
func (o *VirtualizationClusterTypesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
