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

// GrantStatus The state of the particular grant setting
type GrantStatus string

// List of grant-status
const (
	GRANTSTATUS_ACTIVE    GrantStatus = "ACTIVE"
	GRANTSTATUS_SCHEDULED GrantStatus = "SCHEDULED"
	GRANTSTATUS_INACTIVE  GrantStatus = "INACTIVE"
	GRANTSTATUS_EXPIRED   GrantStatus = "EXPIRED"
)

// All allowed values of GrantStatus enum
var AllowedGrantStatusEnumValues = []GrantStatus{
	"ACTIVE",
	"SCHEDULED",
	"INACTIVE",
	"EXPIRED",
}

func (v *GrantStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantStatus(value)
	for _, existing := range AllowedGrantStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantStatus", value)
}

// NewGrantStatusFromValue returns a pointer to a valid GrantStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantStatusFromValue(v string) (*GrantStatus, error) {
	ev := GrantStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantStatus: valid values are %v", v, AllowedGrantStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantStatus) IsValid() bool {
	for _, existing := range AllowedGrantStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grant-status value
func (v GrantStatus) Ptr() *GrantStatus {
	return &v
}

type NullableGrantStatus struct {
	value *GrantStatus
	isSet bool
}

func (v NullableGrantStatus) Get() *GrantStatus {
	return v.value
}

func (v *NullableGrantStatus) Set(val *GrantStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantStatus(val *GrantStatus) *NullableGrantStatus {
	return &NullableGrantStatus{value: val, isSet: true}
}

func (v NullableGrantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
