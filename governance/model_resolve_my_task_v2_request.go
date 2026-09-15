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

// checks if the ResolveMyTaskV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveMyTaskV2Request{}

// ResolveMyTaskV2Request struct for ResolveMyTaskV2Request
type ResolveMyTaskV2Request struct {
	// Specify the task resolution value. Required for `APPROVAL` and `QUESTION` type tasks.   * For `APPROVAL` type tasks, the only supported values are `APPROVED` or `DENIED`.   * For `QUESTION` type tasks, the value is the answer to the question, and can be any string.   * For custom `TODO` type tasks, you don't need to specify a resolution value.
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResolveMyTaskV2Request ResolveMyTaskV2Request

// NewResolveMyTaskV2Request instantiates a new ResolveMyTaskV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveMyTaskV2Request() *ResolveMyTaskV2Request {
	this := ResolveMyTaskV2Request{}
	return &this
}

// NewResolveMyTaskV2RequestWithDefaults instantiates a new ResolveMyTaskV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveMyTaskV2RequestWithDefaults() *ResolveMyTaskV2Request {
	this := ResolveMyTaskV2Request{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResolveMyTaskV2Request) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveMyTaskV2Request) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResolveMyTaskV2Request) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResolveMyTaskV2Request) SetValue(v string) {
	o.Value = &v
}

func (o ResolveMyTaskV2Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveMyTaskV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResolveMyTaskV2Request) UnmarshalJSON(data []byte) (err error) {
	varResolveMyTaskV2Request := _ResolveMyTaskV2Request{}

	err = json.Unmarshal(data, &varResolveMyTaskV2Request)

	if err != nil {
		return err
	}

	*o = ResolveMyTaskV2Request(varResolveMyTaskV2Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResolveMyTaskV2Request struct {
	value *ResolveMyTaskV2Request
	isSet bool
}

func (v NullableResolveMyTaskV2Request) Get() *ResolveMyTaskV2Request {
	return v.value
}

func (v *NullableResolveMyTaskV2Request) Set(val *ResolveMyTaskV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveMyTaskV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveMyTaskV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveMyTaskV2Request(val *ResolveMyTaskV2Request) *NullableResolveMyTaskV2Request {
	return &NullableResolveMyTaskV2Request{value: val, isSet: true}
}

func (v NullableResolveMyTaskV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveMyTaskV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
