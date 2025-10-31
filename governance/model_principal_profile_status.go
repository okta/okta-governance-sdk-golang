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

// PrincipalProfileStatus The status of the principal's profile
type PrincipalProfileStatus string

// List of principal-profile-status
const (
	PRINCIPALPROFILESTATUS_ACTIVE   PrincipalProfileStatus = "ACTIVE"
	PRINCIPALPROFILESTATUS_INACTIVE PrincipalProfileStatus = "INACTIVE"
)

// All allowed values of PrincipalProfileStatus enum
var AllowedPrincipalProfileStatusEnumValues = []PrincipalProfileStatus{
	"ACTIVE",
	"INACTIVE",
}

func (v *PrincipalProfileStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrincipalProfileStatus(value)
	for _, existing := range AllowedPrincipalProfileStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrincipalProfileStatus", value)
}

// NewPrincipalProfileStatusFromValue returns a pointer to a valid PrincipalProfileStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrincipalProfileStatusFromValue(v string) (*PrincipalProfileStatus, error) {
	ev := PrincipalProfileStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrincipalProfileStatus: valid values are %v", v, AllowedPrincipalProfileStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrincipalProfileStatus) IsValid() bool {
	for _, existing := range AllowedPrincipalProfileStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to principal-profile-status value
func (v PrincipalProfileStatus) Ptr() *PrincipalProfileStatus {
	return &v
}

type NullablePrincipalProfileStatus struct {
	value *PrincipalProfileStatus
	isSet bool
}

func (v NullablePrincipalProfileStatus) Get() *PrincipalProfileStatus {
	return v.value
}

func (v *NullablePrincipalProfileStatus) Set(val *PrincipalProfileStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalProfileStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalProfileStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalProfileStatus(val *PrincipalProfileStatus) *NullablePrincipalProfileStatus {
	return &NullablePrincipalProfileStatus{value: val, isSet: true}
}

func (v NullablePrincipalProfileStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalProfileStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
