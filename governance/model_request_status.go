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

// RequestStatus The status of the request
type RequestStatus string

// List of request-status
const (
	REQUESTSTATUS_SUBMITTED RequestStatus = "SUBMITTED"
	REQUESTSTATUS_REJECTED  RequestStatus = "REJECTED"
	REQUESTSTATUS_PENDING   RequestStatus = "PENDING"
	REQUESTSTATUS_APPROVED  RequestStatus = "APPROVED"
	REQUESTSTATUS_DENIED    RequestStatus = "DENIED"
	REQUESTSTATUS_CANCELED  RequestStatus = "CANCELED"
	REQUESTSTATUS_EXPIRED   RequestStatus = "EXPIRED"
)

// All allowed values of RequestStatus enum
var AllowedRequestStatusEnumValues = []RequestStatus{
	"SUBMITTED",
	"REJECTED",
	"PENDING",
	"APPROVED",
	"DENIED",
	"CANCELED",
	"EXPIRED",
}

func (v *RequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestStatus(value)
	for _, existing := range AllowedRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestStatus", value)
}

// NewRequestStatusFromValue returns a pointer to a valid RequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestStatusFromValue(v string) (*RequestStatus, error) {
	ev := RequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestStatus: valid values are %v", v, AllowedRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestStatus) IsValid() bool {
	for _, existing := range AllowedRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-status value
func (v RequestStatus) Ptr() *RequestStatus {
	return &v
}

type NullableRequestStatus struct {
	value *RequestStatus
	isSet bool
}

func (v NullableRequestStatus) Get() *RequestStatus {
	return v.value
}

func (v *NullableRequestStatus) Set(val *RequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatus(val *RequestStatus) *NullableRequestStatus {
	return &NullableRequestStatus{value: val, isSet: true}
}

func (v NullableRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
