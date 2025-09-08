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

// CollectionResourceLinks Links available on a resource within a resource collection
type CollectionResourceLinks struct {
	Entitlements         Link `json:"entitlements"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceLinks CollectionResourceLinks

// NewCollectionResourceLinks instantiates a new CollectionResourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceLinks(entitlements Link) *CollectionResourceLinks {
	this := CollectionResourceLinks{}
	this.Entitlements = entitlements
	return &this
}

// NewCollectionResourceLinksWithDefaults instantiates a new CollectionResourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceLinksWithDefaults() *CollectionResourceLinks {
	this := CollectionResourceLinks{}
	return &this
}

// GetEntitlements returns the Entitlements field value
func (o *CollectionResourceLinks) GetEntitlements() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceLinks) GetEntitlementsOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entitlements, true
}

// SetEntitlements sets field value
func (o *CollectionResourceLinks) SetEntitlements(v Link) {
	o.Entitlements = v
}

func (o CollectionResourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionResourceLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionResourceLinks := _CollectionResourceLinks{}

	err = json.Unmarshal(bytes, &varCollectionResourceLinks)
	if err == nil {
		*o = CollectionResourceLinks(varCollectionResourceLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionResourceLinks struct {
	value *CollectionResourceLinks
	isSet bool
}

func (v NullableCollectionResourceLinks) Get() *CollectionResourceLinks {
	return v.value
}

func (v *NullableCollectionResourceLinks) Set(val *CollectionResourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceLinks(val *CollectionResourceLinks) *NullableCollectionResourceLinks {
	return &NullableCollectionResourceLinks{value: val, isSet: true}
}

func (v NullableCollectionResourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
