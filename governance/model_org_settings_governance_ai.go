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

// checks if the OrgSettingsGovernanceAI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsGovernanceAI{}

// OrgSettingsGovernanceAI Governance AI settings  > **Note:** This feature is excluded from the Okta for AI Agents - Core SKU, > which is the version of Okta for AI Agents available to FedRAMP Moderate and FedRAMP High customers. > Okta for AI Agents - Core isn't available in Okta for US Military cells. > For a current list of features that are excluded from the Okta for AI Agents - Core SKU, > please refer to the [Okta US Public Sector Limitations or Exceptions](https://support.okta.com/help/s/article/okta-us-public-sector-limitations-or-exceptions?language=en_US) documentation.
type OrgSettingsGovernanceAI struct {
	SecurityAccessReview *OrgSettingsGovernanceAISecurityAccessReview `json:"securityAccessReview,omitempty"`
	GovernanceAnalyzer   *OrgSettingsGovernanceAIGovernanceAnalyzer   `json:"governanceAnalyzer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsGovernanceAI OrgSettingsGovernanceAI

// NewOrgSettingsGovernanceAI instantiates a new OrgSettingsGovernanceAI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsGovernanceAI() *OrgSettingsGovernanceAI {
	this := OrgSettingsGovernanceAI{}
	return &this
}

// NewOrgSettingsGovernanceAIWithDefaults instantiates a new OrgSettingsGovernanceAI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsGovernanceAIWithDefaults() *OrgSettingsGovernanceAI {
	this := OrgSettingsGovernanceAI{}
	return &this
}

// GetSecurityAccessReview returns the SecurityAccessReview field value if set, zero value otherwise.
func (o *OrgSettingsGovernanceAI) GetSecurityAccessReview() OrgSettingsGovernanceAISecurityAccessReview {
	if o == nil || IsNil(o.SecurityAccessReview) {
		var ret OrgSettingsGovernanceAISecurityAccessReview
		return ret
	}
	return *o.SecurityAccessReview
}

// GetSecurityAccessReviewOk returns a tuple with the SecurityAccessReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsGovernanceAI) GetSecurityAccessReviewOk() (*OrgSettingsGovernanceAISecurityAccessReview, bool) {
	if o == nil || IsNil(o.SecurityAccessReview) {
		return nil, false
	}
	return o.SecurityAccessReview, true
}

// HasSecurityAccessReview returns a boolean if a field has been set.
func (o *OrgSettingsGovernanceAI) HasSecurityAccessReview() bool {
	if o != nil && !IsNil(o.SecurityAccessReview) {
		return true
	}

	return false
}

// SetSecurityAccessReview gets a reference to the given OrgSettingsGovernanceAISecurityAccessReview and assigns it to the SecurityAccessReview field.
func (o *OrgSettingsGovernanceAI) SetSecurityAccessReview(v OrgSettingsGovernanceAISecurityAccessReview) {
	o.SecurityAccessReview = &v
}

// GetGovernanceAnalyzer returns the GovernanceAnalyzer field value if set, zero value otherwise.
func (o *OrgSettingsGovernanceAI) GetGovernanceAnalyzer() OrgSettingsGovernanceAIGovernanceAnalyzer {
	if o == nil || IsNil(o.GovernanceAnalyzer) {
		var ret OrgSettingsGovernanceAIGovernanceAnalyzer
		return ret
	}
	return *o.GovernanceAnalyzer
}

// GetGovernanceAnalyzerOk returns a tuple with the GovernanceAnalyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsGovernanceAI) GetGovernanceAnalyzerOk() (*OrgSettingsGovernanceAIGovernanceAnalyzer, bool) {
	if o == nil || IsNil(o.GovernanceAnalyzer) {
		return nil, false
	}
	return o.GovernanceAnalyzer, true
}

// HasGovernanceAnalyzer returns a boolean if a field has been set.
func (o *OrgSettingsGovernanceAI) HasGovernanceAnalyzer() bool {
	if o != nil && !IsNil(o.GovernanceAnalyzer) {
		return true
	}

	return false
}

// SetGovernanceAnalyzer gets a reference to the given OrgSettingsGovernanceAIGovernanceAnalyzer and assigns it to the GovernanceAnalyzer field.
func (o *OrgSettingsGovernanceAI) SetGovernanceAnalyzer(v OrgSettingsGovernanceAIGovernanceAnalyzer) {
	o.GovernanceAnalyzer = &v
}

func (o OrgSettingsGovernanceAI) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsGovernanceAI) ToMap() (map[string]interface{}, error) {
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

func (o *OrgSettingsGovernanceAI) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsGovernanceAI := _OrgSettingsGovernanceAI{}

	err = json.Unmarshal(data, &varOrgSettingsGovernanceAI)

	if err != nil {
		return err
	}

	*o = OrgSettingsGovernanceAI(varOrgSettingsGovernanceAI)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "securityAccessReview")
		delete(additionalProperties, "governanceAnalyzer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsGovernanceAI struct {
	value *OrgSettingsGovernanceAI
	isSet bool
}

func (v NullableOrgSettingsGovernanceAI) Get() *OrgSettingsGovernanceAI {
	return v.value
}

func (v *NullableOrgSettingsGovernanceAI) Set(val *OrgSettingsGovernanceAI) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsGovernanceAI) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsGovernanceAI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsGovernanceAI(val *OrgSettingsGovernanceAI) *NullableOrgSettingsGovernanceAI {
	return &NullableOrgSettingsGovernanceAI{value: val, isSet: true}
}

func (v NullableOrgSettingsGovernanceAI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsGovernanceAI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
