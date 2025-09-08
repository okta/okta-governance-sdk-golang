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

// RequestGrantStatus The grant status of the request
type RequestGrantStatus string

// List of request-grant-status
const (
	REQUESTGRANTSTATUS_PENDING         RequestGrantStatus = "PENDING"
	REQUESTGRANTSTATUS_GRANTED         RequestGrantStatus = "GRANTED"
	REQUESTGRANTSTATUS_FAILED          RequestGrantStatus = "FAILED"
	REQUESTGRANTSTATUS_MANUAL_REQUIRED RequestGrantStatus = "MANUAL_REQUIRED"
)

// All allowed values of RequestGrantStatus enum
var AllowedRequestGrantStatusEnumValues = []RequestGrantStatus{
	"PENDING",
	"GRANTED",
	"FAILED",
	"MANUAL_REQUIRED",
}

func (v *RequestGrantStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestGrantStatus(value)
	for _, existing := range AllowedRequestGrantStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestGrantStatus", value)
}

// NewRequestGrantStatusFromValue returns a pointer to a valid RequestGrantStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestGrantStatusFromValue(v string) (*RequestGrantStatus, error) {
	ev := RequestGrantStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestGrantStatus: valid values are %v", v, AllowedRequestGrantStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestGrantStatus) IsValid() bool {
	for _, existing := range AllowedRequestGrantStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-grant-status value
func (v RequestGrantStatus) Ptr() *RequestGrantStatus {
	return &v
}

type NullableRequestGrantStatus struct {
	value *RequestGrantStatus
	isSet bool
}

func (v NullableRequestGrantStatus) Get() *RequestGrantStatus {
	return v.value
}

func (v *NullableRequestGrantStatus) Set(val *RequestGrantStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestGrantStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestGrantStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestGrantStatus(val *RequestGrantStatus) *NullableRequestGrantStatus {
	return &NullableRequestGrantStatus{value: val, isSet: true}
}

func (v NullableRequestGrantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestGrantStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
