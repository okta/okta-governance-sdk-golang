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

// SecurityAccessReviewPrincipal struct for SecurityAccessReviewPrincipal
type SecurityAccessReviewPrincipal struct {
	// The Okta user `id`
	Id string `json:"id"`
	// The Okta user's email address
	Email string `json:"email"`
	// The Okta user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The Okta user's last name
	LastName *string `json:"lastName,omitempty"`
	// The Okta user's login
	Login  *string                `json:"login,omitempty"`
	Status PrincipalProfileStatus `json:"status"`
	// Department
	Department *string `json:"department,omitempty"`
	// Manager
	Manager *string `json:"manager,omitempty"`
	// Role
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
func NewSecurityAccessReviewPrincipal(id string, email string, status PrincipalProfileStatus) *SecurityAccessReviewPrincipal {
	this := SecurityAccessReviewPrincipal{}
	this.Id = id
	this.Email = email
	this.Status = status
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

// GetEmail returns the Email field value
func (o *SecurityAccessReviewPrincipal) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SecurityAccessReviewPrincipal) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
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
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLastName() bool {
	if o != nil && o.LastName != nil {
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
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLogin() bool {
	if o != nil && o.Login != nil {
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

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipal) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasDepartment() bool {
	if o != nil && o.Department != nil {
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
	if o == nil || o.Manager == nil {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetManagerOk() (*string, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasManager() bool {
	if o != nil && o.Manager != nil {
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
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasRole() bool {
	if o != nil && o.Role != nil {
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
	if o == nil || o.HomeLocation == nil {
		var ret SecurityAccessReviewPrincipalLocation
		return ret
	}
	return *o.HomeLocation
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetHomeLocationOk() (*SecurityAccessReviewPrincipalLocation, bool) {
	if o == nil || o.HomeLocation == nil {
		return nil, false
	}
	return o.HomeLocation, true
}

// HasHomeLocation returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasHomeLocation() bool {
	if o != nil && o.HomeLocation != nil {
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
	if o == nil || o.LastLoginInfo == nil {
		var ret SecurityAccessReviewPrincipalLastLoginInfo
		return ret
	}
	return *o.LastLoginInfo
}

// GetLastLoginInfoOk returns a tuple with the LastLoginInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLastLoginInfoOk() (*SecurityAccessReviewPrincipalLastLoginInfo, bool) {
	if o == nil || o.LastLoginInfo == nil {
		return nil, false
	}
	return o.LastLoginInfo, true
}

// HasLastLoginInfo returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLastLoginInfo() bool {
	if o != nil && o.LastLoginInfo != nil {
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
	if o == nil || o.OktaAdminRoles == nil {
		var ret []SecurityAccessReviewPrincipalOktaAdminRole
		return ret
	}
	return o.OktaAdminRoles
}

// GetOktaAdminRolesOk returns a tuple with the OktaAdminRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetOktaAdminRolesOk() ([]SecurityAccessReviewPrincipalOktaAdminRole, bool) {
	if o == nil || o.OktaAdminRoles == nil {
		return nil, false
	}
	return o.OktaAdminRoles, true
}

// HasOktaAdminRoles returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasOktaAdminRoles() bool {
	if o != nil && o.OktaAdminRoles != nil {
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
	if o == nil || o.Links == nil {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipal) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipal) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *SecurityAccessReviewPrincipal) SetLinks(v map[string]Link) {
	o.Links = &v
}

func (o SecurityAccessReviewPrincipal) MarshalJSON() ([]byte, error) {
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
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.HomeLocation != nil {
		toSerialize["homeLocation"] = o.HomeLocation
	}
	if o.LastLoginInfo != nil {
		toSerialize["lastLoginInfo"] = o.LastLoginInfo
	}
	if o.OktaAdminRoles != nil {
		toSerialize["oktaAdminRoles"] = o.OktaAdminRoles
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewPrincipal) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewPrincipal := _SecurityAccessReviewPrincipal{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewPrincipal)
	if err == nil {
		*o = SecurityAccessReviewPrincipal(varSecurityAccessReviewPrincipal)
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
		delete(additionalProperties, "department")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "role")
		delete(additionalProperties, "homeLocation")
		delete(additionalProperties, "lastLoginInfo")
		delete(additionalProperties, "oktaAdminRoles")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
