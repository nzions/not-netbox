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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceWithConfigContext writable device with config context
//
// swagger:model WritableDeviceWithConfigContext
type WritableDeviceWithConfigContext struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Cluster
	Cluster *int64 `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device role
	// Required: true
	DeviceRole *int64 `json:"device_role"`

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Rack face
	// Enum: [front rear]
	Face string `json:"face,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// Name
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// Platform
	Platform *int64 `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	PrimaryIp4 *int64 `json:"primary_ip4,omitempty"`

	// Primary IPv6
	PrimaryIp6 *int64 `json:"primary_ip6,omitempty"`

	// Rack
	Rack *int64 `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Status
	// Enum: [offline active planned staged failed inventory decommissioning]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// Virtual chassis
	VirtualChassis *int64 `json:"virtual_chassis,omitempty"`
}

// Validate validates this writable device with config context
func (m *WritableDeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(*m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeFacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeFacePropEnum = append(writableDeviceWithConfigContextTypeFacePropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextFaceFront captures enum value "front"
	WritableDeviceWithConfigContextFaceFront string = "front"

	// WritableDeviceWithConfigContextFaceRear captures enum value "rear"
	WritableDeviceWithConfigContextFaceRear string = "rear"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateFaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeFacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateFace(formats strfmt.Registry) error {

	if swag.IsZero(m.Face) { // not required
		return nil
	}

	// value enum
	if err := m.validateFaceEnum("face", "body", m.Face); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateParentDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(*m.Position), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", int64(*m.Position), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","inventory","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeStatusPropEnum = append(writableDeviceWithConfigContextTypeStatusPropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextStatusOffline captures enum value "offline"
	WritableDeviceWithConfigContextStatusOffline string = "offline"

	// WritableDeviceWithConfigContextStatusActive captures enum value "active"
	WritableDeviceWithConfigContextStatusActive string = "active"

	// WritableDeviceWithConfigContextStatusPlanned captures enum value "planned"
	WritableDeviceWithConfigContextStatusPlanned string = "planned"

	// WritableDeviceWithConfigContextStatusStaged captures enum value "staged"
	WritableDeviceWithConfigContextStatusStaged string = "staged"

	// WritableDeviceWithConfigContextStatusFailed captures enum value "failed"
	WritableDeviceWithConfigContextStatusFailed string = "failed"

	// WritableDeviceWithConfigContextStatusInventory captures enum value "inventory"
	WritableDeviceWithConfigContextStatusInventory string = "inventory"

	// WritableDeviceWithConfigContextStatusDecommissioning captures enum value "decommissioning"
	WritableDeviceWithConfigContextStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", int64(*m.VcPosition), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", int64(*m.VcPosition), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", int64(*m.VcPriority), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", int64(*m.VcPriority), 255, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableDeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
