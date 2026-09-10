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

// ResourceTypeEnum The type of resource in the inventory
type ResourceTypeEnum string

// List of resource-type-enum
const (
	RESOURCETYPEENUM_APPS                ResourceTypeEnum = "apps"
	RESOURCETYPEENUM_GROUPS              ResourceTypeEnum = "groups"
	RESOURCETYPEENUM_ENTITLEMENT_VALUES  ResourceTypeEnum = "entitlement-values"
	RESOURCETYPEENUM_ENTITLEMENT_BUNDLES ResourceTypeEnum = "entitlement-bundles"
	RESOURCETYPEENUM_COLLECTIONS         ResourceTypeEnum = "collections"
)

// All allowed values of ResourceTypeEnum enum
var AllowedResourceTypeEnumEnumValues = []ResourceTypeEnum{
	"apps",
	"groups",
	"entitlement-values",
	"entitlement-bundles",
	"collections",
}

func (v *ResourceTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceTypeEnum(value)
	for _, existing := range AllowedResourceTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceTypeEnum", value)
}

// NewResourceTypeEnumFromValue returns a pointer to a valid ResourceTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceTypeEnumFromValue(v string) (*ResourceTypeEnum, error) {
	ev := ResourceTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceTypeEnum: valid values are %v", v, AllowedResourceTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceTypeEnum) IsValid() bool {
	for _, existing := range AllowedResourceTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-type-enum value
func (v ResourceTypeEnum) Ptr() *ResourceTypeEnum {
	return &v
}

type NullableResourceTypeEnum struct {
	value *ResourceTypeEnum
	isSet bool
}

func (v NullableResourceTypeEnum) Get() *ResourceTypeEnum {
	return v.value
}

func (v *NullableResourceTypeEnum) Set(val *ResourceTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeEnum(val *ResourceTypeEnum) *NullableResourceTypeEnum {
	return &NullableResourceTypeEnum{value: val, isSet: true}
}

func (v NullableResourceTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
