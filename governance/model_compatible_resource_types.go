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

// CompatibleResourceTypes A resource type the sequence is compatible with.
type CompatibleResourceTypes string

// List of compatible-resource-types
const (
	COMPATIBLERESOURCETYPES_APP   CompatibleResourceTypes = "APP"
	COMPATIBLERESOURCETYPES_GROUP CompatibleResourceTypes = "GROUP"
)

// All allowed values of CompatibleResourceTypes enum
var AllowedCompatibleResourceTypesEnumValues = []CompatibleResourceTypes{
	"APP",
	"GROUP",
}

func (v *CompatibleResourceTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompatibleResourceTypes(value)
	for _, existing := range AllowedCompatibleResourceTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompatibleResourceTypes", value)
}

// NewCompatibleResourceTypesFromValue returns a pointer to a valid CompatibleResourceTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompatibleResourceTypesFromValue(v string) (*CompatibleResourceTypes, error) {
	ev := CompatibleResourceTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompatibleResourceTypes: valid values are %v", v, AllowedCompatibleResourceTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompatibleResourceTypes) IsValid() bool {
	for _, existing := range AllowedCompatibleResourceTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to compatible-resource-types value
func (v CompatibleResourceTypes) Ptr() *CompatibleResourceTypes {
	return &v
}

type NullableCompatibleResourceTypes struct {
	value *CompatibleResourceTypes
	isSet bool
}

func (v NullableCompatibleResourceTypes) Get() *CompatibleResourceTypes {
	return v.value
}

func (v *NullableCompatibleResourceTypes) Set(val *CompatibleResourceTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableCompatibleResourceTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableCompatibleResourceTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompatibleResourceTypes(val *CompatibleResourceTypes) *NullableCompatibleResourceTypes {
	return &NullableCompatibleResourceTypes{value: val, isSet: true}
}

func (v NullableCompatibleResourceTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompatibleResourceTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
