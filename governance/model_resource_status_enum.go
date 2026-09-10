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

// ResourceStatusEnum Current status of the resource
type ResourceStatusEnum string

// List of resource-status-enum
const (
	RESOURCESTATUSENUM_ACTIVE   ResourceStatusEnum = "active"
	RESOURCESTATUSENUM_INACTIVE ResourceStatusEnum = "inactive"
)

// All allowed values of ResourceStatusEnum enum
var AllowedResourceStatusEnumEnumValues = []ResourceStatusEnum{
	"active",
	"inactive",
}

func (v *ResourceStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceStatusEnum(value)
	for _, existing := range AllowedResourceStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceStatusEnum", value)
}

// NewResourceStatusEnumFromValue returns a pointer to a valid ResourceStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceStatusEnumFromValue(v string) (*ResourceStatusEnum, error) {
	ev := ResourceStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceStatusEnum: valid values are %v", v, AllowedResourceStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceStatusEnum) IsValid() bool {
	for _, existing := range AllowedResourceStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-status-enum value
func (v ResourceStatusEnum) Ptr() *ResourceStatusEnum {
	return &v
}

type NullableResourceStatusEnum struct {
	value *ResourceStatusEnum
	isSet bool
}

func (v NullableResourceStatusEnum) Get() *ResourceStatusEnum {
	return v.value
}

func (v *NullableResourceStatusEnum) Set(val *ResourceStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStatusEnum(val *ResourceStatusEnum) *NullableResourceStatusEnum {
	return &NullableResourceStatusEnum{value: val, isSet: true}
}

func (v NullableResourceStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
