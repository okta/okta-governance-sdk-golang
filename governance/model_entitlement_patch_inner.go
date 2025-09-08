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
// EntitlementPatchInner - struct for EntitlementPatchInner
type EntitlementPatchInner struct {
	EntitlementPatchOperation       *EntitlementPatchOperation
	EntitlementValuesPatchOperation *EntitlementValuesPatchOperation
}

// EntitlementPatchOperationAsEntitlementPatchInner is a convenience function that returns EntitlementPatchOperation wrapped in EntitlementPatchInner
func EntitlementPatchOperationAsEntitlementPatchInner(v *EntitlementPatchOperation) EntitlementPatchInner {
	return EntitlementPatchInner{
		EntitlementPatchOperation: v,
	}
}

// EntitlementValuesPatchOperationAsEntitlementPatchInner is a convenience function that returns EntitlementValuesPatchOperation wrapped in EntitlementPatchInner
func EntitlementValuesPatchOperationAsEntitlementPatchInner(v *EntitlementValuesPatchOperation) EntitlementPatchInner {
	return EntitlementPatchInner{
		EntitlementValuesPatchOperation: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *EntitlementPatchInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ENTITLEMENT'
	if jsonDict["refType"] == "ENTITLEMENT" {
		// try to unmarshal JSON data into EntitlementPatchOperation
		err = json.Unmarshal(data, &dst.EntitlementPatchOperation)
		if err == nil {
			return nil // data stored in dst.EntitlementPatchOperation, return on the first match
		} else {
			dst.EntitlementPatchOperation = nil
			return fmt.Errorf("Failed to unmarshal EntitlementPatchInner as EntitlementPatchOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITLEMENT-VALUE'
	if jsonDict["refType"] == "ENTITLEMENT-VALUE" {
		// try to unmarshal JSON data into EntitlementValuesPatchOperation
		err = json.Unmarshal(data, &dst.EntitlementValuesPatchOperation)
		if err == nil {
			return nil // data stored in dst.EntitlementValuesPatchOperation, return on the first match
		} else {
			dst.EntitlementValuesPatchOperation = nil
			return fmt.Errorf("Failed to unmarshal EntitlementPatchInner as EntitlementValuesPatchOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'entitlement-patch-operation'
	if jsonDict["refType"] == "entitlement-patch-operation" {
		// try to unmarshal JSON data into EntitlementPatchOperation
		err = json.Unmarshal(data, &dst.EntitlementPatchOperation)
		if err == nil {
			return nil // data stored in dst.EntitlementPatchOperation, return on the first match
		} else {
			dst.EntitlementPatchOperation = nil
			return fmt.Errorf("Failed to unmarshal EntitlementPatchInner as EntitlementPatchOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'entitlement-values-patch-operation'
	if jsonDict["refType"] == "entitlement-values-patch-operation" {
		// try to unmarshal JSON data into EntitlementValuesPatchOperation
		err = json.Unmarshal(data, &dst.EntitlementValuesPatchOperation)
		if err == nil {
			return nil // data stored in dst.EntitlementValuesPatchOperation, return on the first match
		} else {
			dst.EntitlementValuesPatchOperation = nil
			return fmt.Errorf("Failed to unmarshal EntitlementPatchInner as EntitlementValuesPatchOperation: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntitlementPatchInner) MarshalJSON() ([]byte, error) {
	if src.EntitlementPatchOperation != nil {
		return json.Marshal(&src.EntitlementPatchOperation)
	}

	if src.EntitlementValuesPatchOperation != nil {
		return json.Marshal(&src.EntitlementValuesPatchOperation)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntitlementPatchInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementPatchOperation != nil {
		return obj.EntitlementPatchOperation
	}

	if obj.EntitlementValuesPatchOperation != nil {
		return obj.EntitlementValuesPatchOperation
	}

	// all schemas are nil
	return nil
}

type NullableEntitlementPatchInner struct {
	value *EntitlementPatchInner
	isSet bool
}

func (v NullableEntitlementPatchInner) Get() *EntitlementPatchInner {
	return v.value
}

func (v *NullableEntitlementPatchInner) Set(val *EntitlementPatchInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementPatchInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementPatchInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementPatchInner(val *EntitlementPatchInner) *NullableEntitlementPatchInner {
	return &NullableEntitlementPatchInner{value: val, isSet: true}
}

func (v NullableEntitlementPatchInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementPatchInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
