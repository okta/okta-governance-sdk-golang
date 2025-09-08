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

// RequestersReplacePatch struct for RequestersReplacePatch
type RequestersReplacePatch struct {
	// Operation to replace all existing members with a new set
	Op string `json:"op"`
	// JSON Pointer path that specifies replacing the entire members array
	Path                 string             `json:"path" validate:"regexp=^\\/members$"`
	Value                []RequestersMember `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _RequestersReplacePatch RequestersReplacePatch

// NewRequestersReplacePatch instantiates a new RequestersReplacePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestersReplacePatch(op string, path string, value []RequestersMember) *RequestersReplacePatch {
	this := RequestersReplacePatch{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewRequestersReplacePatchWithDefaults instantiates a new RequestersReplacePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestersReplacePatchWithDefaults() *RequestersReplacePatch {
	this := RequestersReplacePatch{}
	return &this
}

// GetOp returns the Op field value
func (o *RequestersReplacePatch) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *RequestersReplacePatch) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *RequestersReplacePatch) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *RequestersReplacePatch) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RequestersReplacePatch) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RequestersReplacePatch) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *RequestersReplacePatch) GetValue() []RequestersMember {
	if o == nil {
		var ret []RequestersMember
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RequestersReplacePatch) GetValueOk() ([]RequestersMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *RequestersReplacePatch) SetValue(v []RequestersMember) {
	o.Value = v
}

func (o RequestersReplacePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestersReplacePatch) UnmarshalJSON(bytes []byte) (err error) {
	varRequestersReplacePatch := _RequestersReplacePatch{}

	err = json.Unmarshal(bytes, &varRequestersReplacePatch)
	if err == nil {
		*o = RequestersReplacePatch(varRequestersReplacePatch)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestersReplacePatch struct {
	value *RequestersReplacePatch
	isSet bool
}

func (v NullableRequestersReplacePatch) Get() *RequestersReplacePatch {
	return v.value
}

func (v *NullableRequestersReplacePatch) Set(val *RequestersReplacePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestersReplacePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestersReplacePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestersReplacePatch(val *RequestersReplacePatch) *NullableRequestersReplacePatch {
	return &NullableRequestersReplacePatch{value: val, isSet: true}
}

func (v NullableRequestersReplacePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestersReplacePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
