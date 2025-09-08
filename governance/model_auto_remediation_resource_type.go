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

// AutoRemediationResourceType The type of the resource to be automatically remediated. Only GROUP is supported.
type AutoRemediationResourceType string

// List of auto-remediation-resource-type
const (
	AUTOREMEDIATIONRESOURCETYPE_GROUP AutoRemediationResourceType = "GROUP"
)

// All allowed values of AutoRemediationResourceType enum
var AllowedAutoRemediationResourceTypeEnumValues = []AutoRemediationResourceType{
	"GROUP",
}

func (v *AutoRemediationResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutoRemediationResourceType(value)
	for _, existing := range AllowedAutoRemediationResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutoRemediationResourceType", value)
}

// NewAutoRemediationResourceTypeFromValue returns a pointer to a valid AutoRemediationResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoRemediationResourceTypeFromValue(v string) (*AutoRemediationResourceType, error) {
	ev := AutoRemediationResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoRemediationResourceType: valid values are %v", v, AllowedAutoRemediationResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoRemediationResourceType) IsValid() bool {
	for _, existing := range AllowedAutoRemediationResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to auto-remediation-resource-type value
func (v AutoRemediationResourceType) Ptr() *AutoRemediationResourceType {
	return &v
}

type NullableAutoRemediationResourceType struct {
	value *AutoRemediationResourceType
	isSet bool
}

func (v NullableAutoRemediationResourceType) Get() *AutoRemediationResourceType {
	return v.value
}

func (v *NullableAutoRemediationResourceType) Set(val *AutoRemediationResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoRemediationResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoRemediationResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoRemediationResourceType(val *AutoRemediationResourceType) *NullableAutoRemediationResourceType {
	return &NullableAutoRemediationResourceType{value: val, isSet: true}
}

func (v NullableAutoRemediationResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoRemediationResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
