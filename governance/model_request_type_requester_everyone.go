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

// checks if the RequestTypeRequesterEveryone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeRequesterEveryone{}

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
	if o == nil || IsNil(o.RequesterFields) {
		var ret []Field
		return ret
	}
	return o.RequesterFields
}

// GetRequesterFieldsOk returns a tuple with the RequesterFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeRequesterEveryone) GetRequesterFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.RequesterFields) {
		return nil, false
	}
	return o.RequesterFields, true
}

// HasRequesterFields returns a boolean if a field has been set.
func (o *RequestTypeRequesterEveryone) HasRequesterFields() bool {
	if o != nil && !IsNil(o.RequesterFields) {
		return true
	}

	return false
}

// SetRequesterFields gets a reference to the given []Field and assigns it to the RequesterFields field.
func (o *RequestTypeRequesterEveryone) SetRequesterFields(v []Field) {
	o.RequesterFields = v
}

func (o RequestTypeRequesterEveryone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeRequesterEveryone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.RequesterFields) {
		toSerialize["requesterFields"] = o.RequesterFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeRequesterEveryone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varRequestTypeRequesterEveryone := _RequestTypeRequesterEveryone{}

	err = json.Unmarshal(data, &varRequestTypeRequesterEveryone)

	if err != nil {
		return err
	}

	*o = RequestTypeRequesterEveryone(varRequestTypeRequesterEveryone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "requesterFields")
		o.AdditionalProperties = additionalProperties
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
