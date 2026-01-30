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
	"fmt"
)

// checks if the PrincipalProfileEnriched type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalProfileEnriched{}

// PrincipalProfileEnriched A limited set of properties from the principal's profile
type PrincipalProfileEnriched struct {
	// The Okta user `id`
	Id string `json:"id"`
	// The Okta user's email address
	Email *string `json:"email,omitempty"`
	// The Okta user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The Okta user's last name
	LastName *string `json:"lastName,omitempty"`
	// The Okta user's login
	Login                *string                `json:"login,omitempty"`
	Status               PrincipalProfileStatus `json:"status"`
	Type                 PrincipalProfileType   `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalProfileEnriched PrincipalProfileEnriched

// NewPrincipalProfileEnriched instantiates a new PrincipalProfileEnriched object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalProfileEnriched(id string, status PrincipalProfileStatus, type_ PrincipalProfileType) *PrincipalProfileEnriched {
	this := PrincipalProfileEnriched{}
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewPrincipalProfileEnrichedWithDefaults instantiates a new PrincipalProfileEnriched object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalProfileEnrichedWithDefaults() *PrincipalProfileEnriched {
	this := PrincipalProfileEnriched{}
	return &this
}

// GetId returns the Id field value
func (o *PrincipalProfileEnriched) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrincipalProfileEnriched) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PrincipalProfileEnriched) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PrincipalProfileEnriched) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PrincipalProfileEnriched) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PrincipalProfileEnriched) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PrincipalProfileEnriched) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PrincipalProfileEnriched) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PrincipalProfileEnriched) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PrincipalProfileEnriched) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PrincipalProfileEnriched) SetLastName(v string) {
	o.LastName = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *PrincipalProfileEnriched) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *PrincipalProfileEnriched) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *PrincipalProfileEnriched) SetLogin(v string) {
	o.Login = &v
}

// GetStatus returns the Status field value
func (o *PrincipalProfileEnriched) GetStatus() PrincipalProfileStatus {
	if o == nil {
		var ret PrincipalProfileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetStatusOk() (*PrincipalProfileStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PrincipalProfileEnriched) SetStatus(v PrincipalProfileStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *PrincipalProfileEnriched) GetType() PrincipalProfileType {
	if o == nil {
		var ret PrincipalProfileType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrincipalProfileEnriched) GetTypeOk() (*PrincipalProfileType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrincipalProfileEnriched) SetType(v PrincipalProfileType) {
	o.Type = v
}

func (o PrincipalProfileEnriched) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalProfileEnriched) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalProfileEnriched) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPrincipalProfileEnriched := _PrincipalProfileEnriched{}

	err = json.Unmarshal(data, &varPrincipalProfileEnriched)

	if err != nil {
		return err
	}

	*o = PrincipalProfileEnriched(varPrincipalProfileEnriched)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "login")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalProfileEnriched struct {
	value *PrincipalProfileEnriched
	isSet bool
}

func (v NullablePrincipalProfileEnriched) Get() *PrincipalProfileEnriched {
	return v.value
}

func (v *NullablePrincipalProfileEnriched) Set(val *PrincipalProfileEnriched) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalProfileEnriched) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalProfileEnriched) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalProfileEnriched(val *PrincipalProfileEnriched) *NullablePrincipalProfileEnriched {
	return &NullablePrincipalProfileEnriched{value: val, isSet: true}
}

func (v NullablePrincipalProfileEnriched) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalProfileEnriched) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
