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

// RequestActionEnum The type of action taken by the access request service to fulfill the request.
type RequestActionEnum string

// List of request-action-enum
const (
	REQUESTACTIONENUM_ADD_USER_TO_GROUP                  RequestActionEnum = "ADD_USER_TO_GROUP"
	REQUESTACTIONENUM_ASSIGN_APP_TO_USER                 RequestActionEnum = "ASSIGN_APP_TO_USER"
	REQUESTACTIONENUM_GRANT_ACCESS_TO_ENTITLEMENT_BUNDLE RequestActionEnum = "GRANT_ACCESS_TO_ENTITLEMENT_BUNDLE"
)

// All allowed values of RequestActionEnum enum
var AllowedRequestActionEnumEnumValues = []RequestActionEnum{
	"ADD_USER_TO_GROUP",
	"ASSIGN_APP_TO_USER",
	"GRANT_ACCESS_TO_ENTITLEMENT_BUNDLE",
}

func (v *RequestActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestActionEnum(value)
	for _, existing := range AllowedRequestActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestActionEnum", value)
}

// NewRequestActionEnumFromValue returns a pointer to a valid RequestActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestActionEnumFromValue(v string) (*RequestActionEnum, error) {
	ev := RequestActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestActionEnum: valid values are %v", v, AllowedRequestActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestActionEnum) IsValid() bool {
	for _, existing := range AllowedRequestActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-action-enum value
func (v RequestActionEnum) Ptr() *RequestActionEnum {
	return &v
}

type NullableRequestActionEnum struct {
	value *RequestActionEnum
	isSet bool
}

func (v NullableRequestActionEnum) Get() *RequestActionEnum {
	return v.value
}

func (v *NullableRequestActionEnum) Set(val *RequestActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestActionEnum(val *RequestActionEnum) *NullableRequestActionEnum {
	return &NullableRequestActionEnum{value: val, isSet: true}
}

func (v NullableRequestActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
