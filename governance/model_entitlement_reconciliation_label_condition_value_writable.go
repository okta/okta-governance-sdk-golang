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

// checks if the EntitlementReconciliationLabelConditionValueWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationLabelConditionValueWritable{}

// EntitlementReconciliationLabelConditionValueWritable One label value that a label condition is narrowed to
type EntitlementReconciliationLabelConditionValueWritable struct {
	// The ID of a label value
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationLabelConditionValueWritable EntitlementReconciliationLabelConditionValueWritable

// NewEntitlementReconciliationLabelConditionValueWritable instantiates a new EntitlementReconciliationLabelConditionValueWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationLabelConditionValueWritable(id string) *EntitlementReconciliationLabelConditionValueWritable {
	this := EntitlementReconciliationLabelConditionValueWritable{}
	this.Id = id
	return &this
}

// NewEntitlementReconciliationLabelConditionValueWritableWithDefaults instantiates a new EntitlementReconciliationLabelConditionValueWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationLabelConditionValueWritableWithDefaults() *EntitlementReconciliationLabelConditionValueWritable {
	this := EntitlementReconciliationLabelConditionValueWritable{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementReconciliationLabelConditionValueWritable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionValueWritable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementReconciliationLabelConditionValueWritable) SetId(v string) {
	o.Id = v
}

func (o EntitlementReconciliationLabelConditionValueWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationLabelConditionValueWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationLabelConditionValueWritable) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementReconciliationLabelConditionValueWritable := _EntitlementReconciliationLabelConditionValueWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationLabelConditionValueWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationLabelConditionValueWritable(varEntitlementReconciliationLabelConditionValueWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationLabelConditionValueWritable struct {
	value *EntitlementReconciliationLabelConditionValueWritable
	isSet bool
}

func (v NullableEntitlementReconciliationLabelConditionValueWritable) Get() *EntitlementReconciliationLabelConditionValueWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationLabelConditionValueWritable) Set(val *EntitlementReconciliationLabelConditionValueWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationLabelConditionValueWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationLabelConditionValueWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationLabelConditionValueWritable(val *EntitlementReconciliationLabelConditionValueWritable) *NullableEntitlementReconciliationLabelConditionValueWritable {
	return &NullableEntitlementReconciliationLabelConditionValueWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationLabelConditionValueWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationLabelConditionValueWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
