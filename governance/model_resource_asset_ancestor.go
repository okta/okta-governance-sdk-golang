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

// checks if the ResourceAssetAncestor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetAncestor{}

// ResourceAssetAncestor An ancestor asset in the hierarchy, including its own parentId for chain traversal.
type ResourceAssetAncestor struct {
	// Unique identifier for the asset
	Id string `json:"id"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn string `json:"orn"`
	// The display name for a resource asset
	Name string                  `json:"name"`
	Type ResourceAssetTypeSparse `json:"type"`
	// The ID of this ancestor's direct parent, or null if it is a root asset
	ParentId             NullableString      `json:"parentId,omitempty"`
	Links                *ResourceAssetLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetAncestor ResourceAssetAncestor

// NewResourceAssetAncestor instantiates a new ResourceAssetAncestor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetAncestor(id string, orn string, name string, type_ ResourceAssetTypeSparse) *ResourceAssetAncestor {
	this := ResourceAssetAncestor{}
	this.Id = id
	this.Orn = orn
	this.Name = name
	this.Type = type_
	return &this
}

// NewResourceAssetAncestorWithDefaults instantiates a new ResourceAssetAncestor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetAncestorWithDefaults() *ResourceAssetAncestor {
	this := ResourceAssetAncestor{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceAssetAncestor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetAncestor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceAssetAncestor) SetId(v string) {
	o.Id = v
}

// GetOrn returns the Orn field value
func (o *ResourceAssetAncestor) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetAncestor) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ResourceAssetAncestor) SetOrn(v string) {
	o.Orn = v
}

// GetName returns the Name field value
func (o *ResourceAssetAncestor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetAncestor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceAssetAncestor) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ResourceAssetAncestor) GetType() ResourceAssetTypeSparse {
	if o == nil {
		var ret ResourceAssetTypeSparse
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetAncestor) GetTypeOk() (*ResourceAssetTypeSparse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceAssetAncestor) SetType(v ResourceAssetTypeSparse) {
	o.Type = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceAssetAncestor) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceAssetAncestor) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ResourceAssetAncestor) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ResourceAssetAncestor) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ResourceAssetAncestor) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ResourceAssetAncestor) UnsetParentId() {
	o.ParentId.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceAssetAncestor) GetLinks() ResourceAssetLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceAssetLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetAncestor) GetLinksOk() (*ResourceAssetLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceAssetAncestor) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceAssetLinks and assigns it to the Links field.
func (o *ResourceAssetAncestor) SetLinks(v ResourceAssetLinks) {
	o.Links = &v
}

func (o ResourceAssetAncestor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetAncestor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["orn"] = o.Orn
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetAncestor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"orn",
		"name",
		"type",
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

	varResourceAssetAncestor := _ResourceAssetAncestor{}

	err = json.Unmarshal(data, &varResourceAssetAncestor)

	if err != nil {
		return err
	}

	*o = ResourceAssetAncestor(varResourceAssetAncestor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetAncestor struct {
	value *ResourceAssetAncestor
	isSet bool
}

func (v NullableResourceAssetAncestor) Get() *ResourceAssetAncestor {
	return v.value
}

func (v *NullableResourceAssetAncestor) Set(val *ResourceAssetAncestor) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetAncestor) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetAncestor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetAncestor(val *ResourceAssetAncestor) *NullableResourceAssetAncestor {
	return &NullableResourceAssetAncestor{value: val, isSet: true}
}

func (v NullableResourceAssetAncestor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetAncestor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
