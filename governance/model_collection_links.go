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

// checks if the CollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionLinks{}

// CollectionLinks Links available on a resource collection
type CollectionLinks struct {
	Resources            Link `json:"resources"`
	AdditionalProperties map[string]interface{}
}

type _CollectionLinks CollectionLinks

// NewCollectionLinks instantiates a new CollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionLinks(resources Link) *CollectionLinks {
	this := CollectionLinks{}
	this.Resources = resources
	return &this
}

// NewCollectionLinksWithDefaults instantiates a new CollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionLinksWithDefaults() *CollectionLinks {
	this := CollectionLinks{}
	return &this
}

// GetResources returns the Resources field value
func (o *CollectionLinks) GetResources() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetResourcesOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *CollectionLinks) SetResources(v Link) {
	o.Resources = v
}

func (o CollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resources",
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

	varCollectionLinks := _CollectionLinks{}

	err = json.Unmarshal(data, &varCollectionLinks)

	if err != nil {
		return err
	}

	*o = CollectionLinks(varCollectionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionLinks struct {
	value *CollectionLinks
	isSet bool
}

func (v NullableCollectionLinks) Get() *CollectionLinks {
	return v.value
}

func (v *NullableCollectionLinks) Set(val *CollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionLinks(val *CollectionLinks) *NullableCollectionLinks {
	return &NullableCollectionLinks{value: val, isSet: true}
}

func (v NullableCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
