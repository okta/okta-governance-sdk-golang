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

// checks if the ResourceAssetLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetLinks{}

// ResourceAssetLinks Links available in resource asset response
type ResourceAssetLinks struct {
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetLinks ResourceAssetLinks

// NewResourceAssetLinks instantiates a new ResourceAssetLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetLinks(self Link) *ResourceAssetLinks {
	this := ResourceAssetLinks{}
	this.Self = self
	return &this
}

// NewResourceAssetLinksWithDefaults instantiates a new ResourceAssetLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetLinksWithDefaults() *ResourceAssetLinks {
	this := ResourceAssetLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ResourceAssetLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ResourceAssetLinks) SetSelf(v Link) {
	o.Self = v
}

func (o ResourceAssetLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetLinks) UnmarshalJSON(data []byte) (err error) {
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

	varResourceAssetLinks := _ResourceAssetLinks{}

	err = json.Unmarshal(data, &varResourceAssetLinks)

	if err != nil {
		return err
	}

	*o = ResourceAssetLinks(varResourceAssetLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetLinks struct {
	value *ResourceAssetLinks
	isSet bool
}

func (v NullableResourceAssetLinks) Get() *ResourceAssetLinks {
	return v.value
}

func (v *NullableResourceAssetLinks) Set(val *ResourceAssetLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetLinks(val *ResourceAssetLinks) *NullableResourceAssetLinks {
	return &NullableResourceAssetLinks{value: val, isSet: true}
}

func (v NullableResourceAssetLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
