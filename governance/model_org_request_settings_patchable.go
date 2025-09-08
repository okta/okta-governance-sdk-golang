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

// OrgRequestSettingsPatchable Request settings for the org for patch
type OrgRequestSettingsPatchable struct {
	// Whether a customer has acknowledged Access Requests subprocessors
	SubprocessorsAcknowledged bool `json:"subprocessorsAcknowledged"`
	AdditionalProperties      map[string]interface{}
}

type _OrgRequestSettingsPatchable OrgRequestSettingsPatchable

// NewOrgRequestSettingsPatchable instantiates a new OrgRequestSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgRequestSettingsPatchable(subprocessorsAcknowledged bool) *OrgRequestSettingsPatchable {
	this := OrgRequestSettingsPatchable{}
	this.SubprocessorsAcknowledged = subprocessorsAcknowledged
	return &this
}

// NewOrgRequestSettingsPatchableWithDefaults instantiates a new OrgRequestSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgRequestSettingsPatchableWithDefaults() *OrgRequestSettingsPatchable {
	this := OrgRequestSettingsPatchable{}
	return &this
}

// GetSubprocessorsAcknowledged returns the SubprocessorsAcknowledged field value
func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SubprocessorsAcknowledged
}

// GetSubprocessorsAcknowledgedOk returns a tuple with the SubprocessorsAcknowledged field value
// and a boolean to check if the value has been set.
func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledgedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubprocessorsAcknowledged, true
}

// SetSubprocessorsAcknowledged sets field value
func (o *OrgRequestSettingsPatchable) SetSubprocessorsAcknowledged(v bool) {
	o.SubprocessorsAcknowledged = v
}

func (o OrgRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subprocessorsAcknowledged"] = o.SubprocessorsAcknowledged
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgRequestSettingsPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varOrgRequestSettingsPatchable := _OrgRequestSettingsPatchable{}

	err = json.Unmarshal(bytes, &varOrgRequestSettingsPatchable)
	if err == nil {
		*o = OrgRequestSettingsPatchable(varOrgRequestSettingsPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "subprocessorsAcknowledged")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgRequestSettingsPatchable struct {
	value *OrgRequestSettingsPatchable
	isSet bool
}

func (v NullableOrgRequestSettingsPatchable) Get() *OrgRequestSettingsPatchable {
	return v.value
}

func (v *NullableOrgRequestSettingsPatchable) Set(val *OrgRequestSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgRequestSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgRequestSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgRequestSettingsPatchable(val *OrgRequestSettingsPatchable) *NullableOrgRequestSettingsPatchable {
	return &NullableOrgRequestSettingsPatchable{value: val, isSet: true}
}

func (v NullableOrgRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgRequestSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
