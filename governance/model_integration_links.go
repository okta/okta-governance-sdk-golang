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
)

// checks if the IntegrationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationLinks{}

// IntegrationLinks Links to related resources for an integration
type IntegrationLinks struct {
	AuthUrl              *Link `json:"authUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationLinks IntegrationLinks

// NewIntegrationLinks instantiates a new IntegrationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationLinks() *IntegrationLinks {
	this := IntegrationLinks{}
	return &this
}

// NewIntegrationLinksWithDefaults instantiates a new IntegrationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationLinksWithDefaults() *IntegrationLinks {
	this := IntegrationLinks{}
	return &this
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *IntegrationLinks) GetAuthUrl() Link {
	if o == nil || IsNil(o.AuthUrl) {
		var ret Link
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinks) GetAuthUrlOk() (*Link, bool) {
	if o == nil || IsNil(o.AuthUrl) {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *IntegrationLinks) HasAuthUrl() bool {
	if o != nil && !IsNil(o.AuthUrl) {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given Link and assigns it to the AuthUrl field.
func (o *IntegrationLinks) SetAuthUrl(v Link) {
	o.AuthUrl = &v
}

func (o IntegrationLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthUrl) {
		toSerialize["authUrl"] = o.AuthUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationLinks) UnmarshalJSON(data []byte) (err error) {
	varIntegrationLinks := _IntegrationLinks{}

	err = json.Unmarshal(data, &varIntegrationLinks)

	if err != nil {
		return err
	}

	*o = IntegrationLinks(varIntegrationLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationLinks struct {
	value *IntegrationLinks
	isSet bool
}

func (v NullableIntegrationLinks) Get() *IntegrationLinks {
	return v.value
}

func (v *NullableIntegrationLinks) Set(val *IntegrationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationLinks(val *IntegrationLinks) *NullableIntegrationLinks {
	return &NullableIntegrationLinks{value: val, isSet: true}
}

func (v NullableIntegrationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
