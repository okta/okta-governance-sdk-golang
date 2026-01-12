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

// EntitlementSettingsRequestStatus Specify to opt in (enable) or opt out of (disable) entitlement management for the resource
type EntitlementSettingsRequestStatus string

// List of entitlement-settings-request-status
const (
	ENTITLEMENTSETTINGSREQUESTSTATUS_OPTED_IN  EntitlementSettingsRequestStatus = "OPTED_IN"
	ENTITLEMENTSETTINGSREQUESTSTATUS_OPTED_OUT EntitlementSettingsRequestStatus = "OPTED_OUT"
)

// All allowed values of EntitlementSettingsRequestStatus enum
var AllowedEntitlementSettingsRequestStatusEnumValues = []EntitlementSettingsRequestStatus{
	"OPTED_IN",
	"OPTED_OUT",
}

func (v *EntitlementSettingsRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementSettingsRequestStatus(value)
	for _, existing := range AllowedEntitlementSettingsRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementSettingsRequestStatus", value)
}

// NewEntitlementSettingsRequestStatusFromValue returns a pointer to a valid EntitlementSettingsRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementSettingsRequestStatusFromValue(v string) (*EntitlementSettingsRequestStatus, error) {
	ev := EntitlementSettingsRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementSettingsRequestStatus: valid values are %v", v, AllowedEntitlementSettingsRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementSettingsRequestStatus) IsValid() bool {
	for _, existing := range AllowedEntitlementSettingsRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-settings-request-status value
func (v EntitlementSettingsRequestStatus) Ptr() *EntitlementSettingsRequestStatus {
	return &v
}

type NullableEntitlementSettingsRequestStatus struct {
	value *EntitlementSettingsRequestStatus
	isSet bool
}

func (v NullableEntitlementSettingsRequestStatus) Get() *EntitlementSettingsRequestStatus {
	return v.value
}

func (v *NullableEntitlementSettingsRequestStatus) Set(val *EntitlementSettingsRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementSettingsRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementSettingsRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementSettingsRequestStatus(val *EntitlementSettingsRequestStatus) *NullableEntitlementSettingsRequestStatus {
	return &NullableEntitlementSettingsRequestStatus{value: val, isSet: true}
}

func (v NullableEntitlementSettingsRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementSettingsRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
