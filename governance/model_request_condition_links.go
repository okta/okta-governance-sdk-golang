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

// RequestConditionLinks Links available on a single request condition representation.
type RequestConditionLinks struct {
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionLinks RequestConditionLinks

// NewRequestConditionLinks instantiates a new RequestConditionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionLinks(self Link) *RequestConditionLinks {
	this := RequestConditionLinks{}
	this.Self = self
	return &this
}

// NewRequestConditionLinksWithDefaults instantiates a new RequestConditionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionLinksWithDefaults() *RequestConditionLinks {
	this := RequestConditionLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RequestConditionLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestConditionLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestConditionLinks) SetSelf(v Link) {
	o.Self = v
}

func (o RequestConditionLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestConditionLinks) UnmarshalJSON(bytes []byte) (err error) {
	varRequestConditionLinks := _RequestConditionLinks{}

	err = json.Unmarshal(bytes, &varRequestConditionLinks)
	if err == nil {
		*o = RequestConditionLinks(varRequestConditionLinks)
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

type NullableRequestConditionLinks struct {
	value *RequestConditionLinks
	isSet bool
}

func (v NullableRequestConditionLinks) Get() *RequestConditionLinks {
	return v.value
}

func (v *NullableRequestConditionLinks) Set(val *RequestConditionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionLinks(val *RequestConditionLinks) *NullableRequestConditionLinks {
	return &NullableRequestConditionLinks{value: val, isSet: true}
}

func (v NullableRequestConditionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
