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

// checks if the PrincipalEntitlementRowResourceAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementRowResourceAsset{}

// PrincipalEntitlementRowResourceAsset The resource asset that this row represents. Present only for resources that support assets. Omitted when the row is the resource (app) itself.
type PrincipalEntitlementRowResourceAsset struct {
	// The identifier of the resource asset
	Id string `json:"id"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn  string                  `json:"orn"`
	Type ResourceAssetTypeSparse `json:"type"`
	// The display name for a resource asset
	Name string `json:"name"`
	// The description of a resource asset
	Description *string `json:"description,omitempty"`
	// Whether this asset has any children in the resource hierarchy
	HasChildren          bool `json:"hasChildren"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementRowResourceAsset PrincipalEntitlementRowResourceAsset

// NewPrincipalEntitlementRowResourceAsset instantiates a new PrincipalEntitlementRowResourceAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementRowResourceAsset(id string, orn string, type_ ResourceAssetTypeSparse, name string, hasChildren bool) *PrincipalEntitlementRowResourceAsset {
	this := PrincipalEntitlementRowResourceAsset{}
	this.Id = id
	this.Orn = orn
	this.Type = type_
	this.Name = name
	this.HasChildren = hasChildren
	return &this
}

// NewPrincipalEntitlementRowResourceAssetWithDefaults instantiates a new PrincipalEntitlementRowResourceAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementRowResourceAssetWithDefaults() *PrincipalEntitlementRowResourceAsset {
	this := PrincipalEntitlementRowResourceAsset{}
	return &this
}

// GetId returns the Id field value
func (o *PrincipalEntitlementRowResourceAsset) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrincipalEntitlementRowResourceAsset) SetId(v string) {
	o.Id = v
}

// GetOrn returns the Orn field value
func (o *PrincipalEntitlementRowResourceAsset) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *PrincipalEntitlementRowResourceAsset) SetOrn(v string) {
	o.Orn = v
}

// GetType returns the Type field value
func (o *PrincipalEntitlementRowResourceAsset) GetType() ResourceAssetTypeSparse {
	if o == nil {
		var ret ResourceAssetTypeSparse
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetTypeOk() (*ResourceAssetTypeSparse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrincipalEntitlementRowResourceAsset) SetType(v ResourceAssetTypeSparse) {
	o.Type = v
}

// GetName returns the Name field value
func (o *PrincipalEntitlementRowResourceAsset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrincipalEntitlementRowResourceAsset) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrincipalEntitlementRowResourceAsset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrincipalEntitlementRowResourceAsset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrincipalEntitlementRowResourceAsset) SetDescription(v string) {
	o.Description = &v
}

// GetHasChildren returns the HasChildren field value
func (o *PrincipalEntitlementRowResourceAsset) GetHasChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasChildren
}

// GetHasChildrenOk returns a tuple with the HasChildren field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRowResourceAsset) GetHasChildrenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasChildren, true
}

// SetHasChildren sets field value
func (o *PrincipalEntitlementRowResourceAsset) SetHasChildren(v bool) {
	o.HasChildren = v
}

func (o PrincipalEntitlementRowResourceAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementRowResourceAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["orn"] = o.Orn
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["hasChildren"] = o.HasChildren

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementRowResourceAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"orn",
		"type",
		"name",
		"hasChildren",
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

	varPrincipalEntitlementRowResourceAsset := _PrincipalEntitlementRowResourceAsset{}

	err = json.Unmarshal(data, &varPrincipalEntitlementRowResourceAsset)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementRowResourceAsset(varPrincipalEntitlementRowResourceAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hasChildren")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementRowResourceAsset struct {
	value *PrincipalEntitlementRowResourceAsset
	isSet bool
}

func (v NullablePrincipalEntitlementRowResourceAsset) Get() *PrincipalEntitlementRowResourceAsset {
	return v.value
}

func (v *NullablePrincipalEntitlementRowResourceAsset) Set(val *PrincipalEntitlementRowResourceAsset) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementRowResourceAsset) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementRowResourceAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementRowResourceAsset(val *PrincipalEntitlementRowResourceAsset) *NullablePrincipalEntitlementRowResourceAsset {
	return &NullablePrincipalEntitlementRowResourceAsset{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementRowResourceAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementRowResourceAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
