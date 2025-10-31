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

// checks if the TargetPrincipalFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetPrincipalFull{}

// TargetPrincipalFull Representation of a principal
type TargetPrincipalFull struct {
	// The Okta user `id`
	ExternalId           string        `json:"externalId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Type                 PrincipalType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TargetPrincipalFull TargetPrincipalFull

// NewTargetPrincipalFull instantiates a new TargetPrincipalFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetPrincipalFull(externalId string, type_ PrincipalType) *TargetPrincipalFull {
	this := TargetPrincipalFull{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewTargetPrincipalFullWithDefaults instantiates a new TargetPrincipalFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetPrincipalFullWithDefaults() *TargetPrincipalFull {
	this := TargetPrincipalFull{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *TargetPrincipalFull) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *TargetPrincipalFull) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *TargetPrincipalFull) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *TargetPrincipalFull) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetPrincipalFull) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetPrincipalFull) SetType(v PrincipalType) {
	o.Type = v
}

func (o TargetPrincipalFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetPrincipalFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TargetPrincipalFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
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

	varTargetPrincipalFull := _TargetPrincipalFull{}

	err = json.Unmarshal(data, &varTargetPrincipalFull)

	if err != nil {
		return err
	}

	*o = TargetPrincipalFull(varTargetPrincipalFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTargetPrincipalFull struct {
	value *TargetPrincipalFull
	isSet bool
}

func (v NullableTargetPrincipalFull) Get() *TargetPrincipalFull {
	return v.value
}

func (v *NullableTargetPrincipalFull) Set(val *TargetPrincipalFull) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetPrincipalFull) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetPrincipalFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetPrincipalFull(val *TargetPrincipalFull) *NullableTargetPrincipalFull {
	return &NullableTargetPrincipalFull{value: val, isSet: true}
}

func (v NullableTargetPrincipalFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetPrincipalFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
