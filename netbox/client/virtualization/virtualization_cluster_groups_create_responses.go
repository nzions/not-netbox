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

// VirtualizationClusterGroupsCreateReader is a Reader for the VirtualizationClusterGroupsCreate structure.
type VirtualizationClusterGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClusterGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsCreateCreated creates a VirtualizationClusterGroupsCreateCreated with default headers values
func NewVirtualizationClusterGroupsCreateCreated() *VirtualizationClusterGroupsCreateCreated {
	return &VirtualizationClusterGroupsCreateCreated{}
}

/*VirtualizationClusterGroupsCreateCreated handles this case with default header values.

VirtualizationClusterGroupsCreateCreated virtualization cluster groups create created
*/
type VirtualizationClusterGroupsCreateCreated struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-groups/][%d] virtualizationClusterGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *VirtualizationClusterGroupsCreateCreated) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsCreateDefault creates a VirtualizationClusterGroupsCreateDefault with default headers values
func NewVirtualizationClusterGroupsCreateDefault(code int) *VirtualizationClusterGroupsCreateDefault {
	return &VirtualizationClusterGroupsCreateDefault{
		_statusCode: code,
	}
}

/*VirtualizationClusterGroupsCreateDefault handles this case with default header values.

VirtualizationClusterGroupsCreateDefault virtualization cluster groups create default
*/
type VirtualizationClusterGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster groups create default response
func (o *VirtualizationClusterGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-groups/][%d] virtualization_cluster-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
