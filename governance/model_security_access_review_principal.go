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

// checks if the SecurityAccessReviewPrincipal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewPrincipal{}

// SecurityAccessReviewPrincipal struct for SecurityAccessReviewPrincipal
type SecurityAccessReviewPrincipal struct {
	// The Okta user `id`
	Id string `json:"id"`
	// The Okta user's email address
	Email *string `json:"email,omitempty"`
	// The Okta user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The Okta user's last name
	LastName *string `json:"lastName,omitempty"`
	// The Okta user's login
	Login  *string                `json:"login,omitempty"`
	Status PrincipalProfileStatus `json:"status"`
	Type   PrincipalProfileType   `json:"type"`
	// The department of the Okta user
	Department *string `json:"department,omitempty"`
	// The manager of the Okta user
	Manager *string `json:"manager,omitempty"`
	// The Okta user role
	Role          *string                                     `json:"role,omitempty"`
	HomeLocation  *SecurityAccessReviewPrincipalLocation      `json:"homeLocation,omitempty"`
	LastLoginInfo *SecurityAccessReviewPrincipalLastLoginInfo `json:"lastLoginInfo,omitempty"`
	// List of Okta admin roles assigned to the principal
	OktaAdminRoles       []SecurityAccessReviewPrincipalOktaAdminRole `json:"oktaAdminRoles,omitempty"`
	Links                *map[string]Link                             `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewPrincipal SecurityAccessReviewPrincipal

// NewSecurityAccessReviewPrincipal instantiates a new SecurityAccessReviewPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewPrincipal(id string, status PrincipalProfileStatus, type_ PrincipalProfileType) *SecurityAccessReviewPrincipal {
	this := SecurityAccessReviewPrincipal{}
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewSecurityAccessReviewPrincipalWithDefaults instantiates a new SecurityAccessReviewPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewPrincipalWithDefaults() *SecurityAccessReviewPrincipal {
	this := SecurityAccessReviewPrincipal{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewPrincipal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewPrincipal) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SecurityAccessReviewPrincipal) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *SecurityAccessReviewPrincipal) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *SecurityAccessReviewPrincipal) SetLastName(v string) {
	o.LastName = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *SecurityAccessReviewPrincipal) SetLogin(v string) {
	o.Login = &v
}

// GetStatus returns the Status field value
func (o *SecurityAccessReviewPrincipal) GetStatus() PrincipalProfileStatus {
	if o == nil {
		var ret PrincipalProfileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetStatusOk() (*PrincipalProfileStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SecurityAccessReviewPrincipal) SetStatus(v PrincipalProfileStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *SecurityAccessReviewPrincipal) GetType() PrincipalProfileType {
	if o == nil {
		var ret PrincipalProfileType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetTypeOk() (*PrincipalProfileType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewPrincipal) SetType(v PrincipalProfileType) {
	o.Type = v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *SecurityAccessReviewPrincipal) SetDepartment(v string) {
	o.Department = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetManager() string {
	if o == nil || IsNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetManagerOk() (*string, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *SecurityAccessReviewPrincipal) SetManager(v string) {
	o.Manager = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SecurityAccessReviewPrincipal) SetRole(v string) {
	o.Role = &v
}

// GetHomeLocation returns the HomeLocation field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetHomeLocation() SecurityAccessReviewPrincipalLocation {
	if o == nil || IsNil(o.HomeLocation) {
		var ret SecurityAccessReviewPrincipalLocation
		return ret
	}
	return *o.HomeLocation
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetHomeLocationOk() (*SecurityAccessReviewPrincipalLocation, bool) {
	if o == nil || IsNil(o.HomeLocation) {
		return nil, false
	}
	return o.HomeLocation, true
}

// HasHomeLocation returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasHomeLocation() bool {
	if o != nil && !IsNil(o.HomeLocation) {
		return true
	}

	return false
}

// SetHomeLocation gets a reference to the given SecurityAccessReviewPrincipalLocation and assigns it to the HomeLocation field.
func (o *SecurityAccessReviewPrincipal) SetHomeLocation(v SecurityAccessReviewPrincipalLocation) {
	o.HomeLocation = &v
}

// GetLastLoginInfo returns the LastLoginInfo field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetLastLoginInfo() SecurityAccessReviewPrincipalLastLoginInfo {
	if o == nil || IsNil(o.LastLoginInfo) {
		var ret SecurityAccessReviewPrincipalLastLoginInfo
		return ret
	}
	return *o.LastLoginInfo
}

// GetLastLoginInfoOk returns a tuple with the LastLoginInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLastLoginInfoOk() (*SecurityAccessReviewPrincipalLastLoginInfo, bool) {
	if o == nil || IsNil(o.LastLoginInfo) {
		return nil, false
	}
	return o.LastLoginInfo, true
}

// HasLastLoginInfo returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLastLoginInfo() bool {
	if o != nil && !IsNil(o.LastLoginInfo) {
		return true
	}

	return false
}

// SetLastLoginInfo gets a reference to the given SecurityAccessReviewPrincipalLastLoginInfo and assigns it to the LastLoginInfo field.
func (o *SecurityAccessReviewPrincipal) SetLastLoginInfo(v SecurityAccessReviewPrincipalLastLoginInfo) {
	o.LastLoginInfo = &v
}

// GetOktaAdminRoles returns the OktaAdminRoles field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetOktaAdminRoles() []SecurityAccessReviewPrincipalOktaAdminRole {
	if o == nil || IsNil(o.OktaAdminRoles) {
		var ret []SecurityAccessReviewPrincipalOktaAdminRole
		return ret
	}
	return o.OktaAdminRoles
}

// GetOktaAdminRolesOk returns a tuple with the OktaAdminRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetOktaAdminRolesOk() ([]SecurityAccessReviewPrincipalOktaAdminRole, bool) {
	if o == nil || IsNil(o.OktaAdminRoles) {
		return nil, false
	}
	return o.OktaAdminRoles, true
}

// HasOktaAdminRoles returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasOktaAdminRoles() bool {
	if o != nil && !IsNil(o.OktaAdminRoles) {
		return true
	}

	return false
}

// SetOktaAdminRoles gets a reference to the given []SecurityAccessReviewPrincipalOktaAdminRole and assigns it to the OktaAdminRoles field.
func (o *SecurityAccessReviewPrincipal) SetOktaAdminRoles(v []SecurityAccessReviewPrincipalOktaAdminRole) {
	o.OktaAdminRoles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetLinks() map[string]Link {
	if o == nil || IsNil(o.Links) {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *SecurityAccessReviewPrincipal) SetLinks(v map[string]Link) {
	o.Links = &v
}

func (o SecurityAccessReviewPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewPrincipal) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.HomeLocation) {
		toSerialize["homeLocation"] = o.HomeLocation
	}
	if !IsNil(o.LastLoginInfo) {
		toSerialize["lastLoginInfo"] = o.LastLoginInfo
	}
	if !IsNil(o.OktaAdminRoles) {
		toSerialize["oktaAdminRoles"] = o.OktaAdminRoles
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewPrincipal) UnmarshalJSON(data []byte) (err error) {
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

	varSecurityAccessReviewPrincipal := _SecurityAccessReviewPrincipal{}

	err = json.Unmarshal(data, &varSecurityAccessReviewPrincipal)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewPrincipal(varSecurityAccessReviewPrincipal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "login")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "department")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "role")
		delete(additionalProperties, "homeLocation")
		delete(additionalProperties, "lastLoginInfo")
		delete(additionalProperties, "oktaAdminRoles")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewPrincipal struct {
	value *SecurityAccessReviewPrincipal
	isSet bool
}

func (v NullableSecurityAccessReviewPrincipal) Get() *SecurityAccessReviewPrincipal {
	return v.value
}

func (v *NullableSecurityAccessReviewPrincipal) Set(val *SecurityAccessReviewPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewPrincipal(val *SecurityAccessReviewPrincipal) *NullableSecurityAccessReviewPrincipal {
	return &NullableSecurityAccessReviewPrincipal{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
