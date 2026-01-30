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

// checks if the RcarEntryCountsResourceCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcarEntryCountsResourceCounts{}

// RcarEntryCountsResourceCounts Collection resource counts
type RcarEntryCountsResourceCounts struct {
	// Number of app resources in a collection
	Applications         *int32 `json:"applications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntryCountsResourceCounts RcarEntryCountsResourceCounts

// NewRcarEntryCountsResourceCounts instantiates a new RcarEntryCountsResourceCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntryCountsResourceCounts() *RcarEntryCountsResourceCounts {
	this := RcarEntryCountsResourceCounts{}
	return &this
}

// NewRcarEntryCountsResourceCountsWithDefaults instantiates a new RcarEntryCountsResourceCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntryCountsResourceCountsWithDefaults() *RcarEntryCountsResourceCounts {
	this := RcarEntryCountsResourceCounts{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *RcarEntryCountsResourceCounts) GetApplications() int32 {
	if o == nil || IsNil(o.Applications) {
		var ret int32
		return ret
	}
	return *o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryCountsResourceCounts) GetApplicationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *RcarEntryCountsResourceCounts) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given int32 and assigns it to the Applications field.
func (o *RcarEntryCountsResourceCounts) SetApplications(v int32) {
	o.Applications = &v
}

func (o RcarEntryCountsResourceCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcarEntryCountsResourceCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RcarEntryCountsResourceCounts) UnmarshalJSON(data []byte) (err error) {
	varRcarEntryCountsResourceCounts := _RcarEntryCountsResourceCounts{}

	err = json.Unmarshal(data, &varRcarEntryCountsResourceCounts)

	if err != nil {
		return err
	}

	*o = RcarEntryCountsResourceCounts(varRcarEntryCountsResourceCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRcarEntryCountsResourceCounts struct {
	value *RcarEntryCountsResourceCounts
	isSet bool
}

func (v NullableRcarEntryCountsResourceCounts) Get() *RcarEntryCountsResourceCounts {
	return v.value
}

func (v *NullableRcarEntryCountsResourceCounts) Set(val *RcarEntryCountsResourceCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntryCountsResourceCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntryCountsResourceCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntryCountsResourceCounts(val *RcarEntryCountsResourceCounts) *NullableRcarEntryCountsResourceCounts {
	return &NullableRcarEntryCountsResourceCounts{value: val, isSet: true}
}

func (v NullableRcarEntryCountsResourceCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntryCountsResourceCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
