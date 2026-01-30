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

// checks if the DelegateUsersList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegateUsersList{}

// DelegateUsersList struct for DelegateUsersList
type DelegateUsersList struct {
	// Delegate users list
	Data                 []DelegateUser `json:"data,omitempty"`
	Links                *LinkSelf      `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DelegateUsersList DelegateUsersList

// NewDelegateUsersList instantiates a new DelegateUsersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateUsersList() *DelegateUsersList {
	this := DelegateUsersList{}
	return &this
}

// NewDelegateUsersListWithDefaults instantiates a new DelegateUsersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateUsersListWithDefaults() *DelegateUsersList {
	this := DelegateUsersList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DelegateUsersList) GetData() []DelegateUser {
	if o == nil || IsNil(o.Data) {
		var ret []DelegateUser
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateUsersList) GetDataOk() ([]DelegateUser, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DelegateUsersList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DelegateUser and assigns it to the Data field.
func (o *DelegateUsersList) SetData(v []DelegateUser) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DelegateUsersList) GetLinks() LinkSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinkSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateUsersList) GetLinksOk() (*LinkSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DelegateUsersList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinkSelf and assigns it to the Links field.
func (o *DelegateUsersList) SetLinks(v LinkSelf) {
	o.Links = &v
}

func (o DelegateUsersList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegateUsersList) ToMap() (map[string]interface{}, error) {
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

func (o *DelegateUsersList) UnmarshalJSON(data []byte) (err error) {
	varDelegateUsersList := _DelegateUsersList{}

	err = json.Unmarshal(data, &varDelegateUsersList)

	if err != nil {
		return err
	}

	*o = DelegateUsersList(varDelegateUsersList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDelegateUsersList struct {
	value *DelegateUsersList
	isSet bool
}

func (v NullableDelegateUsersList) Get() *DelegateUsersList {
	return v.value
}

func (v *NullableDelegateUsersList) Set(val *DelegateUsersList) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateUsersList) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateUsersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateUsersList(val *DelegateUsersList) *NullableDelegateUsersList {
	return &NullableDelegateUsersList{value: val, isSet: true}
}

func (v NullableDelegateUsersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateUsersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
