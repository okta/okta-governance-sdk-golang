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
)

// checks if the ResolveTaskV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveTaskV2Request{}

// ResolveTaskV2Request struct for ResolveTaskV2Request
type ResolveTaskV2Request struct {
	// The value of the task resolution. Required for `APPROVAL` and `QUESTION` type tasks.   * For `APPROVAL` type tasks, the only supported values are `APPROVED` or `DENIED`.   * For `QUESTION` type tasks, the value is the answer to the question, and can be any string.   * For custom `TODO` type tasks, a value string isn't required.
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResolveTaskV2Request ResolveTaskV2Request

// NewResolveTaskV2Request instantiates a new ResolveTaskV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveTaskV2Request() *ResolveTaskV2Request {
	this := ResolveTaskV2Request{}
	return &this
}

// NewResolveTaskV2RequestWithDefaults instantiates a new ResolveTaskV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveTaskV2RequestWithDefaults() *ResolveTaskV2Request {
	this := ResolveTaskV2Request{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResolveTaskV2Request) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveTaskV2Request) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResolveTaskV2Request) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResolveTaskV2Request) SetValue(v string) {
	o.Value = &v
}

func (o ResolveTaskV2Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveTaskV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResolveTaskV2Request) UnmarshalJSON(data []byte) (err error) {
	varResolveTaskV2Request := _ResolveTaskV2Request{}

	err = json.Unmarshal(data, &varResolveTaskV2Request)

	if err != nil {
		return err
	}

	*o = ResolveTaskV2Request(varResolveTaskV2Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResolveTaskV2Request struct {
	value *ResolveTaskV2Request
	isSet bool
}

func (v NullableResolveTaskV2Request) Get() *ResolveTaskV2Request {
	return v.value
}

func (v *NullableResolveTaskV2Request) Set(val *ResolveTaskV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveTaskV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveTaskV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveTaskV2Request(val *ResolveTaskV2Request) *NullableResolveTaskV2Request {
	return &NullableResolveTaskV2Request{value: val, isSet: true}
}

func (v NullableResolveTaskV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveTaskV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
