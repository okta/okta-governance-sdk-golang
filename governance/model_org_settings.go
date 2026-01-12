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

// checks if the OrgSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettings{}

// OrgSettings struct for OrgSettings
type OrgSettings struct {
	Delegates            *OrgSettingsDelegates    `json:"delegates,omitempty"`
	GovernanceAI         *OrgSettingsGovernanceAI `json:"governanceAI,omitempty"`
	Escalations          *OrgSettingsEscalations  `json:"escalations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettings OrgSettings

// NewOrgSettings instantiates a new OrgSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettings() *OrgSettings {
	this := OrgSettings{}
	return &this
}

// NewOrgSettingsWithDefaults instantiates a new OrgSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsWithDefaults() *OrgSettings {
	this := OrgSettings{}
	return &this
}

// GetDelegates returns the Delegates field value if set, zero value otherwise.
func (o *OrgSettings) GetDelegates() OrgSettingsDelegates {
	if o == nil || IsNil(o.Delegates) {
		var ret OrgSettingsDelegates
		return ret
	}
	return *o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettings) GetDelegatesOk() (*OrgSettingsDelegates, bool) {
	if o == nil || IsNil(o.Delegates) {
		return nil, false
	}
	return o.Delegates, true
}

// HasDelegates returns a boolean if a field has been set.
func (o *OrgSettings) HasDelegates() bool {
	if o != nil && !IsNil(o.Delegates) {
		return true
	}

	return false
}

// SetDelegates gets a reference to the given OrgSettingsDelegates and assigns it to the Delegates field.
func (o *OrgSettings) SetDelegates(v OrgSettingsDelegates) {
	o.Delegates = &v
}

// GetGovernanceAI returns the GovernanceAI field value if set, zero value otherwise.
func (o *OrgSettings) GetGovernanceAI() OrgSettingsGovernanceAI {
	if o == nil || IsNil(o.GovernanceAI) {
		var ret OrgSettingsGovernanceAI
		return ret
	}
	return *o.GovernanceAI
}

// GetGovernanceAIOk returns a tuple with the GovernanceAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettings) GetGovernanceAIOk() (*OrgSettingsGovernanceAI, bool) {
	if o == nil || IsNil(o.GovernanceAI) {
		return nil, false
	}
	return o.GovernanceAI, true
}

// HasGovernanceAI returns a boolean if a field has been set.
func (o *OrgSettings) HasGovernanceAI() bool {
	if o != nil && !IsNil(o.GovernanceAI) {
		return true
	}

	return false
}

// SetGovernanceAI gets a reference to the given OrgSettingsGovernanceAI and assigns it to the GovernanceAI field.
func (o *OrgSettings) SetGovernanceAI(v OrgSettingsGovernanceAI) {
	o.GovernanceAI = &v
}

// GetEscalations returns the Escalations field value if set, zero value otherwise.
func (o *OrgSettings) GetEscalations() OrgSettingsEscalations {
	if o == nil || IsNil(o.Escalations) {
		var ret OrgSettingsEscalations
		return ret
	}
	return *o.Escalations
}

// GetEscalationsOk returns a tuple with the Escalations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettings) GetEscalationsOk() (*OrgSettingsEscalations, bool) {
	if o == nil || IsNil(o.Escalations) {
		return nil, false
	}
	return o.Escalations, true
}

// HasEscalations returns a boolean if a field has been set.
func (o *OrgSettings) HasEscalations() bool {
	if o != nil && !IsNil(o.Escalations) {
		return true
	}

	return false
}

// SetEscalations gets a reference to the given OrgSettingsEscalations and assigns it to the Escalations field.
func (o *OrgSettings) SetEscalations(v OrgSettingsEscalations) {
	o.Escalations = &v
}

func (o OrgSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettings) ToMap() (map[string]interface{}, error) {
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

func (o *OrgSettings) UnmarshalJSON(data []byte) (err error) {
	varOrgSettings := _OrgSettings{}

	err = json.Unmarshal(data, &varOrgSettings)

	if err != nil {
		return err
	}

	*o = OrgSettings(varOrgSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegates")
		delete(additionalProperties, "governanceAI")
		delete(additionalProperties, "escalations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettings struct {
	value *OrgSettings
	isSet bool
}

func (v NullableOrgSettings) Get() *OrgSettings {
	return v.value
}

func (v *NullableOrgSettings) Set(val *OrgSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettings(val *OrgSettings) *NullableOrgSettings {
	return &NullableOrgSettings{value: val, isSet: true}
}

func (v NullableOrgSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
