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

// RequestRequestStatus A request has a lifecycle. See our [request lifecycle documentation](/#request).  A request progresses in its lifecycle state based on requester, approver, and administrator activities across a variety of channels, such as:  - Access Request portal - Slack - Microsoft teams  The following lifecycle operations are available on a single request.  - [Add](/openapi/governance.requests.admin.v1/tag/Requests/#tag/Requests/operation/createRequest)
type RequestRequestStatus string

// List of request-request-status
const (
	REQUESTREQUESTSTATUS_OPEN     RequestRequestStatus = "OPEN"
	REQUESTREQUESTSTATUS_PENDING  RequestRequestStatus = "PENDING"
	REQUESTREQUESTSTATUS_RESOLVED RequestRequestStatus = "RESOLVED"
	REQUESTREQUESTSTATUS_LOCKED   RequestRequestStatus = "LOCKED"
)

// All allowed values of RequestRequestStatus enum
var AllowedRequestRequestStatusEnumValues = []RequestRequestStatus{
	"OPEN",
	"PENDING",
	"RESOLVED",
	"LOCKED",
}

func (v *RequestRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestRequestStatus(value)
	for _, existing := range AllowedRequestRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestRequestStatus", value)
}

// NewRequestRequestStatusFromValue returns a pointer to a valid RequestRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestRequestStatusFromValue(v string) (*RequestRequestStatus, error) {
	ev := RequestRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestRequestStatus: valid values are %v", v, AllowedRequestRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestRequestStatus) IsValid() bool {
	for _, existing := range AllowedRequestRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-request-status value
func (v RequestRequestStatus) Ptr() *RequestRequestStatus {
	return &v
}

type NullableRequestRequestStatus struct {
	value *RequestRequestStatus
	isSet bool
}

func (v NullableRequestRequestStatus) Get() *RequestRequestStatus {
	return v.value
}

func (v *NullableRequestRequestStatus) Set(val *RequestRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestRequestStatus(val *RequestRequestStatus) *NullableRequestRequestStatus {
	return &NullableRequestRequestStatus{value: val, isSet: true}
}

func (v NullableRequestRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
