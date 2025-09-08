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
// RequestTypeApprovalSettingsMutable - There are two `approvalSettings` types. NONE: this request type doesn't need any approval. SERIAL: this request type needs at least one approval.
type RequestTypeApprovalSettingsMutable struct {
	RequestTypeApprovalSettingsNone           *RequestTypeApprovalSettingsNone
	RequestTypeApprovalSettingsSerialWritable *RequestTypeApprovalSettingsSerialWritable
}

// RequestTypeApprovalSettingsNoneAsRequestTypeApprovalSettingsMutable is a convenience function that returns RequestTypeApprovalSettingsNone wrapped in RequestTypeApprovalSettingsMutable
func RequestTypeApprovalSettingsNoneAsRequestTypeApprovalSettingsMutable(v *RequestTypeApprovalSettingsNone) RequestTypeApprovalSettingsMutable {
	return RequestTypeApprovalSettingsMutable{
		RequestTypeApprovalSettingsNone: v,
	}
}

// RequestTypeApprovalSettingsSerialWritableAsRequestTypeApprovalSettingsMutable is a convenience function that returns RequestTypeApprovalSettingsSerialWritable wrapped in RequestTypeApprovalSettingsMutable
func RequestTypeApprovalSettingsSerialWritableAsRequestTypeApprovalSettingsMutable(v *RequestTypeApprovalSettingsSerialWritable) RequestTypeApprovalSettingsMutable {
	return RequestTypeApprovalSettingsMutable{
		RequestTypeApprovalSettingsSerialWritable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestTypeApprovalSettingsMutable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'NONE'
	if jsonDict["type"] == "NONE" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsNone
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsNone)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsNone, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsNone = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsMutable as RequestTypeApprovalSettingsNone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERIAL'
	if jsonDict["type"] == "SERIAL" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsSerialWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsSerialWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsSerialWritable, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsSerialWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsMutable as RequestTypeApprovalSettingsSerialWritable: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsMutable as RequestTypeApprovalSettingsNone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-settings-serial-writable'
	if jsonDict["type"] == "request-type-approval-settings-serial-writable" {
		// try to unmarshal JSON data into RequestTypeApprovalSettingsSerialWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalSettingsSerialWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalSettingsSerialWritable, return on the first match
		} else {
			dst.RequestTypeApprovalSettingsSerialWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalSettingsMutable as RequestTypeApprovalSettingsSerialWritable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeApprovalSettingsMutable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeApprovalSettingsNone != nil {
		return json.Marshal(&src.RequestTypeApprovalSettingsNone)
	}

	if src.RequestTypeApprovalSettingsSerialWritable != nil {
		return json.Marshal(&src.RequestTypeApprovalSettingsSerialWritable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeApprovalSettingsMutable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeApprovalSettingsNone != nil {
		return obj.RequestTypeApprovalSettingsNone
	}

	if obj.RequestTypeApprovalSettingsSerialWritable != nil {
		return obj.RequestTypeApprovalSettingsSerialWritable
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeApprovalSettingsMutable struct {
	value *RequestTypeApprovalSettingsMutable
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsMutable) Get() *RequestTypeApprovalSettingsMutable {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsMutable) Set(val *RequestTypeApprovalSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsMutable(val *RequestTypeApprovalSettingsMutable) *NullableRequestTypeApprovalSettingsMutable {
	return &NullableRequestTypeApprovalSettingsMutable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
