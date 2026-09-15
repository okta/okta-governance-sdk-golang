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

// checks if the EntitlementReconciliationEntitlementConditionWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationEntitlementConditionWritable{}

// EntitlementReconciliationEntitlementConditionWritable Matches drifts on one entitlement, optionally narrowed to specific values of it.  Leave `values` off, or send it empty, to match every value of the entitlement, including values that don't exist yet. Okta stores that as the entitlement itself, not as the list of values it happens to have when you save.
type EntitlementReconciliationEntitlementConditionWritable struct {
	// Identifies this condition as targeting an entitlement
	RefType string `json:"refType"`
	// The `id` property of an entitlement
	EntitlementId string `json:"entitlementId"`
	// The values of the entitlement to narrow to. Omit it, or send an empty array, to match every value of the entitlement.
	Values               []EntitlementReconciliationEntitlementConditionValueWritable `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationEntitlementConditionWritable EntitlementReconciliationEntitlementConditionWritable

// NewEntitlementReconciliationEntitlementConditionWritable instantiates a new EntitlementReconciliationEntitlementConditionWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationEntitlementConditionWritable(refType string, entitlementId string) *EntitlementReconciliationEntitlementConditionWritable {
	this := EntitlementReconciliationEntitlementConditionWritable{}
	this.RefType = refType
	this.EntitlementId = entitlementId
	return &this
}

// NewEntitlementReconciliationEntitlementConditionWritableWithDefaults instantiates a new EntitlementReconciliationEntitlementConditionWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationEntitlementConditionWritableWithDefaults() *EntitlementReconciliationEntitlementConditionWritable {
	this := EntitlementReconciliationEntitlementConditionWritable{}
	return &this
}

// GetRefType returns the RefType field value
func (o *EntitlementReconciliationEntitlementConditionWritable) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionWritable) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementReconciliationEntitlementConditionWritable) SetRefType(v string) {
	o.RefType = v
}

// GetEntitlementId returns the EntitlementId field value
func (o *EntitlementReconciliationEntitlementConditionWritable) GetEntitlementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionWritable) GetEntitlementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementId, true
}

// SetEntitlementId sets field value
func (o *EntitlementReconciliationEntitlementConditionWritable) SetEntitlementId(v string) {
	o.EntitlementId = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementReconciliationEntitlementConditionWritable) GetValues() []EntitlementReconciliationEntitlementConditionValueWritable {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementReconciliationEntitlementConditionValueWritable
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionWritable) GetValuesOk() ([]EntitlementReconciliationEntitlementConditionValueWritable, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementReconciliationEntitlementConditionWritable) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementReconciliationEntitlementConditionValueWritable and assigns it to the Values field.
func (o *EntitlementReconciliationEntitlementConditionWritable) SetValues(v []EntitlementReconciliationEntitlementConditionValueWritable) {
	o.Values = v
}

func (o EntitlementReconciliationEntitlementConditionWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationEntitlementConditionWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refType"] = o.RefType
	toSerialize["entitlementId"] = o.EntitlementId
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationEntitlementConditionWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"refType",
		"entitlementId",
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

	varEntitlementReconciliationEntitlementConditionWritable := _EntitlementReconciliationEntitlementConditionWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationEntitlementConditionWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationEntitlementConditionWritable(varEntitlementReconciliationEntitlementConditionWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "refType")
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationEntitlementConditionWritable struct {
	value *EntitlementReconciliationEntitlementConditionWritable
	isSet bool
}

func (v NullableEntitlementReconciliationEntitlementConditionWritable) Get() *EntitlementReconciliationEntitlementConditionWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationEntitlementConditionWritable) Set(val *EntitlementReconciliationEntitlementConditionWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationEntitlementConditionWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationEntitlementConditionWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationEntitlementConditionWritable(val *EntitlementReconciliationEntitlementConditionWritable) *NullableEntitlementReconciliationEntitlementConditionWritable {
	return &NullableEntitlementReconciliationEntitlementConditionWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationEntitlementConditionWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationEntitlementConditionWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
