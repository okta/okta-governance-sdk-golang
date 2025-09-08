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
// RequestFull - A full representation of a request
type RequestFull struct {
	RequestFullApiCompatible   *RequestFullApiCompatible
	RequestFullApiIncompatible *RequestFullApiIncompatible
}

// RequestFullApiCompatibleAsRequestFull is a convenience function that returns RequestFullApiCompatible wrapped in RequestFull
func RequestFullApiCompatibleAsRequestFull(v *RequestFullApiCompatible) RequestFull {
	return RequestFull{
		RequestFullApiCompatible: v,
	}
}

// RequestFullApiIncompatibleAsRequestFull is a convenience function that returns RequestFullApiIncompatible wrapped in RequestFull
func RequestFullApiIncompatibleAsRequestFull(v *RequestFullApiIncompatible) RequestFull {
	return RequestFull{
		RequestFullApiIncompatible: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestFull) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ACCESS_REQUEST'
	if jsonDict["type"] == "ACCESS_REQUEST" {
		// try to unmarshal JSON data into RequestFullApiCompatible
		err = json.Unmarshal(data, &dst.RequestFullApiCompatible)
		if err == nil {
			return nil // data stored in dst.RequestFullApiCompatible, return on the first match
		} else {
			dst.RequestFullApiCompatible = nil
			return fmt.Errorf("Failed to unmarshal RequestFull as RequestFullApiCompatible: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into RequestFullApiIncompatible
		err = json.Unmarshal(data, &dst.RequestFullApiIncompatible)
		if err == nil {
			return nil // data stored in dst.RequestFullApiIncompatible, return on the first match
		} else {
			dst.RequestFullApiIncompatible = nil
			return fmt.Errorf("Failed to unmarshal RequestFull as RequestFullApiIncompatible: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-full-api-compatible'
	if jsonDict["type"] == "request-full-api-compatible" {
		// try to unmarshal JSON data into RequestFullApiCompatible
		err = json.Unmarshal(data, &dst.RequestFullApiCompatible)
		if err == nil {
			return nil // data stored in dst.RequestFullApiCompatible, return on the first match
		} else {
			dst.RequestFullApiCompatible = nil
			return fmt.Errorf("Failed to unmarshal RequestFull as RequestFullApiCompatible: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-full-api-incompatible'
	if jsonDict["type"] == "request-full-api-incompatible" {
		// try to unmarshal JSON data into RequestFullApiIncompatible
		err = json.Unmarshal(data, &dst.RequestFullApiIncompatible)
		if err == nil {
			return nil // data stored in dst.RequestFullApiIncompatible, return on the first match
		} else {
			dst.RequestFullApiIncompatible = nil
			return fmt.Errorf("Failed to unmarshal RequestFull as RequestFullApiIncompatible: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestFull) MarshalJSON() ([]byte, error) {
	if src.RequestFullApiCompatible != nil {
		return json.Marshal(&src.RequestFullApiCompatible)
	}

	if src.RequestFullApiIncompatible != nil {
		return json.Marshal(&src.RequestFullApiIncompatible)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestFull) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestFullApiCompatible != nil {
		return obj.RequestFullApiCompatible
	}

	if obj.RequestFullApiIncompatible != nil {
		return obj.RequestFullApiIncompatible
	}

	// all schemas are nil
	return nil
}

type NullableRequestFull struct {
	value *RequestFull
	isSet bool
}

func (v NullableRequestFull) Get() *RequestFull {
	return v.value
}

func (v *NullableRequestFull) Set(val *RequestFull) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFull) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFull(val *RequestFull) *NullableRequestFull {
	return &NullableRequestFull{value: val, isSet: true}
}

func (v NullableRequestFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
