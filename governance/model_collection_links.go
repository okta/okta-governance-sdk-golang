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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionLinks := _CollectionLinks{}

	err = json.Unmarshal(bytes, &varCollectionLinks)
	if err == nil {
		*o = CollectionLinks(varCollectionLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
