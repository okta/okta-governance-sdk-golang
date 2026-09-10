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

// checks if the ResourceInventoryStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceInventoryStats{}

// ResourceInventoryStats struct for ResourceInventoryStats
type ResourceInventoryStats struct {
	// The number of app resources matching the search query
	Apps int32 `json:"apps"`
	// The number of group resources matching the search query
	Groups int32 `json:"groups"`
	// The number of entitlement value resources matching the search query
	EntitlementValues int32 `json:"entitlementValues"`
	// The number of entitlement bundle resources matching the search query
	EntitlementBundles int32 `json:"entitlementBundles"`
	// The number of collection resources matching the search query
	Collections          int32 `json:"collections"`
	AdditionalProperties map[string]interface{}
}

type _ResourceInventoryStats ResourceInventoryStats

// NewResourceInventoryStats instantiates a new ResourceInventoryStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceInventoryStats(apps int32, groups int32, entitlementValues int32, entitlementBundles int32, collections int32) *ResourceInventoryStats {
	this := ResourceInventoryStats{}
	this.Apps = apps
	this.Groups = groups
	this.EntitlementValues = entitlementValues
	this.EntitlementBundles = entitlementBundles
	this.Collections = collections
	return &this
}

// NewResourceInventoryStatsWithDefaults instantiates a new ResourceInventoryStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceInventoryStatsWithDefaults() *ResourceInventoryStats {
	this := ResourceInventoryStats{}
	return &this
}

// GetApps returns the Apps field value
func (o *ResourceInventoryStats) GetApps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryStats) GetAppsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apps, true
}

// SetApps sets field value
func (o *ResourceInventoryStats) SetApps(v int32) {
	o.Apps = v
}

// GetGroups returns the Groups field value
func (o *ResourceInventoryStats) GetGroups() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryStats) GetGroupsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *ResourceInventoryStats) SetGroups(v int32) {
	o.Groups = v
}

// GetEntitlementValues returns the EntitlementValues field value
func (o *ResourceInventoryStats) GetEntitlementValues() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntitlementValues
}

// GetEntitlementValuesOk returns a tuple with the EntitlementValues field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryStats) GetEntitlementValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementValues, true
}

// SetEntitlementValues sets field value
func (o *ResourceInventoryStats) SetEntitlementValues(v int32) {
	o.EntitlementValues = v
}

// GetEntitlementBundles returns the EntitlementBundles field value
func (o *ResourceInventoryStats) GetEntitlementBundles() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntitlementBundles
}

// GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryStats) GetEntitlementBundlesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementBundles, true
}

// SetEntitlementBundles sets field value
func (o *ResourceInventoryStats) SetEntitlementBundles(v int32) {
	o.EntitlementBundles = v
}

// GetCollections returns the Collections field value
func (o *ResourceInventoryStats) GetCollections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryStats) GetCollectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collections, true
}

// SetCollections sets field value
func (o *ResourceInventoryStats) SetCollections(v int32) {
	o.Collections = v
}

func (o ResourceInventoryStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceInventoryStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apps"] = o.Apps
	toSerialize["groups"] = o.Groups
	toSerialize["entitlementValues"] = o.EntitlementValues
	toSerialize["entitlementBundles"] = o.EntitlementBundles
	toSerialize["collections"] = o.Collections

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceInventoryStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apps",
		"groups",
		"entitlementValues",
		"entitlementBundles",
		"collections",
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

	varResourceInventoryStats := _ResourceInventoryStats{}

	err = json.Unmarshal(data, &varResourceInventoryStats)

	if err != nil {
		return err
	}

	*o = ResourceInventoryStats(varResourceInventoryStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "entitlementValues")
		delete(additionalProperties, "entitlementBundles")
		delete(additionalProperties, "collections")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceInventoryStats struct {
	value *ResourceInventoryStats
	isSet bool
}

func (v NullableResourceInventoryStats) Get() *ResourceInventoryStats {
	return v.value
}

func (v *NullableResourceInventoryStats) Set(val *ResourceInventoryStats) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceInventoryStats) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceInventoryStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceInventoryStats(val *ResourceInventoryStats) *NullableResourceInventoryStats {
	return &NullableResourceInventoryStats{value: val, isSet: true}
}

func (v NullableResourceInventoryStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceInventoryStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
