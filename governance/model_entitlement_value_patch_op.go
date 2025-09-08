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

// EntitlementValuePatchOp The operation to be performed for update.
type EntitlementValuePatchOp string

// List of entitlement-value-patch-op
const (
	ENTITLEMENTVALUEPATCHOP_ADD     EntitlementValuePatchOp = "ADD"
	ENTITLEMENTVALUEPATCHOP_REMOVE  EntitlementValuePatchOp = "REMOVE"
	ENTITLEMENTVALUEPATCHOP_REPLACE EntitlementValuePatchOp = "REPLACE"
)

// All allowed values of EntitlementValuePatchOp enum
var AllowedEntitlementValuePatchOpEnumValues = []EntitlementValuePatchOp{
	"ADD",
	"REMOVE",
	"REPLACE",
}

func (v *EntitlementValuePatchOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementValuePatchOp(value)
	for _, existing := range AllowedEntitlementValuePatchOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementValuePatchOp", value)
}

// NewEntitlementValuePatchOpFromValue returns a pointer to a valid EntitlementValuePatchOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementValuePatchOpFromValue(v string) (*EntitlementValuePatchOp, error) {
	ev := EntitlementValuePatchOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementValuePatchOp: valid values are %v", v, AllowedEntitlementValuePatchOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementValuePatchOp) IsValid() bool {
	for _, existing := range AllowedEntitlementValuePatchOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-value-patch-op value
func (v EntitlementValuePatchOp) Ptr() *EntitlementValuePatchOp {
	return &v
}

type NullableEntitlementValuePatchOp struct {
	value *EntitlementValuePatchOp
	isSet bool
}

func (v NullableEntitlementValuePatchOp) Get() *EntitlementValuePatchOp {
	return v.value
}

func (v *NullableEntitlementValuePatchOp) Set(val *EntitlementValuePatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuePatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuePatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuePatchOp(val *EntitlementValuePatchOp) *NullableEntitlementValuePatchOp {
	return &NullableEntitlementValuePatchOp{value: val, isSet: true}
}

func (v NullableEntitlementValuePatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuePatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
