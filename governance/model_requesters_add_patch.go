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

// checks if the RequestersAddPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestersAddPatch{}

// RequestersAddPatch struct for RequestersAddPatch
type RequestersAddPatch struct {
	// Operation to add a new member to the existing requester set
	Op string `json:"op"`
	// JSON Pointer path that specifies adding to the members array
	Path                 string           `json:"path" validate:"regexp=^\\/members\\/-$"`
	Value                RequestersMember `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _RequestersAddPatch RequestersAddPatch

// NewRequestersAddPatch instantiates a new RequestersAddPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestersAddPatch(op string, path string, value RequestersMember) *RequestersAddPatch {
	this := RequestersAddPatch{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewRequestersAddPatchWithDefaults instantiates a new RequestersAddPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestersAddPatchWithDefaults() *RequestersAddPatch {
	this := RequestersAddPatch{}
	return &this
}

// GetOp returns the Op field value
func (o *RequestersAddPatch) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *RequestersAddPatch) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *RequestersAddPatch) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *RequestersAddPatch) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RequestersAddPatch) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RequestersAddPatch) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *RequestersAddPatch) GetValue() RequestersMember {
	if o == nil {
		var ret RequestersMember
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RequestersAddPatch) GetValueOk() (*RequestersMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RequestersAddPatch) SetValue(v RequestersMember) {
	o.Value = v
}

func (o RequestersAddPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestersAddPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestersAddPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"value",
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

	varRequestersAddPatch := _RequestersAddPatch{}

	err = json.Unmarshal(data, &varRequestersAddPatch)

	if err != nil {
		return err
	}

	*o = RequestersAddPatch(varRequestersAddPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestersAddPatch struct {
	value *RequestersAddPatch
	isSet bool
}

func (v NullableRequestersAddPatch) Get() *RequestersAddPatch {
	return v.value
}

func (v *NullableRequestersAddPatch) Set(val *RequestersAddPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestersAddPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestersAddPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestersAddPatch(val *RequestersAddPatch) *NullableRequestersAddPatch {
	return &NullableRequestersAddPatch{value: val, isSet: true}
}

func (v NullableRequestersAddPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestersAddPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
