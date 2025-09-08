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

// BaseFieldReadable struct for BaseFieldReadable
type BaseFieldReadable struct {
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _BaseFieldReadable BaseFieldReadable

// NewBaseFieldReadable instantiates a new BaseFieldReadable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFieldReadable(id string) *BaseFieldReadable {
	this := BaseFieldReadable{}
	this.Id = id
	return &this
}

// NewBaseFieldReadableWithDefaults instantiates a new BaseFieldReadable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFieldReadableWithDefaults() *BaseFieldReadable {
	this := BaseFieldReadable{}
	return &this
}

// GetId returns the Id field value
func (o *BaseFieldReadable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseFieldReadable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseFieldReadable) SetId(v string) {
	o.Id = v
}

func (o BaseFieldReadable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseFieldReadable) UnmarshalJSON(bytes []byte) (err error) {
	varBaseFieldReadable := _BaseFieldReadable{}

	err = json.Unmarshal(bytes, &varBaseFieldReadable)
	if err == nil {
		*o = BaseFieldReadable(varBaseFieldReadable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBaseFieldReadable struct {
	value *BaseFieldReadable
	isSet bool
}

func (v NullableBaseFieldReadable) Get() *BaseFieldReadable {
	return v.value
}

func (v *NullableBaseFieldReadable) Set(val *BaseFieldReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFieldReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFieldReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFieldReadable(val *BaseFieldReadable) *NullableBaseFieldReadable {
	return &NullableBaseFieldReadable{value: val, isSet: true}
}

func (v NullableBaseFieldReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFieldReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
