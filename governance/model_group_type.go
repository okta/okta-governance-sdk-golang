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

// GroupType The type of group. It is either 'GROUP' or 'RESOURCE_OWNER'.
type GroupType string

// List of group-type
const (
	GROUPTYPE_GROUP          GroupType = "GROUP"
	GROUPTYPE_RESOURCE_OWNER GroupType = "RESOURCE_OWNER"
)

// All allowed values of GroupType enum
var AllowedGroupTypeEnumValues = []GroupType{
	"GROUP",
	"RESOURCE_OWNER",
}

func (v *GroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupType(value)
	for _, existing := range AllowedGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupType", value)
}

// NewGroupTypeFromValue returns a pointer to a valid GroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupTypeFromValue(v string) (*GroupType, error) {
	ev := GroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupType: valid values are %v", v, AllowedGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupType) IsValid() bool {
	for _, existing := range AllowedGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to group-type value
func (v GroupType) Ptr() *GroupType {
	return &v
}

type NullableGroupType struct {
	value *GroupType
	isSet bool
}

func (v NullableGroupType) Get() *GroupType {
	return v.value
}

func (v *NullableGroupType) Set(val *GroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupType(val *GroupType) *NullableGroupType {
	return &NullableGroupType{value: val, isSet: true}
}

func (v NullableGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
