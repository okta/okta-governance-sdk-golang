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

// checks if the ResourceRequestSettingsPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRequestSettingsPatchable{}

// ResourceRequestSettingsPatchable Resource request settings that specify various configurations and permissions at the resource level.
type ResourceRequestSettingsPatchable struct {
	RequestOnBehalfOfSettings NullableRequestOnBehalfOfSettingsPatchable `json:"requestOnBehalfOfSettings,omitempty"`
	RiskSettings              *RiskSettingsPatchable                     `json:"riskSettings,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _ResourceRequestSettingsPatchable ResourceRequestSettingsPatchable

// NewResourceRequestSettingsPatchable instantiates a new ResourceRequestSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRequestSettingsPatchable() *ResourceRequestSettingsPatchable {
	this := ResourceRequestSettingsPatchable{}
	return &this
}

// NewResourceRequestSettingsPatchableWithDefaults instantiates a new ResourceRequestSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRequestSettingsPatchableWithDefaults() *ResourceRequestSettingsPatchable {
	this := ResourceRequestSettingsPatchable{}
	return &this
}

// GetRequestOnBehalfOfSettings returns the RequestOnBehalfOfSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceRequestSettingsPatchable) GetRequestOnBehalfOfSettings() RequestOnBehalfOfSettingsPatchable {
	if o == nil || IsNil(o.RequestOnBehalfOfSettings.Get()) {
		var ret RequestOnBehalfOfSettingsPatchable
		return ret
	}
	return *o.RequestOnBehalfOfSettings.Get()
}

// GetRequestOnBehalfOfSettingsOk returns a tuple with the RequestOnBehalfOfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceRequestSettingsPatchable) GetRequestOnBehalfOfSettingsOk() (*RequestOnBehalfOfSettingsPatchable, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestOnBehalfOfSettings.Get(), o.RequestOnBehalfOfSettings.IsSet()
}

// HasRequestOnBehalfOfSettings returns a boolean if a field has been set.
func (o *ResourceRequestSettingsPatchable) HasRequestOnBehalfOfSettings() bool {
	if o != nil && o.RequestOnBehalfOfSettings.IsSet() {
		return true
	}

	return false
}

// SetRequestOnBehalfOfSettings gets a reference to the given NullableRequestOnBehalfOfSettingsPatchable and assigns it to the RequestOnBehalfOfSettings field.
func (o *ResourceRequestSettingsPatchable) SetRequestOnBehalfOfSettings(v RequestOnBehalfOfSettingsPatchable) {
	o.RequestOnBehalfOfSettings.Set(&v)
}

// SetRequestOnBehalfOfSettingsNil sets the value for RequestOnBehalfOfSettings to be an explicit nil
func (o *ResourceRequestSettingsPatchable) SetRequestOnBehalfOfSettingsNil() {
	o.RequestOnBehalfOfSettings.Set(nil)
}

// UnsetRequestOnBehalfOfSettings ensures that no value is present for RequestOnBehalfOfSettings, not even an explicit nil
func (o *ResourceRequestSettingsPatchable) UnsetRequestOnBehalfOfSettings() {
	o.RequestOnBehalfOfSettings.Unset()
}

// GetRiskSettings returns the RiskSettings field value if set, zero value otherwise.
func (o *ResourceRequestSettingsPatchable) GetRiskSettings() RiskSettingsPatchable {
	if o == nil || IsNil(o.RiskSettings) {
		var ret RiskSettingsPatchable
		return ret
	}
	return *o.RiskSettings
}

// GetRiskSettingsOk returns a tuple with the RiskSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRequestSettingsPatchable) GetRiskSettingsOk() (*RiskSettingsPatchable, bool) {
	if o == nil || IsNil(o.RiskSettings) {
		return nil, false
	}
	return o.RiskSettings, true
}

// HasRiskSettings returns a boolean if a field has been set.
func (o *ResourceRequestSettingsPatchable) HasRiskSettings() bool {
	if o != nil && !IsNil(o.RiskSettings) {
		return true
	}

	return false
}

// SetRiskSettings gets a reference to the given RiskSettingsPatchable and assigns it to the RiskSettings field.
func (o *ResourceRequestSettingsPatchable) SetRiskSettings(v RiskSettingsPatchable) {
	o.RiskSettings = &v
}

func (o ResourceRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRequestSettingsPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestOnBehalfOfSettings.IsSet() {
		toSerialize["requestOnBehalfOfSettings"] = o.RequestOnBehalfOfSettings.Get()
	}
	if !IsNil(o.RiskSettings) {
		toSerialize["riskSettings"] = o.RiskSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceRequestSettingsPatchable) UnmarshalJSON(data []byte) (err error) {
	varResourceRequestSettingsPatchable := _ResourceRequestSettingsPatchable{}

	err = json.Unmarshal(data, &varResourceRequestSettingsPatchable)

	if err != nil {
		return err
	}

	*o = ResourceRequestSettingsPatchable(varResourceRequestSettingsPatchable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestOnBehalfOfSettings")
		delete(additionalProperties, "riskSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceRequestSettingsPatchable struct {
	value *ResourceRequestSettingsPatchable
	isSet bool
}

func (v NullableResourceRequestSettingsPatchable) Get() *ResourceRequestSettingsPatchable {
	return v.value
}

func (v *NullableResourceRequestSettingsPatchable) Set(val *ResourceRequestSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRequestSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRequestSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRequestSettingsPatchable(val *ResourceRequestSettingsPatchable) *NullableResourceRequestSettingsPatchable {
	return &NullableResourceRequestSettingsPatchable{value: val, isSet: true}
}

func (v NullableResourceRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRequestSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
