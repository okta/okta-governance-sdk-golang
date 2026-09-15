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

// checks if the RequestIntegrations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestIntegrations{}

// RequestIntegrations <x-lifecycle class=\"beta\"></x-lifecycle><br> Access Requests integration settings for the org
type RequestIntegrations struct {
	// Integration settings
	Settings             []SlackIntegrationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestIntegrations RequestIntegrations

// NewRequestIntegrations instantiates a new RequestIntegrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestIntegrations() *RequestIntegrations {
	this := RequestIntegrations{}
	return &this
}

// NewRequestIntegrationsWithDefaults instantiates a new RequestIntegrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestIntegrationsWithDefaults() *RequestIntegrations {
	this := RequestIntegrations{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *RequestIntegrations) GetSettings() []SlackIntegrationSettings {
	if o == nil || IsNil(o.Settings) {
		var ret []SlackIntegrationSettings
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestIntegrations) GetSettingsOk() ([]SlackIntegrationSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *RequestIntegrations) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []SlackIntegrationSettings and assigns it to the Settings field.
func (o *RequestIntegrations) SetSettings(v []SlackIntegrationSettings) {
	o.Settings = v
}

func (o RequestIntegrations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestIntegrations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestIntegrations) UnmarshalJSON(data []byte) (err error) {
	varRequestIntegrations := _RequestIntegrations{}

	err = json.Unmarshal(data, &varRequestIntegrations)

	if err != nil {
		return err
	}

	*o = RequestIntegrations(varRequestIntegrations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestIntegrations struct {
	value *RequestIntegrations
	isSet bool
}

func (v NullableRequestIntegrations) Get() *RequestIntegrations {
	return v.value
}

func (v *NullableRequestIntegrations) Set(val *RequestIntegrations) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestIntegrations(val *RequestIntegrations) *NullableRequestIntegrations {
	return &NullableRequestIntegrations{value: val, isSet: true}
}

func (v NullableRequestIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
