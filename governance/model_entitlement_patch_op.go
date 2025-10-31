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

// EntitlementPatchOp The operation to be performed for update.
type EntitlementPatchOp string

// List of entitlement-patch-op
const (
	ENTITLEMENTPATCHOP_ADD     EntitlementPatchOp = "ADD"
	ENTITLEMENTPATCHOP_REMOVE  EntitlementPatchOp = "REMOVE"
	ENTITLEMENTPATCHOP_REPLACE EntitlementPatchOp = "REPLACE"
)

// All allowed values of EntitlementPatchOp enum
var AllowedEntitlementPatchOpEnumValues = []EntitlementPatchOp{
	"ADD",
	"REMOVE",
	"REPLACE",
}

func (v *EntitlementPatchOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementPatchOp(value)
	for _, existing := range AllowedEntitlementPatchOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementPatchOp", value)
}

// NewEntitlementPatchOpFromValue returns a pointer to a valid EntitlementPatchOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementPatchOpFromValue(v string) (*EntitlementPatchOp, error) {
	ev := EntitlementPatchOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementPatchOp: valid values are %v", v, AllowedEntitlementPatchOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementPatchOp) IsValid() bool {
	for _, existing := range AllowedEntitlementPatchOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-patch-op value
func (v EntitlementPatchOp) Ptr() *EntitlementPatchOp {
	return &v
}

type NullableEntitlementPatchOp struct {
	value *EntitlementPatchOp
	isSet bool
}

func (v NullableEntitlementPatchOp) Get() *EntitlementPatchOp {
	return v.value
}

func (v *NullableEntitlementPatchOp) Set(val *EntitlementPatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementPatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementPatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementPatchOp(val *EntitlementPatchOp) *NullableEntitlementPatchOp {
	return &NullableEntitlementPatchOp{value: val, isSet: true}
}

func (v NullableEntitlementPatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementPatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
