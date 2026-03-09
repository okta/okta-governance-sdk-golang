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

// checks if the RevokePrincipalAccessResourceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessResourceLinks{}

// RevokePrincipalAccessResourceLinks Links to related resources for a revoked resource
type RevokePrincipalAccessResourceLinks struct {
	Links                *RevokePrincipalAccessResourceLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevokePrincipalAccessResourceLinks RevokePrincipalAccessResourceLinks

// NewRevokePrincipalAccessResourceLinks instantiates a new RevokePrincipalAccessResourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessResourceLinks() *RevokePrincipalAccessResourceLinks {
	this := RevokePrincipalAccessResourceLinks{}
	return &this
}

// NewRevokePrincipalAccessResourceLinksWithDefaults instantiates a new RevokePrincipalAccessResourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessResourceLinksWithDefaults() *RevokePrincipalAccessResourceLinks {
	this := RevokePrincipalAccessResourceLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResourceLinks) GetLinks() RevokePrincipalAccessResourceLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret RevokePrincipalAccessResourceLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResourceLinks) GetLinksOk() (*RevokePrincipalAccessResourceLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResourceLinks) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RevokePrincipalAccessResourceLinksLinks and assigns it to the Links field.
func (o *RevokePrincipalAccessResourceLinks) SetLinks(v RevokePrincipalAccessResourceLinksLinks) {
	o.Links = &v
}

func (o RevokePrincipalAccessResourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessResourceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessResourceLinks) UnmarshalJSON(data []byte) (err error) {
	varRevokePrincipalAccessResourceLinks := _RevokePrincipalAccessResourceLinks{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessResourceLinks)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessResourceLinks(varRevokePrincipalAccessResourceLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessResourceLinks struct {
	value *RevokePrincipalAccessResourceLinks
	isSet bool
}

func (v NullableRevokePrincipalAccessResourceLinks) Get() *RevokePrincipalAccessResourceLinks {
	return v.value
}

func (v *NullableRevokePrincipalAccessResourceLinks) Set(val *RevokePrincipalAccessResourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessResourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessResourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessResourceLinks(val *RevokePrincipalAccessResourceLinks) *NullableRevokePrincipalAccessResourceLinks {
	return &NullableRevokePrincipalAccessResourceLinks{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessResourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessResourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
