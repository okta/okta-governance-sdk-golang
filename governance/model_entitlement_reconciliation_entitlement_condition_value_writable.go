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

// checks if the EntitlementReconciliationEntitlementConditionValueWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationEntitlementConditionValueWritable{}

// EntitlementReconciliationEntitlementConditionValueWritable One entitlement value that an entitlement condition is narrowed to
type EntitlementReconciliationEntitlementConditionValueWritable struct {
	// The `id` of the entitlement value
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationEntitlementConditionValueWritable EntitlementReconciliationEntitlementConditionValueWritable

// NewEntitlementReconciliationEntitlementConditionValueWritable instantiates a new EntitlementReconciliationEntitlementConditionValueWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationEntitlementConditionValueWritable(id string) *EntitlementReconciliationEntitlementConditionValueWritable {
	this := EntitlementReconciliationEntitlementConditionValueWritable{}
	this.Id = id
	return &this
}

// NewEntitlementReconciliationEntitlementConditionValueWritableWithDefaults instantiates a new EntitlementReconciliationEntitlementConditionValueWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationEntitlementConditionValueWritableWithDefaults() *EntitlementReconciliationEntitlementConditionValueWritable {
	this := EntitlementReconciliationEntitlementConditionValueWritable{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementReconciliationEntitlementConditionValueWritable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionValueWritable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementReconciliationEntitlementConditionValueWritable) SetId(v string) {
	o.Id = v
}

func (o EntitlementReconciliationEntitlementConditionValueWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationEntitlementConditionValueWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationEntitlementConditionValueWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varEntitlementReconciliationEntitlementConditionValueWritable := _EntitlementReconciliationEntitlementConditionValueWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationEntitlementConditionValueWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationEntitlementConditionValueWritable(varEntitlementReconciliationEntitlementConditionValueWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationEntitlementConditionValueWritable struct {
	value *EntitlementReconciliationEntitlementConditionValueWritable
	isSet bool
}

func (v NullableEntitlementReconciliationEntitlementConditionValueWritable) Get() *EntitlementReconciliationEntitlementConditionValueWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationEntitlementConditionValueWritable) Set(val *EntitlementReconciliationEntitlementConditionValueWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationEntitlementConditionValueWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationEntitlementConditionValueWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationEntitlementConditionValueWritable(val *EntitlementReconciliationEntitlementConditionValueWritable) *NullableEntitlementReconciliationEntitlementConditionValueWritable {
	return &NullableEntitlementReconciliationEntitlementConditionValueWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationEntitlementConditionValueWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationEntitlementConditionValueWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
