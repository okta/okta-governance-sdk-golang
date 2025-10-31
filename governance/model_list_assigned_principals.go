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

// checks if the ListAssignedPrincipals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAssignedPrincipals{}

// ListAssignedPrincipals struct for ListAssignedPrincipals
type ListAssignedPrincipals struct {
	Data                 []AssignedPrincipalDetails `json:"data,omitempty"`
	Links                *PaginationLinks           `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAssignedPrincipals ListAssignedPrincipals

// NewListAssignedPrincipals instantiates a new ListAssignedPrincipals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAssignedPrincipals() *ListAssignedPrincipals {
	this := ListAssignedPrincipals{}
	return &this
}

// NewListAssignedPrincipalsWithDefaults instantiates a new ListAssignedPrincipals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAssignedPrincipalsWithDefaults() *ListAssignedPrincipals {
	this := ListAssignedPrincipals{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAssignedPrincipals) GetData() []AssignedPrincipalDetails {
	if o == nil || IsNil(o.Data) {
		var ret []AssignedPrincipalDetails
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAssignedPrincipals) GetDataOk() ([]AssignedPrincipalDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAssignedPrincipals) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AssignedPrincipalDetails and assigns it to the Data field.
func (o *ListAssignedPrincipals) SetData(v []AssignedPrincipalDetails) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListAssignedPrincipals) GetLinks() PaginationLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAssignedPrincipals) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListAssignedPrincipals) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListAssignedPrincipals) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o ListAssignedPrincipals) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAssignedPrincipals) ToMap() (map[string]interface{}, error) {
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

func (o *ListAssignedPrincipals) UnmarshalJSON(data []byte) (err error) {
	varListAssignedPrincipals := _ListAssignedPrincipals{}

	err = json.Unmarshal(data, &varListAssignedPrincipals)

	if err != nil {
		return err
	}

	*o = ListAssignedPrincipals(varListAssignedPrincipals)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAssignedPrincipals struct {
	value *ListAssignedPrincipals
	isSet bool
}

func (v NullableListAssignedPrincipals) Get() *ListAssignedPrincipals {
	return v.value
}

func (v *NullableListAssignedPrincipals) Set(val *ListAssignedPrincipals) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssignedPrincipals) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssignedPrincipals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssignedPrincipals(val *ListAssignedPrincipals) *NullableListAssignedPrincipals {
	return &NullableListAssignedPrincipals{value: val, isSet: true}
}

func (v NullableListAssignedPrincipals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssignedPrincipals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
