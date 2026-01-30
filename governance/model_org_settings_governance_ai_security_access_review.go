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

// checks if the OrgSettingsGovernanceAISecurityAccessReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsGovernanceAISecurityAccessReview{}

// OrgSettingsGovernanceAISecurityAccessReview Settings for Security Access Review powered by Governance AI
type OrgSettingsGovernanceAISecurityAccessReview struct {
	// Indicates if LLM-generated insights are enabled for security access reviews. If `true`, users can generate AI-powered summaries to help assess risk.
	Enabled              *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsGovernanceAISecurityAccessReview OrgSettingsGovernanceAISecurityAccessReview

// NewOrgSettingsGovernanceAISecurityAccessReview instantiates a new OrgSettingsGovernanceAISecurityAccessReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsGovernanceAISecurityAccessReview() *OrgSettingsGovernanceAISecurityAccessReview {
	this := OrgSettingsGovernanceAISecurityAccessReview{}
	return &this
}

// NewOrgSettingsGovernanceAISecurityAccessReviewWithDefaults instantiates a new OrgSettingsGovernanceAISecurityAccessReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsGovernanceAISecurityAccessReviewWithDefaults() *OrgSettingsGovernanceAISecurityAccessReview {
	this := OrgSettingsGovernanceAISecurityAccessReview{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrgSettingsGovernanceAISecurityAccessReview) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsGovernanceAISecurityAccessReview) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrgSettingsGovernanceAISecurityAccessReview) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrgSettingsGovernanceAISecurityAccessReview) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrgSettingsGovernanceAISecurityAccessReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsGovernanceAISecurityAccessReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsGovernanceAISecurityAccessReview) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsGovernanceAISecurityAccessReview := _OrgSettingsGovernanceAISecurityAccessReview{}

	err = json.Unmarshal(data, &varOrgSettingsGovernanceAISecurityAccessReview)

	if err != nil {
		return err
	}

	*o = OrgSettingsGovernanceAISecurityAccessReview(varOrgSettingsGovernanceAISecurityAccessReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsGovernanceAISecurityAccessReview struct {
	value *OrgSettingsGovernanceAISecurityAccessReview
	isSet bool
}

func (v NullableOrgSettingsGovernanceAISecurityAccessReview) Get() *OrgSettingsGovernanceAISecurityAccessReview {
	return v.value
}

func (v *NullableOrgSettingsGovernanceAISecurityAccessReview) Set(val *OrgSettingsGovernanceAISecurityAccessReview) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsGovernanceAISecurityAccessReview) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsGovernanceAISecurityAccessReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsGovernanceAISecurityAccessReview(val *OrgSettingsGovernanceAISecurityAccessReview) *NullableOrgSettingsGovernanceAISecurityAccessReview {
	return &NullableOrgSettingsGovernanceAISecurityAccessReview{value: val, isSet: true}
}

func (v NullableOrgSettingsGovernanceAISecurityAccessReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsGovernanceAISecurityAccessReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
