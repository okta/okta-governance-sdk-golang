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

// checks if the DelegateSettingsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegateSettingsUser{}

// DelegateSettingsUser struct for DelegateSettingsUser
type DelegateSettingsUser struct {
	// The user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty"`
	// The user's email address
	Email                *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DelegateSettingsUser DelegateSettingsUser

// NewDelegateSettingsUser instantiates a new DelegateSettingsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateSettingsUser() *DelegateSettingsUser {
	this := DelegateSettingsUser{}
	return &this
}

// NewDelegateSettingsUserWithDefaults instantiates a new DelegateSettingsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateSettingsUserWithDefaults() *DelegateSettingsUser {
	this := DelegateSettingsUser{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *DelegateSettingsUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateSettingsUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *DelegateSettingsUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *DelegateSettingsUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *DelegateSettingsUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateSettingsUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *DelegateSettingsUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *DelegateSettingsUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DelegateSettingsUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateSettingsUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DelegateSettingsUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DelegateSettingsUser) SetEmail(v string) {
	o.Email = &v
}

func (o DelegateSettingsUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegateSettingsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DelegateSettingsUser) UnmarshalJSON(data []byte) (err error) {
	varDelegateSettingsUser := _DelegateSettingsUser{}

	err = json.Unmarshal(data, &varDelegateSettingsUser)

	if err != nil {
		return err
	}

	*o = DelegateSettingsUser(varDelegateSettingsUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDelegateSettingsUser struct {
	value *DelegateSettingsUser
	isSet bool
}

func (v NullableDelegateSettingsUser) Get() *DelegateSettingsUser {
	return v.value
}

func (v *NullableDelegateSettingsUser) Set(val *DelegateSettingsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateSettingsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateSettingsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateSettingsUser(val *DelegateSettingsUser) *NullableDelegateSettingsUser {
	return &NullableDelegateSettingsUser{value: val, isSet: true}
}

func (v NullableDelegateSettingsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateSettingsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
