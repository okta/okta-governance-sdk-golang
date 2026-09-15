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

// checks if the OrgCertificationSettingsPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCertificationSettingsPatchable{}

// OrgCertificationSettingsPatchable Update certification settings for the org
type OrgCertificationSettingsPatchable struct {
	Integrations         *Integrations `json:"integrations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCertificationSettingsPatchable OrgCertificationSettingsPatchable

// NewOrgCertificationSettingsPatchable instantiates a new OrgCertificationSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCertificationSettingsPatchable() *OrgCertificationSettingsPatchable {
	this := OrgCertificationSettingsPatchable{}
	return &this
}

// NewOrgCertificationSettingsPatchableWithDefaults instantiates a new OrgCertificationSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCertificationSettingsPatchableWithDefaults() *OrgCertificationSettingsPatchable {
	this := OrgCertificationSettingsPatchable{}
	return &this
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *OrgCertificationSettingsPatchable) GetIntegrations() Integrations {
	if o == nil || IsNil(o.Integrations) {
		var ret Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCertificationSettingsPatchable) GetIntegrationsOk() (*Integrations, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *OrgCertificationSettingsPatchable) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given Integrations and assigns it to the Integrations field.
func (o *OrgCertificationSettingsPatchable) SetIntegrations(v Integrations) {
	o.Integrations = &v
}

func (o OrgCertificationSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCertificationSettingsPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCertificationSettingsPatchable) UnmarshalJSON(data []byte) (err error) {
	varOrgCertificationSettingsPatchable := _OrgCertificationSettingsPatchable{}

	err = json.Unmarshal(data, &varOrgCertificationSettingsPatchable)

	if err != nil {
		return err
	}

	*o = OrgCertificationSettingsPatchable(varOrgCertificationSettingsPatchable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "integrations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCertificationSettingsPatchable struct {
	value *OrgCertificationSettingsPatchable
	isSet bool
}

func (v NullableOrgCertificationSettingsPatchable) Get() *OrgCertificationSettingsPatchable {
	return v.value
}

func (v *NullableOrgCertificationSettingsPatchable) Set(val *OrgCertificationSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCertificationSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCertificationSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCertificationSettingsPatchable(val *OrgCertificationSettingsPatchable) *NullableOrgCertificationSettingsPatchable {
	return &NullableOrgCertificationSettingsPatchable{value: val, isSet: true}
}

func (v NullableOrgCertificationSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCertificationSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
