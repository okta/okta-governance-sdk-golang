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

// checks if the CollectionResourceCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceCreatable{}

// CollectionResourceCreatable The properties expected when adding a new resource to a collection
type CollectionResourceCreatable struct {
	// Collection of entitlements and associated value identifiers
	Entitlements []EntitlementCreatable `json:"entitlements,omitempty"`
	// The ORN identifier for a specific app. Other resource types aren't supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.
	ResourceOrn          string `json:"resourceOrn"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceCreatable CollectionResourceCreatable

// NewCollectionResourceCreatable instantiates a new CollectionResourceCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceCreatable(resourceOrn string) *CollectionResourceCreatable {
	this := CollectionResourceCreatable{}
	this.ResourceOrn = resourceOrn
	return &this
}

// NewCollectionResourceCreatableWithDefaults instantiates a new CollectionResourceCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceCreatableWithDefaults() *CollectionResourceCreatable {
	this := CollectionResourceCreatable{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceCreatable) GetEntitlements() []EntitlementCreatable {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceCreatable) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *CollectionResourceCreatable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceCreatable) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatable) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceCreatable) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

func (o CollectionResourceCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceCreatable) ToMap() (map[string]interface{}, error) {
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

func (o *CollectionResourceCreatable) UnmarshalJSON(data []byte) (err error) {
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

	varCollectionResourceCreatable := _CollectionResourceCreatable{}

	err = json.Unmarshal(data, &varCollectionResourceCreatable)

	if err != nil {
		return err
	}

	*o = CollectionResourceCreatable(varCollectionResourceCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceCreatable struct {
	value *CollectionResourceCreatable
	isSet bool
}

func (v NullableCollectionResourceCreatable) Get() *CollectionResourceCreatable {
	return v.value
}

func (v *NullableCollectionResourceCreatable) Set(val *CollectionResourceCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceCreatable(val *CollectionResourceCreatable) *NullableCollectionResourceCreatable {
	return &NullableCollectionResourceCreatable{value: val, isSet: true}
}

func (v NullableCollectionResourceCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
