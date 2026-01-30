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

// checks if the EntryUsersList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntryUsersList{}

// EntryUsersList struct for EntryUsersList
type EntryUsersList struct {
	// Catalog entry users list
	Data                 []EntryUser         `json:"data,omitempty"`
	Links                *EntryUserListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntryUsersList EntryUsersList

// NewEntryUsersList instantiates a new EntryUsersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryUsersList() *EntryUsersList {
	this := EntryUsersList{}
	return &this
}

// NewEntryUsersListWithDefaults instantiates a new EntryUsersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryUsersListWithDefaults() *EntryUsersList {
	this := EntryUsersList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EntryUsersList) GetData() []EntryUser {
	if o == nil || IsNil(o.Data) {
		var ret []EntryUser
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUsersList) GetDataOk() ([]EntryUser, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EntryUsersList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EntryUser and assigns it to the Data field.
func (o *EntryUsersList) SetData(v []EntryUser) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EntryUsersList) GetLinks() EntryUserListLinks {
	if o == nil || IsNil(o.Links) {
		var ret EntryUserListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUsersList) GetLinksOk() (*EntryUserListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EntryUsersList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EntryUserListLinks and assigns it to the Links field.
func (o *EntryUsersList) SetLinks(v EntryUserListLinks) {
	o.Links = &v
}

func (o EntryUsersList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntryUsersList) ToMap() (map[string]interface{}, error) {
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

func (o *EntryUsersList) UnmarshalJSON(data []byte) (err error) {
	varEntryUsersList := _EntryUsersList{}

	err = json.Unmarshal(data, &varEntryUsersList)

	if err != nil {
		return err
	}

	*o = EntryUsersList(varEntryUsersList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntryUsersList struct {
	value *EntryUsersList
	isSet bool
}

func (v NullableEntryUsersList) Get() *EntryUsersList {
	return v.value
}

func (v *NullableEntryUsersList) Set(val *EntryUsersList) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryUsersList) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryUsersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryUsersList(val *EntryUsersList) *NullableEntryUsersList {
	return &NullableEntryUsersList{value: val, isSet: true}
}

func (v NullableEntryUsersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryUsersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
