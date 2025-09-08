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

// RequestTypeRequesterMemberOf A request where the requester that must be a member of a particular Okta Group.
type RequestTypeRequesterMemberOf struct {
	Type string `json:"type"`
	// Okta groups the user persona must be a member of
	RequesterMemberOf []string `json:"requesterMemberOf"`
	// A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users.
	RequesterFields      []Field `json:"requesterFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeRequesterMemberOf RequestTypeRequesterMemberOf

// NewRequestTypeRequesterMemberOf instantiates a new RequestTypeRequesterMemberOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeRequesterMemberOf(type_ string, requesterMemberOf []string) *RequestTypeRequesterMemberOf {
	this := RequestTypeRequesterMemberOf{}
	this.Type = type_
	this.RequesterMemberOf = requesterMemberOf
	return &this
}

// NewRequestTypeRequesterMemberOfWithDefaults instantiates a new RequestTypeRequesterMemberOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeRequesterMemberOfWithDefaults() *RequestTypeRequesterMemberOf {
	this := RequestTypeRequesterMemberOf{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeRequesterMemberOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterMemberOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeRequesterMemberOf) SetType(v string) {
	o.Type = v
}

// GetRequesterMemberOf returns the RequesterMemberOf field value
func (o *RequestTypeRequesterMemberOf) GetRequesterMemberOf() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequesterMemberOf
}

// GetRequesterMemberOfOk returns a tuple with the RequesterMemberOf field value
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterMemberOf) GetRequesterMemberOfOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterMemberOf, true
}

// SetRequesterMemberOf sets field value
func (o *RequestTypeRequesterMemberOf) SetRequesterMemberOf(v []string) {
	o.RequesterMemberOf = v
}

// GetRequesterFields returns the RequesterFields field value if set, zero value otherwise.
func (o *RequestTypeRequesterMemberOf) GetRequesterFields() []Field {
	if o == nil || o.RequesterFields == nil {
		var ret []Field
		return ret
	}
	return o.RequesterFields
}

// GetRequesterFieldsOk returns a tuple with the RequesterFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterMemberOf) GetRequesterFieldsOk() ([]Field, bool) {
	if o == nil || o.RequesterFields == nil {
		return nil, false
	}
	return o.RequesterFields, true
}

// HasRequesterFields returns a boolean if a field has been set.
func (o *RequestTypeRequesterMemberOf) HasRequesterFields() bool {
	if o != nil && o.RequesterFields != nil {
		return true
	}

	return false
}

// SetRequesterFields gets a reference to the given []Field and assigns it to the RequesterFields field.
func (o *RequestTypeRequesterMemberOf) SetRequesterFields(v []Field) {
	o.RequesterFields = v
}

func (o RequestTypeRequesterMemberOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["requesterMemberOf"] = o.RequesterMemberOf
	}
	if o.RequesterFields != nil {
		toSerialize["requesterFields"] = o.RequesterFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeRequesterMemberOf) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeRequesterMemberOf := _RequestTypeRequesterMemberOf{}

	err = json.Unmarshal(bytes, &varRequestTypeRequesterMemberOf)
	if err == nil {
		*o = RequestTypeRequesterMemberOf(varRequestTypeRequesterMemberOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "requesterMemberOf")
		delete(additionalProperties, "requesterFields")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeRequesterMemberOf struct {
	value *RequestTypeRequesterMemberOf
	isSet bool
}

func (v NullableRequestTypeRequesterMemberOf) Get() *RequestTypeRequesterMemberOf {
	return v.value
}

func (v *NullableRequestTypeRequesterMemberOf) Set(val *RequestTypeRequesterMemberOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeRequesterMemberOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeRequesterMemberOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeRequesterMemberOf(val *RequestTypeRequesterMemberOf) *NullableRequestTypeRequesterMemberOf {
	return &NullableRequestTypeRequesterMemberOf{value: val, isSet: true}
}

func (v NullableRequestTypeRequesterMemberOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeRequesterMemberOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
