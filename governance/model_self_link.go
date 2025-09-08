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

// SelfLink Links available for an object
type SelfLink struct {
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _SelfLink SelfLink

// NewSelfLink instantiates a new SelfLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfLink(self Link) *SelfLink {
	this := SelfLink{}
	this.Self = self
	return &this
}

// NewSelfLinkWithDefaults instantiates a new SelfLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfLinkWithDefaults() *SelfLink {
	this := SelfLink{}
	return &this
}

// GetSelf returns the Self field value
func (o *SelfLink) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *SelfLink) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *SelfLink) SetSelf(v Link) {
	o.Self = v
}

func (o SelfLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfLink) UnmarshalJSON(bytes []byte) (err error) {
	varSelfLink := _SelfLink{}

	err = json.Unmarshal(bytes, &varSelfLink)
	if err == nil {
		*o = SelfLink(varSelfLink)
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

type NullableSelfLink struct {
	value *SelfLink
	isSet bool
}

func (v NullableSelfLink) Get() *SelfLink {
	return v.value
}

func (v *NullableSelfLink) Set(val *SelfLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfLink(val *SelfLink) *NullableSelfLink {
	return &NullableSelfLink{value: val, isSet: true}
}

func (v NullableSelfLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
