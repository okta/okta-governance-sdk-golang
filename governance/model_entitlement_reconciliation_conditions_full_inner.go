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

// EntitlementReconciliationConditionsFullInner - struct for EntitlementReconciliationConditionsFullInner
type EntitlementReconciliationConditionsFullInner struct {
	EntitlementReconciliationEntitlementConditionFull *EntitlementReconciliationEntitlementConditionFull
	EntitlementReconciliationLabelConditionFull       *EntitlementReconciliationLabelConditionFull
}

// EntitlementReconciliationEntitlementConditionFullAsEntitlementReconciliationConditionsFullInner is a convenience function that returns EntitlementReconciliationEntitlementConditionFull wrapped in EntitlementReconciliationConditionsFullInner
func EntitlementReconciliationEntitlementConditionFullAsEntitlementReconciliationConditionsFullInner(v *EntitlementReconciliationEntitlementConditionFull) EntitlementReconciliationConditionsFullInner {
	return EntitlementReconciliationConditionsFullInner{
		EntitlementReconciliationEntitlementConditionFull: v,
	}
}

// EntitlementReconciliationLabelConditionFullAsEntitlementReconciliationConditionsFullInner is a convenience function that returns EntitlementReconciliationLabelConditionFull wrapped in EntitlementReconciliationConditionsFullInner
func EntitlementReconciliationLabelConditionFullAsEntitlementReconciliationConditionsFullInner(v *EntitlementReconciliationLabelConditionFull) EntitlementReconciliationConditionsFullInner {
	return EntitlementReconciliationConditionsFullInner{
		EntitlementReconciliationLabelConditionFull: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntitlementReconciliationConditionsFullInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ENTITLEMENT'
	if jsonDict["refType"] == "ENTITLEMENT" {
		// try to unmarshal JSON data into EntitlementReconciliationEntitlementConditionFull
		err = json.Unmarshal(data, &dst.EntitlementReconciliationEntitlementConditionFull)
		if err == nil {
			return nil // data stored in dst.EntitlementReconciliationEntitlementConditionFull, return on the first match
		} else {
			dst.EntitlementReconciliationEntitlementConditionFull = nil
			return fmt.Errorf("failed to unmarshal EntitlementReconciliationConditionsFullInner as EntitlementReconciliationEntitlementConditionFull: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LABEL'
	if jsonDict["refType"] == "LABEL" {
		// try to unmarshal JSON data into EntitlementReconciliationLabelConditionFull
		err = json.Unmarshal(data, &dst.EntitlementReconciliationLabelConditionFull)
		if err == nil {
			return nil // data stored in dst.EntitlementReconciliationLabelConditionFull, return on the first match
		} else {
			dst.EntitlementReconciliationLabelConditionFull = nil
			return fmt.Errorf("failed to unmarshal EntitlementReconciliationConditionsFullInner as EntitlementReconciliationLabelConditionFull: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntitlementReconciliationConditionsFullInner) MarshalJSON() ([]byte, error) {
	if src.EntitlementReconciliationEntitlementConditionFull != nil {
		return json.Marshal(&src.EntitlementReconciliationEntitlementConditionFull)
	}

	if src.EntitlementReconciliationLabelConditionFull != nil {
		return json.Marshal(&src.EntitlementReconciliationLabelConditionFull)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntitlementReconciliationConditionsFullInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementReconciliationEntitlementConditionFull != nil {
		return obj.EntitlementReconciliationEntitlementConditionFull
	}

	if obj.EntitlementReconciliationLabelConditionFull != nil {
		return obj.EntitlementReconciliationLabelConditionFull
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EntitlementReconciliationConditionsFullInner) GetActualInstanceValue() interface{} {
	if obj.EntitlementReconciliationEntitlementConditionFull != nil {
		return *obj.EntitlementReconciliationEntitlementConditionFull
	}

	if obj.EntitlementReconciliationLabelConditionFull != nil {
		return *obj.EntitlementReconciliationLabelConditionFull
	}

	// all schemas are nil
	return nil
}

type NullableEntitlementReconciliationConditionsFullInner struct {
	value *EntitlementReconciliationConditionsFullInner
	isSet bool
}

func (v NullableEntitlementReconciliationConditionsFullInner) Get() *EntitlementReconciliationConditionsFullInner {
	return v.value
}

func (v *NullableEntitlementReconciliationConditionsFullInner) Set(val *EntitlementReconciliationConditionsFullInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationConditionsFullInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationConditionsFullInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationConditionsFullInner(val *EntitlementReconciliationConditionsFullInner) *NullableEntitlementReconciliationConditionsFullInner {
	return &NullableEntitlementReconciliationConditionsFullInner{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationConditionsFullInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationConditionsFullInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
