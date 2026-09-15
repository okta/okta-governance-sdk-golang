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

// EntitlementDriftType Which way the entitlement drifted, compared to what Okta governed before this import.  * `ADD`: The import has an entitlement that Okta wasn't governing.  * `SUB`: Okta was governing an entitlement that's missing from the import.
type EntitlementDriftType string

// List of entitlement-drift-type
const (
	ENTITLEMENTDRIFTTYPE_ADD EntitlementDriftType = "ADD"
	ENTITLEMENTDRIFTTYPE_SUB EntitlementDriftType = "SUB"
)

// All allowed values of EntitlementDriftType enum
var AllowedEntitlementDriftTypeEnumValues = []EntitlementDriftType{
	"ADD",
	"SUB",
}

func (v *EntitlementDriftType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementDriftType(value)
	for _, existing := range AllowedEntitlementDriftTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementDriftType", value)
}

// NewEntitlementDriftTypeFromValue returns a pointer to a valid EntitlementDriftType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementDriftTypeFromValue(v string) (*EntitlementDriftType, error) {
	ev := EntitlementDriftType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementDriftType: valid values are %v", v, AllowedEntitlementDriftTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementDriftType) IsValid() bool {
	for _, existing := range AllowedEntitlementDriftTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-drift-type value
func (v EntitlementDriftType) Ptr() *EntitlementDriftType {
	return &v
}

type NullableEntitlementDriftType struct {
	value *EntitlementDriftType
	isSet bool
}

func (v NullableEntitlementDriftType) Get() *EntitlementDriftType {
	return v.value
}

func (v *NullableEntitlementDriftType) Set(val *EntitlementDriftType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftType(val *EntitlementDriftType) *NullableEntitlementDriftType {
	return &NullableEntitlementDriftType{value: val, isSet: true}
}

func (v NullableEntitlementDriftType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
