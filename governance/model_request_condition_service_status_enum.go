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

// RequestConditionServiceStatusEnum status describes the request condition service status for the current org.
type RequestConditionServiceStatusEnum string

// List of request-condition-service-status-enum
const (
	REQUESTCONDITIONSERVICESTATUSENUM_READY RequestConditionServiceStatusEnum = "READY"
	REQUESTCONDITIONSERVICESTATUSENUM_LIVE  RequestConditionServiceStatusEnum = "LIVE"
)

// All allowed values of RequestConditionServiceStatusEnum enum
var AllowedRequestConditionServiceStatusEnumEnumValues = []RequestConditionServiceStatusEnum{
	"READY",
	"LIVE",
}

func (v *RequestConditionServiceStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestConditionServiceStatusEnum(value)
	for _, existing := range AllowedRequestConditionServiceStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestConditionServiceStatusEnum", value)
}

// NewRequestConditionServiceStatusEnumFromValue returns a pointer to a valid RequestConditionServiceStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestConditionServiceStatusEnumFromValue(v string) (*RequestConditionServiceStatusEnum, error) {
	ev := RequestConditionServiceStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestConditionServiceStatusEnum: valid values are %v", v, AllowedRequestConditionServiceStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestConditionServiceStatusEnum) IsValid() bool {
	for _, existing := range AllowedRequestConditionServiceStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-condition-service-status-enum value
func (v RequestConditionServiceStatusEnum) Ptr() *RequestConditionServiceStatusEnum {
	return &v
}

type NullableRequestConditionServiceStatusEnum struct {
	value *RequestConditionServiceStatusEnum
	isSet bool
}

func (v NullableRequestConditionServiceStatusEnum) Get() *RequestConditionServiceStatusEnum {
	return v.value
}

func (v *NullableRequestConditionServiceStatusEnum) Set(val *RequestConditionServiceStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionServiceStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionServiceStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionServiceStatusEnum(val *RequestConditionServiceStatusEnum) *NullableRequestConditionServiceStatusEnum {
	return &NullableRequestConditionServiceStatusEnum{value: val, isSet: true}
}

func (v NullableRequestConditionServiceStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionServiceStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
