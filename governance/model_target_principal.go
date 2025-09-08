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

// TargetPrincipal A representation of a principal
type TargetPrincipal struct {
	// The Okta user `id`
	ExternalId           string        `json:"externalId"`
	Type                 PrincipalType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TargetPrincipal TargetPrincipal

// NewTargetPrincipal instantiates a new TargetPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetPrincipal(externalId string, type_ PrincipalType) *TargetPrincipal {
	this := TargetPrincipal{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewTargetPrincipalWithDefaults instantiates a new TargetPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetPrincipalWithDefaults() *TargetPrincipal {
	this := TargetPrincipal{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *TargetPrincipal) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *TargetPrincipal) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *TargetPrincipal) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *TargetPrincipal) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetPrincipal) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetPrincipal) SetType(v PrincipalType) {
	o.Type = v
}

func (o TargetPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["externalId"] = o.ExternalId
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TargetPrincipal) UnmarshalJSON(bytes []byte) (err error) {
	varTargetPrincipal := _TargetPrincipal{}

	err = json.Unmarshal(bytes, &varTargetPrincipal)
	if err == nil {
		*o = TargetPrincipal(varTargetPrincipal)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTargetPrincipal struct {
	value *TargetPrincipal
	isSet bool
}

func (v NullableTargetPrincipal) Get() *TargetPrincipal {
	return v.value
}

func (v *NullableTargetPrincipal) Set(val *TargetPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetPrincipal(val *TargetPrincipal) *NullableTargetPrincipal {
	return &NullableTargetPrincipal{value: val, isSet: true}
}

func (v NullableTargetPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
