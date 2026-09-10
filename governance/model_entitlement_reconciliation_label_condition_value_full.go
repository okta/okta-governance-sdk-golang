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

// checks if the EntitlementReconciliationLabelConditionValueFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationLabelConditionValueFull{}

// EntitlementReconciliationLabelConditionValueFull One label value that a label condition is narrowed to
type EntitlementReconciliationLabelConditionValueFull struct {
	// The ID of a label value
	Id string `json:"id"`
	// The display name of the label value, resolved by Okta on read
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationLabelConditionValueFull EntitlementReconciliationLabelConditionValueFull

// NewEntitlementReconciliationLabelConditionValueFull instantiates a new EntitlementReconciliationLabelConditionValueFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationLabelConditionValueFull(id string) *EntitlementReconciliationLabelConditionValueFull {
	this := EntitlementReconciliationLabelConditionValueFull{}
	this.Id = id
	return &this
}

// NewEntitlementReconciliationLabelConditionValueFullWithDefaults instantiates a new EntitlementReconciliationLabelConditionValueFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationLabelConditionValueFullWithDefaults() *EntitlementReconciliationLabelConditionValueFull {
	this := EntitlementReconciliationLabelConditionValueFull{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementReconciliationLabelConditionValueFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionValueFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementReconciliationLabelConditionValueFull) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementReconciliationLabelConditionValueFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionValueFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementReconciliationLabelConditionValueFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementReconciliationLabelConditionValueFull) SetName(v string) {
	o.Name = &v
}

func (o EntitlementReconciliationLabelConditionValueFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationLabelConditionValueFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationLabelConditionValueFull) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementReconciliationLabelConditionValueFull := _EntitlementReconciliationLabelConditionValueFull{}

	err = json.Unmarshal(data, &varEntitlementReconciliationLabelConditionValueFull)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationLabelConditionValueFull(varEntitlementReconciliationLabelConditionValueFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationLabelConditionValueFull struct {
	value *EntitlementReconciliationLabelConditionValueFull
	isSet bool
}

func (v NullableEntitlementReconciliationLabelConditionValueFull) Get() *EntitlementReconciliationLabelConditionValueFull {
	return v.value
}

func (v *NullableEntitlementReconciliationLabelConditionValueFull) Set(val *EntitlementReconciliationLabelConditionValueFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationLabelConditionValueFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationLabelConditionValueFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationLabelConditionValueFull(val *EntitlementReconciliationLabelConditionValueFull) *NullableEntitlementReconciliationLabelConditionValueFull {
	return &NullableEntitlementReconciliationLabelConditionValueFull{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationLabelConditionValueFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationLabelConditionValueFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
