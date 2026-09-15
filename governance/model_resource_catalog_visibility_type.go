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

// ResourceCatalogVisibilityType The target type for resource catalog visibility
type ResourceCatalogVisibilityType string

// List of resource-catalog-visibility-type
const (
	RESOURCECATALOGVISIBILITYTYPE_GROUP ResourceCatalogVisibilityType = "GROUP"
)

// All allowed values of ResourceCatalogVisibilityType enum
var AllowedResourceCatalogVisibilityTypeEnumValues = []ResourceCatalogVisibilityType{
	"GROUP",
}

func (v *ResourceCatalogVisibilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceCatalogVisibilityType(value)
	for _, existing := range AllowedResourceCatalogVisibilityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceCatalogVisibilityType", value)
}

// NewResourceCatalogVisibilityTypeFromValue returns a pointer to a valid ResourceCatalogVisibilityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceCatalogVisibilityTypeFromValue(v string) (*ResourceCatalogVisibilityType, error) {
	ev := ResourceCatalogVisibilityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceCatalogVisibilityType: valid values are %v", v, AllowedResourceCatalogVisibilityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceCatalogVisibilityType) IsValid() bool {
	for _, existing := range AllowedResourceCatalogVisibilityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-catalog-visibility-type value
func (v ResourceCatalogVisibilityType) Ptr() *ResourceCatalogVisibilityType {
	return &v
}

type NullableResourceCatalogVisibilityType struct {
	value *ResourceCatalogVisibilityType
	isSet bool
}

func (v NullableResourceCatalogVisibilityType) Get() *ResourceCatalogVisibilityType {
	return v.value
}

func (v *NullableResourceCatalogVisibilityType) Set(val *ResourceCatalogVisibilityType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCatalogVisibilityType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCatalogVisibilityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCatalogVisibilityType(val *ResourceCatalogVisibilityType) *NullableResourceCatalogVisibilityType {
	return &NullableResourceCatalogVisibilityType{value: val, isSet: true}
}

func (v NullableResourceCatalogVisibilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCatalogVisibilityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
