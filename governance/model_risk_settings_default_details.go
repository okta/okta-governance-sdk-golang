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
// RiskSettingsDefaultDetails - Default risk settings that are valid for an access request when a risk has been detected for the resource and requesting user. Request submission indicates whether users are allowed or restricted to submit a request when there is any risk conflict.  An approval sequence must be selected with optional access duration settings here for `ALLOWED_WITH_OVERRIDES`, while the  settings from the matching request condition would apply for `ALLOWED_WITH_NO_OVERRIDES`. For `RESTRICTED`, access duration settings can only be present when there is an error.
type RiskSettingsDefaultDetails struct {
	RiskSettingsDefaultAllowedWithNoOverridesDetails *RiskSettingsDefaultAllowedWithNoOverridesDetails
	RiskSettingsDefaultAllowedWithOverridesDetails   *RiskSettingsDefaultAllowedWithOverridesDetails
	RiskSettingsDefaultRestrictedDetails             *RiskSettingsDefaultRestrictedDetails
}

// RiskSettingsDefaultAllowedWithNoOverridesDetailsAsRiskSettingsDefaultDetails is a convenience function that returns RiskSettingsDefaultAllowedWithNoOverridesDetails wrapped in RiskSettingsDefaultDetails
func RiskSettingsDefaultAllowedWithNoOverridesDetailsAsRiskSettingsDefaultDetails(v *RiskSettingsDefaultAllowedWithNoOverridesDetails) RiskSettingsDefaultDetails {
	return RiskSettingsDefaultDetails{
		RiskSettingsDefaultAllowedWithNoOverridesDetails: v,
	}
}

// RiskSettingsDefaultAllowedWithOverridesDetailsAsRiskSettingsDefaultDetails is a convenience function that returns RiskSettingsDefaultAllowedWithOverridesDetails wrapped in RiskSettingsDefaultDetails
func RiskSettingsDefaultAllowedWithOverridesDetailsAsRiskSettingsDefaultDetails(v *RiskSettingsDefaultAllowedWithOverridesDetails) RiskSettingsDefaultDetails {
	return RiskSettingsDefaultDetails{
		RiskSettingsDefaultAllowedWithOverridesDetails: v,
	}
}

// RiskSettingsDefaultRestrictedDetailsAsRiskSettingsDefaultDetails is a convenience function that returns RiskSettingsDefaultRestrictedDetails wrapped in RiskSettingsDefaultDetails
func RiskSettingsDefaultRestrictedDetailsAsRiskSettingsDefaultDetails(v *RiskSettingsDefaultRestrictedDetails) RiskSettingsDefaultDetails {
	return RiskSettingsDefaultDetails{
		RiskSettingsDefaultRestrictedDetails: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RiskSettingsDefaultDetails) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ALLOWED_WITH_NO_OVERRIDES'
	if jsonDict["requestSubmissionType"] == "ALLOWED_WITH_NO_OVERRIDES" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithNoOverridesDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithNoOverridesDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithNoOverridesDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithNoOverridesDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultAllowedWithNoOverridesDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ALLOWED_WITH_OVERRIDES'
	if jsonDict["requestSubmissionType"] == "ALLOWED_WITH_OVERRIDES" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithOverridesDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithOverridesDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithOverridesDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithOverridesDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultAllowedWithOverridesDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESTRICTED'
	if jsonDict["requestSubmissionType"] == "RESTRICTED" {
		// try to unmarshal JSON data into RiskSettingsDefaultRestrictedDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultRestrictedDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultRestrictedDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultRestrictedDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultRestrictedDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-allowed-with-no-overrides-details'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-allowed-with-no-overrides-details" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithNoOverridesDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithNoOverridesDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithNoOverridesDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithNoOverridesDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultAllowedWithNoOverridesDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-allowed-with-overrides-details'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-allowed-with-overrides-details" {
		// try to unmarshal JSON data into RiskSettingsDefaultAllowedWithOverridesDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultAllowedWithOverridesDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultAllowedWithOverridesDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultAllowedWithOverridesDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultAllowedWithOverridesDetails: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risk-settings-default-restricted-details'
	if jsonDict["requestSubmissionType"] == "risk-settings-default-restricted-details" {
		// try to unmarshal JSON data into RiskSettingsDefaultRestrictedDetails
		err = json.Unmarshal(data, &dst.RiskSettingsDefaultRestrictedDetails)
		if err == nil {
			return nil // data stored in dst.RiskSettingsDefaultRestrictedDetails, return on the first match
		} else {
			dst.RiskSettingsDefaultRestrictedDetails = nil
			return fmt.Errorf("Failed to unmarshal RiskSettingsDefaultDetails as RiskSettingsDefaultRestrictedDetails: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskSettingsDefaultDetails) MarshalJSON() ([]byte, error) {
	if src.RiskSettingsDefaultAllowedWithNoOverridesDetails != nil {
		return json.Marshal(&src.RiskSettingsDefaultAllowedWithNoOverridesDetails)
	}

	if src.RiskSettingsDefaultAllowedWithOverridesDetails != nil {
		return json.Marshal(&src.RiskSettingsDefaultAllowedWithOverridesDetails)
	}

	if src.RiskSettingsDefaultRestrictedDetails != nil {
		return json.Marshal(&src.RiskSettingsDefaultRestrictedDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskSettingsDefaultDetails) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RiskSettingsDefaultAllowedWithNoOverridesDetails != nil {
		return obj.RiskSettingsDefaultAllowedWithNoOverridesDetails
	}

	if obj.RiskSettingsDefaultAllowedWithOverridesDetails != nil {
		return obj.RiskSettingsDefaultAllowedWithOverridesDetails
	}

	if obj.RiskSettingsDefaultRestrictedDetails != nil {
		return obj.RiskSettingsDefaultRestrictedDetails
	}

	// all schemas are nil
	return nil
}

type NullableRiskSettingsDefaultDetails struct {
	value *RiskSettingsDefaultDetails
	isSet bool
}

func (v NullableRiskSettingsDefaultDetails) Get() *RiskSettingsDefaultDetails {
	return v.value
}

func (v *NullableRiskSettingsDefaultDetails) Set(val *RiskSettingsDefaultDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultDetails(val *RiskSettingsDefaultDetails) *NullableRiskSettingsDefaultDetails {
	return &NullableRiskSettingsDefaultDetails{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
