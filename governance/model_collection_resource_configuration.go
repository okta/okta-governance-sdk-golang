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

// checks if the CollectionResourceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceConfiguration{}

// CollectionResourceConfiguration Server-resolved configuration details for this resource within the collection
type CollectionResourceConfiguration struct {
	// Indicates whether Entitlement Management is enabled for this resource
	EntitlementManagementEnabled bool `json:"entitlementManagementEnabled"`
	// Indicates whether this group resource has push group mappings to applications (is a push group). Only present when `type` is `GROUP`.
	HasPushMapping *bool `json:"hasPushMapping,omitempty"`
	// The type of resource
	Type                 string                    `json:"type"`
	Counts               *CollectionResourceCounts `json:"counts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceConfiguration CollectionResourceConfiguration

// NewCollectionResourceConfiguration instantiates a new CollectionResourceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceConfiguration(entitlementManagementEnabled bool, type_ string) *CollectionResourceConfiguration {
	this := CollectionResourceConfiguration{}
	this.EntitlementManagementEnabled = entitlementManagementEnabled
	this.Type = type_
	return &this
}

// NewCollectionResourceConfigurationWithDefaults instantiates a new CollectionResourceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceConfigurationWithDefaults() *CollectionResourceConfiguration {
	this := CollectionResourceConfiguration{}
	return &this
}

// GetEntitlementManagementEnabled returns the EntitlementManagementEnabled field value
func (o *CollectionResourceConfiguration) GetEntitlementManagementEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EntitlementManagementEnabled
}

// GetEntitlementManagementEnabledOk returns a tuple with the EntitlementManagementEnabled field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceConfiguration) GetEntitlementManagementEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementManagementEnabled, true
}

// SetEntitlementManagementEnabled sets field value
func (o *CollectionResourceConfiguration) SetEntitlementManagementEnabled(v bool) {
	o.EntitlementManagementEnabled = v
}

// GetHasPushMapping returns the HasPushMapping field value if set, zero value otherwise.
func (o *CollectionResourceConfiguration) GetHasPushMapping() bool {
	if o == nil || IsNil(o.HasPushMapping) {
		var ret bool
		return ret
	}
	return *o.HasPushMapping
}

// GetHasPushMappingOk returns a tuple with the HasPushMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceConfiguration) GetHasPushMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPushMapping) {
		return nil, false
	}
	return o.HasPushMapping, true
}

// HasHasPushMapping returns a boolean if a field has been set.
func (o *CollectionResourceConfiguration) HasHasPushMapping() bool {
	if o != nil && !IsNil(o.HasPushMapping) {
		return true
	}

	return false
}

// SetHasPushMapping gets a reference to the given bool and assigns it to the HasPushMapping field.
func (o *CollectionResourceConfiguration) SetHasPushMapping(v bool) {
	o.HasPushMapping = &v
}

// GetType returns the Type field value
func (o *CollectionResourceConfiguration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceConfiguration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CollectionResourceConfiguration) SetType(v string) {
	o.Type = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *CollectionResourceConfiguration) GetCounts() CollectionResourceCounts {
	if o == nil || IsNil(o.Counts) {
		var ret CollectionResourceCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceConfiguration) GetCountsOk() (*CollectionResourceCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *CollectionResourceConfiguration) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given CollectionResourceCounts and assigns it to the Counts field.
func (o *CollectionResourceConfiguration) SetCounts(v CollectionResourceCounts) {
	o.Counts = &v
}

func (o CollectionResourceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entitlementManagementEnabled"] = o.EntitlementManagementEnabled
	if !IsNil(o.HasPushMapping) {
		toSerialize["hasPushMapping"] = o.HasPushMapping
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entitlementManagementEnabled",
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

	varCollectionResourceConfiguration := _CollectionResourceConfiguration{}

	err = json.Unmarshal(data, &varCollectionResourceConfiguration)

	if err != nil {
		return err
	}

	*o = CollectionResourceConfiguration(varCollectionResourceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementManagementEnabled")
		delete(additionalProperties, "hasPushMapping")
		delete(additionalProperties, "type")
		delete(additionalProperties, "counts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceConfiguration struct {
	value *CollectionResourceConfiguration
	isSet bool
}

func (v NullableCollectionResourceConfiguration) Get() *CollectionResourceConfiguration {
	return v.value
}

func (v *NullableCollectionResourceConfiguration) Set(val *CollectionResourceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceConfiguration(val *CollectionResourceConfiguration) *NullableCollectionResourceConfiguration {
	return &NullableCollectionResourceConfiguration{value: val, isSet: true}
}

func (v NullableCollectionResourceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
