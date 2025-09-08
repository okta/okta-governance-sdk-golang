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

// RequestResourceCatalogEntryCreatable A representation of a requestable resource
type RequestResourceCatalogEntryCreatable struct {
	Type string `json:"type"`
	// The ID of the resource catalog entry
	EntryId              string `json:"entryId"`
	AdditionalProperties map[string]interface{}
}

type _RequestResourceCatalogEntryCreatable RequestResourceCatalogEntryCreatable

// NewRequestResourceCatalogEntryCreatable instantiates a new RequestResourceCatalogEntryCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResourceCatalogEntryCreatable(type_ string, entryId string) *RequestResourceCatalogEntryCreatable {
	this := RequestResourceCatalogEntryCreatable{}
	this.Type = type_
	this.EntryId = entryId
	return &this
}

// NewRequestResourceCatalogEntryCreatableWithDefaults instantiates a new RequestResourceCatalogEntryCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResourceCatalogEntryCreatableWithDefaults() *RequestResourceCatalogEntryCreatable {
	this := RequestResourceCatalogEntryCreatable{}
	return &this
}

// GetType returns the Type field value
func (o *RequestResourceCatalogEntryCreatable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestResourceCatalogEntryCreatable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestResourceCatalogEntryCreatable) SetType(v string) {
	o.Type = v
}

// GetEntryId returns the EntryId field value
func (o *RequestResourceCatalogEntryCreatable) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *RequestResourceCatalogEntryCreatable) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *RequestResourceCatalogEntryCreatable) SetEntryId(v string) {
	o.EntryId = v
}

func (o RequestResourceCatalogEntryCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestResourceCatalogEntryCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestResourceCatalogEntryCreatable := _RequestResourceCatalogEntryCreatable{}

	err = json.Unmarshal(bytes, &varRequestResourceCatalogEntryCreatable)
	if err == nil {
		*o = RequestResourceCatalogEntryCreatable(varRequestResourceCatalogEntryCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "entryId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestResourceCatalogEntryCreatable struct {
	value *RequestResourceCatalogEntryCreatable
	isSet bool
}

func (v NullableRequestResourceCatalogEntryCreatable) Get() *RequestResourceCatalogEntryCreatable {
	return v.value
}

func (v *NullableRequestResourceCatalogEntryCreatable) Set(val *RequestResourceCatalogEntryCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResourceCatalogEntryCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResourceCatalogEntryCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResourceCatalogEntryCreatable(val *RequestResourceCatalogEntryCreatable) *NullableRequestResourceCatalogEntryCreatable {
	return &NullableRequestResourceCatalogEntryCreatable{value: val, isSet: true}
}

func (v NullableRequestResourceCatalogEntryCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResourceCatalogEntryCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
