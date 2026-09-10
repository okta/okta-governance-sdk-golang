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

// checks if the ResourceAssetUpdatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetUpdatable{}

// ResourceAssetUpdatable The updatable properties of a resource asset.
type ResourceAssetUpdatable struct {
	// The description of a resource asset
	Description          NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetUpdatable ResourceAssetUpdatable

// NewResourceAssetUpdatable instantiates a new ResourceAssetUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetUpdatable() *ResourceAssetUpdatable {
	this := ResourceAssetUpdatable{}
	return &this
}

// NewResourceAssetUpdatableWithDefaults instantiates a new ResourceAssetUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetUpdatableWithDefaults() *ResourceAssetUpdatable {
	this := ResourceAssetUpdatable{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceAssetUpdatable) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceAssetUpdatable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceAssetUpdatable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResourceAssetUpdatable) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResourceAssetUpdatable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResourceAssetUpdatable) UnsetDescription() {
	o.Description.Unset()
}

func (o ResourceAssetUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetUpdatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetUpdatable) UnmarshalJSON(data []byte) (err error) {
	varResourceAssetUpdatable := _ResourceAssetUpdatable{}

	err = json.Unmarshal(data, &varResourceAssetUpdatable)

	if err != nil {
		return err
	}

	*o = ResourceAssetUpdatable(varResourceAssetUpdatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetUpdatable struct {
	value *ResourceAssetUpdatable
	isSet bool
}

func (v NullableResourceAssetUpdatable) Get() *ResourceAssetUpdatable {
	return v.value
}

func (v *NullableResourceAssetUpdatable) Set(val *ResourceAssetUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetUpdatable(val *ResourceAssetUpdatable) *NullableResourceAssetUpdatable {
	return &NullableResourceAssetUpdatable{value: val, isSet: true}
}

func (v NullableResourceAssetUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
