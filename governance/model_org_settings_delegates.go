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

// checks if the OrgSettingsDelegates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsDelegates{}

// OrgSettingsDelegates Delegate settings
type OrgSettingsDelegates struct {
	Enduser              *OrgSettingsDelegatesEnduser `json:"enduser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsDelegates OrgSettingsDelegates

// NewOrgSettingsDelegates instantiates a new OrgSettingsDelegates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsDelegates() *OrgSettingsDelegates {
	this := OrgSettingsDelegates{}
	return &this
}

// NewOrgSettingsDelegatesWithDefaults instantiates a new OrgSettingsDelegates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsDelegatesWithDefaults() *OrgSettingsDelegates {
	this := OrgSettingsDelegates{}
	return &this
}

// GetEnduser returns the Enduser field value if set, zero value otherwise.
func (o *OrgSettingsDelegates) GetEnduser() OrgSettingsDelegatesEnduser {
	if o == nil || IsNil(o.Enduser) {
		var ret OrgSettingsDelegatesEnduser
		return ret
	}
	return *o.Enduser
}

// GetEnduserOk returns a tuple with the Enduser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsDelegates) GetEnduserOk() (*OrgSettingsDelegatesEnduser, bool) {
	if o == nil || IsNil(o.Enduser) {
		return nil, false
	}
	return o.Enduser, true
}

// HasEnduser returns a boolean if a field has been set.
func (o *OrgSettingsDelegates) HasEnduser() bool {
	if o != nil && !IsNil(o.Enduser) {
		return true
	}

	return false
}

// SetEnduser gets a reference to the given OrgSettingsDelegatesEnduser and assigns it to the Enduser field.
func (o *OrgSettingsDelegates) SetEnduser(v OrgSettingsDelegatesEnduser) {
	o.Enduser = &v
}

func (o OrgSettingsDelegates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsDelegates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enduser) {
		toSerialize["enduser"] = o.Enduser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsDelegates) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsDelegates := _OrgSettingsDelegates{}

	err = json.Unmarshal(data, &varOrgSettingsDelegates)

	if err != nil {
		return err
	}

	*o = OrgSettingsDelegates(varOrgSettingsDelegates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enduser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsDelegates struct {
	value *OrgSettingsDelegates
	isSet bool
}

func (v NullableOrgSettingsDelegates) Get() *OrgSettingsDelegates {
	return v.value
}

func (v *NullableOrgSettingsDelegates) Set(val *OrgSettingsDelegates) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsDelegates) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsDelegates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsDelegates(val *OrgSettingsDelegates) *NullableOrgSettingsDelegates {
	return &NullableOrgSettingsDelegates{value: val, isSet: true}
}

func (v NullableOrgSettingsDelegates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsDelegates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
