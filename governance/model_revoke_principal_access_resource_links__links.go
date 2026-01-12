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

// checks if the RevokePrincipalAccessResourceLinksLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessResourceLinksLinks{}

// RevokePrincipalAccessResourceLinksLinks Links to related resources for a revoked-access resource
type RevokePrincipalAccessResourceLinksLinks struct {
	PrincipalAccess       *RevokePrincipalAccessResourceLinksLinksPrincipalAccess       `json:"principal-access,omitempty"`
	PrincipalEntitlements *RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements `json:"principal-entitlements,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _RevokePrincipalAccessResourceLinksLinks RevokePrincipalAccessResourceLinksLinks

// NewRevokePrincipalAccessResourceLinksLinks instantiates a new RevokePrincipalAccessResourceLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessResourceLinksLinks() *RevokePrincipalAccessResourceLinksLinks {
	this := RevokePrincipalAccessResourceLinksLinks{}
	return &this
}

// NewRevokePrincipalAccessResourceLinksLinksWithDefaults instantiates a new RevokePrincipalAccessResourceLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessResourceLinksLinksWithDefaults() *RevokePrincipalAccessResourceLinksLinks {
	this := RevokePrincipalAccessResourceLinksLinks{}
	return &this
}

// GetPrincipalAccess returns the PrincipalAccess field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResourceLinksLinks) GetPrincipalAccess() RevokePrincipalAccessResourceLinksLinksPrincipalAccess {
	if o == nil || IsNil(o.PrincipalAccess) {
		var ret RevokePrincipalAccessResourceLinksLinksPrincipalAccess
		return ret
	}
	return *o.PrincipalAccess
}

// GetPrincipalAccessOk returns a tuple with the PrincipalAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResourceLinksLinks) GetPrincipalAccessOk() (*RevokePrincipalAccessResourceLinksLinksPrincipalAccess, bool) {
	if o == nil || IsNil(o.PrincipalAccess) {
		return nil, false
	}
	return o.PrincipalAccess, true
}

// HasPrincipalAccess returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResourceLinksLinks) HasPrincipalAccess() bool {
	if o != nil && !IsNil(o.PrincipalAccess) {
		return true
	}

	return false
}

// SetPrincipalAccess gets a reference to the given RevokePrincipalAccessResourceLinksLinksPrincipalAccess and assigns it to the PrincipalAccess field.
func (o *RevokePrincipalAccessResourceLinksLinks) SetPrincipalAccess(v RevokePrincipalAccessResourceLinksLinksPrincipalAccess) {
	o.PrincipalAccess = &v
}

// GetPrincipalEntitlements returns the PrincipalEntitlements field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResourceLinksLinks) GetPrincipalEntitlements() RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements {
	if o == nil || IsNil(o.PrincipalEntitlements) {
		var ret RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements
		return ret
	}
	return *o.PrincipalEntitlements
}

// GetPrincipalEntitlementsOk returns a tuple with the PrincipalEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResourceLinksLinks) GetPrincipalEntitlementsOk() (*RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements, bool) {
	if o == nil || IsNil(o.PrincipalEntitlements) {
		return nil, false
	}
	return o.PrincipalEntitlements, true
}

// HasPrincipalEntitlements returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResourceLinksLinks) HasPrincipalEntitlements() bool {
	if o != nil && !IsNil(o.PrincipalEntitlements) {
		return true
	}

	return false
}

// SetPrincipalEntitlements gets a reference to the given RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements and assigns it to the PrincipalEntitlements field.
func (o *RevokePrincipalAccessResourceLinksLinks) SetPrincipalEntitlements(v RevokePrincipalAccessResourceLinksLinksPrincipalEntitlements) {
	o.PrincipalEntitlements = &v
}

func (o RevokePrincipalAccessResourceLinksLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessResourceLinksLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrincipalAccess) {
		toSerialize["principal-access"] = o.PrincipalAccess
	}
	if !IsNil(o.PrincipalEntitlements) {
		toSerialize["principal-entitlements"] = o.PrincipalEntitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessResourceLinksLinks) UnmarshalJSON(data []byte) (err error) {
	varRevokePrincipalAccessResourceLinksLinks := _RevokePrincipalAccessResourceLinksLinks{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessResourceLinksLinks)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessResourceLinksLinks(varRevokePrincipalAccessResourceLinksLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principal-access")
		delete(additionalProperties, "principal-entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessResourceLinksLinks struct {
	value *RevokePrincipalAccessResourceLinksLinks
	isSet bool
}

func (v NullableRevokePrincipalAccessResourceLinksLinks) Get() *RevokePrincipalAccessResourceLinksLinks {
	return v.value
}

func (v *NullableRevokePrincipalAccessResourceLinksLinks) Set(val *RevokePrincipalAccessResourceLinksLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessResourceLinksLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessResourceLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessResourceLinksLinks(val *RevokePrincipalAccessResourceLinksLinks) *NullableRevokePrincipalAccessResourceLinksLinks {
	return &NullableRevokePrincipalAccessResourceLinksLinks{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessResourceLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessResourceLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
