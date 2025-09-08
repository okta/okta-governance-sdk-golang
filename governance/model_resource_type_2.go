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

// ResourceType2 The type of resource
type ResourceType2 string

// List of resource-type-2
const (
	RESOURCETYPE2_APPLICATION ResourceType2 = "APPLICATION"
)

// All allowed values of ResourceType2 enum
var AllowedResourceType2EnumValues = []ResourceType2{
	"APPLICATION",
}

func (v *ResourceType2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceType2(value)
	for _, existing := range AllowedResourceType2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceType2", value)
}

// NewResourceType2FromValue returns a pointer to a valid ResourceType2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceType2FromValue(v string) (*ResourceType2, error) {
	ev := ResourceType2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceType2: valid values are %v", v, AllowedResourceType2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceType2) IsValid() bool {
	for _, existing := range AllowedResourceType2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-type-2 value
func (v ResourceType2) Ptr() *ResourceType2 {
	return &v
}

type NullableResourceType2 struct {
	value *ResourceType2
	isSet bool
}

func (v NullableResourceType2) Get() *ResourceType2 {
	return v.value
}

func (v *NullableResourceType2) Set(val *ResourceType2) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType2) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType2(val *ResourceType2) *NullableResourceType2 {
	return &NullableResourceType2{value: val, isSet: true}
}

func (v NullableResourceType2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
