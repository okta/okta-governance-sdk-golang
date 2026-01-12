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

// EntitlementPropertyDatatype The data type of the entitlement property
type EntitlementPropertyDatatype string

// List of entitlement-property-datatype
const (
	ENTITLEMENTPROPERTYDATATYPE_STRING EntitlementPropertyDatatype = "string"
	ENTITLEMENTPROPERTYDATATYPE_ARRAY  EntitlementPropertyDatatype = "array"
)

// All allowed values of EntitlementPropertyDatatype enum
var AllowedEntitlementPropertyDatatypeEnumValues = []EntitlementPropertyDatatype{
	"string",
	"array",
}

func (v *EntitlementPropertyDatatype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementPropertyDatatype(value)
	for _, existing := range AllowedEntitlementPropertyDatatypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementPropertyDatatype", value)
}

// NewEntitlementPropertyDatatypeFromValue returns a pointer to a valid EntitlementPropertyDatatype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementPropertyDatatypeFromValue(v string) (*EntitlementPropertyDatatype, error) {
	ev := EntitlementPropertyDatatype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementPropertyDatatype: valid values are %v", v, AllowedEntitlementPropertyDatatypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementPropertyDatatype) IsValid() bool {
	for _, existing := range AllowedEntitlementPropertyDatatypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-property-datatype value
func (v EntitlementPropertyDatatype) Ptr() *EntitlementPropertyDatatype {
	return &v
}

type NullableEntitlementPropertyDatatype struct {
	value *EntitlementPropertyDatatype
	isSet bool
}

func (v NullableEntitlementPropertyDatatype) Get() *EntitlementPropertyDatatype {
	return v.value
}

func (v *NullableEntitlementPropertyDatatype) Set(val *EntitlementPropertyDatatype) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementPropertyDatatype) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementPropertyDatatype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementPropertyDatatype(val *EntitlementPropertyDatatype) *NullableEntitlementPropertyDatatype {
	return &NullableEntitlementPropertyDatatype{value: val, isSet: true}
}

func (v NullableEntitlementPropertyDatatype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementPropertyDatatype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
