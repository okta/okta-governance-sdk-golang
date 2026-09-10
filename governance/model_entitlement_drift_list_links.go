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

// checks if the EntitlementDriftListLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDriftListLinks{}

// EntitlementDriftListLinks Links for navigating the list of drifts
type EntitlementDriftListLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDriftListLinks EntitlementDriftListLinks

// NewEntitlementDriftListLinks instantiates a new EntitlementDriftListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDriftListLinks(self Link) *EntitlementDriftListLinks {
	this := EntitlementDriftListLinks{}
	this.Self = self
	return &this
}

// NewEntitlementDriftListLinksWithDefaults instantiates a new EntitlementDriftListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDriftListLinksWithDefaults() *EntitlementDriftListLinks {
	this := EntitlementDriftListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *EntitlementDriftListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EntitlementDriftListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *EntitlementDriftListLinks) GetNext() Link {
	if o == nil || IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftListLinks) GetNextOk() (*Link, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *EntitlementDriftListLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *EntitlementDriftListLinks) SetNext(v Link) {
	o.Next = &v
}

func (o EntitlementDriftListLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDriftListLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDriftListLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varEntitlementDriftListLinks := _EntitlementDriftListLinks{}

	err = json.Unmarshal(data, &varEntitlementDriftListLinks)

	if err != nil {
		return err
	}

	*o = EntitlementDriftListLinks(varEntitlementDriftListLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDriftListLinks struct {
	value *EntitlementDriftListLinks
	isSet bool
}

func (v NullableEntitlementDriftListLinks) Get() *EntitlementDriftListLinks {
	return v.value
}

func (v *NullableEntitlementDriftListLinks) Set(val *EntitlementDriftListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftListLinks(val *EntitlementDriftListLinks) *NullableEntitlementDriftListLinks {
	return &NullableEntitlementDriftListLinks{value: val, isSet: true}
}

func (v NullableEntitlementDriftListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
