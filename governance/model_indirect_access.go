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

// checks if the IndirectAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndirectAccess{}

// IndirectAccess Access obtained indirectly through group membership or resource asset hierarchy including the entitlements
type IndirectAccess struct {
	Grant GrantDetails `json:"grant"`
	// Entitlements obtained through this grant grouped by resource asset scope
	ScopedEntitlements   []ScopedEntitlementItem `json:"scopedEntitlements"`
	AdditionalProperties map[string]interface{}
}

type _IndirectAccess IndirectAccess

// NewIndirectAccess instantiates a new IndirectAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndirectAccess(grant GrantDetails, scopedEntitlements []ScopedEntitlementItem) *IndirectAccess {
	this := IndirectAccess{}
	this.Grant = grant
	this.ScopedEntitlements = scopedEntitlements
	return &this
}

// NewIndirectAccessWithDefaults instantiates a new IndirectAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndirectAccessWithDefaults() *IndirectAccess {
	this := IndirectAccess{}
	return &this
}

// GetGrant returns the Grant field value
func (o *IndirectAccess) GetGrant() GrantDetails {
	if o == nil {
		var ret GrantDetails
		return ret
	}

	return o.Grant
}

// GetGrantOk returns a tuple with the Grant field value
// and a boolean to check if the value has been set.
func (o *IndirectAccess) GetGrantOk() (*GrantDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grant, true
}

// SetGrant sets field value
func (o *IndirectAccess) SetGrant(v GrantDetails) {
	o.Grant = v
}

// GetScopedEntitlements returns the ScopedEntitlements field value
func (o *IndirectAccess) GetScopedEntitlements() []ScopedEntitlementItem {
	if o == nil {
		var ret []ScopedEntitlementItem
		return ret
	}

	return o.ScopedEntitlements
}

// GetScopedEntitlementsOk returns a tuple with the ScopedEntitlements field value
// and a boolean to check if the value has been set.
func (o *IndirectAccess) GetScopedEntitlementsOk() ([]ScopedEntitlementItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopedEntitlements, true
}

// SetScopedEntitlements sets field value
func (o *IndirectAccess) SetScopedEntitlements(v []ScopedEntitlementItem) {
	o.ScopedEntitlements = v
}

func (o IndirectAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndirectAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grant"] = o.Grant
	toSerialize["scopedEntitlements"] = o.ScopedEntitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndirectAccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grant",
		"scopedEntitlements",
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

	varIndirectAccess := _IndirectAccess{}

	err = json.Unmarshal(data, &varIndirectAccess)

	if err != nil {
		return err
	}

	*o = IndirectAccess(varIndirectAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grant")
		delete(additionalProperties, "scopedEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndirectAccess struct {
	value *IndirectAccess
	isSet bool
}

func (v NullableIndirectAccess) Get() *IndirectAccess {
	return v.value
}

func (v *NullableIndirectAccess) Set(val *IndirectAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableIndirectAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableIndirectAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndirectAccess(val *IndirectAccess) *NullableIndirectAccess {
	return &NullableIndirectAccess{value: val, isSet: true}
}

func (v NullableIndirectAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndirectAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
