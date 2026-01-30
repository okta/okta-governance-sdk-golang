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

// checks if the RequestConditionServiceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConditionServiceStatus{}

// RequestConditionServiceStatus An object describes the status of request condition service.
type RequestConditionServiceStatus struct {
	Status               RequestConditionServiceStatusEnum `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionServiceStatus RequestConditionServiceStatus

// NewRequestConditionServiceStatus instantiates a new RequestConditionServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionServiceStatus(status RequestConditionServiceStatusEnum) *RequestConditionServiceStatus {
	this := RequestConditionServiceStatus{}
	this.Status = status
	return &this
}

// NewRequestConditionServiceStatusWithDefaults instantiates a new RequestConditionServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionServiceStatusWithDefaults() *RequestConditionServiceStatus {
	this := RequestConditionServiceStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestConditionServiceStatus) GetStatus() RequestConditionServiceStatusEnum {
	if o == nil {
		var ret RequestConditionServiceStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestConditionServiceStatus) GetStatusOk() (*RequestConditionServiceStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestConditionServiceStatus) SetStatus(v RequestConditionServiceStatusEnum) {
	o.Status = v
}

func (o RequestConditionServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConditionServiceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestConditionServiceStatus) UnmarshalJSON(data []byte) (err error) {
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

	varRequestConditionServiceStatus := _RequestConditionServiceStatus{}

	err = json.Unmarshal(data, &varRequestConditionServiceStatus)

	if err != nil {
		return err
	}

	*o = RequestConditionServiceStatus(varRequestConditionServiceStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestConditionServiceStatus struct {
	value *RequestConditionServiceStatus
	isSet bool
}

func (v NullableRequestConditionServiceStatus) Get() *RequestConditionServiceStatus {
	return v.value
}

func (v *NullableRequestConditionServiceStatus) Set(val *RequestConditionServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionServiceStatus(val *RequestConditionServiceStatus) *NullableRequestConditionServiceStatus {
	return &NullableRequestConditionServiceStatus{value: val, isSet: true}
}

func (v NullableRequestConditionServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
