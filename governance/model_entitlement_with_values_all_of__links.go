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

// checks if the EntitlementWithValuesAllOfLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementWithValuesAllOfLinks{}

// EntitlementWithValuesAllOfLinks Link relations for this entitlement's values. `next` is present only when more effective values remain beyond the current page.
type EntitlementWithValuesAllOfLinks struct {
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementWithValuesAllOfLinks EntitlementWithValuesAllOfLinks

// NewEntitlementWithValuesAllOfLinks instantiates a new EntitlementWithValuesAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementWithValuesAllOfLinks() *EntitlementWithValuesAllOfLinks {
	this := EntitlementWithValuesAllOfLinks{}
	return &this
}

// NewEntitlementWithValuesAllOfLinksWithDefaults instantiates a new EntitlementWithValuesAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithValuesAllOfLinksWithDefaults() *EntitlementWithValuesAllOfLinks {
	this := EntitlementWithValuesAllOfLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *EntitlementWithValuesAllOfLinks) GetNext() Link {
	if o == nil || IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValuesAllOfLinks) GetNextOk() (*Link, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *EntitlementWithValuesAllOfLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *EntitlementWithValuesAllOfLinks) SetNext(v Link) {
	o.Next = &v
}

func (o EntitlementWithValuesAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementWithValuesAllOfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementWithValuesAllOfLinks) UnmarshalJSON(data []byte) (err error) {
	varEntitlementWithValuesAllOfLinks := _EntitlementWithValuesAllOfLinks{}

	err = json.Unmarshal(data, &varEntitlementWithValuesAllOfLinks)

	if err != nil {
		return err
	}

	*o = EntitlementWithValuesAllOfLinks(varEntitlementWithValuesAllOfLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementWithValuesAllOfLinks struct {
	value *EntitlementWithValuesAllOfLinks
	isSet bool
}

func (v NullableEntitlementWithValuesAllOfLinks) Get() *EntitlementWithValuesAllOfLinks {
	return v.value
}

func (v *NullableEntitlementWithValuesAllOfLinks) Set(val *EntitlementWithValuesAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementWithValuesAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementWithValuesAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementWithValuesAllOfLinks(val *EntitlementWithValuesAllOfLinks) *NullableEntitlementWithValuesAllOfLinks {
	return &NullableEntitlementWithValuesAllOfLinks{value: val, isSet: true}
}

func (v NullableEntitlementWithValuesAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementWithValuesAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
