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

// RequestRevocationStatus The revocation status of the request
type RequestRevocationStatus string

// List of request-revocation-status
const (
	REQUESTREVOCATIONSTATUS_PENDING RequestRevocationStatus = "PENDING"
	REQUESTREVOCATIONSTATUS_REVOKED RequestRevocationStatus = "REVOKED"
	REQUESTREVOCATIONSTATUS_FAILED  RequestRevocationStatus = "FAILED"
)

// All allowed values of RequestRevocationStatus enum
var AllowedRequestRevocationStatusEnumValues = []RequestRevocationStatus{
	"PENDING",
	"REVOKED",
	"FAILED",
}

func (v *RequestRevocationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestRevocationStatus(value)
	for _, existing := range AllowedRequestRevocationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestRevocationStatus", value)
}

// NewRequestRevocationStatusFromValue returns a pointer to a valid RequestRevocationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestRevocationStatusFromValue(v string) (*RequestRevocationStatus, error) {
	ev := RequestRevocationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestRevocationStatus: valid values are %v", v, AllowedRequestRevocationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestRevocationStatus) IsValid() bool {
	for _, existing := range AllowedRequestRevocationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-revocation-status value
func (v RequestRevocationStatus) Ptr() *RequestRevocationStatus {
	return &v
}

type NullableRequestRevocationStatus struct {
	value *RequestRevocationStatus
	isSet bool
}

func (v NullableRequestRevocationStatus) Get() *RequestRevocationStatus {
	return v.value
}

func (v *NullableRequestRevocationStatus) Set(val *RequestRevocationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestRevocationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestRevocationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestRevocationStatus(val *RequestRevocationStatus) *NullableRequestRevocationStatus {
	return &NullableRequestRevocationStatus{value: val, isSet: true}
}

func (v NullableRequestRevocationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestRevocationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
