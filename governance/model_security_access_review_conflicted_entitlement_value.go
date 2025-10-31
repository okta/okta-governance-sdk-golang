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

// checks if the SecurityAccessReviewConflictedEntitlementValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewConflictedEntitlementValue{}

// SecurityAccessReviewConflictedEntitlementValue struct for SecurityAccessReviewConflictedEntitlementValue
type SecurityAccessReviewConflictedEntitlementValue struct {
	// ID of entitlement value
	Id string `json:"id"`
	// Entitlement value name
	Name string `json:"name"`
	// A true value indicates that the associated resource is for this entitlement value or is a bundle that contains this entitlement value.
	SourceOfConflict     bool `json:"sourceOfConflict"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewConflictedEntitlementValue SecurityAccessReviewConflictedEntitlementValue

// NewSecurityAccessReviewConflictedEntitlementValue instantiates a new SecurityAccessReviewConflictedEntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewConflictedEntitlementValue(id string, name string, sourceOfConflict bool) *SecurityAccessReviewConflictedEntitlementValue {
	this := SecurityAccessReviewConflictedEntitlementValue{}
	this.Id = id
	this.Name = name
	this.SourceOfConflict = sourceOfConflict
	return &this
}

// NewSecurityAccessReviewConflictedEntitlementValueWithDefaults instantiates a new SecurityAccessReviewConflictedEntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewConflictedEntitlementValueWithDefaults() *SecurityAccessReviewConflictedEntitlementValue {
	this := SecurityAccessReviewConflictedEntitlementValue{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewConflictedEntitlementValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlementValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewConflictedEntitlementValue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewConflictedEntitlementValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlementValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewConflictedEntitlementValue) SetName(v string) {
	o.Name = v
}

// GetSourceOfConflict returns the SourceOfConflict field value
func (o *SecurityAccessReviewConflictedEntitlementValue) GetSourceOfConflict() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SourceOfConflict
}

// GetSourceOfConflictOk returns a tuple with the SourceOfConflict field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewConflictedEntitlementValue) GetSourceOfConflictOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOfConflict, true
}

// SetSourceOfConflict sets field value
func (o *SecurityAccessReviewConflictedEntitlementValue) SetSourceOfConflict(v bool) {
	o.SourceOfConflict = v
}

func (o SecurityAccessReviewConflictedEntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewConflictedEntitlementValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["sourceOfConflict"] = o.SourceOfConflict

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewConflictedEntitlementValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"sourceOfConflict",
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

	varSecurityAccessReviewConflictedEntitlementValue := _SecurityAccessReviewConflictedEntitlementValue{}

	err = json.Unmarshal(data, &varSecurityAccessReviewConflictedEntitlementValue)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewConflictedEntitlementValue(varSecurityAccessReviewConflictedEntitlementValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sourceOfConflict")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewConflictedEntitlementValue struct {
	value *SecurityAccessReviewConflictedEntitlementValue
	isSet bool
}

func (v NullableSecurityAccessReviewConflictedEntitlementValue) Get() *SecurityAccessReviewConflictedEntitlementValue {
	return v.value
}

func (v *NullableSecurityAccessReviewConflictedEntitlementValue) Set(val *SecurityAccessReviewConflictedEntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewConflictedEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewConflictedEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewConflictedEntitlementValue(val *SecurityAccessReviewConflictedEntitlementValue) *NullableSecurityAccessReviewConflictedEntitlementValue {
	return &NullableSecurityAccessReviewConflictedEntitlementValue{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewConflictedEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewConflictedEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
