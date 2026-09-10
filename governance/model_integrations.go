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

// checks if the Integrations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Integrations{}

// Integrations Access certification integration settings for the org
type Integrations struct {
	// Integration settings
	Settings             []IntegrationsSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Integrations Integrations

// NewIntegrations instantiates a new Integrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrations() *Integrations {
	this := Integrations{}
	return &this
}

// NewIntegrationsWithDefaults instantiates a new Integrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationsWithDefaults() *Integrations {
	this := Integrations{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Integrations) GetSettings() []IntegrationsSettings {
	if o == nil || IsNil(o.Settings) {
		var ret []IntegrationsSettings
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integrations) GetSettingsOk() ([]IntegrationsSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Integrations) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []IntegrationsSettings and assigns it to the Settings field.
func (o *Integrations) SetSettings(v []IntegrationsSettings) {
	o.Settings = v
}

func (o Integrations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Integrations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Integrations) UnmarshalJSON(data []byte) (err error) {
	varIntegrations := _Integrations{}

	err = json.Unmarshal(data, &varIntegrations)

	if err != nil {
		return err
	}

	*o = Integrations(varIntegrations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrations struct {
	value *Integrations
	isSet bool
}

func (v NullableIntegrations) Get() *Integrations {
	return v.value
}

func (v *NullableIntegrations) Set(val *Integrations) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrations(val *Integrations) *NullableIntegrations {
	return &NullableIntegrations{value: val, isSet: true}
}

func (v NullableIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
