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

// checks if the OrgSettingsIntegrations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsIntegrations{}

// OrgSettingsIntegrations Integration settings
type OrgSettingsIntegrations struct {
	// Supported integrations in this org
	Supported            []OrgSettingsIntegrationsSupportedInner `json:"supported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsIntegrations OrgSettingsIntegrations

// NewOrgSettingsIntegrations instantiates a new OrgSettingsIntegrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsIntegrations() *OrgSettingsIntegrations {
	this := OrgSettingsIntegrations{}
	return &this
}

// NewOrgSettingsIntegrationsWithDefaults instantiates a new OrgSettingsIntegrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsIntegrationsWithDefaults() *OrgSettingsIntegrations {
	this := OrgSettingsIntegrations{}
	return &this
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *OrgSettingsIntegrations) GetSupported() []OrgSettingsIntegrationsSupportedInner {
	if o == nil || IsNil(o.Supported) {
		var ret []OrgSettingsIntegrationsSupportedInner
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsIntegrations) GetSupportedOk() ([]OrgSettingsIntegrationsSupportedInner, bool) {
	if o == nil || IsNil(o.Supported) {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *OrgSettingsIntegrations) HasSupported() bool {
	if o != nil && !IsNil(o.Supported) {
		return true
	}

	return false
}

// SetSupported gets a reference to the given []OrgSettingsIntegrationsSupportedInner and assigns it to the Supported field.
func (o *OrgSettingsIntegrations) SetSupported(v []OrgSettingsIntegrationsSupportedInner) {
	o.Supported = v
}

func (o OrgSettingsIntegrations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsIntegrations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supported) {
		toSerialize["supported"] = o.Supported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsIntegrations) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsIntegrations := _OrgSettingsIntegrations{}

	err = json.Unmarshal(data, &varOrgSettingsIntegrations)

	if err != nil {
		return err
	}

	*o = OrgSettingsIntegrations(varOrgSettingsIntegrations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "supported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsIntegrations struct {
	value *OrgSettingsIntegrations
	isSet bool
}

func (v NullableOrgSettingsIntegrations) Get() *OrgSettingsIntegrations {
	return v.value
}

func (v *NullableOrgSettingsIntegrations) Set(val *OrgSettingsIntegrations) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsIntegrations(val *OrgSettingsIntegrations) *NullableOrgSettingsIntegrations {
	return &NullableOrgSettingsIntegrations{value: val, isSet: true}
}

func (v NullableOrgSettingsIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
