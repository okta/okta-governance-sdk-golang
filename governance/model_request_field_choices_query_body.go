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

// checks if the RequestFieldChoicesQueryBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestFieldChoicesQueryBody{}

// RequestFieldChoicesQueryBody Request body for field choices queries, providing the current values of other fields to filter the available choices.
type RequestFieldChoicesQueryBody struct {
	// Previously selected field values, keyed by field ID.
	FieldValues map[string]RequestFieldValuesMapValue `json:"fieldValues,omitempty"`
	// SCIM filter to narrow the returned choices. Supported attribute: `label`. Supported operators: `co` (contains), `sw` (starts with).
	Filter               *string `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestFieldChoicesQueryBody RequestFieldChoicesQueryBody

// NewRequestFieldChoicesQueryBody instantiates a new RequestFieldChoicesQueryBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFieldChoicesQueryBody() *RequestFieldChoicesQueryBody {
	this := RequestFieldChoicesQueryBody{}
	return &this
}

// NewRequestFieldChoicesQueryBodyWithDefaults instantiates a new RequestFieldChoicesQueryBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFieldChoicesQueryBodyWithDefaults() *RequestFieldChoicesQueryBody {
	this := RequestFieldChoicesQueryBody{}
	return &this
}

// GetFieldValues returns the FieldValues field value if set, zero value otherwise.
func (o *RequestFieldChoicesQueryBody) GetFieldValues() map[string]RequestFieldValuesMapValue {
	if o == nil || IsNil(o.FieldValues) {
		var ret map[string]RequestFieldValuesMapValue
		return ret
	}
	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldChoicesQueryBody) GetFieldValuesOk() (map[string]RequestFieldValuesMapValue, bool) {
	if o == nil || IsNil(o.FieldValues) {
		return map[string]RequestFieldValuesMapValue{}, false
	}
	return o.FieldValues, true
}

// HasFieldValues returns a boolean if a field has been set.
func (o *RequestFieldChoicesQueryBody) HasFieldValues() bool {
	if o != nil && !IsNil(o.FieldValues) {
		return true
	}

	return false
}

// SetFieldValues gets a reference to the given map[string]RequestFieldValuesMapValue and assigns it to the FieldValues field.
func (o *RequestFieldChoicesQueryBody) SetFieldValues(v map[string]RequestFieldValuesMapValue) {
	o.FieldValues = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RequestFieldChoicesQueryBody) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldChoicesQueryBody) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RequestFieldChoicesQueryBody) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *RequestFieldChoicesQueryBody) SetFilter(v string) {
	o.Filter = &v
}

func (o RequestFieldChoicesQueryBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestFieldChoicesQueryBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldValues) {
		toSerialize["fieldValues"] = o.FieldValues
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestFieldChoicesQueryBody) UnmarshalJSON(data []byte) (err error) {
	varRequestFieldChoicesQueryBody := _RequestFieldChoicesQueryBody{}

	err = json.Unmarshal(data, &varRequestFieldChoicesQueryBody)

	if err != nil {
		return err
	}

	*o = RequestFieldChoicesQueryBody(varRequestFieldChoicesQueryBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fieldValues")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestFieldChoicesQueryBody struct {
	value *RequestFieldChoicesQueryBody
	isSet bool
}

func (v NullableRequestFieldChoicesQueryBody) Get() *RequestFieldChoicesQueryBody {
	return v.value
}

func (v *NullableRequestFieldChoicesQueryBody) Set(val *RequestFieldChoicesQueryBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFieldChoicesQueryBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFieldChoicesQueryBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFieldChoicesQueryBody(val *RequestFieldChoicesQueryBody) *NullableRequestFieldChoicesQueryBody {
	return &NullableRequestFieldChoicesQueryBody{value: val, isSet: true}
}

func (v NullableRequestFieldChoicesQueryBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFieldChoicesQueryBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
