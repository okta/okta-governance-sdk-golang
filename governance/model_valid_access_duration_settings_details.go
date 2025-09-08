/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

// ValidAccessDurationSettingsDetails Admin specified access duration settings added to a request condition or risk settings for the specified resource
type ValidAccessDurationSettingsDetails struct {
	// Access duration settings that are eligible to be added to a request condition or risk settings for the specified resource.
	SupportedTypes []ValidAccessDurationType `json:"supportedTypes,omitempty"`
	// Whether `accessDurationSetting` must be included in the request conditions or risk settings for the specified resource.
	Required *bool `json:"required,omitempty"`
	// The minimum and maximum values allowed in `accessDurationSettings.duration` expression for a request condition or risk settings, when that expression contains day units. For example: `P5D` is valid, and `P91D` is not.
	MaximumDays *float32 `json:"maximumDays,omitempty"`
	// The minimum and maximum values allowed in `accessDurationSettings.duration` expression for a request condition or risk settings, when that expression contains hour units. For example: `P5H` is valid, and `P73H` is not.
	MaximumHours *float32 `json:"maximumHours,omitempty"`
	// The minimum and maximum values allowed in `accessDurationSettings.duration` expression for a request condition or risk settings, when that expression contains week units. For example: `P5W` is valid, and `P13W` is not.
	MaximumWeeks         *float32 `json:"maximumWeeks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidAccessDurationSettingsDetails ValidAccessDurationSettingsDetails

// NewValidAccessDurationSettingsDetails instantiates a new ValidAccessDurationSettingsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidAccessDurationSettingsDetails() *ValidAccessDurationSettingsDetails {
	this := ValidAccessDurationSettingsDetails{}
	return &this
}

// NewValidAccessDurationSettingsDetailsWithDefaults instantiates a new ValidAccessDurationSettingsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidAccessDurationSettingsDetailsWithDefaults() *ValidAccessDurationSettingsDetails {
	this := ValidAccessDurationSettingsDetails{}
	return &this
}

// GetSupportedTypes returns the SupportedTypes field value if set, zero value otherwise.
func (o *ValidAccessDurationSettingsDetails) GetSupportedTypes() []ValidAccessDurationType {
	if o == nil || o.SupportedTypes == nil {
		var ret []ValidAccessDurationType
		return ret
	}
	return o.SupportedTypes
}

// GetSupportedTypesOk returns a tuple with the SupportedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationSettingsDetails) GetSupportedTypesOk() ([]ValidAccessDurationType, bool) {
	if o == nil || o.SupportedTypes == nil {
		return nil, false
	}
	return o.SupportedTypes, true
}

// HasSupportedTypes returns a boolean if a field has been set.
func (o *ValidAccessDurationSettingsDetails) HasSupportedTypes() bool {
	if o != nil && o.SupportedTypes != nil {
		return true
	}

	return false
}

// SetSupportedTypes gets a reference to the given []ValidAccessDurationType and assigns it to the SupportedTypes field.
func (o *ValidAccessDurationSettingsDetails) SetSupportedTypes(v []ValidAccessDurationType) {
	o.SupportedTypes = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ValidAccessDurationSettingsDetails) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationSettingsDetails) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ValidAccessDurationSettingsDetails) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ValidAccessDurationSettingsDetails) SetRequired(v bool) {
	o.Required = &v
}

// GetMaximumDays returns the MaximumDays field value if set, zero value otherwise.
func (o *ValidAccessDurationSettingsDetails) GetMaximumDays() float32 {
	if o == nil || o.MaximumDays == nil {
		var ret float32
		return ret
	}
	return *o.MaximumDays
}

// GetMaximumDaysOk returns a tuple with the MaximumDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationSettingsDetails) GetMaximumDaysOk() (*float32, bool) {
	if o == nil || o.MaximumDays == nil {
		return nil, false
	}
	return o.MaximumDays, true
}

// HasMaximumDays returns a boolean if a field has been set.
func (o *ValidAccessDurationSettingsDetails) HasMaximumDays() bool {
	if o != nil && o.MaximumDays != nil {
		return true
	}

	return false
}

// SetMaximumDays gets a reference to the given float32 and assigns it to the MaximumDays field.
func (o *ValidAccessDurationSettingsDetails) SetMaximumDays(v float32) {
	o.MaximumDays = &v
}

// GetMaximumHours returns the MaximumHours field value if set, zero value otherwise.
func (o *ValidAccessDurationSettingsDetails) GetMaximumHours() float32 {
	if o == nil || o.MaximumHours == nil {
		var ret float32
		return ret
	}
	return *o.MaximumHours
}

// GetMaximumHoursOk returns a tuple with the MaximumHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationSettingsDetails) GetMaximumHoursOk() (*float32, bool) {
	if o == nil || o.MaximumHours == nil {
		return nil, false
	}
	return o.MaximumHours, true
}

// HasMaximumHours returns a boolean if a field has been set.
func (o *ValidAccessDurationSettingsDetails) HasMaximumHours() bool {
	if o != nil && o.MaximumHours != nil {
		return true
	}

	return false
}

// SetMaximumHours gets a reference to the given float32 and assigns it to the MaximumHours field.
func (o *ValidAccessDurationSettingsDetails) SetMaximumHours(v float32) {
	o.MaximumHours = &v
}

// GetMaximumWeeks returns the MaximumWeeks field value if set, zero value otherwise.
func (o *ValidAccessDurationSettingsDetails) GetMaximumWeeks() float32 {
	if o == nil || o.MaximumWeeks == nil {
		var ret float32
		return ret
	}
	return *o.MaximumWeeks
}

// GetMaximumWeeksOk returns a tuple with the MaximumWeeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationSettingsDetails) GetMaximumWeeksOk() (*float32, bool) {
	if o == nil || o.MaximumWeeks == nil {
		return nil, false
	}
	return o.MaximumWeeks, true
}

// HasMaximumWeeks returns a boolean if a field has been set.
func (o *ValidAccessDurationSettingsDetails) HasMaximumWeeks() bool {
	if o != nil && o.MaximumWeeks != nil {
		return true
	}

	return false
}

// SetMaximumWeeks gets a reference to the given float32 and assigns it to the MaximumWeeks field.
func (o *ValidAccessDurationSettingsDetails) SetMaximumWeeks(v float32) {
	o.MaximumWeeks = &v
}

func (o ValidAccessDurationSettingsDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportedTypes != nil {
		toSerialize["supportedTypes"] = o.SupportedTypes
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.MaximumDays != nil {
		toSerialize["maximumDays"] = o.MaximumDays
	}
	if o.MaximumHours != nil {
		toSerialize["maximumHours"] = o.MaximumHours
	}
	if o.MaximumWeeks != nil {
		toSerialize["maximumWeeks"] = o.MaximumWeeks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidAccessDurationSettingsDetails) UnmarshalJSON(bytes []byte) (err error) {
	varValidAccessDurationSettingsDetails := _ValidAccessDurationSettingsDetails{}

	err = json.Unmarshal(bytes, &varValidAccessDurationSettingsDetails)
	if err == nil {
		*o = ValidAccessDurationSettingsDetails(varValidAccessDurationSettingsDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "supportedTypes")
		delete(additionalProperties, "required")
		delete(additionalProperties, "maximumDays")
		delete(additionalProperties, "maximumHours")
		delete(additionalProperties, "maximumWeeks")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableValidAccessDurationSettingsDetails struct {
	value *ValidAccessDurationSettingsDetails
	isSet bool
}

func (v NullableValidAccessDurationSettingsDetails) Get() *ValidAccessDurationSettingsDetails {
	return v.value
}

func (v *NullableValidAccessDurationSettingsDetails) Set(val *ValidAccessDurationSettingsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableValidAccessDurationSettingsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableValidAccessDurationSettingsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidAccessDurationSettingsDetails(val *ValidAccessDurationSettingsDetails) *NullableValidAccessDurationSettingsDetails {
	return &NullableValidAccessDurationSettingsDetails{value: val, isSet: true}
}

func (v NullableValidAccessDurationSettingsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidAccessDurationSettingsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
