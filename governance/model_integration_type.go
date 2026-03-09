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

// IntegrationType Integration type (eg. SLACK)
type IntegrationType string

// List of integration-type
const (
	INTEGRATIONTYPE_SLACK IntegrationType = "SLACK"
)

// All allowed values of IntegrationType enum
var AllowedIntegrationTypeEnumValues = []IntegrationType{
	"SLACK",
}

func (v *IntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationType(value)
	for _, existing := range AllowedIntegrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationType", value)
}

// NewIntegrationTypeFromValue returns a pointer to a valid IntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationTypeFromValue(v string) (*IntegrationType, error) {
	ev := IntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationType: valid values are %v", v, AllowedIntegrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationType) IsValid() bool {
	for _, existing := range AllowedIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to integration-type value
func (v IntegrationType) Ptr() *IntegrationType {
	return &v
}

type NullableIntegrationType struct {
	value *IntegrationType
	isSet bool
}

func (v NullableIntegrationType) Get() *IntegrationType {
	return v.value
}

func (v *NullableIntegrationType) Set(val *IntegrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationType(val *IntegrationType) *NullableIntegrationType {
	return &NullableIntegrationType{value: val, isSet: true}
}

func (v NullableIntegrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
