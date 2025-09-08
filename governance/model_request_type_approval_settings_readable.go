/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance

import (
	"encoding/json"
	"fmt"
)

// model_oneof.mustache
// RequestTypeApprovalSettingsReadable - struct for RequestTypeApprovalSettingsReadable
type RequestTypeApprovalSettingsReadable struct {
	RequestTypeApprovalSettingsCustom *RequestTypeApprovalSettingsCustom
	RequestTypeApprovalSettingsNone   *RequestTypeApprovalSettingsNone
	RequestTypeApprovalSettingsSerial *RequestTypeApprovalSettingsSerial
}

// RequestTypeApprovalSettingsCustomAsRequestTypeApprovalSettingsReadable is a convenience function that returns RequestTypeApprovalSettingsCustom wrapped in RequestTypeApprovalSettingsReadable
func RequestTypeApprovalSettingsCustomAsRequestTypeApprovalSettingsReadable(v *RequestTypeApprovalSettingsCustom) RequestTypeApprovalSettingsReadable {
	return RequestTypeApprovalSettingsReadable{
		RequestTypeApprovalSettingsCustom: v,
	}
}

// RequestTypeApprovalSettingsNoneAsRequestTypeApprovalSettingsReadable is a convenience function that returns RequestTypeApprovalSettingsNone wrapped in RequestTypeApprovalSettingsReadable
func RequestTypeApprovalSettingsNoneAsRequestTypeApprovalSettingsReadable(v *RequestTypeApprovalSettingsNone) RequestTypeApprovalSettingsReadable {
	return RequestTypeApprovalSettingsReadable{
		RequestTypeApprovalSettingsNone: v,
	}
}

// RequestTypeApprovalSettingsSerialAsRequestTypeApprovalSettingsReadable is a convenience function that returns RequestTypeApprovalSettingsSerial wrapped in RequestTypeApprovalSettingsReadable
func RequestTypeApprovalSettingsSerialAsRequestTypeApprovalSettingsReadable(v *RequestTypeApprovalSettingsSerial) RequestTypeApprovalSettingsReadable {
	return RequestTypeApprovalSettingsReadable{
		RequestTypeApprovalSettingsSerial: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestTypeApprovalSettingsReadable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsCustom
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsCustom)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsCustom, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsCustom = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsCustom: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NONE'
	if jsonDict["type"] == "NONE" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsNone
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsNone)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsNone, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsNone = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsNone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERIAL'
	if jsonDict["type"] == "SERIAL" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsSerial
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsSerial)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsSerial, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsSerial = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsSerial: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-settings-custom'
	if jsonDict["type"] == "request-type-approval-settings-custom" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsCustom
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsCustom)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsCustom, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsCustom = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsCustom: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-settings-none'
	if jsonDict["type"] == "request-type-approval-settings-none" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsNone
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsNone)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsNone, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsNone = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsNone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-settings-serial'
	if jsonDict["type"] == "request-type-approval-settings-serial" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsSerial
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsSerial)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsSerial, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsSerial = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsReadable as RequestTypeApprovalSettingsSerial: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeApprovalSettingsReadable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeApprovalSettingsCustom != nil {
		return json.Marshal(&src.RequestTypeApprovalSettingsCustom)
	}

	if src.RequestTypeApprovalSettingsNone != nil {
		return json.Marshal(&src.RequestTypeApprovalSettingsNone)
	}

	if src.RequestTypeApprovalSettingsSerial != nil {
		return json.Marshal(&src.RequestTypeApprovalSettingsSerial)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeApprovalSettingsReadable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeApprovalSettingsCustom != nil {
		return obj.RequestTypeApprovalSettingsCustom
	}

	if obj.RequestTypeApprovalSettingsNone != nil {
		return obj.RequestTypeApprovalSettingsNone
	}

	if obj.RequestTypeApprovalSettingsSerial != nil {
		return obj.RequestTypeApprovalSettingsSerial
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeApprovalSettingsReadable struct {
	value *RequestTypeApprovalSettingsReadable
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsReadable) Get() *RequestTypeApprovalSettingsReadable {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsReadable) Set(val *RequestTypeApprovalSettingsReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsReadable(val *RequestTypeApprovalSettingsReadable) *NullableRequestTypeApprovalSettingsReadable {
	return &NullableRequestTypeApprovalSettingsReadable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
