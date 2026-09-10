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

// checks if the CollectionResourceCreatableV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceCreatableV2{}

// CollectionResourceCreatableV2 The properties expected when adding a new resource to a collection
type CollectionResourceCreatableV2 struct {
	// Collection of entitlements and associated value identifiers
	Entitlements []EntitlementCreatable `json:"entitlements,omitempty"`
	// The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.
	ResourceOrn          string `json:"resourceOrn"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceCreatableV2 CollectionResourceCreatableV2

// NewCollectionResourceCreatableV2 instantiates a new CollectionResourceCreatableV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceCreatableV2(resourceOrn string) *CollectionResourceCreatableV2 {
	this := CollectionResourceCreatableV2{}
	this.ResourceOrn = resourceOrn
	return &this
}

// NewCollectionResourceCreatableV2WithDefaults instantiates a new CollectionResourceCreatableV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceCreatableV2WithDefaults() *CollectionResourceCreatableV2 {
	this := CollectionResourceCreatableV2{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceCreatableV2) GetEntitlements() []EntitlementCreatable {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatableV2) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceCreatableV2) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *CollectionResourceCreatableV2) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceCreatableV2) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatableV2) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceCreatableV2) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

func (o CollectionResourceCreatableV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceCreatableV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	toSerialize["resourceOrn"] = o.ResourceOrn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceCreatableV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceOrn",
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

	varCollectionResourceCreatableV2 := _CollectionResourceCreatableV2{}

	err = json.Unmarshal(data, &varCollectionResourceCreatableV2)

	if err != nil {
		return err
	}

	*o = CollectionResourceCreatableV2(varCollectionResourceCreatableV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceCreatableV2 struct {
	value *CollectionResourceCreatableV2
	isSet bool
}

func (v NullableCollectionResourceCreatableV2) Get() *CollectionResourceCreatableV2 {
	return v.value
}

func (v *NullableCollectionResourceCreatableV2) Set(val *CollectionResourceCreatableV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceCreatableV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceCreatableV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceCreatableV2(val *CollectionResourceCreatableV2) *NullableCollectionResourceCreatableV2 {
	return &NullableCollectionResourceCreatableV2{value: val, isSet: true}
}

func (v NullableCollectionResourceCreatableV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceCreatableV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
