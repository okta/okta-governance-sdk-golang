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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RequestUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestUpdate{}

// RequestUpdate struct for RequestUpdate
type RequestUpdate struct {
	Status RequestStatusPatchable `json:"status"`
}

type _RequestUpdate RequestUpdate

// NewRequestUpdate instantiates a new RequestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUpdate(status RequestStatusPatchable) *RequestUpdate {
	this := RequestUpdate{}
	this.Status = status
	return &this
}

// NewRequestUpdateWithDefaults instantiates a new RequestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUpdateWithDefaults() *RequestUpdate {
	this := RequestUpdate{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestUpdate) GetStatus() RequestStatusPatchable {
	if o == nil {
		var ret RequestStatusPatchable
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestUpdate) GetStatusOk() (*RequestStatusPatchable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestUpdate) SetStatus(v RequestStatusPatchable) {
	o.Status = v
}

func (o RequestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *RequestUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varRequestUpdate := _RequestUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestUpdate)

	if err != nil {
		return err
	}

	*o = RequestUpdate(varRequestUpdate)

	return err
}

type NullableRequestUpdate struct {
	value *RequestUpdate
	isSet bool
}

func (v NullableRequestUpdate) Get() *RequestUpdate {
	return v.value
}

func (v *NullableRequestUpdate) Set(val *RequestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUpdate(val *RequestUpdate) *NullableRequestUpdate {
	return &NullableRequestUpdate{value: val, isSet: true}
}

func (v NullableRequestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
