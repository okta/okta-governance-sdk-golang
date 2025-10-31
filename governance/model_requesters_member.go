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

// checks if the RequestersMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestersMember{}

// RequestersMember struct for RequestersMember
type RequestersMember struct {
	Type                 string `json:"type"`
	ExternalId           string `json:"externalId"`
	AdditionalProperties map[string]interface{}
}

type _RequestersMember RequestersMember

// NewRequestersMember instantiates a new RequestersMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestersMember(type_ string, externalId string) *RequestersMember {
	this := RequestersMember{}
	this.Type = type_
	this.ExternalId = externalId
	return &this
}

// NewRequestersMemberWithDefaults instantiates a new RequestersMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestersMemberWithDefaults() *RequestersMember {
	this := RequestersMember{}
	return &this
}

// GetType returns the Type field value
func (o *RequestersMember) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestersMember) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestersMember) SetType(v string) {
	o.Type = v
}

// GetExternalId returns the ExternalId field value
func (o *RequestersMember) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *RequestersMember) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *RequestersMember) SetExternalId(v string) {
	o.ExternalId = v
}

func (o RequestersMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestersMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["externalId"] = o.ExternalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestersMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"externalId",
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

	varRequestersMember := _RequestersMember{}

	err = json.Unmarshal(data, &varRequestersMember)

	if err != nil {
		return err
	}

	*o = RequestersMember(varRequestersMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "externalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestersMember struct {
	value *RequestersMember
	isSet bool
}

func (v NullableRequestersMember) Get() *RequestersMember {
	return v.value
}

func (v *NullableRequestersMember) Set(val *RequestersMember) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestersMember) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestersMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestersMember(val *RequestersMember) *NullableRequestersMember {
	return &NullableRequestersMember{value: val, isSet: true}
}

func (v NullableRequestersMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestersMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
