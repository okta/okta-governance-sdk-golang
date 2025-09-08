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

// EntitlementBundlesArrayFullInner Entitlement bundle id and entitlement bundle name that is made requestable
type EntitlementBundlesArrayFullInner struct {
	// The entitlement bundle `id`
	Id                   string `json:"id" validate:"regexp=enb[0-9a-zA-Z]+"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBundlesArrayFullInner EntitlementBundlesArrayFullInner

// NewEntitlementBundlesArrayFullInner instantiates a new EntitlementBundlesArrayFullInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundlesArrayFullInner(id string) *EntitlementBundlesArrayFullInner {
	this := EntitlementBundlesArrayFullInner{}
	this.Id = id
	return &this
}

// NewEntitlementBundlesArrayFullInnerWithDefaults instantiates a new EntitlementBundlesArrayFullInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundlesArrayFullInnerWithDefaults() *EntitlementBundlesArrayFullInner {
	this := EntitlementBundlesArrayFullInner{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementBundlesArrayFullInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundlesArrayFullInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementBundlesArrayFullInner) SetId(v string) {
	o.Id = v
}

func (o EntitlementBundlesArrayFullInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementBundlesArrayFullInner) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBundlesArrayFullInner := _EntitlementBundlesArrayFullInner{}

	err = json.Unmarshal(bytes, &varEntitlementBundlesArrayFullInner)
	if err == nil {
		*o = EntitlementBundlesArrayFullInner(varEntitlementBundlesArrayFullInner)
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

type NullableEntitlementBundlesArrayFullInner struct {
	value *EntitlementBundlesArrayFullInner
	isSet bool
}

func (v NullableEntitlementBundlesArrayFullInner) Get() *EntitlementBundlesArrayFullInner {
	return v.value
}

func (v *NullableEntitlementBundlesArrayFullInner) Set(val *EntitlementBundlesArrayFullInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundlesArrayFullInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundlesArrayFullInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundlesArrayFullInner(val *EntitlementBundlesArrayFullInner) *NullableEntitlementBundlesArrayFullInner {
	return &NullableEntitlementBundlesArrayFullInner{value: val, isSet: true}
}

func (v NullableEntitlementBundlesArrayFullInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundlesArrayFullInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
