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

// ResourceTypeExclude The type of Okta resource to exclude
type ResourceTypeExclude string

// List of resource-type-exclude
const (
	RESOURCETYPEEXCLUDE_GROUP       ResourceTypeExclude = "GROUP"
	RESOURCETYPEEXCLUDE_APPLICATION ResourceTypeExclude = "APPLICATION"
)

// All allowed values of ResourceTypeExclude enum
var AllowedResourceTypeExcludeEnumValues = []ResourceTypeExclude{
	"GROUP",
	"APPLICATION",
}

func (v *ResourceTypeExclude) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceTypeExclude(value)
	for _, existing := range AllowedResourceTypeExcludeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceTypeExclude", value)
}

// NewResourceTypeExcludeFromValue returns a pointer to a valid ResourceTypeExclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceTypeExcludeFromValue(v string) (*ResourceTypeExclude, error) {
	ev := ResourceTypeExclude(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceTypeExclude: valid values are %v", v, AllowedResourceTypeExcludeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceTypeExclude) IsValid() bool {
	for _, existing := range AllowedResourceTypeExcludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-type-exclude value
func (v ResourceTypeExclude) Ptr() *ResourceTypeExclude {
	return &v
}

type NullableResourceTypeExclude struct {
	value *ResourceTypeExclude
	isSet bool
}

func (v NullableResourceTypeExclude) Get() *ResourceTypeExclude {
	return v.value
}

func (v *NullableResourceTypeExclude) Set(val *ResourceTypeExclude) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeExclude) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeExclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeExclude(val *ResourceTypeExclude) *NullableResourceTypeExclude {
	return &NullableResourceTypeExclude{value: val, isSet: true}
}

func (v NullableResourceTypeExclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeExclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
