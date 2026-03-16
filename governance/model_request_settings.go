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
	"fmt"
)

// checks if the RequestSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestSettings{}

// RequestSettings Request settings specify what conditions are valid for the given resource
type RequestSettings struct {
	// Access scope settings that are eligible to be added to a request condition for the specified resource
	ValidAccessScopeSettings []ValidAccessDetail `json:"validAccessScopeSettings"`
	// Request scope settings that are eligible to be added to a request condition for the specified resource
	ValidRequesterSettings      []ValidRequesterSetting            `json:"validRequesterSettings"`
	ValidAccessDurationSettings ValidAccessDurationSettingsDetails `json:"validAccessDurationSettings"`
	RequestOnBehalfOfSettings   *RequestOnBehalfOfSettingsDetails  `json:"requestOnBehalfOfSettings,omitempty"`
	ValidRiskSettings           *ValidRiskSettingsDetails          `json:"validRiskSettings,omitempty"`
	RiskSettings                *RiskSettingsDetails               `json:"riskSettings,omitempty"`
	// Request on behalf of settings that are eligible to be added to a request condition for the specified resource
	ValidRequestOnBehalfOfSettings []ValidRequestOnBehalfOfSetting `json:"validRequestOnBehalfOfSettings,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _RequestSettings RequestSettings

// NewRequestSettings instantiates a new RequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSettings(validAccessScopeSettings []ValidAccessDetail, validRequesterSettings []ValidRequesterSetting, validAccessDurationSettings ValidAccessDurationSettingsDetails) *RequestSettings {
	this := RequestSettings{}
	this.ValidAccessScopeSettings = validAccessScopeSettings
	this.ValidRequesterSettings = validRequesterSettings
	this.ValidAccessDurationSettings = validAccessDurationSettings
	return &this
}

// NewRequestSettingsWithDefaults instantiates a new RequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSettingsWithDefaults() *RequestSettings {
	this := RequestSettings{}
	return &this
}

// GetValidAccessScopeSettings returns the ValidAccessScopeSettings field value
func (o *RequestSettings) GetValidAccessScopeSettings() []ValidAccessDetail {
	if o == nil {
		var ret []ValidAccessDetail
		return ret
	}

	return o.ValidAccessScopeSettings
}

// GetValidAccessScopeSettingsOk returns a tuple with the ValidAccessScopeSettings field value
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetValidAccessScopeSettingsOk() ([]ValidAccessDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidAccessScopeSettings, true
}

// SetValidAccessScopeSettings sets field value
func (o *RequestSettings) SetValidAccessScopeSettings(v []ValidAccessDetail) {
	o.ValidAccessScopeSettings = v
}

// GetValidRequesterSettings returns the ValidRequesterSettings field value
func (o *RequestSettings) GetValidRequesterSettings() []ValidRequesterSetting {
	if o == nil {
		var ret []ValidRequesterSetting
		return ret
	}

	return o.ValidRequesterSettings
}

// GetValidRequesterSettingsOk returns a tuple with the ValidRequesterSettings field value
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetValidRequesterSettingsOk() ([]ValidRequesterSetting, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidRequesterSettings, true
}

// SetValidRequesterSettings sets field value
func (o *RequestSettings) SetValidRequesterSettings(v []ValidRequesterSetting) {
	o.ValidRequesterSettings = v
}

// GetValidAccessDurationSettings returns the ValidAccessDurationSettings field value
func (o *RequestSettings) GetValidAccessDurationSettings() ValidAccessDurationSettingsDetails {
	if o == nil {
		var ret ValidAccessDurationSettingsDetails
		return ret
	}

	return o.ValidAccessDurationSettings
}

// GetValidAccessDurationSettingsOk returns a tuple with the ValidAccessDurationSettings field value
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetValidAccessDurationSettingsOk() (*ValidAccessDurationSettingsDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidAccessDurationSettings, true
}

// SetValidAccessDurationSettings sets field value
func (o *RequestSettings) SetValidAccessDurationSettings(v ValidAccessDurationSettingsDetails) {
	o.ValidAccessDurationSettings = v
}

// GetRequestOnBehalfOfSettings returns the RequestOnBehalfOfSettings field value if set, zero value otherwise.
func (o *RequestSettings) GetRequestOnBehalfOfSettings() RequestOnBehalfOfSettingsDetails {
	if o == nil || IsNil(o.RequestOnBehalfOfSettings) {
		var ret RequestOnBehalfOfSettingsDetails
		return ret
	}
	return *o.RequestOnBehalfOfSettings
}

// GetRequestOnBehalfOfSettingsOk returns a tuple with the RequestOnBehalfOfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetRequestOnBehalfOfSettingsOk() (*RequestOnBehalfOfSettingsDetails, bool) {
	if o == nil || IsNil(o.RequestOnBehalfOfSettings) {
		return nil, false
	}
	return o.RequestOnBehalfOfSettings, true
}

// HasRequestOnBehalfOfSettings returns a boolean if a field has been set.
func (o *RequestSettings) HasRequestOnBehalfOfSettings() bool {
	if o != nil && !IsNil(o.RequestOnBehalfOfSettings) {
		return true
	}

	return false
}

// SetRequestOnBehalfOfSettings gets a reference to the given RequestOnBehalfOfSettingsDetails and assigns it to the RequestOnBehalfOfSettings field.
func (o *RequestSettings) SetRequestOnBehalfOfSettings(v RequestOnBehalfOfSettingsDetails) {
	o.RequestOnBehalfOfSettings = &v
}

// GetValidRiskSettings returns the ValidRiskSettings field value if set, zero value otherwise.
func (o *RequestSettings) GetValidRiskSettings() ValidRiskSettingsDetails {
	if o == nil || IsNil(o.ValidRiskSettings) {
		var ret ValidRiskSettingsDetails
		return ret
	}
	return *o.ValidRiskSettings
}

// GetValidRiskSettingsOk returns a tuple with the ValidRiskSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetValidRiskSettingsOk() (*ValidRiskSettingsDetails, bool) {
	if o == nil || IsNil(o.ValidRiskSettings) {
		return nil, false
	}
	return o.ValidRiskSettings, true
}

// HasValidRiskSettings returns a boolean if a field has been set.
func (o *RequestSettings) HasValidRiskSettings() bool {
	if o != nil && !IsNil(o.ValidRiskSettings) {
		return true
	}

	return false
}

// SetValidRiskSettings gets a reference to the given ValidRiskSettingsDetails and assigns it to the ValidRiskSettings field.
func (o *RequestSettings) SetValidRiskSettings(v ValidRiskSettingsDetails) {
	o.ValidRiskSettings = &v
}

// GetRiskSettings returns the RiskSettings field value if set, zero value otherwise.
func (o *RequestSettings) GetRiskSettings() RiskSettingsDetails {
	if o == nil || IsNil(o.RiskSettings) {
		var ret RiskSettingsDetails
		return ret
	}
	return *o.RiskSettings
}

// GetRiskSettingsOk returns a tuple with the RiskSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetRiskSettingsOk() (*RiskSettingsDetails, bool) {
	if o == nil || IsNil(o.RiskSettings) {
		return nil, false
	}
	return o.RiskSettings, true
}

// HasRiskSettings returns a boolean if a field has been set.
func (o *RequestSettings) HasRiskSettings() bool {
	if o != nil && !IsNil(o.RiskSettings) {
		return true
	}

	return false
}

// SetRiskSettings gets a reference to the given RiskSettingsDetails and assigns it to the RiskSettings field.
func (o *RequestSettings) SetRiskSettings(v RiskSettingsDetails) {
	o.RiskSettings = &v
}

// GetValidRequestOnBehalfOfSettings returns the ValidRequestOnBehalfOfSettings field value if set, zero value otherwise.
func (o *RequestSettings) GetValidRequestOnBehalfOfSettings() []ValidRequestOnBehalfOfSetting {
	if o == nil || IsNil(o.ValidRequestOnBehalfOfSettings) {
		var ret []ValidRequestOnBehalfOfSetting
		return ret
	}
	return o.ValidRequestOnBehalfOfSettings
}

// GetValidRequestOnBehalfOfSettingsOk returns a tuple with the ValidRequestOnBehalfOfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetValidRequestOnBehalfOfSettingsOk() ([]ValidRequestOnBehalfOfSetting, bool) {
	if o == nil || IsNil(o.ValidRequestOnBehalfOfSettings) {
		return nil, false
	}
	return o.ValidRequestOnBehalfOfSettings, true
}

// HasValidRequestOnBehalfOfSettings returns a boolean if a field has been set.
func (o *RequestSettings) HasValidRequestOnBehalfOfSettings() bool {
	if o != nil && !IsNil(o.ValidRequestOnBehalfOfSettings) {
		return true
	}

	return false
}

// SetValidRequestOnBehalfOfSettings gets a reference to the given []ValidRequestOnBehalfOfSetting and assigns it to the ValidRequestOnBehalfOfSettings field.
func (o *RequestSettings) SetValidRequestOnBehalfOfSettings(v []ValidRequestOnBehalfOfSetting) {
	o.ValidRequestOnBehalfOfSettings = v
}

func (o RequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validAccessScopeSettings"] = o.ValidAccessScopeSettings
	toSerialize["validRequesterSettings"] = o.ValidRequesterSettings
	toSerialize["validAccessDurationSettings"] = o.ValidAccessDurationSettings
	if !IsNil(o.RequestOnBehalfOfSettings) {
		toSerialize["requestOnBehalfOfSettings"] = o.RequestOnBehalfOfSettings
	}
	if !IsNil(o.ValidRiskSettings) {
		toSerialize["validRiskSettings"] = o.ValidRiskSettings
	}
	if !IsNil(o.RiskSettings) {
		toSerialize["riskSettings"] = o.RiskSettings
	}
	if !IsNil(o.ValidRequestOnBehalfOfSettings) {
		toSerialize["validRequestOnBehalfOfSettings"] = o.ValidRequestOnBehalfOfSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"validAccessScopeSettings",
		"validRequesterSettings",
		"validAccessDurationSettings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRequestSettings := _RequestSettings{}

	err = json.Unmarshal(data, &varRequestSettings)

	if err != nil {
		return err
	}

	*o = RequestSettings(varRequestSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "validAccessScopeSettings")
		delete(additionalProperties, "validRequesterSettings")
		delete(additionalProperties, "validAccessDurationSettings")
		delete(additionalProperties, "requestOnBehalfOfSettings")
		delete(additionalProperties, "validRiskSettings")
		delete(additionalProperties, "riskSettings")
		delete(additionalProperties, "validRequestOnBehalfOfSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSettings struct {
	value *RequestSettings
	isSet bool
}

func (v NullableRequestSettings) Get() *RequestSettings {
	return v.value
}

func (v *NullableRequestSettings) Set(val *RequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSettings(val *RequestSettings) *NullableRequestSettings {
	return &NullableRequestSettings{value: val, isSet: true}
}

func (v NullableRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
