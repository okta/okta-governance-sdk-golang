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

// checks if the PrincipalScopeSettingsMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalScopeSettingsMutable{}

// PrincipalScopeSettingsMutable User scope specific settings. If all the users of a resource under review are not part of the scope of certification, provide the scope of the user by means of a user expression.
type PrincipalScopeSettingsMutable struct {
	Type PrincipalScopeType `json:"type"`
	// The Okta expression language user expression on the `resourceSettings` to include users in the campaign.
	UserScopeExpression *string `json:"userScopeExpression,omitempty"`
	// An array of Okta user IDs excluded from access certification or the campaign. This field is optional. A maximum of 50 users can be specified in the array.
	ExcludedUserIds []string `json:"excludedUserIds,omitempty"`
	// An array of Okta user IDs included from access certification or the campaign. `userIds`, `groupIds` or `userScopeExpression` is required if campaign type is `USER`. A maximum of 100 users can be specified in the array.
	UserIds []string `json:"userIds,omitempty"`
	// An array of Okta group IDs included from access certification or the campaign. `userIds`, `groupIds` or `userScopeExpression` is required if campaign type is `USER`. A maximum of 5 groups can be specified in the array.
	GroupIds []string `json:"groupIds,omitempty"`
	// If set to `true`, only active Okta users are included in the campaign
	IncludeOnlyActiveUsers       *bool                                 `json:"includeOnlyActiveUsers,omitempty"`
	PredefinedInactiveUsersScope *PredefinedInactiveUsersScopeSettings `json:"predefinedInactiveUsersScope,omitempty"`
	// If set to `true`, only includes users that have at least one SOD conflict that was caused due to entitlement(s) within Campaign scope
	OnlyIncludeUsersWithSODConflicts *bool `json:"onlyIncludeUsersWithSODConflicts,omitempty"`
	AdditionalProperties             map[string]interface{}
}

type _PrincipalScopeSettingsMutable PrincipalScopeSettingsMutable

// NewPrincipalScopeSettingsMutable instantiates a new PrincipalScopeSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalScopeSettingsMutable(type_ PrincipalScopeType) *PrincipalScopeSettingsMutable {
	this := PrincipalScopeSettingsMutable{}
	this.Type = type_
	return &this
}

// NewPrincipalScopeSettingsMutableWithDefaults instantiates a new PrincipalScopeSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalScopeSettingsMutableWithDefaults() *PrincipalScopeSettingsMutable {
	this := PrincipalScopeSettingsMutable{}
	return &this
}

// GetType returns the Type field value
func (o *PrincipalScopeSettingsMutable) GetType() PrincipalScopeType {
	if o == nil {
		var ret PrincipalScopeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetTypeOk() (*PrincipalScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrincipalScopeSettingsMutable) SetType(v PrincipalScopeType) {
	o.Type = v
}

// GetUserScopeExpression returns the UserScopeExpression field value if set, zero value otherwise.
func (o *PrincipalScopeSettingsMutable) GetUserScopeExpression() string {
	if o == nil || IsNil(o.UserScopeExpression) {
		var ret string
		return ret
	}
	return *o.UserScopeExpression
}

// GetUserScopeExpressionOk returns a tuple with the UserScopeExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetUserScopeExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.UserScopeExpression) {
		return nil, false
	}
	return o.UserScopeExpression, true
}

// HasUserScopeExpression returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasUserScopeExpression() bool {
	if o != nil && !IsNil(o.UserScopeExpression) {
		return true
	}

	return false
}

// SetUserScopeExpression gets a reference to the given string and assigns it to the UserScopeExpression field.
func (o *PrincipalScopeSettingsMutable) SetUserScopeExpression(v string) {
	o.UserScopeExpression = &v
}

// GetExcludedUserIds returns the ExcludedUserIds field value if set, zero value otherwise.
func (o *PrincipalScopeSettingsMutable) GetExcludedUserIds() []string {
	if o == nil || IsNil(o.ExcludedUserIds) {
		var ret []string
		return ret
	}
	return o.ExcludedUserIds
}

// GetExcludedUserIdsOk returns a tuple with the ExcludedUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetExcludedUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedUserIds) {
		return nil, false
	}
	return o.ExcludedUserIds, true
}

// HasExcludedUserIds returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasExcludedUserIds() bool {
	if o != nil && !IsNil(o.ExcludedUserIds) {
		return true
	}

	return false
}

// SetExcludedUserIds gets a reference to the given []string and assigns it to the ExcludedUserIds field.
func (o *PrincipalScopeSettingsMutable) SetExcludedUserIds(v []string) {
	o.ExcludedUserIds = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalScopeSettingsMutable) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalScopeSettingsMutable) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *PrincipalScopeSettingsMutable) SetUserIds(v []string) {
	o.UserIds = v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalScopeSettingsMutable) GetGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalScopeSettingsMutable) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *PrincipalScopeSettingsMutable) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetIncludeOnlyActiveUsers returns the IncludeOnlyActiveUsers field value if set, zero value otherwise.
func (o *PrincipalScopeSettingsMutable) GetIncludeOnlyActiveUsers() bool {
	if o == nil || IsNil(o.IncludeOnlyActiveUsers) {
		var ret bool
		return ret
	}
	return *o.IncludeOnlyActiveUsers
}

// GetIncludeOnlyActiveUsersOk returns a tuple with the IncludeOnlyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetIncludeOnlyActiveUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeOnlyActiveUsers) {
		return nil, false
	}
	return o.IncludeOnlyActiveUsers, true
}

// HasIncludeOnlyActiveUsers returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasIncludeOnlyActiveUsers() bool {
	if o != nil && !IsNil(o.IncludeOnlyActiveUsers) {
		return true
	}

	return false
}

// SetIncludeOnlyActiveUsers gets a reference to the given bool and assigns it to the IncludeOnlyActiveUsers field.
func (o *PrincipalScopeSettingsMutable) SetIncludeOnlyActiveUsers(v bool) {
	o.IncludeOnlyActiveUsers = &v
}

// GetPredefinedInactiveUsersScope returns the PredefinedInactiveUsersScope field value if set, zero value otherwise.
func (o *PrincipalScopeSettingsMutable) GetPredefinedInactiveUsersScope() PredefinedInactiveUsersScopeSettings {
	if o == nil || IsNil(o.PredefinedInactiveUsersScope) {
		var ret PredefinedInactiveUsersScopeSettings
		return ret
	}
	return *o.PredefinedInactiveUsersScope
}

// GetPredefinedInactiveUsersScopeOk returns a tuple with the PredefinedInactiveUsersScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetPredefinedInactiveUsersScopeOk() (*PredefinedInactiveUsersScopeSettings, bool) {
	if o == nil || IsNil(o.PredefinedInactiveUsersScope) {
		return nil, false
	}
	return o.PredefinedInactiveUsersScope, true
}

// HasPredefinedInactiveUsersScope returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasPredefinedInactiveUsersScope() bool {
	if o != nil && !IsNil(o.PredefinedInactiveUsersScope) {
		return true
	}

	return false
}

// SetPredefinedInactiveUsersScope gets a reference to the given PredefinedInactiveUsersScopeSettings and assigns it to the PredefinedInactiveUsersScope field.
func (o *PrincipalScopeSettingsMutable) SetPredefinedInactiveUsersScope(v PredefinedInactiveUsersScopeSettings) {
	o.PredefinedInactiveUsersScope = &v
}

// GetOnlyIncludeUsersWithSODConflicts returns the OnlyIncludeUsersWithSODConflicts field value if set, zero value otherwise.
func (o *PrincipalScopeSettingsMutable) GetOnlyIncludeUsersWithSODConflicts() bool {
	if o == nil || IsNil(o.OnlyIncludeUsersWithSODConflicts) {
		var ret bool
		return ret
	}
	return *o.OnlyIncludeUsersWithSODConflicts
}

// GetOnlyIncludeUsersWithSODConflictsOk returns a tuple with the OnlyIncludeUsersWithSODConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalScopeSettingsMutable) GetOnlyIncludeUsersWithSODConflictsOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyIncludeUsersWithSODConflicts) {
		return nil, false
	}
	return o.OnlyIncludeUsersWithSODConflicts, true
}

// HasOnlyIncludeUsersWithSODConflicts returns a boolean if a field has been set.
func (o *PrincipalScopeSettingsMutable) HasOnlyIncludeUsersWithSODConflicts() bool {
	if o != nil && !IsNil(o.OnlyIncludeUsersWithSODConflicts) {
		return true
	}

	return false
}

// SetOnlyIncludeUsersWithSODConflicts gets a reference to the given bool and assigns it to the OnlyIncludeUsersWithSODConflicts field.
func (o *PrincipalScopeSettingsMutable) SetOnlyIncludeUsersWithSODConflicts(v bool) {
	o.OnlyIncludeUsersWithSODConflicts = &v
}

func (o PrincipalScopeSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalScopeSettingsMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.UserScopeExpression) {
		toSerialize["userScopeExpression"] = o.UserScopeExpression
	}
	if !IsNil(o.ExcludedUserIds) {
		toSerialize["excludedUserIds"] = o.ExcludedUserIds
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	if !IsNil(o.IncludeOnlyActiveUsers) {
		toSerialize["includeOnlyActiveUsers"] = o.IncludeOnlyActiveUsers
	}
	if !IsNil(o.PredefinedInactiveUsersScope) {
		toSerialize["predefinedInactiveUsersScope"] = o.PredefinedInactiveUsersScope
	}
	if !IsNil(o.OnlyIncludeUsersWithSODConflicts) {
		toSerialize["onlyIncludeUsersWithSODConflicts"] = o.OnlyIncludeUsersWithSODConflicts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalScopeSettingsMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPrincipalScopeSettingsMutable := _PrincipalScopeSettingsMutable{}

	err = json.Unmarshal(data, &varPrincipalScopeSettingsMutable)

	if err != nil {
		return err
	}

	*o = PrincipalScopeSettingsMutable(varPrincipalScopeSettingsMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "userScopeExpression")
		delete(additionalProperties, "excludedUserIds")
		delete(additionalProperties, "userIds")
		delete(additionalProperties, "groupIds")
		delete(additionalProperties, "includeOnlyActiveUsers")
		delete(additionalProperties, "predefinedInactiveUsersScope")
		delete(additionalProperties, "onlyIncludeUsersWithSODConflicts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalScopeSettingsMutable struct {
	value *PrincipalScopeSettingsMutable
	isSet bool
}

func (v NullablePrincipalScopeSettingsMutable) Get() *PrincipalScopeSettingsMutable {
	return v.value
}

func (v *NullablePrincipalScopeSettingsMutable) Set(val *PrincipalScopeSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalScopeSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalScopeSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalScopeSettingsMutable(val *PrincipalScopeSettingsMutable) *NullablePrincipalScopeSettingsMutable {
	return &NullablePrincipalScopeSettingsMutable{value: val, isSet: true}
}

func (v NullablePrincipalScopeSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalScopeSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
