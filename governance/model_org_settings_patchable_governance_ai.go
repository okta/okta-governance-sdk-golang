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
)

// checks if the OrgSettingsPatchableGovernanceAI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsPatchableGovernanceAI{}

// OrgSettingsPatchableGovernanceAI Governance AI settings  > **Note:** This feature is excluded from the Okta for AI Agents - Core SKU, > which is the version of Okta for AI Agents available to FedRAMP Moderate and FedRAMP High customers. > Okta for AI Agents - Core isn't available in Okta for US Military cells. > For a current list of features that are excluded from the Okta for AI Agents - Core SKU, > please refer to the [Okta US Public Sector Limitations or Exceptions](https://support.okta.com/help/s/article/okta-us-public-sector-limitations-or-exceptions?language=en_US) documentation.
type OrgSettingsPatchableGovernanceAI struct {
	SecurityAccessReview *OrgSettingsGovernanceAISecurityAccessReview `json:"securityAccessReview,omitempty"`
	GovernanceAnalyzer   *OrgSettingsGovernanceAIGovernanceAnalyzer   `json:"governanceAnalyzer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsPatchableGovernanceAI OrgSettingsPatchableGovernanceAI

// NewOrgSettingsPatchableGovernanceAI instantiates a new OrgSettingsPatchableGovernanceAI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsPatchableGovernanceAI() *OrgSettingsPatchableGovernanceAI {
	this := OrgSettingsPatchableGovernanceAI{}
	return &this
}

// NewOrgSettingsPatchableGovernanceAIWithDefaults instantiates a new OrgSettingsPatchableGovernanceAI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsPatchableGovernanceAIWithDefaults() *OrgSettingsPatchableGovernanceAI {
	this := OrgSettingsPatchableGovernanceAI{}
	return &this
}

// GetSecurityAccessReview returns the SecurityAccessReview field value if set, zero value otherwise.
func (o *OrgSettingsPatchableGovernanceAI) GetSecurityAccessReview() OrgSettingsGovernanceAISecurityAccessReview {
	if o == nil || IsNil(o.SecurityAccessReview) {
		var ret OrgSettingsGovernanceAISecurityAccessReview
		return ret
	}
	return *o.SecurityAccessReview
}

// GetSecurityAccessReviewOk returns a tuple with the SecurityAccessReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchableGovernanceAI) GetSecurityAccessReviewOk() (*OrgSettingsGovernanceAISecurityAccessReview, bool) {
	if o == nil || IsNil(o.SecurityAccessReview) {
		return nil, false
	}
	return o.SecurityAccessReview, true
}

// HasSecurityAccessReview returns a boolean if a field has been set.
func (o *OrgSettingsPatchableGovernanceAI) HasSecurityAccessReview() bool {
	if o != nil && !IsNil(o.SecurityAccessReview) {
		return true
	}

	return false
}

// SetSecurityAccessReview gets a reference to the given OrgSettingsGovernanceAISecurityAccessReview and assigns it to the SecurityAccessReview field.
func (o *OrgSettingsPatchableGovernanceAI) SetSecurityAccessReview(v OrgSettingsGovernanceAISecurityAccessReview) {
	o.SecurityAccessReview = &v
}

// GetGovernanceAnalyzer returns the GovernanceAnalyzer field value if set, zero value otherwise.
func (o *OrgSettingsPatchableGovernanceAI) GetGovernanceAnalyzer() OrgSettingsGovernanceAIGovernanceAnalyzer {
	if o == nil || IsNil(o.GovernanceAnalyzer) {
		var ret OrgSettingsGovernanceAIGovernanceAnalyzer
		return ret
	}
	return *o.GovernanceAnalyzer
}

// GetGovernanceAnalyzerOk returns a tuple with the GovernanceAnalyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchableGovernanceAI) GetGovernanceAnalyzerOk() (*OrgSettingsGovernanceAIGovernanceAnalyzer, bool) {
	if o == nil || IsNil(o.GovernanceAnalyzer) {
		return nil, false
	}
	return o.GovernanceAnalyzer, true
}

// HasGovernanceAnalyzer returns a boolean if a field has been set.
func (o *OrgSettingsPatchableGovernanceAI) HasGovernanceAnalyzer() bool {
	if o != nil && !IsNil(o.GovernanceAnalyzer) {
		return true
	}

	return false
}

// SetGovernanceAnalyzer gets a reference to the given OrgSettingsGovernanceAIGovernanceAnalyzer and assigns it to the GovernanceAnalyzer field.
func (o *OrgSettingsPatchableGovernanceAI) SetGovernanceAnalyzer(v OrgSettingsGovernanceAIGovernanceAnalyzer) {
	o.GovernanceAnalyzer = &v
}

func (o OrgSettingsPatchableGovernanceAI) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsPatchableGovernanceAI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecurityAccessReview) {
		toSerialize["securityAccessReview"] = o.SecurityAccessReview
	}
	if !IsNil(o.GovernanceAnalyzer) {
		toSerialize["governanceAnalyzer"] = o.GovernanceAnalyzer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsPatchableGovernanceAI) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsPatchableGovernanceAI := _OrgSettingsPatchableGovernanceAI{}

	err = json.Unmarshal(data, &varOrgSettingsPatchableGovernanceAI)

	if err != nil {
		return err
	}

	*o = OrgSettingsPatchableGovernanceAI(varOrgSettingsPatchableGovernanceAI)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "securityAccessReview")
		delete(additionalProperties, "governanceAnalyzer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsPatchableGovernanceAI struct {
	value *OrgSettingsPatchableGovernanceAI
	isSet bool
}

func (v NullableOrgSettingsPatchableGovernanceAI) Get() *OrgSettingsPatchableGovernanceAI {
	return v.value
}

func (v *NullableOrgSettingsPatchableGovernanceAI) Set(val *OrgSettingsPatchableGovernanceAI) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsPatchableGovernanceAI) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsPatchableGovernanceAI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsPatchableGovernanceAI(val *OrgSettingsPatchableGovernanceAI) *NullableOrgSettingsPatchableGovernanceAI {
	return &NullableOrgSettingsPatchableGovernanceAI{value: val, isSet: true}
}

func (v NullableOrgSettingsPatchableGovernanceAI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsPatchableGovernanceAI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
