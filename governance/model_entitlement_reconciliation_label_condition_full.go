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

// checks if the EntitlementReconciliationLabelConditionFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationLabelConditionFull{}

// EntitlementReconciliationLabelConditionFull Matches drifts on entitlements carrying one label, optionally narrowed to specific values of it.  When `values` is absent or empty the condition matches every value of the label, including values that don't exist yet. Okta stores the label itself rather than the list of values it happened to have when the configuration was saved, so this keeps matching as the label's values change.  `labelName` and the `name` on each value are conveniences for rendering the condition. Okta resolves them on read, and they aren't part of the writable model.
type EntitlementReconciliationLabelConditionFull struct {
	// Identifies this condition as targeting a label
	RefType string `json:"refType"`
	// The ID of a label
	LabelId string `json:"labelId"`
	// The display name of the label, resolved by Okta on read
	LabelName *string `json:"labelName,omitempty"`
	// The values of the label the condition is narrowed to. Absent or empty when the condition matches every value.
	Values               []EntitlementReconciliationLabelConditionValueFull `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationLabelConditionFull EntitlementReconciliationLabelConditionFull

// NewEntitlementReconciliationLabelConditionFull instantiates a new EntitlementReconciliationLabelConditionFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationLabelConditionFull(refType string, labelId string) *EntitlementReconciliationLabelConditionFull {
	this := EntitlementReconciliationLabelConditionFull{}
	this.RefType = refType
	this.LabelId = labelId
	return &this
}

// NewEntitlementReconciliationLabelConditionFullWithDefaults instantiates a new EntitlementReconciliationLabelConditionFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationLabelConditionFullWithDefaults() *EntitlementReconciliationLabelConditionFull {
	this := EntitlementReconciliationLabelConditionFull{}
	return &this
}

// GetRefType returns the RefType field value
func (o *EntitlementReconciliationLabelConditionFull) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionFull) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementReconciliationLabelConditionFull) SetRefType(v string) {
	o.RefType = v
}

// GetLabelId returns the LabelId field value
func (o *EntitlementReconciliationLabelConditionFull) GetLabelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionFull) GetLabelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelId, true
}

// SetLabelId sets field value
func (o *EntitlementReconciliationLabelConditionFull) SetLabelId(v string) {
	o.LabelId = v
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *EntitlementReconciliationLabelConditionFull) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionFull) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *EntitlementReconciliationLabelConditionFull) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *EntitlementReconciliationLabelConditionFull) SetLabelName(v string) {
	o.LabelName = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementReconciliationLabelConditionFull) GetValues() []EntitlementReconciliationLabelConditionValueFull {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementReconciliationLabelConditionValueFull
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationLabelConditionFull) GetValuesOk() ([]EntitlementReconciliationLabelConditionValueFull, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementReconciliationLabelConditionFull) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementReconciliationLabelConditionValueFull and assigns it to the Values field.
func (o *EntitlementReconciliationLabelConditionFull) SetValues(v []EntitlementReconciliationLabelConditionValueFull) {
	o.Values = v
}

func (o EntitlementReconciliationLabelConditionFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationLabelConditionFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refType"] = o.RefType
	toSerialize["labelId"] = o.LabelId
	if !IsNil(o.LabelName) {
		toSerialize["labelName"] = o.LabelName
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationLabelConditionFull) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementReconciliationLabelConditionFull := _EntitlementReconciliationLabelConditionFull{}

	err = json.Unmarshal(data, &varEntitlementReconciliationLabelConditionFull)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationLabelConditionFull(varEntitlementReconciliationLabelConditionFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "refType")
		delete(additionalProperties, "labelId")
		delete(additionalProperties, "labelName")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationLabelConditionFull struct {
	value *EntitlementReconciliationLabelConditionFull
	isSet bool
}

func (v NullableEntitlementReconciliationLabelConditionFull) Get() *EntitlementReconciliationLabelConditionFull {
	return v.value
}

func (v *NullableEntitlementReconciliationLabelConditionFull) Set(val *EntitlementReconciliationLabelConditionFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationLabelConditionFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationLabelConditionFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationLabelConditionFull(val *EntitlementReconciliationLabelConditionFull) *NullableEntitlementReconciliationLabelConditionFull {
	return &NullableEntitlementReconciliationLabelConditionFull{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationLabelConditionFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationLabelConditionFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
