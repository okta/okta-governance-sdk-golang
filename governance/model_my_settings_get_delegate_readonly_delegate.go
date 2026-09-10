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

// checks if the MySettingsGetDelegateReadonlyDelegate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MySettingsGetDelegateReadonlyDelegate{}

// MySettingsGetDelegateReadonlyDelegate struct for MySettingsGetDelegateReadonlyDelegate
type MySettingsGetDelegateReadonlyDelegate struct {
	// The Okta user `id`
	ExternalId string        `json:"externalId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Type       PrincipalType `json:"type"`
	// The user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty"`
	// The user's email address
	Email                *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MySettingsGetDelegateReadonlyDelegate MySettingsGetDelegateReadonlyDelegate

// NewMySettingsGetDelegateReadonlyDelegate instantiates a new MySettingsGetDelegateReadonlyDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettingsGetDelegateReadonlyDelegate(externalId string, type_ PrincipalType) *MySettingsGetDelegateReadonlyDelegate {
	this := MySettingsGetDelegateReadonlyDelegate{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewMySettingsGetDelegateReadonlyDelegateWithDefaults instantiates a new MySettingsGetDelegateReadonlyDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsGetDelegateReadonlyDelegateWithDefaults() *MySettingsGetDelegateReadonlyDelegate {
	this := MySettingsGetDelegateReadonlyDelegate{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *MySettingsGetDelegateReadonlyDelegate) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *MySettingsGetDelegateReadonlyDelegate) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *MySettingsGetDelegateReadonlyDelegate) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MySettingsGetDelegateReadonlyDelegate) SetType(v PrincipalType) {
	o.Type = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonlyDelegate) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *MySettingsGetDelegateReadonlyDelegate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonlyDelegate) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *MySettingsGetDelegateReadonlyDelegate) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonlyDelegate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonlyDelegate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MySettingsGetDelegateReadonlyDelegate) SetEmail(v string) {
	o.Email = &v
}

func (o MySettingsGetDelegateReadonlyDelegate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MySettingsGetDelegateReadonlyDelegate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type
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

func (o *MySettingsGetDelegateReadonlyDelegate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
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

	varMySettingsGetDelegateReadonlyDelegate := _MySettingsGetDelegateReadonlyDelegate{}

	err = json.Unmarshal(data, &varMySettingsGetDelegateReadonlyDelegate)

	if err != nil {
		return err
	}

	*o = MySettingsGetDelegateReadonlyDelegate(varMySettingsGetDelegateReadonlyDelegate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMySettingsGetDelegateReadonlyDelegate struct {
	value *MySettingsGetDelegateReadonlyDelegate
	isSet bool
}

func (v NullableMySettingsGetDelegateReadonlyDelegate) Get() *MySettingsGetDelegateReadonlyDelegate {
	return v.value
}

func (v *NullableMySettingsGetDelegateReadonlyDelegate) Set(val *MySettingsGetDelegateReadonlyDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullableMySettingsGetDelegateReadonlyDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullableMySettingsGetDelegateReadonlyDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMySettingsGetDelegateReadonlyDelegate(val *MySettingsGetDelegateReadonlyDelegate) *NullableMySettingsGetDelegateReadonlyDelegate {
	return &NullableMySettingsGetDelegateReadonlyDelegate{value: val, isSet: true}
}

func (v NullableMySettingsGetDelegateReadonlyDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMySettingsGetDelegateReadonlyDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
