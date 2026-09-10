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

// ResourceTargetType The type of Okta resource to target
type ResourceTargetType string

// List of resource-target-type
const (
	RESOURCETARGETTYPE_GROUP                ResourceTargetType = "GROUP"
	RESOURCETARGETTYPE_APPLICATION          ResourceTargetType = "APPLICATION"
	RESOURCETARGETTYPE_APP_SERVICE_ACCOUNT  ResourceTargetType = "APP_SERVICE_ACCOUNT"
	RESOURCETARGETTYPE_OKTA_SERVICE_ACCOUNT ResourceTargetType = "OKTA_SERVICE_ACCOUNT"
	RESOURCETARGETTYPE_COLLECTION           ResourceTargetType = "COLLECTION"
	RESOURCETARGETTYPE_AI_AGENT_CONNECTION  ResourceTargetType = "AI_AGENT_CONNECTION"
)

// All allowed values of ResourceTargetType enum
var AllowedResourceTargetTypeEnumValues = []ResourceTargetType{
	"GROUP",
	"APPLICATION",
	"APP_SERVICE_ACCOUNT",
	"OKTA_SERVICE_ACCOUNT",
	"COLLECTION",
	"AI_AGENT_CONNECTION",
}

func (v *ResourceTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceTargetType(value)
	for _, existing := range AllowedResourceTargetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceTargetType", value)
}

// NewResourceTargetTypeFromValue returns a pointer to a valid ResourceTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceTargetTypeFromValue(v string) (*ResourceTargetType, error) {
	ev := ResourceTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceTargetType: valid values are %v", v, AllowedResourceTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceTargetType) IsValid() bool {
	for _, existing := range AllowedResourceTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resource-target-type value
func (v ResourceTargetType) Ptr() *ResourceTargetType {
	return &v
}

type NullableResourceTargetType struct {
	value *ResourceTargetType
	isSet bool
}

func (v NullableResourceTargetType) Get() *ResourceTargetType {
	return v.value
}

func (v *NullableResourceTargetType) Set(val *ResourceTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTargetType(val *ResourceTargetType) *NullableResourceTargetType {
	return &NullableResourceTargetType{value: val, isSet: true}
}

func (v NullableResourceTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
