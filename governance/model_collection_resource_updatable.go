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

// checks if the CollectionResourceUpdatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceUpdatable{}

// CollectionResourceUpdatable the updatable properties of a collection resource
type CollectionResourceUpdatable struct {
	// Collection of entitlements and associated value identifiers
	Entitlements         []EntitlementCreatable `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceUpdatable CollectionResourceUpdatable

// NewCollectionResourceUpdatable instantiates a new CollectionResourceUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceUpdatable() *CollectionResourceUpdatable {
	this := CollectionResourceUpdatable{}
	return &this
}

// NewCollectionResourceUpdatableWithDefaults instantiates a new CollectionResourceUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceUpdatableWithDefaults() *CollectionResourceUpdatable {
	this := CollectionResourceUpdatable{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceUpdatable) GetEntitlements() []EntitlementCreatable {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceUpdatable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceUpdatable) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *CollectionResourceUpdatable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

func (o CollectionResourceUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceUpdatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceUpdatable) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourceUpdatable := _CollectionResourceUpdatable{}

	err = json.Unmarshal(data, &varCollectionResourceUpdatable)

	if err != nil {
		return err
	}

	*o = CollectionResourceUpdatable(varCollectionResourceUpdatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceUpdatable struct {
	value *CollectionResourceUpdatable
	isSet bool
}

func (v NullableCollectionResourceUpdatable) Get() *CollectionResourceUpdatable {
	return v.value
}

func (v *NullableCollectionResourceUpdatable) Set(val *CollectionResourceUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceUpdatable(val *CollectionResourceUpdatable) *NullableCollectionResourceUpdatable {
	return &NullableCollectionResourceUpdatable{value: val, isSet: true}
}

func (v NullableCollectionResourceUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
