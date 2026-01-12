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

// IntegrationStatus Integration status
type IntegrationStatus string

// List of integration-status
const (
	INTEGRATIONSTATUS_CONNECTED   IntegrationStatus = "CONNECTED"
	INTEGRATIONSTATUS_IN_PROGRESS IntegrationStatus = "IN_PROGRESS"
)

// All allowed values of IntegrationStatus enum
var AllowedIntegrationStatusEnumValues = []IntegrationStatus{
	"CONNECTED",
	"IN_PROGRESS",
}

func (v *IntegrationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationStatus(value)
	for _, existing := range AllowedIntegrationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationStatus", value)
}

// NewIntegrationStatusFromValue returns a pointer to a valid IntegrationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationStatusFromValue(v string) (*IntegrationStatus, error) {
	ev := IntegrationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationStatus: valid values are %v", v, AllowedIntegrationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationStatus) IsValid() bool {
	for _, existing := range AllowedIntegrationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to integration-status value
func (v IntegrationStatus) Ptr() *IntegrationStatus {
	return &v
}

type NullableIntegrationStatus struct {
	value *IntegrationStatus
	isSet bool
}

func (v NullableIntegrationStatus) Get() *IntegrationStatus {
	return v.value
}

func (v *NullableIntegrationStatus) Set(val *IntegrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationStatus(val *IntegrationStatus) *NullableIntegrationStatus {
	return &NullableIntegrationStatus{value: val, isSet: true}
}

func (v NullableIntegrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
