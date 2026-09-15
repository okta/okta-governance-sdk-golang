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

// checks if the ScopedEntitlementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopedEntitlementItem{}

// ScopedEntitlementItem A scoped entitlement grouping entitlements under a specific resource asset
type ScopedEntitlementItem struct {
	// The identifier of the resource asset associated with this scoped entitlement
	ResourceAssetId string                         `json:"resourceAssetId"`
	ResourceAsset   ScopedEntitlementResourceAsset `json:"resourceAsset"`
	// Entitlements associated with this resource asset
	Entitlements         []GrantedEntitlements `json:"entitlements"`
	AdditionalProperties map[string]interface{}
}

type _ScopedEntitlementItem ScopedEntitlementItem

// NewScopedEntitlementItem instantiates a new ScopedEntitlementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopedEntitlementItem(resourceAssetId string, resourceAsset ScopedEntitlementResourceAsset, entitlements []GrantedEntitlements) *ScopedEntitlementItem {
	this := ScopedEntitlementItem{}
	this.ResourceAssetId = resourceAssetId
	this.ResourceAsset = resourceAsset
	this.Entitlements = entitlements
	return &this
}

// NewScopedEntitlementItemWithDefaults instantiates a new ScopedEntitlementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopedEntitlementItemWithDefaults() *ScopedEntitlementItem {
	this := ScopedEntitlementItem{}
	return &this
}

// GetResourceAssetId returns the ResourceAssetId field value
func (o *ScopedEntitlementItem) GetResourceAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceAssetId
}

// GetResourceAssetIdOk returns a tuple with the ResourceAssetId field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementItem) GetResourceAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceAssetId, true
}

// SetResourceAssetId sets field value
func (o *ScopedEntitlementItem) SetResourceAssetId(v string) {
	o.ResourceAssetId = v
}

// GetResourceAsset returns the ResourceAsset field value
func (o *ScopedEntitlementItem) GetResourceAsset() ScopedEntitlementResourceAsset {
	if o == nil {
		var ret ScopedEntitlementResourceAsset
		return ret
	}

	return o.ResourceAsset
}

// GetResourceAssetOk returns a tuple with the ResourceAsset field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementItem) GetResourceAssetOk() (*ScopedEntitlementResourceAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceAsset, true
}

// SetResourceAsset sets field value
func (o *ScopedEntitlementItem) SetResourceAsset(v ScopedEntitlementResourceAsset) {
	o.ResourceAsset = v
}

// GetEntitlements returns the Entitlements field value
func (o *ScopedEntitlementItem) GetEntitlements() []GrantedEntitlements {
	if o == nil {
		var ret []GrantedEntitlements
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *ScopedEntitlementItem) GetEntitlementsOk() ([]GrantedEntitlements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *ScopedEntitlementItem) SetEntitlements(v []GrantedEntitlements) {
	o.Entitlements = v
}

func (o ScopedEntitlementItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopedEntitlementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceAssetId"] = o.ResourceAssetId
	toSerialize["resourceAsset"] = o.ResourceAsset
	toSerialize["entitlements"] = o.Entitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScopedEntitlementItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceAssetId",
		"resourceAsset",
		"entitlements",
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

	varScopedEntitlementItem := _ScopedEntitlementItem{}

	err = json.Unmarshal(data, &varScopedEntitlementItem)

	if err != nil {
		return err
	}

	*o = ScopedEntitlementItem(varScopedEntitlementItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceAssetId")
		delete(additionalProperties, "resourceAsset")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScopedEntitlementItem struct {
	value *ScopedEntitlementItem
	isSet bool
}

func (v NullableScopedEntitlementItem) Get() *ScopedEntitlementItem {
	return v.value
}

func (v *NullableScopedEntitlementItem) Set(val *ScopedEntitlementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedEntitlementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedEntitlementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedEntitlementItem(val *ScopedEntitlementItem) *NullableScopedEntitlementItem {
	return &NullableScopedEntitlementItem{value: val, isSet: true}
}

func (v NullableScopedEntitlementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedEntitlementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
