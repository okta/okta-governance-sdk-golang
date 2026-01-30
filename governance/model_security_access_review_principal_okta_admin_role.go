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

// checks if the SecurityAccessReviewPrincipalOktaAdminRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewPrincipalOktaAdminRole{}

// SecurityAccessReviewPrincipalOktaAdminRole struct for SecurityAccessReviewPrincipalOktaAdminRole
type SecurityAccessReviewPrincipalOktaAdminRole struct {
	// Unique identifier for the Okta admin role
	RoleId string `json:"roleId"`
	// Name of the Okta admin role
	RoleName             string `json:"roleName"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewPrincipalOktaAdminRole SecurityAccessReviewPrincipalOktaAdminRole

// NewSecurityAccessReviewPrincipalOktaAdminRole instantiates a new SecurityAccessReviewPrincipalOktaAdminRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewPrincipalOktaAdminRole(roleId string, roleName string) *SecurityAccessReviewPrincipalOktaAdminRole {
	this := SecurityAccessReviewPrincipalOktaAdminRole{}
	this.RoleId = roleId
	this.RoleName = roleName
	return &this
}

// NewSecurityAccessReviewPrincipalOktaAdminRoleWithDefaults instantiates a new SecurityAccessReviewPrincipalOktaAdminRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewPrincipalOktaAdminRoleWithDefaults() *SecurityAccessReviewPrincipalOktaAdminRole {
	this := SecurityAccessReviewPrincipalOktaAdminRole{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *SecurityAccessReviewPrincipalOktaAdminRole) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalOktaAdminRole) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *SecurityAccessReviewPrincipalOktaAdminRole) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *SecurityAccessReviewPrincipalOktaAdminRole) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalOktaAdminRole) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *SecurityAccessReviewPrincipalOktaAdminRole) SetRoleName(v string) {
	o.RoleName = v
}

func (o SecurityAccessReviewPrincipalOktaAdminRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewPrincipalOktaAdminRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	toSerialize["roleName"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewPrincipalOktaAdminRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleId",
		"roleName",
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

	varSecurityAccessReviewPrincipalOktaAdminRole := _SecurityAccessReviewPrincipalOktaAdminRole{}

	err = json.Unmarshal(data, &varSecurityAccessReviewPrincipalOktaAdminRole)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewPrincipalOktaAdminRole(varSecurityAccessReviewPrincipalOktaAdminRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "roleId")
		delete(additionalProperties, "roleName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewPrincipalOktaAdminRole struct {
	value *SecurityAccessReviewPrincipalOktaAdminRole
	isSet bool
}

func (v NullableSecurityAccessReviewPrincipalOktaAdminRole) Get() *SecurityAccessReviewPrincipalOktaAdminRole {
	return v.value
}

func (v *NullableSecurityAccessReviewPrincipalOktaAdminRole) Set(val *SecurityAccessReviewPrincipalOktaAdminRole) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewPrincipalOktaAdminRole) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewPrincipalOktaAdminRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewPrincipalOktaAdminRole(val *SecurityAccessReviewPrincipalOktaAdminRole) *NullableSecurityAccessReviewPrincipalOktaAdminRole {
	return &NullableSecurityAccessReviewPrincipalOktaAdminRole{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewPrincipalOktaAdminRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewPrincipalOktaAdminRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
