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

// checks if the EntitlementReconciliationEntitlementConditionFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationEntitlementConditionFull{}

// EntitlementReconciliationEntitlementConditionFull Matches drifts on one entitlement, optionally narrowed to specific values of it.  When `values` is absent or empty the condition matches every value of the entitlement, including values that don't exist yet. Okta stores the entitlement itself rather than the list of values it happened to have when the configuration was saved, so this keeps matching as the app's entitlement values change.  `entitlementName` and the `name` on each value are conveniences for rendering the condition. Okta resolves them on read, and they aren't part of the writable model.
type EntitlementReconciliationEntitlementConditionFull struct {
	// Identifies this condition as targeting an entitlement
	RefType string `json:"refType"`
	// The `id` property of an entitlement
	EntitlementId string `json:"entitlementId"`
	// The display name for an entitlement property
	EntitlementName *string `json:"entitlementName,omitempty"`
	// The values of the entitlement the condition is narrowed to. Absent or empty when the condition matches every value.
	Values               []EntitlementReconciliationEntitlementConditionValueFull `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationEntitlementConditionFull EntitlementReconciliationEntitlementConditionFull

// NewEntitlementReconciliationEntitlementConditionFull instantiates a new EntitlementReconciliationEntitlementConditionFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationEntitlementConditionFull(refType string, entitlementId string) *EntitlementReconciliationEntitlementConditionFull {
	this := EntitlementReconciliationEntitlementConditionFull{}
	this.RefType = refType
	this.EntitlementId = entitlementId
	return &this
}

// NewEntitlementReconciliationEntitlementConditionFullWithDefaults instantiates a new EntitlementReconciliationEntitlementConditionFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationEntitlementConditionFullWithDefaults() *EntitlementReconciliationEntitlementConditionFull {
	this := EntitlementReconciliationEntitlementConditionFull{}
	return &this
}

// GetRefType returns the RefType field value
func (o *EntitlementReconciliationEntitlementConditionFull) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementReconciliationEntitlementConditionFull) SetRefType(v string) {
	o.RefType = v
}

// GetEntitlementId returns the EntitlementId field value
func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementId, true
}

// SetEntitlementId sets field value
func (o *EntitlementReconciliationEntitlementConditionFull) SetEntitlementId(v string) {
	o.EntitlementId = v
}

// GetEntitlementName returns the EntitlementName field value if set, zero value otherwise.
func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementName() string {
	if o == nil || IsNil(o.EntitlementName) {
		var ret string
		return ret
	}
	return *o.EntitlementName
}

// GetEntitlementNameOk returns a tuple with the EntitlementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementName) {
		return nil, false
	}
	return o.EntitlementName, true
}

// HasEntitlementName returns a boolean if a field has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) HasEntitlementName() bool {
	if o != nil && !IsNil(o.EntitlementName) {
		return true
	}

	return false
}

// SetEntitlementName gets a reference to the given string and assigns it to the EntitlementName field.
func (o *EntitlementReconciliationEntitlementConditionFull) SetEntitlementName(v string) {
	o.EntitlementName = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementReconciliationEntitlementConditionFull) GetValues() []EntitlementReconciliationEntitlementConditionValueFull {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementReconciliationEntitlementConditionValueFull
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) GetValuesOk() ([]EntitlementReconciliationEntitlementConditionValueFull, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementReconciliationEntitlementConditionFull) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementReconciliationEntitlementConditionValueFull and assigns it to the Values field.
func (o *EntitlementReconciliationEntitlementConditionFull) SetValues(v []EntitlementReconciliationEntitlementConditionValueFull) {
	o.Values = v
}

func (o EntitlementReconciliationEntitlementConditionFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationEntitlementConditionFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refType"] = o.RefType
	toSerialize["entitlementId"] = o.EntitlementId
	if !IsNil(o.EntitlementName) {
		toSerialize["entitlementName"] = o.EntitlementName
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationEntitlementConditionFull) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementReconciliationEntitlementConditionFull := _EntitlementReconciliationEntitlementConditionFull{}

	err = json.Unmarshal(data, &varEntitlementReconciliationEntitlementConditionFull)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationEntitlementConditionFull(varEntitlementReconciliationEntitlementConditionFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "refType")
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "entitlementName")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationEntitlementConditionFull struct {
	value *EntitlementReconciliationEntitlementConditionFull
	isSet bool
}

func (v NullableEntitlementReconciliationEntitlementConditionFull) Get() *EntitlementReconciliationEntitlementConditionFull {
	return v.value
}

func (v *NullableEntitlementReconciliationEntitlementConditionFull) Set(val *EntitlementReconciliationEntitlementConditionFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationEntitlementConditionFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationEntitlementConditionFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationEntitlementConditionFull(val *EntitlementReconciliationEntitlementConditionFull) *NullableEntitlementReconciliationEntitlementConditionFull {
	return &NullableEntitlementReconciliationEntitlementConditionFull{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationEntitlementConditionFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationEntitlementConditionFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
