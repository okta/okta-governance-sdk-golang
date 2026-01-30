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

// checks if the RequestConditionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConditionLinks{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConditionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestConditionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varRequestConditionLinks := _RequestConditionLinks{}

	err = json.Unmarshal(data, &varRequestConditionLinks)

	if err != nil {
		return err
	}

	*o = RequestConditionLinks(varRequestConditionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
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
