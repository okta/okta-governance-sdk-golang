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

// checks if the ScopedEntitlementResourceAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopedEntitlementResourceAsset{}

// ScopedEntitlementResourceAsset A resource asset associated with a scoped entitlement
type ScopedEntitlementResourceAsset struct {
	// The identifier of the resource asset
	Id string `json:"id"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn  string                  `json:"orn"`
	Type ResourceAssetTypeSparse `json:"type"`
	// The display name for a resource asset
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ScopedEntitlementResourceAsset ScopedEntitlementResourceAsset

// NewScopedEntitlementResourceAsset instantiates a new ScopedEntitlementResourceAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopedEntitlementResourceAsset(id string, orn string, type_ ResourceAssetTypeSparse, name string) *ScopedEntitlementResourceAsset {
	this := ScopedEntitlementResourceAsset{}
	this.Id = id
	this.Orn = orn
	this.Type = type_
	this.Name = name
	return &this
}

// NewScopedEntitlementResourceAssetWithDefaults instantiates a new ScopedEntitlementResourceAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopedEntitlementResourceAssetWithDefaults() *ScopedEntitlementResourceAsset {
	this := ScopedEntitlementResourceAsset{}
	return &this
}

// GetId returns the Id field value
func (o *ScopedEntitlementResourceAsset) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementResourceAsset) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScopedEntitlementResourceAsset) SetId(v string) {
	o.Id = v
}

// GetOrn returns the Orn field value
func (o *ScopedEntitlementResourceAsset) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementResourceAsset) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ScopedEntitlementResourceAsset) SetOrn(v string) {
	o.Orn = v
}

// GetType returns the Type field value
func (o *ScopedEntitlementResourceAsset) GetType() ResourceAssetTypeSparse {
	if o == nil {
		var ret ResourceAssetTypeSparse
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementResourceAsset) GetTypeOk() (*ResourceAssetTypeSparse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScopedEntitlementResourceAsset) SetType(v ResourceAssetTypeSparse) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ScopedEntitlementResourceAsset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementResourceAsset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopedEntitlementResourceAsset) SetName(v string) {
	o.Name = v
}

func (o ScopedEntitlementResourceAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopedEntitlementResourceAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["orn"] = o.Orn
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScopedEntitlementResourceAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"orn",
		"type",
		"name",
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

	varScopedEntitlementResourceAsset := _ScopedEntitlementResourceAsset{}

	err = json.Unmarshal(data, &varScopedEntitlementResourceAsset)

	if err != nil {
		return err
	}

	*o = ScopedEntitlementResourceAsset(varScopedEntitlementResourceAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScopedEntitlementResourceAsset struct {
	value *ScopedEntitlementResourceAsset
	isSet bool
}

func (v NullableScopedEntitlementResourceAsset) Get() *ScopedEntitlementResourceAsset {
	return v.value
}

func (v *NullableScopedEntitlementResourceAsset) Set(val *ScopedEntitlementResourceAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedEntitlementResourceAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedEntitlementResourceAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedEntitlementResourceAsset(val *ScopedEntitlementResourceAsset) *NullableScopedEntitlementResourceAsset {
	return &NullableScopedEntitlementResourceAsset{value: val, isSet: true}
}

func (v NullableScopedEntitlementResourceAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedEntitlementResourceAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
