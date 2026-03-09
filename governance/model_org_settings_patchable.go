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

// checks if the OrgSettingsPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsPatchable{}

// OrgSettingsPatchable struct for OrgSettingsPatchable
type OrgSettingsPatchable struct {
	Delegates            *OrgSettingsPatchableDelegates `json:"delegates,omitempty"`
	GovernanceAI         *OrgSettingsGovernanceAI       `json:"governanceAI,omitempty"`
	Escalations          *OrgSettingsEscalations        `json:"escalations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsPatchable OrgSettingsPatchable

// NewOrgSettingsPatchable instantiates a new OrgSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsPatchable() *OrgSettingsPatchable {
	this := OrgSettingsPatchable{}
	return &this
}

// NewOrgSettingsPatchableWithDefaults instantiates a new OrgSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsPatchableWithDefaults() *OrgSettingsPatchable {
	this := OrgSettingsPatchable{}
	return &this
}

// GetDelegates returns the Delegates field value if set, zero value otherwise.
func (o *OrgSettingsPatchable) GetDelegates() OrgSettingsPatchableDelegates {
	if o == nil || IsNil(o.Delegates) {
		var ret OrgSettingsPatchableDelegates
		return ret
	}
	return *o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchable) GetDelegatesOk() (*OrgSettingsPatchableDelegates, bool) {
	if o == nil || IsNil(o.Delegates) {
		return nil, false
	}
	return o.Delegates, true
}

// HasDelegates returns a boolean if a field has been set.
func (o *OrgSettingsPatchable) HasDelegates() bool {
	if o != nil && !IsNil(o.Delegates) {
		return true
	}

	return false
}

// SetDelegates gets a reference to the given OrgSettingsPatchableDelegates and assigns it to the Delegates field.
func (o *OrgSettingsPatchable) SetDelegates(v OrgSettingsPatchableDelegates) {
	o.Delegates = &v
}

// GetGovernanceAI returns the GovernanceAI field value if set, zero value otherwise.
func (o *OrgSettingsPatchable) GetGovernanceAI() OrgSettingsGovernanceAI {
	if o == nil || IsNil(o.GovernanceAI) {
		var ret OrgSettingsGovernanceAI
		return ret
	}
	return *o.GovernanceAI
}

// GetGovernanceAIOk returns a tuple with the GovernanceAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchable) GetGovernanceAIOk() (*OrgSettingsGovernanceAI, bool) {
	if o == nil || IsNil(o.GovernanceAI) {
		return nil, false
	}
	return o.GovernanceAI, true
}

// HasGovernanceAI returns a boolean if a field has been set.
func (o *OrgSettingsPatchable) HasGovernanceAI() bool {
	if o != nil && !IsNil(o.GovernanceAI) {
		return true
	}

	return false
}

// SetGovernanceAI gets a reference to the given OrgSettingsGovernanceAI and assigns it to the GovernanceAI field.
func (o *OrgSettingsPatchable) SetGovernanceAI(v OrgSettingsGovernanceAI) {
	o.GovernanceAI = &v
}

// GetEscalations returns the Escalations field value if set, zero value otherwise.
func (o *OrgSettingsPatchable) GetEscalations() OrgSettingsEscalations {
	if o == nil || IsNil(o.Escalations) {
		var ret OrgSettingsEscalations
		return ret
	}
	return *o.Escalations
}

// GetEscalationsOk returns a tuple with the Escalations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchable) GetEscalationsOk() (*OrgSettingsEscalations, bool) {
	if o == nil || IsNil(o.Escalations) {
		return nil, false
	}
	return o.Escalations, true
}

// HasEscalations returns a boolean if a field has been set.
func (o *OrgSettingsPatchable) HasEscalations() bool {
	if o != nil && !IsNil(o.Escalations) {
		return true
	}

	return false
}

// SetEscalations gets a reference to the given OrgSettingsEscalations and assigns it to the Escalations field.
func (o *OrgSettingsPatchable) SetEscalations(v OrgSettingsEscalations) {
	o.Escalations = &v
}

func (o OrgSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Delegates) {
		toSerialize["delegates"] = o.Delegates
	}
	if !IsNil(o.GovernanceAI) {
		toSerialize["governanceAI"] = o.GovernanceAI
	}
	if !IsNil(o.Escalations) {
		toSerialize["escalations"] = o.Escalations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsPatchable) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsPatchable := _OrgSettingsPatchable{}

	err = json.Unmarshal(data, &varOrgSettingsPatchable)

	if err != nil {
		return err
	}

	*o = OrgSettingsPatchable(varOrgSettingsPatchable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegates")
		delete(additionalProperties, "governanceAI")
		delete(additionalProperties, "escalations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsPatchable struct {
	value *OrgSettingsPatchable
	isSet bool
}

func (v NullableOrgSettingsPatchable) Get() *OrgSettingsPatchable {
	return v.value
}

func (v *NullableOrgSettingsPatchable) Set(val *OrgSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsPatchable(val *OrgSettingsPatchable) *NullableOrgSettingsPatchable {
	return &NullableOrgSettingsPatchable{value: val, isSet: true}
}

func (v NullableOrgSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
