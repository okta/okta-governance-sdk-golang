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

// checks if the OrgSettingsPatchableDelegates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsPatchableDelegates{}

// OrgSettingsPatchableDelegates Delegate settings
type OrgSettingsPatchableDelegates struct {
	Enduser              *OrgSettingsPatchableDelegatesEnduser `json:"enduser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsPatchableDelegates OrgSettingsPatchableDelegates

// NewOrgSettingsPatchableDelegates instantiates a new OrgSettingsPatchableDelegates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsPatchableDelegates() *OrgSettingsPatchableDelegates {
	this := OrgSettingsPatchableDelegates{}
	return &this
}

// NewOrgSettingsPatchableDelegatesWithDefaults instantiates a new OrgSettingsPatchableDelegates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsPatchableDelegatesWithDefaults() *OrgSettingsPatchableDelegates {
	this := OrgSettingsPatchableDelegates{}
	return &this
}

// GetEnduser returns the Enduser field value if set, zero value otherwise.
func (o *OrgSettingsPatchableDelegates) GetEnduser() OrgSettingsPatchableDelegatesEnduser {
	if o == nil || IsNil(o.Enduser) {
		var ret OrgSettingsPatchableDelegatesEnduser
		return ret
	}
	return *o.Enduser
}

// GetEnduserOk returns a tuple with the Enduser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsPatchableDelegates) GetEnduserOk() (*OrgSettingsPatchableDelegatesEnduser, bool) {
	if o == nil || IsNil(o.Enduser) {
		return nil, false
	}
	return o.Enduser, true
}

// HasEnduser returns a boolean if a field has been set.
func (o *OrgSettingsPatchableDelegates) HasEnduser() bool {
	if o != nil && !IsNil(o.Enduser) {
		return true
	}

	return false
}

// SetEnduser gets a reference to the given OrgSettingsPatchableDelegatesEnduser and assigns it to the Enduser field.
func (o *OrgSettingsPatchableDelegates) SetEnduser(v OrgSettingsPatchableDelegatesEnduser) {
	o.Enduser = &v
}

func (o OrgSettingsPatchableDelegates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsPatchableDelegates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enduser) {
		toSerialize["enduser"] = o.Enduser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsPatchableDelegates) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsPatchableDelegates := _OrgSettingsPatchableDelegates{}

	err = json.Unmarshal(data, &varOrgSettingsPatchableDelegates)

	if err != nil {
		return err
	}

	*o = OrgSettingsPatchableDelegates(varOrgSettingsPatchableDelegates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enduser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsPatchableDelegates struct {
	value *OrgSettingsPatchableDelegates
	isSet bool
}

func (v NullableOrgSettingsPatchableDelegates) Get() *OrgSettingsPatchableDelegates {
	return v.value
}

func (v *NullableOrgSettingsPatchableDelegates) Set(val *OrgSettingsPatchableDelegates) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsPatchableDelegates) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsPatchableDelegates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsPatchableDelegates(val *OrgSettingsPatchableDelegates) *NullableOrgSettingsPatchableDelegates {
	return &NullableOrgSettingsPatchableDelegates{value: val, isSet: true}
}

func (v NullableOrgSettingsPatchableDelegates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsPatchableDelegates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
