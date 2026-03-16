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

// checks if the OrgSettingsEscalations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsEscalations{}

// OrgSettingsEscalations Escalation settings
type OrgSettingsEscalations struct {
	AccessRequests       *OrgSettingsEscalationsAccessRequests `json:"accessRequests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsEscalations OrgSettingsEscalations

// NewOrgSettingsEscalations instantiates a new OrgSettingsEscalations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsEscalations() *OrgSettingsEscalations {
	this := OrgSettingsEscalations{}
	return &this
}

// NewOrgSettingsEscalationsWithDefaults instantiates a new OrgSettingsEscalations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsEscalationsWithDefaults() *OrgSettingsEscalations {
	this := OrgSettingsEscalations{}
	return &this
}

// GetAccessRequests returns the AccessRequests field value if set, zero value otherwise.
func (o *OrgSettingsEscalations) GetAccessRequests() OrgSettingsEscalationsAccessRequests {
	if o == nil || IsNil(o.AccessRequests) {
		var ret OrgSettingsEscalationsAccessRequests
		return ret
	}
	return *o.AccessRequests
}

// GetAccessRequestsOk returns a tuple with the AccessRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsEscalations) GetAccessRequestsOk() (*OrgSettingsEscalationsAccessRequests, bool) {
	if o == nil || IsNil(o.AccessRequests) {
		return nil, false
	}
	return o.AccessRequests, true
}

// HasAccessRequests returns a boolean if a field has been set.
func (o *OrgSettingsEscalations) HasAccessRequests() bool {
	if o != nil && !IsNil(o.AccessRequests) {
		return true
	}

	return false
}

// SetAccessRequests gets a reference to the given OrgSettingsEscalationsAccessRequests and assigns it to the AccessRequests field.
func (o *OrgSettingsEscalations) SetAccessRequests(v OrgSettingsEscalationsAccessRequests) {
	o.AccessRequests = &v
}

func (o OrgSettingsEscalations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsEscalations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessRequests) {
		toSerialize["accessRequests"] = o.AccessRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsEscalations) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsEscalations := _OrgSettingsEscalations{}

	err = json.Unmarshal(data, &varOrgSettingsEscalations)

	if err != nil {
		return err
	}

	*o = OrgSettingsEscalations(varOrgSettingsEscalations)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessRequests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsEscalations struct {
	value *OrgSettingsEscalations
	isSet bool
}

func (v NullableOrgSettingsEscalations) Get() *OrgSettingsEscalations {
	return v.value
}

func (v *NullableOrgSettingsEscalations) Set(val *OrgSettingsEscalations) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsEscalations) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsEscalations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsEscalations(val *OrgSettingsEscalations) *NullableOrgSettingsEscalations {
	return &NullableOrgSettingsEscalations{value: val, isSet: true}
}

func (v NullableOrgSettingsEscalations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsEscalations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
