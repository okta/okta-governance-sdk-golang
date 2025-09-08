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

// RequestTypeCreatableStatus Whether the request type starts with a status of DRAFT or ACTIVE. Starting as ACTIVE is preferable to avoid a publish operation when no manual inspection of the request type is necessary before publication.
type RequestTypeCreatableStatus string

// List of request-type-creatable-status
const (
	REQUESTTYPECREATABLESTATUS_DRAFT  RequestTypeCreatableStatus = "DRAFT"
	REQUESTTYPECREATABLESTATUS_ACTIVE RequestTypeCreatableStatus = "ACTIVE"
)

// All allowed values of RequestTypeCreatableStatus enum
var AllowedRequestTypeCreatableStatusEnumValues = []RequestTypeCreatableStatus{
	"DRAFT",
	"ACTIVE",
}

func (v *RequestTypeCreatableStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestTypeCreatableStatus(value)
	for _, existing := range AllowedRequestTypeCreatableStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestTypeCreatableStatus", value)
}

// NewRequestTypeCreatableStatusFromValue returns a pointer to a valid RequestTypeCreatableStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeCreatableStatusFromValue(v string) (*RequestTypeCreatableStatus, error) {
	ev := RequestTypeCreatableStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestTypeCreatableStatus: valid values are %v", v, AllowedRequestTypeCreatableStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestTypeCreatableStatus) IsValid() bool {
	for _, existing := range AllowedRequestTypeCreatableStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-type-creatable-status value
func (v RequestTypeCreatableStatus) Ptr() *RequestTypeCreatableStatus {
	return &v
}

type NullableRequestTypeCreatableStatus struct {
	value *RequestTypeCreatableStatus
	isSet bool
}

func (v NullableRequestTypeCreatableStatus) Get() *RequestTypeCreatableStatus {
	return v.value
}

func (v *NullableRequestTypeCreatableStatus) Set(val *RequestTypeCreatableStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeCreatableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeCreatableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeCreatableStatus(val *RequestTypeCreatableStatus) *NullableRequestTypeCreatableStatus {
	return &NullableRequestTypeCreatableStatus{value: val, isSet: true}
}

func (v NullableRequestTypeCreatableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeCreatableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
