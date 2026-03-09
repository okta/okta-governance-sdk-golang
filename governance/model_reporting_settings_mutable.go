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

// checks if the ReportingSettingsMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingSettingsMutable{}

// ReportingSettingsMutable Reporting properties for processing post-completed campaigns
type ReportingSettingsMutable struct {
	// If `true`, a post-campaign reporting package is created
	CreateReportingPackageEnabled *bool `json:"createReportingPackageEnabled,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _ReportingSettingsMutable ReportingSettingsMutable

// NewReportingSettingsMutable instantiates a new ReportingSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingSettingsMutable() *ReportingSettingsMutable {
	this := ReportingSettingsMutable{}
	var createReportingPackageEnabled bool = false
	this.CreateReportingPackageEnabled = &createReportingPackageEnabled
	return &this
}

// NewReportingSettingsMutableWithDefaults instantiates a new ReportingSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingSettingsMutableWithDefaults() *ReportingSettingsMutable {
	this := ReportingSettingsMutable{}
	var createReportingPackageEnabled bool = false
	this.CreateReportingPackageEnabled = &createReportingPackageEnabled
	return &this
}

// GetCreateReportingPackageEnabled returns the CreateReportingPackageEnabled field value if set, zero value otherwise.
func (o *ReportingSettingsMutable) GetCreateReportingPackageEnabled() bool {
	if o == nil || IsNil(o.CreateReportingPackageEnabled) {
		var ret bool
		return ret
	}
	return *o.CreateReportingPackageEnabled
}

// GetCreateReportingPackageEnabledOk returns a tuple with the CreateReportingPackageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSettingsMutable) GetCreateReportingPackageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateReportingPackageEnabled) {
		return nil, false
	}
	return o.CreateReportingPackageEnabled, true
}

// HasCreateReportingPackageEnabled returns a boolean if a field has been set.
func (o *ReportingSettingsMutable) HasCreateReportingPackageEnabled() bool {
	if o != nil && !IsNil(o.CreateReportingPackageEnabled) {
		return true
	}

	return false
}

// SetCreateReportingPackageEnabled gets a reference to the given bool and assigns it to the CreateReportingPackageEnabled field.
func (o *ReportingSettingsMutable) SetCreateReportingPackageEnabled(v bool) {
	o.CreateReportingPackageEnabled = &v
}

func (o ReportingSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingSettingsMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateReportingPackageEnabled) {
		toSerialize["createReportingPackageEnabled"] = o.CreateReportingPackageEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportingSettingsMutable) UnmarshalJSON(data []byte) (err error) {
	varReportingSettingsMutable := _ReportingSettingsMutable{}

	err = json.Unmarshal(data, &varReportingSettingsMutable)

	if err != nil {
		return err
	}

	*o = ReportingSettingsMutable(varReportingSettingsMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createReportingPackageEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportingSettingsMutable struct {
	value *ReportingSettingsMutable
	isSet bool
}

func (v NullableReportingSettingsMutable) Get() *ReportingSettingsMutable {
	return v.value
}

func (v *NullableReportingSettingsMutable) Set(val *ReportingSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingSettingsMutable(val *ReportingSettingsMutable) *NullableReportingSettingsMutable {
	return &NullableReportingSettingsMutable{value: val, isSet: true}
}

func (v NullableReportingSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
