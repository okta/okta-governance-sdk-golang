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

// checks if the RequestConditionsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConditionsList{}

// RequestConditionsList struct for RequestConditionsList
type RequestConditionsList struct {
	// All request conditions
	Data                 []RequestConditionSparse `json:"data"`
	Links                RequestConditionsLinks   `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionsList RequestConditionsList

// NewRequestConditionsList instantiates a new RequestConditionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionsList(data []RequestConditionSparse, links RequestConditionsLinks) *RequestConditionsList {
	this := RequestConditionsList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewRequestConditionsListWithDefaults instantiates a new RequestConditionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionsListWithDefaults() *RequestConditionsList {
	this := RequestConditionsList{}
	return &this
}

// GetData returns the Data field value
func (o *RequestConditionsList) GetData() []RequestConditionSparse {
	if o == nil {
		var ret []RequestConditionSparse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RequestConditionsList) GetDataOk() ([]RequestConditionSparse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *RequestConditionsList) SetData(v []RequestConditionSparse) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *RequestConditionsList) GetLinks() RequestConditionsLinks {
	if o == nil {
		var ret RequestConditionsLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestConditionsList) GetLinksOk() (*RequestConditionsLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestConditionsList) SetLinks(v RequestConditionsLinks) {
	o.Links = v
}

func (o RequestConditionsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConditionsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestConditionsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
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

	varRequestConditionsList := _RequestConditionsList{}

	err = json.Unmarshal(data, &varRequestConditionsList)

	if err != nil {
		return err
	}

	*o = RequestConditionsList(varRequestConditionsList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestConditionsList struct {
	value *RequestConditionsList
	isSet bool
}

func (v NullableRequestConditionsList) Get() *RequestConditionsList {
	return v.value
}

func (v *NullableRequestConditionsList) Set(val *RequestConditionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionsList(val *RequestConditionsList) *NullableRequestConditionsList {
	return &NullableRequestConditionsList{value: val, isSet: true}
}

func (v NullableRequestConditionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
