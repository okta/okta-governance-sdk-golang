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

// checks if the GrantPrincipalEntitlementGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalEntitlementGroup{}

// GrantPrincipalEntitlementGroup A group of entitlement values optionally scoped to a resource asset. When `resourceAssetOrn` is omitted, the entitlements are at the app level. When `resourceAssetOrn` is present, the entitlements are scoped to that specific resource asset.
type GrantPrincipalEntitlementGroup struct {
	// The ORN of the resource asset to scope the entitlements to. Omit this field for app-level entitlements that aren't tied to a specific resource asset.
	ResourceAssetOrn *string `json:"resourceAssetOrn,omitempty"`
	// List of entitlement value ORNs to grant within this scope
	EntitlementValueOrns []string `json:"entitlementValueOrns"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalEntitlementGroup GrantPrincipalEntitlementGroup

// NewGrantPrincipalEntitlementGroup instantiates a new GrantPrincipalEntitlementGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalEntitlementGroup(entitlementValueOrns []string) *GrantPrincipalEntitlementGroup {
	this := GrantPrincipalEntitlementGroup{}
	this.EntitlementValueOrns = entitlementValueOrns
	return &this
}

// NewGrantPrincipalEntitlementGroupWithDefaults instantiates a new GrantPrincipalEntitlementGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalEntitlementGroupWithDefaults() *GrantPrincipalEntitlementGroup {
	this := GrantPrincipalEntitlementGroup{}
	return &this
}

// GetResourceAssetOrn returns the ResourceAssetOrn field value if set, zero value otherwise.
func (o *GrantPrincipalEntitlementGroup) GetResourceAssetOrn() string {
	if o == nil || IsNil(o.ResourceAssetOrn) {
		var ret string
		return ret
	}
	return *o.ResourceAssetOrn
}

// GetResourceAssetOrnOk returns a tuple with the ResourceAssetOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementGroup) GetResourceAssetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceAssetOrn) {
		return nil, false
	}
	return o.ResourceAssetOrn, true
}

// HasResourceAssetOrn returns a boolean if a field has been set.
func (o *GrantPrincipalEntitlementGroup) HasResourceAssetOrn() bool {
	if o != nil && !IsNil(o.ResourceAssetOrn) {
		return true
	}

	return false
}

// SetResourceAssetOrn gets a reference to the given string and assigns it to the ResourceAssetOrn field.
func (o *GrantPrincipalEntitlementGroup) SetResourceAssetOrn(v string) {
	o.ResourceAssetOrn = &v
}

// GetEntitlementValueOrns returns the EntitlementValueOrns field value
func (o *GrantPrincipalEntitlementGroup) GetEntitlementValueOrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EntitlementValueOrns
}

// GetEntitlementValueOrnsOk returns a tuple with the EntitlementValueOrns field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementGroup) GetEntitlementValueOrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementValueOrns, true
}

// SetEntitlementValueOrns sets field value
func (o *GrantPrincipalEntitlementGroup) SetEntitlementValueOrns(v []string) {
	o.EntitlementValueOrns = v
}

func (o GrantPrincipalEntitlementGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalEntitlementGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceAssetOrn) {
		toSerialize["resourceAssetOrn"] = o.ResourceAssetOrn
	}
	toSerialize["entitlementValueOrns"] = o.EntitlementValueOrns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalEntitlementGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entitlementValueOrns",
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

	varGrantPrincipalEntitlementGroup := _GrantPrincipalEntitlementGroup{}

	err = json.Unmarshal(data, &varGrantPrincipalEntitlementGroup)

	if err != nil {
		return err
	}

	*o = GrantPrincipalEntitlementGroup(varGrantPrincipalEntitlementGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceAssetOrn")
		delete(additionalProperties, "entitlementValueOrns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalEntitlementGroup struct {
	value *GrantPrincipalEntitlementGroup
	isSet bool
}

func (v NullableGrantPrincipalEntitlementGroup) Get() *GrantPrincipalEntitlementGroup {
	return v.value
}

func (v *NullableGrantPrincipalEntitlementGroup) Set(val *GrantPrincipalEntitlementGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalEntitlementGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalEntitlementGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalEntitlementGroup(val *GrantPrincipalEntitlementGroup) *NullableGrantPrincipalEntitlementGroup {
	return &NullableGrantPrincipalEntitlementGroup{value: val, isSet: true}
}

func (v NullableGrantPrincipalEntitlementGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalEntitlementGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
