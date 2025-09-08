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

// ListCatalogUsers struct for ListCatalogUsers
type ListCatalogUsers struct {
	Data                 []PrincipalProfile `json:"data,omitempty"`
	Links                *SelfLink          `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCatalogUsers ListCatalogUsers

// NewListCatalogUsers instantiates a new ListCatalogUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCatalogUsers() *ListCatalogUsers {
	this := ListCatalogUsers{}
	return &this
}

// NewListCatalogUsersWithDefaults instantiates a new ListCatalogUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCatalogUsersWithDefaults() *ListCatalogUsers {
	this := ListCatalogUsers{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListCatalogUsers) GetData() []PrincipalProfile {
	if o == nil || o.Data == nil {
		var ret []PrincipalProfile
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogUsers) GetDataOk() ([]PrincipalProfile, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCatalogUsers) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrincipalProfile and assigns it to the Data field.
func (o *ListCatalogUsers) SetData(v []PrincipalProfile) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListCatalogUsers) GetLinks() SelfLink {
	if o == nil || o.Links == nil {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogUsers) GetLinksOk() (*SelfLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListCatalogUsers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *ListCatalogUsers) SetLinks(v SelfLink) {
	o.Links = &v
}

func (o ListCatalogUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListCatalogUsers) UnmarshalJSON(bytes []byte) (err error) {
	varListCatalogUsers := _ListCatalogUsers{}

	err = json.Unmarshal(bytes, &varListCatalogUsers)
	if err == nil {
		*o = ListCatalogUsers(varListCatalogUsers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableListCatalogUsers struct {
	value *ListCatalogUsers
	isSet bool
}

func (v NullableListCatalogUsers) Get() *ListCatalogUsers {
	return v.value
}

func (v *NullableListCatalogUsers) Set(val *ListCatalogUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableListCatalogUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableListCatalogUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCatalogUsers(val *ListCatalogUsers) *NullableListCatalogUsers {
	return &NullableListCatalogUsers{value: val, isSet: true}
}

func (v NullableListCatalogUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCatalogUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
