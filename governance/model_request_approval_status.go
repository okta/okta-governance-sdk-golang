/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

// RequestApprovalStatus The approver's decision.
type RequestApprovalStatus string

// List of request-approval-status
const (
	REQUESTAPPROVALSTATUS_APPROVED RequestApprovalStatus = "APPROVED"
	REQUESTAPPROVALSTATUS_DENIED   RequestApprovalStatus = "DENIED"
)

// All allowed values of RequestApprovalStatus enum
var AllowedRequestApprovalStatusEnumValues = []RequestApprovalStatus{
	"APPROVED",
	"DENIED",
}

func (v *RequestApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestApprovalStatus(value)
	for _, existing := range AllowedRequestApprovalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestApprovalStatus", value)
}

// NewRequestApprovalStatusFromValue returns a pointer to a valid RequestApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestApprovalStatusFromValue(v string) (*RequestApprovalStatus, error) {
	ev := RequestApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestApprovalStatus: valid values are %v", v, AllowedRequestApprovalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestApprovalStatus) IsValid() bool {
	for _, existing := range AllowedRequestApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-approval-status value
func (v RequestApprovalStatus) Ptr() *RequestApprovalStatus {
	return &v
}

type NullableRequestApprovalStatus struct {
	value *RequestApprovalStatus
	isSet bool
}

func (v NullableRequestApprovalStatus) Get() *RequestApprovalStatus {
	return v.value
}

func (v *NullableRequestApprovalStatus) Set(val *RequestApprovalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalStatus(val *RequestApprovalStatus) *NullableRequestApprovalStatus {
	return &NullableRequestApprovalStatus{value: val, isSet: true}
}

func (v NullableRequestApprovalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
