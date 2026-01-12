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

// checks if the EntitlementsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementsInner{}

// EntitlementsInner struct for EntitlementsInner
type EntitlementsInner struct {
	// The entitlement `id`
	Id string `json:"id"`
	// Indicates whether to include all entitlement values:   * If `true`, all entitlement values are included.   * If `false`, you must specify the `values` property.
	IncludeAllValues *bool `json:"includeAllValues,omitempty"`
	// Entitlement value IDs
	Values               []EntitlementValue `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementsInner EntitlementsInner

// NewEntitlementsInner instantiates a new EntitlementsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsInner(id string) *EntitlementsInner {
	this := EntitlementsInner{}
	this.Id = id
	return &this
}

// NewEntitlementsInnerWithDefaults instantiates a new EntitlementsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsInnerWithDefaults() *EntitlementsInner {
	this := EntitlementsInner{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementsInner) SetId(v string) {
	o.Id = v
}

// GetIncludeAllValues returns the IncludeAllValues field value if set, zero value otherwise.
func (o *EntitlementsInner) GetIncludeAllValues() bool {
	if o == nil || IsNil(o.IncludeAllValues) {
		var ret bool
		return ret
	}
	return *o.IncludeAllValues
}

// GetIncludeAllValuesOk returns a tuple with the IncludeAllValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsInner) GetIncludeAllValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAllValues) {
		return nil, false
	}
	return o.IncludeAllValues, true
}

// HasIncludeAllValues returns a boolean if a field has been set.
func (o *EntitlementsInner) HasIncludeAllValues() bool {
	if o != nil && !IsNil(o.IncludeAllValues) {
		return true
	}

	return false
}

// SetIncludeAllValues gets a reference to the given bool and assigns it to the IncludeAllValues field.
func (o *EntitlementsInner) SetIncludeAllValues(v bool) {
	o.IncludeAllValues = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementsInner) GetValues() []EntitlementValue {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsInner) GetValuesOk() ([]EntitlementValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementsInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementValue and assigns it to the Values field.
func (o *EntitlementsInner) SetValues(v []EntitlementValue) {
	o.Values = v
}

func (o EntitlementsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IncludeAllValues) {
		toSerialize["includeAllValues"] = o.IncludeAllValues
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementsInner) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementsInner := _EntitlementsInner{}

	err = json.Unmarshal(data, &varEntitlementsInner)

	if err != nil {
		return err
	}

	*o = EntitlementsInner(varEntitlementsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "includeAllValues")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementsInner struct {
	value *EntitlementsInner
	isSet bool
}

func (v NullableEntitlementsInner) Get() *EntitlementsInner {
	return v.value
}

func (v *NullableEntitlementsInner) Set(val *EntitlementsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsInner(val *EntitlementsInner) *NullableEntitlementsInner {
	return &NullableEntitlementsInner{value: val, isSet: true}
}

func (v NullableEntitlementsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
