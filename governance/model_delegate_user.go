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

// DelegateUser A limited set of properties from the user's profile
type DelegateUser struct {
	// The Okta user `id`
	Id string `json:"id" validate:"regexp=00u[0-9a-zA-Z]+"`
	// The user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty"`
	// The user's email address
	Email                string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _DelegateUser DelegateUser

// NewDelegateUser instantiates a new DelegateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateUser(id string, email string) *DelegateUser {
	this := DelegateUser{}
	this.Id = id
	this.Email = email
	return &this
}

// NewDelegateUserWithDefaults instantiates a new DelegateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateUserWithDefaults() *DelegateUser {
	this := DelegateUser{}
	return &this
}

// GetId returns the Id field value
func (o *DelegateUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelegateUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelegateUser) SetId(v string) {
	o.Id = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *DelegateUser) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateUser) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *DelegateUser) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *DelegateUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *DelegateUser) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateUser) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *DelegateUser) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *DelegateUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value
func (o *DelegateUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DelegateUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DelegateUser) SetEmail(v string) {
	o.Email = v
}

func (o DelegateUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DelegateUser) UnmarshalJSON(bytes []byte) (err error) {
	varDelegateUser := _DelegateUser{}

	err = json.Unmarshal(bytes, &varDelegateUser)
	if err == nil {
		*o = DelegateUser(varDelegateUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDelegateUser struct {
	value *DelegateUser
	isSet bool
}

func (v NullableDelegateUser) Get() *DelegateUser {
	return v.value
}

func (v *NullableDelegateUser) Set(val *DelegateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateUser(val *DelegateUser) *NullableDelegateUser {
	return &NullableDelegateUser{value: val, isSet: true}
}

func (v NullableDelegateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
