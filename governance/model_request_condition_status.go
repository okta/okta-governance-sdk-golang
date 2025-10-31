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

// RequestConditionStatus status indicating if this condition is active or not. Default status is INACTIVE
type RequestConditionStatus string

// List of request-condition-status
const (
	REQUESTCONDITIONSTATUS_ACTIVE   RequestConditionStatus = "ACTIVE"
	REQUESTCONDITIONSTATUS_INACTIVE RequestConditionStatus = "INACTIVE"
	REQUESTCONDITIONSTATUS_INVALID  RequestConditionStatus = "INVALID"
	REQUESTCONDITIONSTATUS_DELETED  RequestConditionStatus = "DELETED"
)

// All allowed values of RequestConditionStatus enum
var AllowedRequestConditionStatusEnumValues = []RequestConditionStatus{
	"ACTIVE",
	"INACTIVE",
	"INVALID",
	"DELETED",
}

func (v *RequestConditionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestConditionStatus(value)
	for _, existing := range AllowedRequestConditionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestConditionStatus", value)
}

// NewRequestConditionStatusFromValue returns a pointer to a valid RequestConditionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestConditionStatusFromValue(v string) (*RequestConditionStatus, error) {
	ev := RequestConditionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestConditionStatus: valid values are %v", v, AllowedRequestConditionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestConditionStatus) IsValid() bool {
	for _, existing := range AllowedRequestConditionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-condition-status value
func (v RequestConditionStatus) Ptr() *RequestConditionStatus {
	return &v
}

type NullableRequestConditionStatus struct {
	value *RequestConditionStatus
	isSet bool
}

func (v NullableRequestConditionStatus) Get() *RequestConditionStatus {
	return v.value
}

func (v *NullableRequestConditionStatus) Set(val *RequestConditionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionStatus(val *RequestConditionStatus) *NullableRequestConditionStatus {
	return &NullableRequestConditionStatus{value: val, isSet: true}
}

func (v NullableRequestConditionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
