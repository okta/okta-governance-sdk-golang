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

// ResourceType3 The requested resource type
type ResourceType3 string

// List of resource-type-3
const (
	RESOURCETYPE3_APPLICATION  ResourceType3 = "APPLICATION"
	RESOURCETYPE3_REQUEST_TYPE ResourceType3 = "REQUEST_TYPE"
)

// All allowed values of ResourceType3 enum
var AllowedResourceType3EnumValues = []ResourceType3{
	"APPLICATION",
	"REQUEST_TYPE",
}

func (v *ResourceType3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceType3(value)
	for _, existing := range AllowedResourceType3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceType3", value)
}

// NewResourceType3FromValue returns a pointer to a valid ResourceType3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceType3FromValue(v string) (*ResourceType3, error) {
	ev := ResourceType3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceType3: valid values are %v", v, AllowedResourceType3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceType3) IsValid() bool {
	for _, existing := range AllowedResourceType3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-type-3 value
func (v ResourceType3) Ptr() *ResourceType3 {
	return &v
}

type NullableResourceType3 struct {
	value *ResourceType3
	isSet bool
}

func (v NullableResourceType3) Get() *ResourceType3 {
	return v.value
}

func (v *NullableResourceType3) Set(val *ResourceType3) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType3) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType3(val *ResourceType3) *NullableResourceType3 {
	return &NullableResourceType3{value: val, isSet: true}
}

func (v NullableResourceType3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
