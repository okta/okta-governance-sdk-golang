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

// checks if the GrantPrincipalEntitlementDataBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalEntitlementDataBundle{}

// GrantPrincipalEntitlementDataBundle Entitlement data for entitlement bundle assignment. The server resolves the bundle contents to determine the actual entitlements to grant.
type GrantPrincipalEntitlementDataBundle struct {
	// Indicates entitlement bundle assignment
	Type string `json:"type"`
	// The ORN of the entitlement bundle to grant
	EntitlementBundleOrn string          `json:"entitlementBundleOrn"`
	AccessDuration       *AccessDuration `json:"accessDuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalEntitlementDataBundle GrantPrincipalEntitlementDataBundle

// NewGrantPrincipalEntitlementDataBundle instantiates a new GrantPrincipalEntitlementDataBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalEntitlementDataBundle(type_ string, entitlementBundleOrn string) *GrantPrincipalEntitlementDataBundle {
	this := GrantPrincipalEntitlementDataBundle{}
	this.Type = type_
	this.EntitlementBundleOrn = entitlementBundleOrn
	return &this
}

// NewGrantPrincipalEntitlementDataBundleWithDefaults instantiates a new GrantPrincipalEntitlementDataBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalEntitlementDataBundleWithDefaults() *GrantPrincipalEntitlementDataBundle {
	this := GrantPrincipalEntitlementDataBundle{}
	return &this
}

// GetType returns the Type field value
func (o *GrantPrincipalEntitlementDataBundle) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataBundle) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GrantPrincipalEntitlementDataBundle) SetType(v string) {
	o.Type = v
}

// GetEntitlementBundleOrn returns the EntitlementBundleOrn field value
func (o *GrantPrincipalEntitlementDataBundle) GetEntitlementBundleOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementBundleOrn
}

// GetEntitlementBundleOrnOk returns a tuple with the EntitlementBundleOrn field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataBundle) GetEntitlementBundleOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementBundleOrn, true
}

// SetEntitlementBundleOrn sets field value
func (o *GrantPrincipalEntitlementDataBundle) SetEntitlementBundleOrn(v string) {
	o.EntitlementBundleOrn = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise.
func (o *GrantPrincipalEntitlementDataBundle) GetAccessDuration() AccessDuration {
	if o == nil || IsNil(o.AccessDuration) {
		var ret AccessDuration
		return ret
	}
	return *o.AccessDuration
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalEntitlementDataBundle) GetAccessDurationOk() (*AccessDuration, bool) {
	if o == nil || IsNil(o.AccessDuration) {
		return nil, false
	}
	return o.AccessDuration, true
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *GrantPrincipalEntitlementDataBundle) HasAccessDuration() bool {
	if o != nil && !IsNil(o.AccessDuration) {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given AccessDuration and assigns it to the AccessDuration field.
func (o *GrantPrincipalEntitlementDataBundle) SetAccessDuration(v AccessDuration) {
	o.AccessDuration = &v
}

func (o GrantPrincipalEntitlementDataBundle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalEntitlementDataBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["entitlementBundleOrn"] = o.EntitlementBundleOrn
	if !IsNil(o.AccessDuration) {
		toSerialize["accessDuration"] = o.AccessDuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalEntitlementDataBundle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"entitlementBundleOrn",
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

	varGrantPrincipalEntitlementDataBundle := _GrantPrincipalEntitlementDataBundle{}

	err = json.Unmarshal(data, &varGrantPrincipalEntitlementDataBundle)

	if err != nil {
		return err
	}

	*o = GrantPrincipalEntitlementDataBundle(varGrantPrincipalEntitlementDataBundle)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "entitlementBundleOrn")
		delete(additionalProperties, "accessDuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalEntitlementDataBundle struct {
	value *GrantPrincipalEntitlementDataBundle
	isSet bool
}

func (v NullableGrantPrincipalEntitlementDataBundle) Get() *GrantPrincipalEntitlementDataBundle {
	return v.value
}

func (v *NullableGrantPrincipalEntitlementDataBundle) Set(val *GrantPrincipalEntitlementDataBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalEntitlementDataBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalEntitlementDataBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalEntitlementDataBundle(val *GrantPrincipalEntitlementDataBundle) *NullableGrantPrincipalEntitlementDataBundle {
	return &NullableGrantPrincipalEntitlementDataBundle{value: val, isSet: true}
}

func (v NullableGrantPrincipalEntitlementDataBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalEntitlementDataBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
