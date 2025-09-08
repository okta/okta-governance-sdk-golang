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

// RequestTypeRequesterEveryone A request where the requester may be any Okta user in the Okta organization.
type RequestTypeRequesterEveryone struct {
	Type string `json:"type"`
	// A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users.
	RequesterFields      []Field `json:"requesterFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeRequesterEveryone RequestTypeRequesterEveryone

// NewRequestTypeRequesterEveryone instantiates a new RequestTypeRequesterEveryone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeRequesterEveryone(type_ string) *RequestTypeRequesterEveryone {
	this := RequestTypeRequesterEveryone{}
	this.Type = type_
	return &this
}

// NewRequestTypeRequesterEveryoneWithDefaults instantiates a new RequestTypeRequesterEveryone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeRequesterEveryoneWithDefaults() *RequestTypeRequesterEveryone {
	this := RequestTypeRequesterEveryone{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeRequesterEveryone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterEveryone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeRequesterEveryone) SetType(v string) {
	o.Type = v
}

// GetRequesterFields returns the RequesterFields field value if set, zero value otherwise.
func (o *RequestTypeRequesterEveryone) GetRequesterFields() []Field {
	if o == nil || o.RequesterFields == nil {
		var ret []Field
		return ret
	}
	return o.RequesterFields
}

// GetRequesterFieldsOk returns a tuple with the RequesterFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterEveryone) GetRequesterFieldsOk() ([]Field, bool) {
	if o == nil || o.RequesterFields == nil {
		return nil, false
	}
	return o.RequesterFields, true
}

// HasRequesterFields returns a boolean if a field has been set.
func (o *RequestTypeRequesterEveryone) HasRequesterFields() bool {
	if o != nil && o.RequesterFields != nil {
		return true
	}

	return false
}

// SetRequesterFields gets a reference to the given []Field and assigns it to the RequesterFields field.
func (o *RequestTypeRequesterEveryone) SetRequesterFields(v []Field) {
	o.RequesterFields = v
}

func (o RequestTypeRequesterEveryone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.RequesterFields != nil {
		toSerialize["requesterFields"] = o.RequesterFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeRequesterEveryone) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeRequesterEveryone := _RequestTypeRequesterEveryone{}

	err = json.Unmarshal(bytes, &varRequestTypeRequesterEveryone)
	if err == nil {
		*o = RequestTypeRequesterEveryone(varRequestTypeRequesterEveryone)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "requesterFields")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeRequesterEveryone struct {
	value *RequestTypeRequesterEveryone
	isSet bool
}

func (v NullableRequestTypeRequesterEveryone) Get() *RequestTypeRequesterEveryone {
	return v.value
}

func (v *NullableRequestTypeRequesterEveryone) Set(val *RequestTypeRequesterEveryone) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeRequesterEveryone) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeRequesterEveryone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeRequesterEveryone(val *RequestTypeRequesterEveryone) *NullableRequestTypeRequesterEveryone {
	return &NullableRequestTypeRequesterEveryone{value: val, isSet: true}
}

func (v NullableRequestTypeRequesterEveryone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeRequesterEveryone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
