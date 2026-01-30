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

// checks if the SecurityAccessReviewSodConflictConflictedEntitlementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewSodConflictConflictedEntitlementInfo{}

// SecurityAccessReviewSodConflictConflictedEntitlementInfo struct for SecurityAccessReviewSodConflictConflictedEntitlementInfo
type SecurityAccessReviewSodConflictConflictedEntitlementInfo struct {
	// The entitlement list name
	ListName string                                             `json:"listName"`
	Scope    SecurityAccessReviewConflictedEntitlementListScope `json:"scope"`
	// List of entitlements
	Entitlements         []SecurityAccessReviewConflictedEntitlement `json:"entitlements"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewSodConflictConflictedEntitlementInfo SecurityAccessReviewSodConflictConflictedEntitlementInfo

// NewSecurityAccessReviewSodConflictConflictedEntitlementInfo instantiates a new SecurityAccessReviewSodConflictConflictedEntitlementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSodConflictConflictedEntitlementInfo(listName string, scope SecurityAccessReviewConflictedEntitlementListScope, entitlements []SecurityAccessReviewConflictedEntitlement) *SecurityAccessReviewSodConflictConflictedEntitlementInfo {
	this := SecurityAccessReviewSodConflictConflictedEntitlementInfo{}
	this.ListName = listName
	this.Scope = scope
	this.Entitlements = entitlements
	return &this
}

// NewSecurityAccessReviewSodConflictConflictedEntitlementInfoWithDefaults instantiates a new SecurityAccessReviewSodConflictConflictedEntitlementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSodConflictConflictedEntitlementInfoWithDefaults() *SecurityAccessReviewSodConflictConflictedEntitlementInfo {
	this := SecurityAccessReviewSodConflictConflictedEntitlementInfo{}
	return &this
}

// GetListName returns the ListName field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetListName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListName
}

// GetListNameOk returns a tuple with the ListName field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetListNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListName, true
}

// SetListName sets field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetListName(v string) {
	o.ListName = v
}

// GetScope returns the Scope field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetScope() SecurityAccessReviewConflictedEntitlementListScope {
	if o == nil {
		var ret SecurityAccessReviewConflictedEntitlementListScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetScopeOk() (*SecurityAccessReviewConflictedEntitlementListScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetScope(v SecurityAccessReviewConflictedEntitlementListScope) {
	o.Scope = v
}

// GetEntitlements returns the Entitlements field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetEntitlements() []SecurityAccessReviewConflictedEntitlement {
	if o == nil {
		var ret []SecurityAccessReviewConflictedEntitlement
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetEntitlementsOk() ([]SecurityAccessReviewConflictedEntitlement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetEntitlements(v []SecurityAccessReviewConflictedEntitlement) {
	o.Entitlements = v
}

func (o SecurityAccessReviewSodConflictConflictedEntitlementInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewSodConflictConflictedEntitlementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listName"] = o.ListName
	toSerialize["scope"] = o.Scope
	toSerialize["entitlements"] = o.Entitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listName",
		"scope",
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

	varSecurityAccessReviewSodConflictConflictedEntitlementInfo := _SecurityAccessReviewSodConflictConflictedEntitlementInfo{}

	err = json.Unmarshal(data, &varSecurityAccessReviewSodConflictConflictedEntitlementInfo)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewSodConflictConflictedEntitlementInfo(varSecurityAccessReviewSodConflictConflictedEntitlementInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listName")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo struct {
	value *SecurityAccessReviewSodConflictConflictedEntitlementInfo
	isSet bool
}

func (v NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) Get() *SecurityAccessReviewSodConflictConflictedEntitlementInfo {
	return v.value
}

func (v *NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) Set(val *SecurityAccessReviewSodConflictConflictedEntitlementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSodConflictConflictedEntitlementInfo(val *SecurityAccessReviewSodConflictConflictedEntitlementInfo) *NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo {
	return &NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSodConflictConflictedEntitlementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
