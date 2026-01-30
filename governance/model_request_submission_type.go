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

// RequestSubmissionType Whether request submission is allowed or restricted in the risk settings.
type RequestSubmissionType string

// List of request-submission-type
const (
	REQUESTSUBMISSIONTYPE_ALLOWED_WITH_NO_OVERRIDES RequestSubmissionType = "ALLOWED_WITH_NO_OVERRIDES"
	REQUESTSUBMISSIONTYPE_ALLOWED_WITH_OVERRIDES    RequestSubmissionType = "ALLOWED_WITH_OVERRIDES"
	REQUESTSUBMISSIONTYPE_RESTRICTED                RequestSubmissionType = "RESTRICTED"
)

// All allowed values of RequestSubmissionType enum
var AllowedRequestSubmissionTypeEnumValues = []RequestSubmissionType{
	"ALLOWED_WITH_NO_OVERRIDES",
	"ALLOWED_WITH_OVERRIDES",
	"RESTRICTED",
}

func (v *RequestSubmissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestSubmissionType(value)
	for _, existing := range AllowedRequestSubmissionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestSubmissionType", value)
}

// NewRequestSubmissionTypeFromValue returns a pointer to a valid RequestSubmissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestSubmissionTypeFromValue(v string) (*RequestSubmissionType, error) {
	ev := RequestSubmissionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestSubmissionType: valid values are %v", v, AllowedRequestSubmissionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestSubmissionType) IsValid() bool {
	for _, existing := range AllowedRequestSubmissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-submission-type value
func (v RequestSubmissionType) Ptr() *RequestSubmissionType {
	return &v
}

type NullableRequestSubmissionType struct {
	value *RequestSubmissionType
	isSet bool
}

func (v NullableRequestSubmissionType) Get() *RequestSubmissionType {
	return v.value
}

func (v *NullableRequestSubmissionType) Set(val *RequestSubmissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmissionType(val *RequestSubmissionType) *NullableRequestSubmissionType {
	return &NullableRequestSubmissionType{value: val, isSet: true}
}

func (v NullableRequestSubmissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
