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

// ProvisioningStatus The org provisioning status in Access Requests
type ProvisioningStatus string

// List of provisioning-status
const (
	PROVISIONINGSTATUS_NOT_PROVISIONED     ProvisioningStatus = "NOT_PROVISIONED"
	PROVISIONINGSTATUS_PROVISIONING_FAILED ProvisioningStatus = "PROVISIONING_FAILED"
	PROVISIONINGSTATUS_PROVISIONED         ProvisioningStatus = "PROVISIONED"
)

// All allowed values of ProvisioningStatus enum
var AllowedProvisioningStatusEnumValues = []ProvisioningStatus{
	"NOT_PROVISIONED",
	"PROVISIONING_FAILED",
	"PROVISIONED",
}

func (v *ProvisioningStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningStatus(value)
	for _, existing := range AllowedProvisioningStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningStatus", value)
}

// NewProvisioningStatusFromValue returns a pointer to a valid ProvisioningStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningStatusFromValue(v string) (*ProvisioningStatus, error) {
	ev := ProvisioningStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningStatus: valid values are %v", v, AllowedProvisioningStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningStatus) IsValid() bool {
	for _, existing := range AllowedProvisioningStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to provisioning-status value
func (v ProvisioningStatus) Ptr() *ProvisioningStatus {
	return &v
}

type NullableProvisioningStatus struct {
	value *ProvisioningStatus
	isSet bool
}

func (v NullableProvisioningStatus) Get() *ProvisioningStatus {
	return v.value
}

func (v *NullableProvisioningStatus) Set(val *ProvisioningStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningStatus(val *ProvisioningStatus) *NullableProvisioningStatus {
	return &NullableProvisioningStatus{value: val, isSet: true}
}

func (v NullableProvisioningStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
