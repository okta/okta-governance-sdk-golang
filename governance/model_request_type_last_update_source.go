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

// RequestTypeLastUpdateSource Whether the request type was last updated through the API or the user interface.
type RequestTypeLastUpdateSource string

// List of request-type-last-update-source
const (
	REQUESTTYPELASTUPDATESOURCE_API RequestTypeLastUpdateSource = "API"
	REQUESTTYPELASTUPDATESOURCE_WEB RequestTypeLastUpdateSource = "WEB"
)

// All allowed values of RequestTypeLastUpdateSource enum
var AllowedRequestTypeLastUpdateSourceEnumValues = []RequestTypeLastUpdateSource{
	"API",
	"WEB",
}

func (v *RequestTypeLastUpdateSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestTypeLastUpdateSource(value)
	for _, existing := range AllowedRequestTypeLastUpdateSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestTypeLastUpdateSource", value)
}

// NewRequestTypeLastUpdateSourceFromValue returns a pointer to a valid RequestTypeLastUpdateSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeLastUpdateSourceFromValue(v string) (*RequestTypeLastUpdateSource, error) {
	ev := RequestTypeLastUpdateSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestTypeLastUpdateSource: valid values are %v", v, AllowedRequestTypeLastUpdateSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestTypeLastUpdateSource) IsValid() bool {
	for _, existing := range AllowedRequestTypeLastUpdateSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-type-last-update-source value
func (v RequestTypeLastUpdateSource) Ptr() *RequestTypeLastUpdateSource {
	return &v
}

type NullableRequestTypeLastUpdateSource struct {
	value *RequestTypeLastUpdateSource
	isSet bool
}

func (v NullableRequestTypeLastUpdateSource) Get() *RequestTypeLastUpdateSource {
	return v.value
}

func (v *NullableRequestTypeLastUpdateSource) Set(val *RequestTypeLastUpdateSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeLastUpdateSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeLastUpdateSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeLastUpdateSource(val *RequestTypeLastUpdateSource) *NullableRequestTypeLastUpdateSource {
	return &NullableRequestTypeLastUpdateSource{value: val, isSet: true}
}

func (v NullableRequestTypeLastUpdateSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeLastUpdateSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
