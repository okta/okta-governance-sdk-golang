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

// checks if the ClientCredentialPrincipal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientCredentialPrincipal{}

// ClientCredentialPrincipal A representation of a principal
type ClientCredentialPrincipal struct {
	// The principal ID
	ExternalId string `json:"externalId" validate:"regexp=(00u|pca)[0-9a-zA-Z]+"`
	// The type of principal
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ClientCredentialPrincipal ClientCredentialPrincipal

// NewClientCredentialPrincipal instantiates a new ClientCredentialPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCredentialPrincipal(externalId string, type_ string) *ClientCredentialPrincipal {
	this := ClientCredentialPrincipal{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewClientCredentialPrincipalWithDefaults instantiates a new ClientCredentialPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientCredentialPrincipalWithDefaults() *ClientCredentialPrincipal {
	this := ClientCredentialPrincipal{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *ClientCredentialPrincipal) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ClientCredentialPrincipal) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ClientCredentialPrincipal) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *ClientCredentialPrincipal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClientCredentialPrincipal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClientCredentialPrincipal) SetType(v string) {
	o.Type = v
}

func (o ClientCredentialPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCredentialPrincipal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientCredentialPrincipal) UnmarshalJSON(data []byte) (err error) {
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

	varClientCredentialPrincipal := _ClientCredentialPrincipal{}

	err = json.Unmarshal(data, &varClientCredentialPrincipal)

	if err != nil {
		return err
	}

	*o = ClientCredentialPrincipal(varClientCredentialPrincipal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientCredentialPrincipal struct {
	value *ClientCredentialPrincipal
	isSet bool
}

func (v NullableClientCredentialPrincipal) Get() *ClientCredentialPrincipal {
	return v.value
}

func (v *NullableClientCredentialPrincipal) Set(val *ClientCredentialPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCredentialPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCredentialPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCredentialPrincipal(val *ClientCredentialPrincipal) *NullableClientCredentialPrincipal {
	return &NullableClientCredentialPrincipal{value: val, isSet: true}
}

func (v NullableClientCredentialPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCredentialPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
