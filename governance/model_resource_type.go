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

// ResourceType The type of Okta resource
type ResourceType string

// List of resource-type
const (
	RESOURCETYPE_GROUP                  ResourceType = "GROUP"
	RESOURCETYPE_APPLICATION            ResourceType = "APPLICATION"
	RESOURCETYPE_GOVERNANCE_LABEL_VALUE ResourceType = "GOVERNANCE_LABEL_VALUE"
	RESOURCETYPE_APP_SERVICE_ACCOUNT    ResourceType = "APP_SERVICE_ACCOUNT"
	RESOURCETYPE_OKTA_SERVICE_ACCOUNT   ResourceType = "OKTA_SERVICE_ACCOUNT"
)

// All allowed values of ResourceType enum
var AllowedResourceTypeEnumValues = []ResourceType{
	"GROUP",
	"APPLICATION",
	"GOVERNANCE_LABEL_VALUE",
	"APP_SERVICE_ACCOUNT",
	"OKTA_SERVICE_ACCOUNT",
}

func (v *ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceType(value)
	for _, existing := range AllowedResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceType", value)
}

// NewResourceTypeFromValue returns a pointer to a valid ResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceTypeFromValue(v string) (*ResourceType, error) {
	ev := ResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceType: valid values are %v", v, AllowedResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceType) IsValid() bool {
	for _, existing := range AllowedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-type value
func (v ResourceType) Ptr() *ResourceType {
	return &v
}

type NullableResourceType struct {
	value *ResourceType
	isSet bool
}

func (v NullableResourceType) Get() *ResourceType {
	return v.value
}

func (v *NullableResourceType) Set(val *ResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType(val *ResourceType) *NullableResourceType {
	return &NullableResourceType{value: val, isSet: true}
}

func (v NullableResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
