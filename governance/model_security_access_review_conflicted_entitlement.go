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

// checks if the SecurityAccessReviewConflictedEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewConflictedEntitlement{}

// SecurityAccessReviewConflictedEntitlement struct for SecurityAccessReviewConflictedEntitlement
type SecurityAccessReviewConflictedEntitlement struct {
	// Entitlement ID
	Id string `json:"id"`
	// Entitlement name
	Name string `json:"name"`
	// List of entitlement values
	Values               []SecurityAccessReviewConflictedEntitlementValue `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewConflictedEntitlement SecurityAccessReviewConflictedEntitlement

// NewSecurityAccessReviewConflictedEntitlement instantiates a new SecurityAccessReviewConflictedEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewConflictedEntitlement(id string, name string, values []SecurityAccessReviewConflictedEntitlementValue) *SecurityAccessReviewConflictedEntitlement {
	this := SecurityAccessReviewConflictedEntitlement{}
	this.Id = id
	this.Name = name
	this.Values = values
	return &this
}

// NewSecurityAccessReviewConflictedEntitlementWithDefaults instantiates a new SecurityAccessReviewConflictedEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewConflictedEntitlementWithDefaults() *SecurityAccessReviewConflictedEntitlement {
	this := SecurityAccessReviewConflictedEntitlement{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewConflictedEntitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewConflictedEntitlement) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewConflictedEntitlement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewConflictedEntitlement) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *SecurityAccessReviewConflictedEntitlement) GetValues() []SecurityAccessReviewConflictedEntitlementValue {
	if o == nil {
		var ret []SecurityAccessReviewConflictedEntitlementValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlement) GetValuesOk() ([]SecurityAccessReviewConflictedEntitlementValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *SecurityAccessReviewConflictedEntitlement) SetValues(v []SecurityAccessReviewConflictedEntitlementValue) {
	o.Values = v
}

func (o SecurityAccessReviewConflictedEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewConflictedEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewConflictedEntitlement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"values",
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

	varSecurityAccessReviewConflictedEntitlement := _SecurityAccessReviewConflictedEntitlement{}

	err = json.Unmarshal(data, &varSecurityAccessReviewConflictedEntitlement)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewConflictedEntitlement(varSecurityAccessReviewConflictedEntitlement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewConflictedEntitlement struct {
	value *SecurityAccessReviewConflictedEntitlement
	isSet bool
}

func (v NullableSecurityAccessReviewConflictedEntitlement) Get() *SecurityAccessReviewConflictedEntitlement {
	return v.value
}

func (v *NullableSecurityAccessReviewConflictedEntitlement) Set(val *SecurityAccessReviewConflictedEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewConflictedEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewConflictedEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewConflictedEntitlement(val *SecurityAccessReviewConflictedEntitlement) *NullableSecurityAccessReviewConflictedEntitlement {
	return &NullableSecurityAccessReviewConflictedEntitlement{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewConflictedEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewConflictedEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
