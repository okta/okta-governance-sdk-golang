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
// RequestTypeRequestSettingsReadable - The `requestSettings` control:  - Which users can request access using this request type. - What fields the user may or must fill out in order to request access.
type RequestTypeRequestSettingsReadable struct {
	RequestTypeRequesterCustom   *RequestTypeRequesterCustom
	RequestTypeRequesterEveryone *RequestTypeRequesterEveryone
	RequestTypeRequesterMemberOf *RequestTypeRequesterMemberOf
}

// RequestTypeRequesterCustomAsRequestTypeRequestSettingsReadable is a convenience function that returns RequestTypeRequesterCustom wrapped in RequestTypeRequestSettingsReadable
func RequestTypeRequesterCustomAsRequestTypeRequestSettingsReadable(v *RequestTypeRequesterCustom) RequestTypeRequestSettingsReadable {
	return RequestTypeRequestSettingsReadable{
		RequestTypeRequesterCustom: v,
	}
}

// RequestTypeRequesterEveryoneAsRequestTypeRequestSettingsReadable is a convenience function that returns RequestTypeRequesterEveryone wrapped in RequestTypeRequestSettingsReadable
func RequestTypeRequesterEveryoneAsRequestTypeRequestSettingsReadable(v *RequestTypeRequesterEveryone) RequestTypeRequestSettingsReadable {
	return RequestTypeRequestSettingsReadable{
		RequestTypeRequesterEveryone: v,
	}
}

// RequestTypeRequesterMemberOfAsRequestTypeRequestSettingsReadable is a convenience function that returns RequestTypeRequesterMemberOf wrapped in RequestTypeRequestSettingsReadable
func RequestTypeRequesterMemberOfAsRequestTypeRequestSettingsReadable(v *RequestTypeRequesterMemberOf) RequestTypeRequestSettingsReadable {
	return RequestTypeRequestSettingsReadable{
		RequestTypeRequesterMemberOf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestTypeRequestSettingsReadable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into RequestTypeRequesterCustom
		err = json.Unmarshal(data, &dst.RequestTypeRequesterCustom)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterCustom, return on the first match
		} else {
			dst.RequestTypeRequesterCustom = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterCustom: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVERYONE'
	if jsonDict["type"] == "EVERYONE" {
		// try to unmarshal JSON data into RequestTypeRequesterEveryone
		err = json.Unmarshal(data, &dst.RequestTypeRequesterEveryone)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterEveryone, return on the first match
		} else {
			dst.RequestTypeRequesterEveryone = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterEveryone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MEMBER_OF'
	if jsonDict["type"] == "MEMBER_OF" {
		// try to unmarshal JSON data into RequestTypeRequesterMemberOf
		err = json.Unmarshal(data, &dst.RequestTypeRequesterMemberOf)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterMemberOf, return on the first match
		} else {
			dst.RequestTypeRequesterMemberOf = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterMemberOf: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-requester-custom'
	if jsonDict["type"] == "request-type-requester-custom" {
		// try to unmarshal JSON data into RequestTypeRequesterCustom
		err = json.Unmarshal(data, &dst.RequestTypeRequesterCustom)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterCustom, return on the first match
		} else {
			dst.RequestTypeRequesterCustom = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterCustom: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-requester-everyone'
	if jsonDict["type"] == "request-type-requester-everyone" {
		// try to unmarshal JSON data into RequestTypeRequesterEveryone
		err = json.Unmarshal(data, &dst.RequestTypeRequesterEveryone)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterEveryone, return on the first match
		} else {
			dst.RequestTypeRequesterEveryone = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterEveryone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-requester-member-of'
	if jsonDict["type"] == "request-type-requester-member-of" {
		// try to unmarshal JSON data into RequestTypeRequesterMemberOf
		err = json.Unmarshal(data, &dst.RequestTypeRequesterMemberOf)
		if err == nil {
			return nil // data stored in dst.RequestTypeRequesterMemberOf, return on the first match
		} else {
			dst.RequestTypeRequesterMemberOf = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeRequestSettingsReadable as RequestTypeRequesterMemberOf: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeRequestSettingsReadable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeRequesterCustom != nil {
		return json.Marshal(&src.RequestTypeRequesterCustom)
	}

	if src.RequestTypeRequesterEveryone != nil {
		return json.Marshal(&src.RequestTypeRequesterEveryone)
	}

	if src.RequestTypeRequesterMemberOf != nil {
		return json.Marshal(&src.RequestTypeRequesterMemberOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeRequestSettingsReadable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeRequesterCustom != nil {
		return obj.RequestTypeRequesterCustom
	}

	if obj.RequestTypeRequesterEveryone != nil {
		return obj.RequestTypeRequesterEveryone
	}

	if obj.RequestTypeRequesterMemberOf != nil {
		return obj.RequestTypeRequesterMemberOf
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeRequestSettingsReadable struct {
	value *RequestTypeRequestSettingsReadable
	isSet bool
}

func (v NullableRequestTypeRequestSettingsReadable) Get() *RequestTypeRequestSettingsReadable {
	return v.value
}

func (v *NullableRequestTypeRequestSettingsReadable) Set(val *RequestTypeRequestSettingsReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeRequestSettingsReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeRequestSettingsReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeRequestSettingsReadable(val *RequestTypeRequestSettingsReadable) *NullableRequestTypeRequestSettingsReadable {
	return &NullableRequestTypeRequestSettingsReadable{value: val, isSet: true}
}

func (v NullableRequestTypeRequestSettingsReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeRequestSettingsReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
