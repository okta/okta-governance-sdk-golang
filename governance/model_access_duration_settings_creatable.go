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

// AccessDurationSettingsCreatable - Settings that control who may specify the access duration allowed by this request condition or risk settings, as well as what duration may be requested.  **Note:** The resource request settings affect what access duration settings are valid. See the `validAccessDurationSettings` property.
type AccessDurationSettingsCreatable struct {
	AccessDurationSettingsAdminFixedDuration         *AccessDurationSettingsAdminFixedDuration
	AccessDurationSettingsRequesterSpecifiedDuration *AccessDurationSettingsRequesterSpecifiedDuration
}

// AccessDurationSettingsAdminFixedDurationAsAccessDurationSettingsCreatable is a convenience function that returns AccessDurationSettingsAdminFixedDuration wrapped in AccessDurationSettingsCreatable
func AccessDurationSettingsAdminFixedDurationAsAccessDurationSettingsCreatable(v *AccessDurationSettingsAdminFixedDuration) AccessDurationSettingsCreatable {
	return AccessDurationSettingsCreatable{
		AccessDurationSettingsAdminFixedDuration: v,
	}
}

// AccessDurationSettingsRequesterSpecifiedDurationAsAccessDurationSettingsCreatable is a convenience function that returns AccessDurationSettingsRequesterSpecifiedDuration wrapped in AccessDurationSettingsCreatable
func AccessDurationSettingsRequesterSpecifiedDurationAsAccessDurationSettingsCreatable(v *AccessDurationSettingsRequesterSpecifiedDuration) AccessDurationSettingsCreatable {
	return AccessDurationSettingsCreatable{
		AccessDurationSettingsRequesterSpecifiedDuration: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccessDurationSettingsCreatable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ADMIN_FIXED_DURATION'
	if jsonDict["type"] == "ADMIN_FIXED_DURATION" {
		// try to unmarshal JSON data into AccessDurationSettingsAdminFixedDuration
		err = json.Unmarshal(data, &dst.AccessDurationSettingsAdminFixedDuration)
		if err == nil {
			return nil // data stored in dst.AccessDurationSettingsAdminFixedDuration, return on the first match
		} else {
			dst.AccessDurationSettingsAdminFixedDuration = nil
			return fmt.Errorf("failed to unmarshal AccessDurationSettingsCreatable as AccessDurationSettingsAdminFixedDuration: %s", err.Error())
		}
	}

	// check if the discriminator value is 'REQUESTER_SPECIFIED_DURATION'
	if jsonDict["type"] == "REQUESTER_SPECIFIED_DURATION" {
		// try to unmarshal JSON data into AccessDurationSettingsRequesterSpecifiedDuration
		err = json.Unmarshal(data, &dst.AccessDurationSettingsRequesterSpecifiedDuration)
		if err == nil {
			return nil // data stored in dst.AccessDurationSettingsRequesterSpecifiedDuration, return on the first match
		} else {
			dst.AccessDurationSettingsRequesterSpecifiedDuration = nil
			return fmt.Errorf("failed to unmarshal AccessDurationSettingsCreatable as AccessDurationSettingsRequesterSpecifiedDuration: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccessDurationSettingsCreatable) MarshalJSON() ([]byte, error) {
	if src.AccessDurationSettingsAdminFixedDuration != nil {
		return json.Marshal(&src.AccessDurationSettingsAdminFixedDuration)
	}

	if src.AccessDurationSettingsRequesterSpecifiedDuration != nil {
		return json.Marshal(&src.AccessDurationSettingsRequesterSpecifiedDuration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccessDurationSettingsCreatable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessDurationSettingsAdminFixedDuration != nil {
		return obj.AccessDurationSettingsAdminFixedDuration
	}

	if obj.AccessDurationSettingsRequesterSpecifiedDuration != nil {
		return obj.AccessDurationSettingsRequesterSpecifiedDuration
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AccessDurationSettingsCreatable) GetActualInstanceValue() interface{} {
	if obj.AccessDurationSettingsAdminFixedDuration != nil {
		return *obj.AccessDurationSettingsAdminFixedDuration
	}

	if obj.AccessDurationSettingsRequesterSpecifiedDuration != nil {
		return *obj.AccessDurationSettingsRequesterSpecifiedDuration
	}

	// all schemas are nil
	return nil
}

type NullableAccessDurationSettingsCreatable struct {
	value *AccessDurationSettingsCreatable
	isSet bool
}

func (v NullableAccessDurationSettingsCreatable) Get() *AccessDurationSettingsCreatable {
	return v.value
}

func (v *NullableAccessDurationSettingsCreatable) Set(val *AccessDurationSettingsCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDurationSettingsCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDurationSettingsCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDurationSettingsCreatable(val *AccessDurationSettingsCreatable) *NullableAccessDurationSettingsCreatable {
	return &NullableAccessDurationSettingsCreatable{value: val, isSet: true}
}

func (v NullableAccessDurationSettingsCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDurationSettingsCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
