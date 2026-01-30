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

// RequestersMembersPatchOperation - struct for RequestersMembersPatchOperation
type RequestersMembersPatchOperation struct {
	RequestersAddPatch     *RequestersAddPatch
	RequestersReplacePatch *RequestersReplacePatch
}

// RequestersAddPatchAsRequestersMembersPatchOperation is a convenience function that returns RequestersAddPatch wrapped in RequestersMembersPatchOperation
func RequestersAddPatchAsRequestersMembersPatchOperation(v *RequestersAddPatch) RequestersMembersPatchOperation {
	return RequestersMembersPatchOperation{
		RequestersAddPatch: v,
	}
}

// RequestersReplacePatchAsRequestersMembersPatchOperation is a convenience function that returns RequestersReplacePatch wrapped in RequestersMembersPatchOperation
func RequestersReplacePatchAsRequestersMembersPatchOperation(v *RequestersReplacePatch) RequestersMembersPatchOperation {
	return RequestersMembersPatchOperation{
		RequestersReplacePatch: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestersMembersPatchOperation) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'add'
	if jsonDict["op"] == "add" {
		// try to unmarshal JSON data into RequestersAddPatch
		err = json.Unmarshal(data, &dst.RequestersAddPatch)
		if err == nil {
			return nil // data stored in dst.RequestersAddPatch, return on the first match
		} else {
			dst.RequestersAddPatch = nil
			return fmt.Errorf("failed to unmarshal RequestersMembersPatchOperation as RequestersAddPatch: %s", err.Error())
		}
	}

	// check if the discriminator value is 'replace'
	if jsonDict["op"] == "replace" {
		// try to unmarshal JSON data into RequestersReplacePatch
		err = json.Unmarshal(data, &dst.RequestersReplacePatch)
		if err == nil {
			return nil // data stored in dst.RequestersReplacePatch, return on the first match
		} else {
			dst.RequestersReplacePatch = nil
			return fmt.Errorf("failed to unmarshal RequestersMembersPatchOperation as RequestersReplacePatch: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestersMembersPatchOperation) MarshalJSON() ([]byte, error) {
	if src.RequestersAddPatch != nil {
		return json.Marshal(&src.RequestersAddPatch)
	}

	if src.RequestersReplacePatch != nil {
		return json.Marshal(&src.RequestersReplacePatch)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestersMembersPatchOperation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestersAddPatch != nil {
		return obj.RequestersAddPatch
	}

	if obj.RequestersReplacePatch != nil {
		return obj.RequestersReplacePatch
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestersMembersPatchOperation) GetActualInstanceValue() interface{} {
	if obj.RequestersAddPatch != nil {
		return *obj.RequestersAddPatch
	}

	if obj.RequestersReplacePatch != nil {
		return *obj.RequestersReplacePatch
	}

	// all schemas are nil
	return nil
}

type NullableRequestersMembersPatchOperation struct {
	value *RequestersMembersPatchOperation
	isSet bool
}

func (v NullableRequestersMembersPatchOperation) Get() *RequestersMembersPatchOperation {
	return v.value
}

func (v *NullableRequestersMembersPatchOperation) Set(val *RequestersMembersPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestersMembersPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestersMembersPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestersMembersPatchOperation(val *RequestersMembersPatchOperation) *NullableRequestersMembersPatchOperation {
	return &NullableRequestersMembersPatchOperation{value: val, isSet: true}
}

func (v NullableRequestersMembersPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestersMembersPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
