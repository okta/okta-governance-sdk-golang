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

// EntitlementLink Links available in entitlements list response
type EntitlementLink struct {
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementLink EntitlementLink

// NewEntitlementLink instantiates a new EntitlementLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementLink(self Link) *EntitlementLink {
	this := EntitlementLink{}
	this.Self = self
	return &this
}

// NewEntitlementLinkWithDefaults instantiates a new EntitlementLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementLinkWithDefaults() *EntitlementLink {
	this := EntitlementLink{}
	return &this
}

// GetSelf returns the Self field value
func (o *EntitlementLink) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EntitlementLink) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EntitlementLink) SetSelf(v Link) {
	o.Self = v
}

func (o EntitlementLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementLink) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementLink := _EntitlementLink{}

	err = json.Unmarshal(bytes, &varEntitlementLink)
	if err == nil {
		*o = EntitlementLink(varEntitlementLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementLink struct {
	value *EntitlementLink
	isSet bool
}

func (v NullableEntitlementLink) Get() *EntitlementLink {
	return v.value
}

func (v *NullableEntitlementLink) Set(val *EntitlementLink) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementLink) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementLink(val *EntitlementLink) *NullableEntitlementLink {
	return &NullableEntitlementLink{value: val, isSet: true}
}

func (v NullableEntitlementLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
