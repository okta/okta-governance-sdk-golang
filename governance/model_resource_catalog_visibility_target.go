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

// checks if the ResourceCatalogVisibilityTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceCatalogVisibilityTarget{}

// ResourceCatalogVisibilityTarget A typed visibility target used to scope resource catalog access
type ResourceCatalogVisibilityTarget struct {
	Type ResourceCatalogVisibilityType `json:"type"`
	// The ID for the given target type (for example, a group ID for `GROUP`)
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ResourceCatalogVisibilityTarget ResourceCatalogVisibilityTarget

// NewResourceCatalogVisibilityTarget instantiates a new ResourceCatalogVisibilityTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCatalogVisibilityTarget(type_ ResourceCatalogVisibilityType, id string) *ResourceCatalogVisibilityTarget {
	this := ResourceCatalogVisibilityTarget{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewResourceCatalogVisibilityTargetWithDefaults instantiates a new ResourceCatalogVisibilityTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCatalogVisibilityTargetWithDefaults() *ResourceCatalogVisibilityTarget {
	this := ResourceCatalogVisibilityTarget{}
	return &this
}

// GetType returns the Type field value
func (o *ResourceCatalogVisibilityTarget) GetType() ResourceCatalogVisibilityType {
	if o == nil {
		var ret ResourceCatalogVisibilityType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceCatalogVisibilityTarget) GetTypeOk() (*ResourceCatalogVisibilityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceCatalogVisibilityTarget) SetType(v ResourceCatalogVisibilityType) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ResourceCatalogVisibilityTarget) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceCatalogVisibilityTarget) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceCatalogVisibilityTarget) SetId(v string) {
	o.Id = v
}

func (o ResourceCatalogVisibilityTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceCatalogVisibilityTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceCatalogVisibilityTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varResourceCatalogVisibilityTarget := _ResourceCatalogVisibilityTarget{}

	err = json.Unmarshal(data, &varResourceCatalogVisibilityTarget)

	if err != nil {
		return err
	}

	*o = ResourceCatalogVisibilityTarget(varResourceCatalogVisibilityTarget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceCatalogVisibilityTarget struct {
	value *ResourceCatalogVisibilityTarget
	isSet bool
}

func (v NullableResourceCatalogVisibilityTarget) Get() *ResourceCatalogVisibilityTarget {
	return v.value
}

func (v *NullableResourceCatalogVisibilityTarget) Set(val *ResourceCatalogVisibilityTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCatalogVisibilityTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCatalogVisibilityTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCatalogVisibilityTarget(val *ResourceCatalogVisibilityTarget) *NullableResourceCatalogVisibilityTarget {
	return &NullableResourceCatalogVisibilityTarget{value: val, isSet: true}
}

func (v NullableResourceCatalogVisibilityTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCatalogVisibilityTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
