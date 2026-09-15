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

// checks if the OrgCertificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCertificationSettings{}

// OrgCertificationSettings Certification settings for the org
type OrgCertificationSettings struct {
	Integrations         *Integrations `json:"integrations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCertificationSettings OrgCertificationSettings

// NewOrgCertificationSettings instantiates a new OrgCertificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCertificationSettings() *OrgCertificationSettings {
	this := OrgCertificationSettings{}
	return &this
}

// NewOrgCertificationSettingsWithDefaults instantiates a new OrgCertificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCertificationSettingsWithDefaults() *OrgCertificationSettings {
	this := OrgCertificationSettings{}
	return &this
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *OrgCertificationSettings) GetIntegrations() Integrations {
	if o == nil || IsNil(o.Integrations) {
		var ret Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCertificationSettings) GetIntegrationsOk() (*Integrations, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *OrgCertificationSettings) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given Integrations and assigns it to the Integrations field.
func (o *OrgCertificationSettings) SetIntegrations(v Integrations) {
	o.Integrations = &v
}

func (o OrgCertificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCertificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCertificationSettings) UnmarshalJSON(data []byte) (err error) {
	varOrgCertificationSettings := _OrgCertificationSettings{}

	err = json.Unmarshal(data, &varOrgCertificationSettings)

	if err != nil {
		return err
	}

	*o = OrgCertificationSettings(varOrgCertificationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "integrations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCertificationSettings struct {
	value *OrgCertificationSettings
	isSet bool
}

func (v NullableOrgCertificationSettings) Get() *OrgCertificationSettings {
	return v.value
}

func (v *NullableOrgCertificationSettings) Set(val *OrgCertificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCertificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCertificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCertificationSettings(val *OrgCertificationSettings) *NullableOrgCertificationSettings {
	return &NullableOrgCertificationSettings{value: val, isSet: true}
}

func (v NullableOrgCertificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCertificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
