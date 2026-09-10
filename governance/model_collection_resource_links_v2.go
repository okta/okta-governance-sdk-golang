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

// checks if the CollectionResourceLinksV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceLinksV2{}

// CollectionResourceLinksV2 Links available on a resource within a resource collection
type CollectionResourceLinksV2 struct {
	Resource             Link  `json:"resource"`
	Entitlements         *Link `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceLinksV2 CollectionResourceLinksV2

// NewCollectionResourceLinksV2 instantiates a new CollectionResourceLinksV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceLinksV2(resource Link) *CollectionResourceLinksV2 {
	this := CollectionResourceLinksV2{}
	this.Resource = resource
	return &this
}

// NewCollectionResourceLinksV2WithDefaults instantiates a new CollectionResourceLinksV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceLinksV2WithDefaults() *CollectionResourceLinksV2 {
	this := CollectionResourceLinksV2{}
	return &this
}

// GetResource returns the Resource field value
func (o *CollectionResourceLinksV2) GetResource() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceLinksV2) GetResourceOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CollectionResourceLinksV2) SetResource(v Link) {
	o.Resource = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceLinksV2) GetEntitlements() Link {
	if o == nil || IsNil(o.Entitlements) {
		var ret Link
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceLinksV2) GetEntitlementsOk() (*Link, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceLinksV2) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given Link and assigns it to the Entitlements field.
func (o *CollectionResourceLinksV2) SetEntitlements(v Link) {
	o.Entitlements = &v
}

func (o CollectionResourceLinksV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceLinksV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceLinksV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource",
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

	varCollectionResourceLinksV2 := _CollectionResourceLinksV2{}

	err = json.Unmarshal(data, &varCollectionResourceLinksV2)

	if err != nil {
		return err
	}

	*o = CollectionResourceLinksV2(varCollectionResourceLinksV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resource")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceLinksV2 struct {
	value *CollectionResourceLinksV2
	isSet bool
}

func (v NullableCollectionResourceLinksV2) Get() *CollectionResourceLinksV2 {
	return v.value
}

func (v *NullableCollectionResourceLinksV2) Set(val *CollectionResourceLinksV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceLinksV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceLinksV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceLinksV2(val *CollectionResourceLinksV2) *NullableCollectionResourceLinksV2 {
	return &NullableCollectionResourceLinksV2{value: val, isSet: true}
}

func (v NullableCollectionResourceLinksV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceLinksV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
