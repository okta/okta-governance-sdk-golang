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

// checks if the EntitlementReconciliationConfigWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationConfigWritable{}

// EntitlementReconciliationConfigWritable The fields that you can set on a reconciliation configuration.  `additive` covers entitlements added directly in the app, the drifts reported with `driftType` of `ADD`. `subtractive` covers entitlements removed there, reported as `SUB`.  An upsert is a pure replace. Whatever you send becomes the stored configuration, and anything you leave out is dropped.  `additive` and `subtractive` are both required when `mode` is `ENABLED`. When `mode` is `DISABLED` they're independently optional: send both to keep a configuration on file for when reconciliation is turned back on, send one to keep only that direction, or leave both out to wipe it.
type EntitlementReconciliationConfigWritable struct {
	Mode                 ReconciliationMode                          `json:"mode"`
	Additive             *EntitlementReconciliationDirectionWritable `json:"additive,omitempty"`
	Subtractive          *EntitlementReconciliationDirectionWritable `json:"subtractive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationConfigWritable EntitlementReconciliationConfigWritable

// NewEntitlementReconciliationConfigWritable instantiates a new EntitlementReconciliationConfigWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationConfigWritable(mode ReconciliationMode) *EntitlementReconciliationConfigWritable {
	this := EntitlementReconciliationConfigWritable{}
	this.Mode = mode
	return &this
}

// NewEntitlementReconciliationConfigWritableWithDefaults instantiates a new EntitlementReconciliationConfigWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationConfigWritableWithDefaults() *EntitlementReconciliationConfigWritable {
	this := EntitlementReconciliationConfigWritable{}
	return &this
}

// GetMode returns the Mode field value
func (o *EntitlementReconciliationConfigWritable) GetMode() ReconciliationMode {
	if o == nil {
		var ret ReconciliationMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigWritable) GetModeOk() (*ReconciliationMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *EntitlementReconciliationConfigWritable) SetMode(v ReconciliationMode) {
	o.Mode = v
}

// GetAdditive returns the Additive field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigWritable) GetAdditive() EntitlementReconciliationDirectionWritable {
	if o == nil || IsNil(o.Additive) {
		var ret EntitlementReconciliationDirectionWritable
		return ret
	}
	return *o.Additive
}

// GetAdditiveOk returns a tuple with the Additive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigWritable) GetAdditiveOk() (*EntitlementReconciliationDirectionWritable, bool) {
	if o == nil || IsNil(o.Additive) {
		return nil, false
	}
	return o.Additive, true
}

// HasAdditive returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigWritable) HasAdditive() bool {
	if o != nil && !IsNil(o.Additive) {
		return true
	}

	return false
}

// SetAdditive gets a reference to the given EntitlementReconciliationDirectionWritable and assigns it to the Additive field.
func (o *EntitlementReconciliationConfigWritable) SetAdditive(v EntitlementReconciliationDirectionWritable) {
	o.Additive = &v
}

// GetSubtractive returns the Subtractive field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigWritable) GetSubtractive() EntitlementReconciliationDirectionWritable {
	if o == nil || IsNil(o.Subtractive) {
		var ret EntitlementReconciliationDirectionWritable
		return ret
	}
	return *o.Subtractive
}

// GetSubtractiveOk returns a tuple with the Subtractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigWritable) GetSubtractiveOk() (*EntitlementReconciliationDirectionWritable, bool) {
	if o == nil || IsNil(o.Subtractive) {
		return nil, false
	}
	return o.Subtractive, true
}

// HasSubtractive returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigWritable) HasSubtractive() bool {
	if o != nil && !IsNil(o.Subtractive) {
		return true
	}

	return false
}

// SetSubtractive gets a reference to the given EntitlementReconciliationDirectionWritable and assigns it to the Subtractive field.
func (o *EntitlementReconciliationConfigWritable) SetSubtractive(v EntitlementReconciliationDirectionWritable) {
	o.Subtractive = &v
}

func (o EntitlementReconciliationConfigWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationConfigWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	if !IsNil(o.Additive) {
		toSerialize["additive"] = o.Additive
	}
	if !IsNil(o.Subtractive) {
		toSerialize["subtractive"] = o.Subtractive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationConfigWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varEntitlementReconciliationConfigWritable := _EntitlementReconciliationConfigWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationConfigWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationConfigWritable(varEntitlementReconciliationConfigWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "additive")
		delete(additionalProperties, "subtractive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationConfigWritable struct {
	value *EntitlementReconciliationConfigWritable
	isSet bool
}

func (v NullableEntitlementReconciliationConfigWritable) Get() *EntitlementReconciliationConfigWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationConfigWritable) Set(val *EntitlementReconciliationConfigWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationConfigWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationConfigWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationConfigWritable(val *EntitlementReconciliationConfigWritable) *NullableEntitlementReconciliationConfigWritable {
	return &NullableEntitlementReconciliationConfigWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationConfigWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationConfigWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
