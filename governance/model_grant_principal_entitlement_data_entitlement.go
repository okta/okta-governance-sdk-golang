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

// checks if the GrantPrincipalEntitlementDataEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalEntitlementDataEntitlement{}

// GrantPrincipalEntitlementDataEntitlement Entitlement data for direct entitlement assignment. Entitlements can be grouped by resource asset scope. Each group contains one or more entitlement values to grant.
type GrantPrincipalEntitlementDataEntitlement struct {
	// Indicates direct entitlement assignment
	Type string `json:"type"`
	// List of entitlement data. Each entry optionally specifies a resource asset scope. Entry without `resourceAssetOrn` represent app-level entitlements.
	Entitlements         []GrantPrincipalEntitlementGroup `json:"entitlements"`
	AccessDuration       *AccessDuration                  `json:"accessDuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalEntitlementDataEntitlement GrantPrincipalEntitlementDataEntitlement

// NewGrantPrincipalEntitlementDataEntitlement instantiates a new GrantPrincipalEntitlementDataEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalEntitlementDataEntitlement(type_ string, entitlements []GrantPrincipalEntitlementGroup) *GrantPrincipalEntitlementDataEntitlement {
	this := GrantPrincipalEntitlementDataEntitlement{}
	this.Type = type_
	this.Entitlements = entitlements
	return &this
}

// NewGrantPrincipalEntitlementDataEntitlementWithDefaults instantiates a new GrantPrincipalEntitlementDataEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalEntitlementDataEntitlementWithDefaults() *GrantPrincipalEntitlementDataEntitlement {
	this := GrantPrincipalEntitlementDataEntitlement{}
	return &this
}

// GetType returns the Type field value
func (o *GrantPrincipalEntitlementDataEntitlement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataEntitlement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GrantPrincipalEntitlementDataEntitlement) SetType(v string) {
	o.Type = v
}

// GetEntitlements returns the Entitlements field value
func (o *GrantPrincipalEntitlementDataEntitlement) GetEntitlements() []GrantPrincipalEntitlementGroup {
	if o == nil {
		var ret []GrantPrincipalEntitlementGroup
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataEntitlement) GetEntitlementsOk() ([]GrantPrincipalEntitlementGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *GrantPrincipalEntitlementDataEntitlement) SetEntitlements(v []GrantPrincipalEntitlementGroup) {
	o.Entitlements = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise.
func (o *GrantPrincipalEntitlementDataEntitlement) GetAccessDuration() AccessDuration {
	if o == nil || IsNil(o.AccessDuration) {
		var ret AccessDuration
		return ret
	}
	return *o.AccessDuration
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataEntitlement) GetAccessDurationOk() (*AccessDuration, bool) {
	if o == nil || IsNil(o.AccessDuration) {
		return nil, false
	}
	return o.AccessDuration, true
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *GrantPrincipalEntitlementDataEntitlement) HasAccessDuration() bool {
	if o != nil && !IsNil(o.AccessDuration) {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given AccessDuration and assigns it to the AccessDuration field.
func (o *GrantPrincipalEntitlementDataEntitlement) SetAccessDuration(v AccessDuration) {
	o.AccessDuration = &v
}

func (o GrantPrincipalEntitlementDataEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalEntitlementDataEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["entitlements"] = o.Entitlements
	if !IsNil(o.AccessDuration) {
		toSerialize["accessDuration"] = o.AccessDuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalEntitlementDataEntitlement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varGrantPrincipalEntitlementDataEntitlement := _GrantPrincipalEntitlementDataEntitlement{}

	err = json.Unmarshal(data, &varGrantPrincipalEntitlementDataEntitlement)

	if err != nil {
		return err
	}

	*o = GrantPrincipalEntitlementDataEntitlement(varGrantPrincipalEntitlementDataEntitlement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "accessDuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalEntitlementDataEntitlement struct {
	value *GrantPrincipalEntitlementDataEntitlement
	isSet bool
}

func (v NullableGrantPrincipalEntitlementDataEntitlement) Get() *GrantPrincipalEntitlementDataEntitlement {
	return v.value
}

func (v *NullableGrantPrincipalEntitlementDataEntitlement) Set(val *GrantPrincipalEntitlementDataEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalEntitlementDataEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalEntitlementDataEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalEntitlementDataEntitlement(val *GrantPrincipalEntitlementDataEntitlement) *NullableGrantPrincipalEntitlementDataEntitlement {
	return &NullableGrantPrincipalEntitlementDataEntitlement{value: val, isSet: true}
}

func (v NullableGrantPrincipalEntitlementDataEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalEntitlementDataEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
