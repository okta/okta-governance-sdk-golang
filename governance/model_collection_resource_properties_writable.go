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

// CollectionResourcePropertiesWritable struct for CollectionResourcePropertiesWritable
type CollectionResourcePropertiesWritable struct {
	// The ORN identifier for a specific app. Other resource types aren't supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.
	ResourceOrn          *string `json:"resourceOrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourcePropertiesWritable CollectionResourcePropertiesWritable

// NewCollectionResourcePropertiesWritable instantiates a new CollectionResourcePropertiesWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourcePropertiesWritable() *CollectionResourcePropertiesWritable {
	this := CollectionResourcePropertiesWritable{}
	return &this
}

// NewCollectionResourcePropertiesWritableWithDefaults instantiates a new CollectionResourcePropertiesWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourcePropertiesWritableWithDefaults() *CollectionResourcePropertiesWritable {
	this := CollectionResourcePropertiesWritable{}
	return &this
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *CollectionResourcePropertiesWritable) GetResourceOrn() string {
	if o == nil || o.ResourceOrn == nil {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourcePropertiesWritable) GetResourceOrnOk() (*string, bool) {
	if o == nil || o.ResourceOrn == nil {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *CollectionResourcePropertiesWritable) HasResourceOrn() bool {
	if o != nil && o.ResourceOrn != nil {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *CollectionResourcePropertiesWritable) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

func (o CollectionResourcePropertiesWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceOrn != nil {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionResourcePropertiesWritable) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionResourcePropertiesWritable := _CollectionResourcePropertiesWritable{}

	err = json.Unmarshal(bytes, &varCollectionResourcePropertiesWritable)
	if err == nil {
		*o = CollectionResourcePropertiesWritable(varCollectionResourcePropertiesWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionResourcePropertiesWritable struct {
	value *CollectionResourcePropertiesWritable
	isSet bool
}

func (v NullableCollectionResourcePropertiesWritable) Get() *CollectionResourcePropertiesWritable {
	return v.value
}

func (v *NullableCollectionResourcePropertiesWritable) Set(val *CollectionResourcePropertiesWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourcePropertiesWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourcePropertiesWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourcePropertiesWritable(val *CollectionResourcePropertiesWritable) *NullableCollectionResourcePropertiesWritable {
	return &NullableCollectionResourcePropertiesWritable{value: val, isSet: true}
}

func (v NullableCollectionResourcePropertiesWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourcePropertiesWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
