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

// checks if the RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements{}

// RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements Principal entitlement resources
type RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements struct {
	// Link to principal entitlements
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements

// NewRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements instantiates a new RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements() *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements {
	this := RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements{}
	return &this
}

// NewRevokePrincipalAccessResourceLinksLinksPrincipalEntitlementsWithDefaults instantiates a new RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessResourceLinksLinksPrincipalEntitlementsWithDefaults() *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements {
	this := RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) SetHref(v string) {
	o.Href = &v
}

func (o RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) UnmarshalJSON(data []byte) (err error) {
	varRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements := _RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements(varRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements struct {
	value *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements
	isSet bool
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) Get() *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements {
	return v.value
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) Set(val *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements(val *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) *NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements {
	return &NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
