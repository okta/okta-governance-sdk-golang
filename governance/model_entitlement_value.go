/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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
)

// EntitlementValue list of entitlement value ids
type EntitlementValue struct {
	// The entitlement value id
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValue EntitlementValue

// NewEntitlementValue instantiates a new EntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValue(id string) *EntitlementValue {
	this := EntitlementValue{}
	this.Id = id
	return &this
}

// NewEntitlementValueWithDefaults instantiates a new EntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueWithDefaults() *EntitlementValue {
	this := EntitlementValue{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementValue) SetId(v string) {
	o.Id = v
}

func (o EntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValue) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValue := _EntitlementValue{}

	err = json.Unmarshal(bytes, &varEntitlementValue)
	if err == nil {
		*o = EntitlementValue(varEntitlementValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValue struct {
	value *EntitlementValue
	isSet bool
}

func (v NullableEntitlementValue) Get() *EntitlementValue {
	return v.value
}

func (v *NullableEntitlementValue) Set(val *EntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValue(val *EntitlementValue) *NullableEntitlementValue {
	return &NullableEntitlementValue{value: val, isSet: true}
}

func (v NullableEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
