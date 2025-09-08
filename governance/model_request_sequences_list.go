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

// RequestSequencesList struct for RequestSequencesList
type RequestSequencesList struct {
	// All request sequences
	Data                 []RequestSequence          `json:"data,omitempty"`
	Links                *RequestSequencesListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestSequencesList RequestSequencesList

// NewRequestSequencesList instantiates a new RequestSequencesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSequencesList() *RequestSequencesList {
	this := RequestSequencesList{}
	return &this
}

// NewRequestSequencesListWithDefaults instantiates a new RequestSequencesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSequencesListWithDefaults() *RequestSequencesList {
	this := RequestSequencesList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestSequencesList) GetData() []RequestSequence {
	if o == nil || o.Data == nil {
		var ret []RequestSequence
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSequencesList) GetDataOk() ([]RequestSequence, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestSequencesList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestSequence and assigns it to the Data field.
func (o *RequestSequencesList) SetData(v []RequestSequence) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestSequencesList) GetLinks() RequestSequencesListLinks {
	if o == nil || o.Links == nil {
		var ret RequestSequencesListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSequencesList) GetLinksOk() (*RequestSequencesListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestSequencesList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestSequencesListLinks and assigns it to the Links field.
func (o *RequestSequencesList) SetLinks(v RequestSequencesListLinks) {
	o.Links = &v
}

func (o RequestSequencesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestSequencesList) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSequencesList := _RequestSequencesList{}

	err = json.Unmarshal(bytes, &varRequestSequencesList)
	if err == nil {
		*o = RequestSequencesList(varRequestSequencesList)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestSequencesList struct {
	value *RequestSequencesList
	isSet bool
}

func (v NullableRequestSequencesList) Get() *RequestSequencesList {
	return v.value
}

func (v *NullableRequestSequencesList) Set(val *RequestSequencesList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSequencesList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSequencesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSequencesList(val *RequestSequencesList) *NullableRequestSequencesList {
	return &NullableRequestSequencesList{value: val, isSet: true}
}

func (v NullableRequestSequencesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSequencesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
