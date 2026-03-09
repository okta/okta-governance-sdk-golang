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

// checks if the RevokePrincipalAccessResourceLinksLinksPrincipalAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessResourceLinksLinksPrincipalAccess{}

// RevokePrincipalAccessResourceLinksLinksPrincipalAccess Principal access resources
type RevokePrincipalAccessResourceLinksLinksPrincipalAccess struct {
	// Link to principal access resources
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevokePrincipalAccessResourceLinksLinksPrincipalAccess RevokePrincipalAccessResourceLinksLinksPrincipalAccess

// NewRevokePrincipalAccessResourceLinksLinksPrincipalAccess instantiates a new RevokePrincipalAccessResourceLinksLinksPrincipalAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessResourceLinksLinksPrincipalAccess() *RevokePrincipalAccessResourceLinksLinksPrincipalAccess {
	this := RevokePrincipalAccessResourceLinksLinksPrincipalAccess{}
	return &this
}

// NewRevokePrincipalAccessResourceLinksLinksPrincipalAccessWithDefaults instantiates a new RevokePrincipalAccessResourceLinksLinksPrincipalAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessResourceLinksLinksPrincipalAccessWithDefaults() *RevokePrincipalAccessResourceLinksLinksPrincipalAccess {
	this := RevokePrincipalAccessResourceLinksLinksPrincipalAccess{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) SetHref(v string) {
	o.Href = &v
}

func (o RevokePrincipalAccessResourceLinksLinksPrincipalAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessResourceLinksLinksPrincipalAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) UnmarshalJSON(data []byte) (err error) {
	varRevokePrincipalAccessResourceLinksLinksPrincipalAccess := _RevokePrincipalAccessResourceLinksLinksPrincipalAccess{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessResourceLinksLinksPrincipalAccess)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessResourceLinksLinksPrincipalAccess(varRevokePrincipalAccessResourceLinksLinksPrincipalAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess struct {
	value *RevokePrincipalAccessResourceLinksLinksPrincipalAccess
	isSet bool
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) Get() *RevokePrincipalAccessResourceLinksLinksPrincipalAccess {
	return v.value
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) Set(val *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess(val *RevokePrincipalAccessResourceLinksLinksPrincipalAccess) *NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess {
	return &NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
