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

// checks if the RequestUpdate2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestUpdate2{}

// RequestUpdate2 struct for RequestUpdate2
type RequestUpdate2 struct {
	Status               *RequestStatusPatchable `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestUpdate2 RequestUpdate2

// NewRequestUpdate2 instantiates a new RequestUpdate2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUpdate2() *RequestUpdate2 {
	this := RequestUpdate2{}
	return &this
}

// NewRequestUpdate2WithDefaults instantiates a new RequestUpdate2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUpdate2WithDefaults() *RequestUpdate2 {
	this := RequestUpdate2{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestUpdate2) GetStatus() RequestStatusPatchable {
	if o == nil || IsNil(o.Status) {
		var ret RequestStatusPatchable
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpdate2) GetStatusOk() (*RequestStatusPatchable, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestUpdate2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RequestStatusPatchable and assigns it to the Status field.
func (o *RequestUpdate2) SetStatus(v RequestStatusPatchable) {
	o.Status = &v
}

func (o RequestUpdate2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestUpdate2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestUpdate2) UnmarshalJSON(data []byte) (err error) {
	varRequestUpdate2 := _RequestUpdate2{}

	err = json.Unmarshal(data, &varRequestUpdate2)

	if err != nil {
		return err
	}

	*o = RequestUpdate2(varRequestUpdate2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestUpdate2 struct {
	value *RequestUpdate2
	isSet bool
}

func (v NullableRequestUpdate2) Get() *RequestUpdate2 {
	return v.value
}

func (v *NullableRequestUpdate2) Set(val *RequestUpdate2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUpdate2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUpdate2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUpdate2(val *RequestUpdate2) *NullableRequestUpdate2 {
	return &NullableRequestUpdate2{value: val, isSet: true}
}

func (v NullableRequestUpdate2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUpdate2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
