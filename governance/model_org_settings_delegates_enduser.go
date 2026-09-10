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

// checks if the OrgSettingsDelegatesEnduser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsDelegatesEnduser{}

// OrgSettingsDelegatesEnduser Delegate permission settings for end users
type OrgSettingsDelegatesEnduser struct {
	// The permission that applies to this setting:  | Permission | Description | |------|------| | `READ` | Allow end users to view their delegates | | `WRITE` | Allow end users to set their own delegates |
	Permissions []string `json:"permissions,omitempty"`
	// Restricts the scope of delegate assignment for end users when `permissions` is set to `WRITE`. If this list is empty, end users can assign any user as their delegate.
	OnlyFor              []DelegateScope `json:"onlyFor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsDelegatesEnduser OrgSettingsDelegatesEnduser

// NewOrgSettingsDelegatesEnduser instantiates a new OrgSettingsDelegatesEnduser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsDelegatesEnduser() *OrgSettingsDelegatesEnduser {
	this := OrgSettingsDelegatesEnduser{}
	return &this
}

// NewOrgSettingsDelegatesEnduserWithDefaults instantiates a new OrgSettingsDelegatesEnduser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsDelegatesEnduserWithDefaults() *OrgSettingsDelegatesEnduser {
	this := OrgSettingsDelegatesEnduser{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *OrgSettingsDelegatesEnduser) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsDelegatesEnduser) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *OrgSettingsDelegatesEnduser) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *OrgSettingsDelegatesEnduser) SetPermissions(v []string) {
	o.Permissions = v
}

// GetOnlyFor returns the OnlyFor field value if set, zero value otherwise.
func (o *OrgSettingsDelegatesEnduser) GetOnlyFor() []DelegateScope {
	if o == nil || IsNil(o.OnlyFor) {
		var ret []DelegateScope
		return ret
	}
	return o.OnlyFor
}

// GetOnlyForOk returns a tuple with the OnlyFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsDelegatesEnduser) GetOnlyForOk() ([]DelegateScope, bool) {
	if o == nil || IsNil(o.OnlyFor) {
		return nil, false
	}
	return o.OnlyFor, true
}

// HasOnlyFor returns a boolean if a field has been set.
func (o *OrgSettingsDelegatesEnduser) HasOnlyFor() bool {
	if o != nil && !IsNil(o.OnlyFor) {
		return true
	}

	return false
}

// SetOnlyFor gets a reference to the given []DelegateScope and assigns it to the OnlyFor field.
func (o *OrgSettingsDelegatesEnduser) SetOnlyFor(v []DelegateScope) {
	o.OnlyFor = v
}

func (o OrgSettingsDelegatesEnduser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsDelegatesEnduser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.OnlyFor) {
		toSerialize["onlyFor"] = o.OnlyFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsDelegatesEnduser) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsDelegatesEnduser := _OrgSettingsDelegatesEnduser{}

	err = json.Unmarshal(data, &varOrgSettingsDelegatesEnduser)

	if err != nil {
		return err
	}

	*o = OrgSettingsDelegatesEnduser(varOrgSettingsDelegatesEnduser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "onlyFor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsDelegatesEnduser struct {
	value *OrgSettingsDelegatesEnduser
	isSet bool
}

func (v NullableOrgSettingsDelegatesEnduser) Get() *OrgSettingsDelegatesEnduser {
	return v.value
}

func (v *NullableOrgSettingsDelegatesEnduser) Set(val *OrgSettingsDelegatesEnduser) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsDelegatesEnduser) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsDelegatesEnduser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsDelegatesEnduser(val *OrgSettingsDelegatesEnduser) *NullableOrgSettingsDelegatesEnduser {
	return &NullableOrgSettingsDelegatesEnduser{value: val, isSet: true}
}

func (v NullableOrgSettingsDelegatesEnduser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsDelegatesEnduser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
