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

// EntitlementReconciliationConditionsWritableInner - struct for EntitlementReconciliationConditionsWritableInner
type EntitlementReconciliationConditionsWritableInner struct {
	EntitlementReconciliationEntitlementConditionWritable *EntitlementReconciliationEntitlementConditionWritable
	EntitlementReconciliationLabelConditionWritable       *EntitlementReconciliationLabelConditionWritable
}

// EntitlementReconciliationEntitlementConditionWritableAsEntitlementReconciliationConditionsWritableInner is a convenience function that returns EntitlementReconciliationEntitlementConditionWritable wrapped in EntitlementReconciliationConditionsWritableInner
func EntitlementReconciliationEntitlementConditionWritableAsEntitlementReconciliationConditionsWritableInner(v *EntitlementReconciliationEntitlementConditionWritable) EntitlementReconciliationConditionsWritableInner {
	return EntitlementReconciliationConditionsWritableInner{
		EntitlementReconciliationEntitlementConditionWritable: v,
	}
}

// EntitlementReconciliationLabelConditionWritableAsEntitlementReconciliationConditionsWritableInner is a convenience function that returns EntitlementReconciliationLabelConditionWritable wrapped in EntitlementReconciliationConditionsWritableInner
func EntitlementReconciliationLabelConditionWritableAsEntitlementReconciliationConditionsWritableInner(v *EntitlementReconciliationLabelConditionWritable) EntitlementReconciliationConditionsWritableInner {
	return EntitlementReconciliationConditionsWritableInner{
		EntitlementReconciliationLabelConditionWritable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntitlementReconciliationConditionsWritableInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ENTITLEMENT'
	if jsonDict["refType"] == "ENTITLEMENT" {
		// try to unmarshal JSON data into EntitlementReconciliationEntitlementConditionWritable
		err = json.Unmarshal(data, &dst.EntitlementReconciliationEntitlementConditionWritable)
		if err == nil {
			return nil // data stored in dst.EntitlementReconciliationEntitlementConditionWritable, return on the first match
		} else {
			dst.EntitlementReconciliationEntitlementConditionWritable = nil
			return fmt.Errorf("failed to unmarshal EntitlementReconciliationConditionsWritableInner as EntitlementReconciliationEntitlementConditionWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LABEL'
	if jsonDict["refType"] == "LABEL" {
		// try to unmarshal JSON data into EntitlementReconciliationLabelConditionWritable
		err = json.Unmarshal(data, &dst.EntitlementReconciliationLabelConditionWritable)
		if err == nil {
			return nil // data stored in dst.EntitlementReconciliationLabelConditionWritable, return on the first match
		} else {
			dst.EntitlementReconciliationLabelConditionWritable = nil
			return fmt.Errorf("failed to unmarshal EntitlementReconciliationConditionsWritableInner as EntitlementReconciliationLabelConditionWritable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntitlementReconciliationConditionsWritableInner) MarshalJSON() ([]byte, error) {
	if src.EntitlementReconciliationEntitlementConditionWritable != nil {
		return json.Marshal(&src.EntitlementReconciliationEntitlementConditionWritable)
	}

	if src.EntitlementReconciliationLabelConditionWritable != nil {
		return json.Marshal(&src.EntitlementReconciliationLabelConditionWritable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntitlementReconciliationConditionsWritableInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementReconciliationEntitlementConditionWritable != nil {
		return obj.EntitlementReconciliationEntitlementConditionWritable
	}

	if obj.EntitlementReconciliationLabelConditionWritable != nil {
		return obj.EntitlementReconciliationLabelConditionWritable
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EntitlementReconciliationConditionsWritableInner) GetActualInstanceValue() interface{} {
	if obj.EntitlementReconciliationEntitlementConditionWritable != nil {
		return *obj.EntitlementReconciliationEntitlementConditionWritable
	}

	if obj.EntitlementReconciliationLabelConditionWritable != nil {
		return *obj.EntitlementReconciliationLabelConditionWritable
	}

	// all schemas are nil
	return nil
}

type NullableEntitlementReconciliationConditionsWritableInner struct {
	value *EntitlementReconciliationConditionsWritableInner
	isSet bool
}

func (v NullableEntitlementReconciliationConditionsWritableInner) Get() *EntitlementReconciliationConditionsWritableInner {
	return v.value
}

func (v *NullableEntitlementReconciliationConditionsWritableInner) Set(val *EntitlementReconciliationConditionsWritableInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationConditionsWritableInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationConditionsWritableInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationConditionsWritableInner(val *EntitlementReconciliationConditionsWritableInner) *NullableEntitlementReconciliationConditionsWritableInner {
	return &NullableEntitlementReconciliationConditionsWritableInner{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationConditionsWritableInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationConditionsWritableInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
