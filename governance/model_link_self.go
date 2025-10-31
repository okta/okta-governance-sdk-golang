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

// checks if the LinkSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkSelf{}

// LinkSelf Links available for an object
type LinkSelf struct {
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _LinkSelf LinkSelf

// NewLinkSelf instantiates a new LinkSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSelf(self Link) *LinkSelf {
	this := LinkSelf{}
	this.Self = self
	return &this
}

// NewLinkSelfWithDefaults instantiates a new LinkSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSelfWithDefaults() *LinkSelf {
	this := LinkSelf{}
	return &this
}

// GetSelf returns the Self field value
func (o *LinkSelf) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *LinkSelf) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *LinkSelf) SetSelf(v Link) {
	o.Self = v
}

func (o LinkSelf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkSelf) UnmarshalJSON(data []byte) (err error) {
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

	varLinkSelf := _LinkSelf{}

	err = json.Unmarshal(data, &varLinkSelf)

	if err != nil {
		return err
	}

	*o = LinkSelf(varLinkSelf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkSelf struct {
	value *LinkSelf
	isSet bool
}

func (v NullableLinkSelf) Get() *LinkSelf {
	return v.value
}

func (v *NullableLinkSelf) Set(val *LinkSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSelf(val *LinkSelf) *NullableLinkSelf {
	return &NullableLinkSelf{value: val, isSet: true}
}

func (v NullableLinkSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
