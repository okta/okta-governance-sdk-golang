/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

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

// RequestTypeRequestSettingsMutable - The `requestSettings` control:  - Which users can request access using this request type. - What fields the user may or must fill out in order to request access.
type RequestTypeRequestSettingsMutable struct {
	RequestTypeRequesterEveryoneWritable *RequestTypeRequesterEveryoneWritable
	RequestTypeRequesterMemberOfWritable *RequestTypeRequesterMemberOfWritable
}

// RequestTypeRequesterEveryoneWritableAsRequestTypeRequestSettingsMutable is a convenience function that returns RequestTypeRequesterEveryoneWritable wrapped in RequestTypeRequestSettingsMutable
func RequestTypeRequesterEveryoneWritableAsRequestTypeRequestSettingsMutable(v *RequestTypeRequesterEveryoneWritable) RequestTypeRequestSettingsMutable {
	return RequestTypeRequestSettingsMutable{
		RequestTypeRequesterEveryoneWritable: v,
	}
}

// RequestTypeRequesterMemberOfWritableAsRequestTypeRequestSettingsMutable is a convenience function that returns RequestTypeRequesterMemberOfWritable wrapped in RequestTypeRequestSettingsMutable
func RequestTypeRequesterMemberOfWritableAsRequestTypeRequestSettingsMutable(v *RequestTypeRequesterMemberOfWritable) RequestTypeRequestSettingsMutable {
	return RequestTypeRequestSettingsMutable{
		RequestTypeRequesterMemberOfWritable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTypeRequestSettingsMutable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVERYONE'
	if jsonDict["type"] == "EVERYONE" {
		// try to unmarshal JSON data into RequestTypeRequesterEveryoneWritable
		err = json.Unmarshal(data, &dst.RequestTypeRequesterEveryoneWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterEveryoneWritable, return on the first match
		} else {
			dst.RequestTypeRequesterEveryoneWritable = nil
			return fmt.Errorf("failed to unmarshal RequestTypeRequestSettingsMutable as RequestTypeRequesterEveryoneWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MEMBER_OF'
	if jsonDict["type"] == "MEMBER_OF" {
		// try to unmarshal JSON data into RequestTypeRequesterMemberOfWritable
		err = json.Unmarshal(data, &dst.RequestTypeRequesterMemberOfWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterMemberOfWritable, return on the first match
		} else {
			dst.RequestTypeRequesterMemberOfWritable = nil
			return fmt.Errorf("failed to unmarshal RequestTypeRequestSettingsMutable as RequestTypeRequesterMemberOfWritable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeRequestSettingsMutable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeRequesterEveryoneWritable != nil {
		return json.Marshal(&src.RequestTypeRequesterEveryoneWritable)
	}

	if src.RequestTypeRequesterMemberOfWritable != nil {
		return json.Marshal(&src.RequestTypeRequesterMemberOfWritable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeRequestSettingsMutable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeRequesterEveryoneWritable != nil {
		return obj.RequestTypeRequesterEveryoneWritable
	}

	if obj.RequestTypeRequesterMemberOfWritable != nil {
		return obj.RequestTypeRequesterMemberOfWritable
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestTypeRequestSettingsMutable) GetActualInstanceValue() interface{} {
	if obj.RequestTypeRequesterEveryoneWritable != nil {
		return *obj.RequestTypeRequesterEveryoneWritable
	}

	if obj.RequestTypeRequesterMemberOfWritable != nil {
		return *obj.RequestTypeRequesterMemberOfWritable
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeRequestSettingsMutable struct {
	value *RequestTypeRequestSettingsMutable
	isSet bool
}

func (v NullableRequestTypeRequestSettingsMutable) Get() *RequestTypeRequestSettingsMutable {
	return v.value
}

func (v *NullableRequestTypeRequestSettingsMutable) Set(val *RequestTypeRequestSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeRequestSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeRequestSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeRequestSettingsMutable(val *RequestTypeRequestSettingsMutable) *NullableRequestTypeRequestSettingsMutable {
	return &NullableRequestTypeRequestSettingsMutable{value: val, isSet: true}
}

func (v NullableRequestTypeRequestSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeRequestSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
