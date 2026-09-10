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

// RequestStatusPatchable Specify the status of the request
type RequestStatusPatchable string

// List of request-status-patchable
const (
	REQUESTSTATUSPATCHABLE_CANCELED RequestStatusPatchable = "CANCELED"
)

// All allowed values of RequestStatusPatchable enum
var AllowedRequestStatusPatchableEnumValues = []RequestStatusPatchable{
	"CANCELED",
}

func (v *RequestStatusPatchable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestStatusPatchable(value)
	for _, existing := range AllowedRequestStatusPatchableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestStatusPatchable", value)
}

// NewRequestStatusPatchableFromValue returns a pointer to a valid RequestStatusPatchable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestStatusPatchableFromValue(v string) (*RequestStatusPatchable, error) {
	ev := RequestStatusPatchable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestStatusPatchable: valid values are %v", v, AllowedRequestStatusPatchableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestStatusPatchable) IsValid() bool {
	for _, existing := range AllowedRequestStatusPatchableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-status-patchable value
func (v RequestStatusPatchable) Ptr() *RequestStatusPatchable {
	return &v
}

type NullableRequestStatusPatchable struct {
	value *RequestStatusPatchable
	isSet bool
}

func (v NullableRequestStatusPatchable) Get() *RequestStatusPatchable {
	return v.value
}

func (v *NullableRequestStatusPatchable) Set(val *RequestStatusPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatusPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatusPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatusPatchable(val *RequestStatusPatchable) *NullableRequestStatusPatchable {
	return &NullableRequestStatusPatchable{value: val, isSet: true}
}

func (v NullableRequestStatusPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatusPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
