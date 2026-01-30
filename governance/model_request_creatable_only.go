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

// checks if the RequestCreatableOnly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestCreatableOnly{}

// RequestCreatableOnly struct for RequestCreatableOnly
type RequestCreatableOnly struct {
	// Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.
	RequesterFieldValues []FieldValueWritable `json:"requesterFieldValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestCreatableOnly RequestCreatableOnly

// NewRequestCreatableOnly instantiates a new RequestCreatableOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreatableOnly() *RequestCreatableOnly {
	this := RequestCreatableOnly{}
	return &this
}

// NewRequestCreatableOnlyWithDefaults instantiates a new RequestCreatableOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreatableOnlyWithDefaults() *RequestCreatableOnly {
	this := RequestCreatableOnly{}
	return &this
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *RequestCreatableOnly) GetRequesterFieldValues() []FieldValueWritable {
	if o == nil || IsNil(o.RequesterFieldValues) {
		var ret []FieldValueWritable
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreatableOnly) GetRequesterFieldValuesOk() ([]FieldValueWritable, bool) {
	if o == nil || IsNil(o.RequesterFieldValues) {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestCreatableOnly) HasRequesterFieldValues() bool {
	if o != nil && !IsNil(o.RequesterFieldValues) {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []FieldValueWritable and assigns it to the RequesterFieldValues field.
func (o *RequestCreatableOnly) SetRequesterFieldValues(v []FieldValueWritable) {
	o.RequesterFieldValues = v
}

func (o RequestCreatableOnly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestCreatableOnly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequesterFieldValues) {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestCreatableOnly) UnmarshalJSON(data []byte) (err error) {
	varRequestCreatableOnly := _RequestCreatableOnly{}

	err = json.Unmarshal(data, &varRequestCreatableOnly)

	if err != nil {
		return err
	}

	*o = RequestCreatableOnly(varRequestCreatableOnly)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requesterFieldValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestCreatableOnly struct {
	value *RequestCreatableOnly
	isSet bool
}

func (v NullableRequestCreatableOnly) Get() *RequestCreatableOnly {
	return v.value
}

func (v *NullableRequestCreatableOnly) Set(val *RequestCreatableOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreatableOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreatableOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreatableOnly(val *RequestCreatableOnly) *NullableRequestCreatableOnly {
	return &NullableRequestCreatableOnly{value: val, isSet: true}
}

func (v NullableRequestCreatableOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreatableOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
