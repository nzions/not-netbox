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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nzions/not-netbox/netbox/models"
)

// SecretsSecretsUpdateReader is a Reader for the SecretsSecretsUpdate structure.
type SecretsSecretsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsSecretsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretsUpdateOK creates a SecretsSecretsUpdateOK with default headers values
func NewSecretsSecretsUpdateOK() *SecretsSecretsUpdateOK {
	return &SecretsSecretsUpdateOK{}
}

/*SecretsSecretsUpdateOK handles this case with default header values.

SecretsSecretsUpdateOK secrets secrets update o k
*/
type SecretsSecretsUpdateOK struct {
	Payload *models.Secret
}

func (o *SecretsSecretsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/secrets/{id}/][%d] secretsSecretsUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretsUpdateOK) GetPayload() *models.Secret {
	return o.Payload
}

func (o *SecretsSecretsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsSecretsUpdateDefault creates a SecretsSecretsUpdateDefault with default headers values
func NewSecretsSecretsUpdateDefault(code int) *SecretsSecretsUpdateDefault {
	return &SecretsSecretsUpdateDefault{
		_statusCode: code,
	}
}

/*SecretsSecretsUpdateDefault handles this case with default header values.

SecretsSecretsUpdateDefault secrets secrets update default
*/
type SecretsSecretsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secrets update default response
func (o *SecretsSecretsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/secrets/{id}/][%d] secrets_secrets_update default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
