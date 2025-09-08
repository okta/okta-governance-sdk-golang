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

// RcarEntryCounts Entry count metadata
type RcarEntryCounts struct {
	ResourceCounts       *RcarEntryCountsResourceCounts `json:"resourceCounts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntryCounts RcarEntryCounts

// NewRcarEntryCounts instantiates a new RcarEntryCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntryCounts() *RcarEntryCounts {
	this := RcarEntryCounts{}
	return &this
}

// NewRcarEntryCountsWithDefaults instantiates a new RcarEntryCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntryCountsWithDefaults() *RcarEntryCounts {
	this := RcarEntryCounts{}
	return &this
}

// GetResourceCounts returns the ResourceCounts field value if set, zero value otherwise.
func (o *RcarEntryCounts) GetResourceCounts() RcarEntryCountsResourceCounts {
	if o == nil || o.ResourceCounts == nil {
		var ret RcarEntryCountsResourceCounts
		return ret
	}
	return *o.ResourceCounts
}

// GetResourceCountsOk returns a tuple with the ResourceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryCounts) GetResourceCountsOk() (*RcarEntryCountsResourceCounts, bool) {
	if o == nil || o.ResourceCounts == nil {
		return nil, false
	}
	return o.ResourceCounts, true
}

// HasResourceCounts returns a boolean if a field has been set.
func (o *RcarEntryCounts) HasResourceCounts() bool {
	if o != nil && o.ResourceCounts != nil {
		return true
	}

	return false
}

// SetResourceCounts gets a reference to the given RcarEntryCountsResourceCounts and assigns it to the ResourceCounts field.
func (o *RcarEntryCounts) SetResourceCounts(v RcarEntryCountsResourceCounts) {
	o.ResourceCounts = &v
}

func (o RcarEntryCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceCounts != nil {
		toSerialize["resourceCounts"] = o.ResourceCounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RcarEntryCounts) UnmarshalJSON(bytes []byte) (err error) {
	varRcarEntryCounts := _RcarEntryCounts{}

	err = json.Unmarshal(bytes, &varRcarEntryCounts)
	if err == nil {
		*o = RcarEntryCounts(varRcarEntryCounts)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceCounts")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRcarEntryCounts struct {
	value *RcarEntryCounts
	isSet bool
}

func (v NullableRcarEntryCounts) Get() *RcarEntryCounts {
	return v.value
}

func (v *NullableRcarEntryCounts) Set(val *RcarEntryCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntryCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntryCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntryCounts(val *RcarEntryCounts) *NullableRcarEntryCounts {
	return &NullableRcarEntryCounts{value: val, isSet: true}
}

func (v NullableRcarEntryCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntryCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
