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

// RequestExperience the model 'RequestExperience'
type RequestExperience string

// List of request-experience
const (
	REQUESTEXPERIENCE_ADMIN_ROLES RequestExperience = "ADMIN_ROLES"
	REQUESTEXPERIENCE_CATALOG     RequestExperience = "CATALOG"
)

// All allowed values of RequestExperience enum
var AllowedRequestExperienceEnumValues = []RequestExperience{
	"ADMIN_ROLES",
	"CATALOG",
}

func (v *RequestExperience) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestExperience(value)
	for _, existing := range AllowedRequestExperienceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestExperience", value)
}

// NewRequestExperienceFromValue returns a pointer to a valid RequestExperience
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestExperienceFromValue(v string) (*RequestExperience, error) {
	ev := RequestExperience(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestExperience: valid values are %v", v, AllowedRequestExperienceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestExperience) IsValid() bool {
	for _, existing := range AllowedRequestExperienceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-experience value
func (v RequestExperience) Ptr() *RequestExperience {
	return &v
}

type NullableRequestExperience struct {
	value *RequestExperience
	isSet bool
}

func (v NullableRequestExperience) Get() *RequestExperience {
	return v.value
}

func (v *NullableRequestExperience) Set(val *RequestExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestExperience(val *RequestExperience) *NullableRequestExperience {
	return &NullableRequestExperience{value: val, isSet: true}
}

func (v NullableRequestExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
