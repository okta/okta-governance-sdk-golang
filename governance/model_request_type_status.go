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

// RequestTypeStatus A request type has a status lifecycle described in our [request type lifecycle documentation](/#request-type-status-lifecycle).  A request type may start with a status of DRAFT or ACTIVE, based on the status property in the Add request body.  The publish, unpublish, and delete operations all affect a request type's status.  Additionally, if any operation detects invalid settings, such as an approver who is no longer in the organization, the request type automatically transitions to disabled.
type RequestTypeStatus string

// List of request-type-status
const (
	REQUESTTYPESTATUS_DRAFT    RequestTypeStatus = "DRAFT"
	REQUESTTYPESTATUS_ACTIVE   RequestTypeStatus = "ACTIVE"
	REQUESTTYPESTATUS_DISABLED RequestTypeStatus = "DISABLED"
)

// All allowed values of RequestTypeStatus enum
var AllowedRequestTypeStatusEnumValues = []RequestTypeStatus{
	"DRAFT",
	"ACTIVE",
	"DISABLED",
}

func (v *RequestTypeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestTypeStatus(value)
	for _, existing := range AllowedRequestTypeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestTypeStatus", value)
}

// NewRequestTypeStatusFromValue returns a pointer to a valid RequestTypeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeStatusFromValue(v string) (*RequestTypeStatus, error) {
	ev := RequestTypeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestTypeStatus: valid values are %v", v, AllowedRequestTypeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestTypeStatus) IsValid() bool {
	for _, existing := range AllowedRequestTypeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-type-status value
func (v RequestTypeStatus) Ptr() *RequestTypeStatus {
	return &v
}

type NullableRequestTypeStatus struct {
	value *RequestTypeStatus
	isSet bool
}

func (v NullableRequestTypeStatus) Get() *RequestTypeStatus {
	return v.value
}

func (v *NullableRequestTypeStatus) Set(val *RequestTypeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeStatus(val *RequestTypeStatus) *NullableRequestTypeStatus {
	return &NullableRequestTypeStatus{value: val, isSet: true}
}

func (v NullableRequestTypeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
