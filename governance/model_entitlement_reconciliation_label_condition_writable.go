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

// checks if the EntitlementReconciliationLabelConditionWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationLabelConditionWritable{}

// EntitlementReconciliationLabelConditionWritable Matches drifts on entitlements carrying one label, optionally narrowed to specific values of it.  Leave `values` off, or send it empty, to match every value of the label, including values that don't exist yet. Okta stores that as the label itself, not as the list of values the label happens to have when you save.
type EntitlementReconciliationLabelConditionWritable struct {
	// Identifies this condition as targeting a label
	RefType string `json:"refType"`
	// The ID of a label
	LabelId string `json:"labelId"`
	// The values of the label to narrow to. Omit it, or send an empty array, to match every value of the label.
	Values               []EntitlementReconciliationLabelConditionValueWritable `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationLabelConditionWritable EntitlementReconciliationLabelConditionWritable

// NewEntitlementReconciliationLabelConditionWritable instantiates a new EntitlementReconciliationLabelConditionWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationLabelConditionWritable(refType string, labelId string) *EntitlementReconciliationLabelConditionWritable {
	this := EntitlementReconciliationLabelConditionWritable{}
	this.RefType = refType
	this.LabelId = labelId
	return &this
}

// NewEntitlementReconciliationLabelConditionWritableWithDefaults instantiates a new EntitlementReconciliationLabelConditionWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationLabelConditionWritableWithDefaults() *EntitlementReconciliationLabelConditionWritable {
	this := EntitlementReconciliationLabelConditionWritable{}
	return &this
}

// GetRefType returns the RefType field value
func (o *EntitlementReconciliationLabelConditionWritable) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionWritable) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementReconciliationLabelConditionWritable) SetRefType(v string) {
	o.RefType = v
}

// GetLabelId returns the LabelId field value
func (o *EntitlementReconciliationLabelConditionWritable) GetLabelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionWritable) GetLabelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelId, true
}

// SetLabelId sets field value
func (o *EntitlementReconciliationLabelConditionWritable) SetLabelId(v string) {
	o.LabelId = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementReconciliationLabelConditionWritable) GetValues() []EntitlementReconciliationLabelConditionValueWritable {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementReconciliationLabelConditionValueWritable
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionWritable) GetValuesOk() ([]EntitlementReconciliationLabelConditionValueWritable, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementReconciliationLabelConditionWritable) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementReconciliationLabelConditionValueWritable and assigns it to the Values field.
func (o *EntitlementReconciliationLabelConditionWritable) SetValues(v []EntitlementReconciliationLabelConditionValueWritable) {
	o.Values = v
}

func (o EntitlementReconciliationLabelConditionWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationLabelConditionWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refType"] = o.RefType
	toSerialize["labelId"] = o.LabelId
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationLabelConditionWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"refType",
		"labelId",
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

	varEntitlementReconciliationLabelConditionWritable := _EntitlementReconciliationLabelConditionWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationLabelConditionWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationLabelConditionWritable(varEntitlementReconciliationLabelConditionWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "refType")
		delete(additionalProperties, "labelId")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationLabelConditionWritable struct {
	value *EntitlementReconciliationLabelConditionWritable
	isSet bool
}

func (v NullableEntitlementReconciliationLabelConditionWritable) Get() *EntitlementReconciliationLabelConditionWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationLabelConditionWritable) Set(val *EntitlementReconciliationLabelConditionWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationLabelConditionWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationLabelConditionWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationLabelConditionWritable(val *EntitlementReconciliationLabelConditionWritable) *NullableEntitlementReconciliationLabelConditionWritable {
	return &NullableEntitlementReconciliationLabelConditionWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationLabelConditionWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationLabelConditionWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
