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

// PrincipalProfile A limited set of properties from the principal's profile
type PrincipalProfile struct {
	// The Okta user `id`
	Id string `json:"id"`
	// The Okta user's email address
	Email string `json:"email"`
	// The Okta user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The Okta user's last name
	LastName *string `json:"lastName,omitempty"`
	// The Okta user's login
	Login                *string                `json:"login,omitempty"`
	Status               PrincipalProfileStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalProfile PrincipalProfile

// NewPrincipalProfile instantiates a new PrincipalProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalProfile(id string, email string, status PrincipalProfileStatus) *PrincipalProfile {
	this := PrincipalProfile{}
	this.Id = id
	this.Email = email
	this.Status = status
	return &this
}

// NewPrincipalProfileWithDefaults instantiates a new PrincipalProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalProfileWithDefaults() *PrincipalProfile {
	this := PrincipalProfile{}
	return &this
}

// GetId returns the Id field value
func (o *PrincipalProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrincipalProfile) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *PrincipalProfile) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PrincipalProfile) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PrincipalProfile) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PrincipalProfile) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PrincipalProfile) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PrincipalProfile) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PrincipalProfile) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PrincipalProfile) SetLastName(v string) {
	o.LastName = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *PrincipalProfile) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *PrincipalProfile) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *PrincipalProfile) SetLogin(v string) {
	o.Login = &v
}

// GetStatus returns the Status field value
func (o *PrincipalProfile) GetStatus() PrincipalProfileStatus {
	if o == nil {
		var ret PrincipalProfileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfile) GetStatusOk() (*PrincipalProfileStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PrincipalProfile) SetStatus(v PrincipalProfileStatus) {
	o.Status = v
}

func (o PrincipalProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalProfile) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalProfile := _PrincipalProfile{}

	err = json.Unmarshal(bytes, &varPrincipalProfile)
	if err == nil {
		*o = PrincipalProfile(varPrincipalProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "login")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrincipalProfile struct {
	value *PrincipalProfile
	isSet bool
}

func (v NullablePrincipalProfile) Get() *PrincipalProfile {
	return v.value
}

func (v *NullablePrincipalProfile) Set(val *PrincipalProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalProfile(val *PrincipalProfile) *NullablePrincipalProfile {
	return &NullablePrincipalProfile{value: val, isSet: true}
}

func (v NullablePrincipalProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
