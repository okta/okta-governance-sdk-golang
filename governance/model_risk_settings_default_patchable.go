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

// model_oneof.mustache
// RiskSettingsDefaultPatchable - Default risk settings that are valid for an access request when a risk has been detected for the resource and requesting user. Request submission indicates whether users are allowed or restricted to submit a request when there is any risk conflict. An approval sequence must be selected with optional access duration settings here for `ALLOWED_WITH_OVERRIDES`, while the settings from the matching request condition would apply for `ALLOWED_WITH_NO_OVERRIDES`. For `RESTRICTED` and `ALLOWED_WITH_NO_OVERRIDES`, access duration settings can only be present when there is an error.
type RiskSettingsDefaultPatchable struct {
	RiskSettingsDefaultAllowedWithNoOverridesPatchable *RiskSettingsDefaultAllowedWithNoOverridesPatchable
	RiskSettingsDefaultAllowedWithOverridesPatchable   *RiskSettingsDefaultAllowedWithOverridesPatchable
	RiskSettingsDefaultRestrictedPatchable             *RiskSettingsDefaultRestrictedPatchable
}

// RiskSettingsDefaultAllowedWithNoOverridesPatchableAsRiskSettingsDefaultPatchable is a convenience function that returns RiskSettingsDefaultAllowedWithNoOverridesPatchable wrapped in RiskSettingsDefaultPatchable
func RiskSettingsDefaultAllowedWithNoOverridesPatchableAsRiskSettingsDefaultPatchable(v *RiskSettingsDefaultAllowedWithNoOverridesPatchable) RiskSettingsDefaultPatchable {
	return RiskSettingsDefaultPatchable{
		RiskSettingsDefaultAllowedWithNoOverridesPatchable: v,
	}
}

// RiskSettingsDefaultAllowedWithOverridesPatchableAsRiskSettingsDefaultPatchable is a convenience function that returns RiskSettingsDefaultAllowedWithOverridesPatchable wrapped in RiskSettingsDefaultPatchable
func RiskSettingsDefaultAllowedWithOverridesPatchableAsRiskSettingsDefaultPatchable(v *RiskSettingsDefaultAllowedWithOverridesPatchable) RiskSettingsDefaultPatchable {
	return RiskSettingsDefaultPatchable{
		RiskSettingsDefaultAllowedWithOverridesPatchable: v,
	}
}

// RiskSettingsDefaultRestrictedPatchableAsRiskSettingsDefaultPatchable is a convenience function that returns RiskSettingsDefaultRestrictedPatchable wrapped in RiskSettingsDefaultPatchable
func RiskSettingsDefaultRestrictedPatchableAsRiskSettingsDefaultPatchable(v *RiskSettingsDefaultRestrictedPatchable) RiskSettingsDefaultPatchable {
	return RiskSettingsDefaultPatchable{
		RiskSettingsDefaultRestrictedPatchable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RiskSettingsDefaultPatchable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ALLOWED_WITH_NO_OVERRIDES'
	if jsonDict["requestSubmissionType"] == "ALLOWED_WITH_NO_OVERRIDES" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithNoOverridesPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultAllowedWithNoOverridesPatchable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ALLOWED_WITH_OVERRIDES'
	if jsonDict["requestSubmissionType"] == "ALLOWED_WITH_OVERRIDES" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithOverridesPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithOverridesPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithOverridesPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithOverridesPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultAllowedWithOverridesPatchable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESTRICTED'
	if jsonDict["requestSubmissionType"] == "RESTRICTED" {
		// try to unmarshal JSON data into RiskSettingsDefaultRestrictedPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultRestrictedPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultRestrictedPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultRestrictedPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultRestrictedPatchable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-allowed-with-no-overrides-patchable'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-allowed-with-no-overrides-patchable" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithNoOverridesPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithNoOverridesPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultAllowedWithNoOverridesPatchable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-allowed-with-overrides-patchable'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-allowed-with-overrides-patchable" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithOverridesPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithOverridesPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithOverridesPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithOverridesPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultAllowedWithOverridesPatchable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-restricted-patchable'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-restricted-patchable" {
		// try to unmarshal JSON data into RiskSettingsDefaultRestrictedPatchable
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultRestrictedPatchable)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultRestrictedPatchable, return on the first match
		} else {
			dst.RiskSettingsDefaultRestrictedPatchable = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultPatchable as RiskSettingsDefaultRestrictedPatchable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskSettingsDefaultPatchable) MarshalJSON() ([]byte, error) {
	if src.RiskSettingsDefaultAllowedWithNoOverridesPatchable != nil {
		return json.Marshal(&src.RiskSettingsDefaultAllowedWithNoOverridesPatchable)
	}

	if src.RiskSettingsDefaultAllowedWithOverridesPatchable != nil {
		return json.Marshal(&src.RiskSettingsDefaultAllowedWithOverridesPatchable)
	}

	if src.RiskSettingsDefaultRestrictedPatchable != nil {
		return json.Marshal(&src.RiskSettingsDefaultRestrictedPatchable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskSettingsDefaultPatchable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RiskSettingsDefaultAllowedWithNoOverridesPatchable != nil {
		return obj.RiskSettingsDefaultAllowedWithNoOverridesPatchable
	}

	if obj.RiskSettingsDefaultAllowedWithOverridesPatchable != nil {
		return obj.RiskSettingsDefaultAllowedWithOverridesPatchable
	}

	if obj.RiskSettingsDefaultRestrictedPatchable != nil {
		return obj.RiskSettingsDefaultRestrictedPatchable
	}

	// all schemas are nil
	return nil
}

type NullableRiskSettingsDefaultPatchable struct {
	value *RiskSettingsDefaultPatchable
	isSet bool
}

func (v NullableRiskSettingsDefaultPatchable) Get() *RiskSettingsDefaultPatchable {
	return v.value
}

func (v *NullableRiskSettingsDefaultPatchable) Set(val *RiskSettingsDefaultPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultPatchable(val *RiskSettingsDefaultPatchable) *NullableRiskSettingsDefaultPatchable {
	return &NullableRiskSettingsDefaultPatchable{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
