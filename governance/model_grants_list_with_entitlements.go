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

// checks if the GrantsListWithEntitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantsListWithEntitlements{}

// GrantsListWithEntitlements struct for GrantsListWithEntitlements
type GrantsListWithEntitlements struct {
	Data                 []GrantFullWithEntitlements `json:"data,omitempty"`
	Links                *GrantListLinks             `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantsListWithEntitlements GrantsListWithEntitlements

// NewGrantsListWithEntitlements instantiates a new GrantsListWithEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantsListWithEntitlements() *GrantsListWithEntitlements {
	this := GrantsListWithEntitlements{}
	return &this
}

// NewGrantsListWithEntitlementsWithDefaults instantiates a new GrantsListWithEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantsListWithEntitlementsWithDefaults() *GrantsListWithEntitlements {
	this := GrantsListWithEntitlements{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GrantsListWithEntitlements) GetData() []GrantFullWithEntitlements {
	if o == nil || IsNil(o.Data) {
		var ret []GrantFullWithEntitlements
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantsListWithEntitlements) GetDataOk() ([]GrantFullWithEntitlements, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GrantsListWithEntitlements) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GrantFullWithEntitlements and assigns it to the Data field.
func (o *GrantsListWithEntitlements) SetData(v []GrantFullWithEntitlements) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GrantsListWithEntitlements) GetLinks() GrantListLinks {
	if o == nil || IsNil(o.Links) {
		var ret GrantListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantsListWithEntitlements) GetLinksOk() (*GrantListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GrantsListWithEntitlements) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GrantListLinks and assigns it to the Links field.
func (o *GrantsListWithEntitlements) SetLinks(v GrantListLinks) {
	o.Links = &v
}

func (o GrantsListWithEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantsListWithEntitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantsListWithEntitlements) UnmarshalJSON(data []byte) (err error) {
	varGrantsListWithEntitlements := _GrantsListWithEntitlements{}

	err = json.Unmarshal(data, &varGrantsListWithEntitlements)

	if err != nil {
		return err
	}

	*o = GrantsListWithEntitlements(varGrantsListWithEntitlements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantsListWithEntitlements struct {
	value *GrantsListWithEntitlements
	isSet bool
}

func (v NullableGrantsListWithEntitlements) Get() *GrantsListWithEntitlements {
	return v.value
}

func (v *NullableGrantsListWithEntitlements) Set(val *GrantsListWithEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantsListWithEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantsListWithEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantsListWithEntitlements(val *GrantsListWithEntitlements) *NullableGrantsListWithEntitlements {
	return &NullableGrantsListWithEntitlements{value: val, isSet: true}
}

func (v NullableGrantsListWithEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantsListWithEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
