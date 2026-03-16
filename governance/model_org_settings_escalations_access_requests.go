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

// checks if the OrgSettingsEscalationsAccessRequests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsEscalationsAccessRequests{}

// OrgSettingsEscalationsAccessRequests Access request escalation settings
type OrgSettingsEscalationsAccessRequests struct {
	// Indicates if access request escalations are enabled for end users. If `true`, end users can escalate requests to higher-level approvers.
	EndUserCanEscalate   *bool `json:"endUserCanEscalate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsEscalationsAccessRequests OrgSettingsEscalationsAccessRequests

// NewOrgSettingsEscalationsAccessRequests instantiates a new OrgSettingsEscalationsAccessRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsEscalationsAccessRequests() *OrgSettingsEscalationsAccessRequests {
	this := OrgSettingsEscalationsAccessRequests{}
	return &this
}

// NewOrgSettingsEscalationsAccessRequestsWithDefaults instantiates a new OrgSettingsEscalationsAccessRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsEscalationsAccessRequestsWithDefaults() *OrgSettingsEscalationsAccessRequests {
	this := OrgSettingsEscalationsAccessRequests{}
	return &this
}

// GetEndUserCanEscalate returns the EndUserCanEscalate field value if set, zero value otherwise.
func (o *OrgSettingsEscalationsAccessRequests) GetEndUserCanEscalate() bool {
	if o == nil || IsNil(o.EndUserCanEscalate) {
		var ret bool
		return ret
	}
	return *o.EndUserCanEscalate
}

// GetEndUserCanEscalateOk returns a tuple with the EndUserCanEscalate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsEscalationsAccessRequests) GetEndUserCanEscalateOk() (*bool, bool) {
	if o == nil || IsNil(o.EndUserCanEscalate) {
		return nil, false
	}
	return o.EndUserCanEscalate, true
}

// HasEndUserCanEscalate returns a boolean if a field has been set.
func (o *OrgSettingsEscalationsAccessRequests) HasEndUserCanEscalate() bool {
	if o != nil && !IsNil(o.EndUserCanEscalate) {
		return true
	}

	return false
}

// SetEndUserCanEscalate gets a reference to the given bool and assigns it to the EndUserCanEscalate field.
func (o *OrgSettingsEscalationsAccessRequests) SetEndUserCanEscalate(v bool) {
	o.EndUserCanEscalate = &v
}

func (o OrgSettingsEscalationsAccessRequests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsEscalationsAccessRequests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndUserCanEscalate) {
		toSerialize["endUserCanEscalate"] = o.EndUserCanEscalate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsEscalationsAccessRequests) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsEscalationsAccessRequests := _OrgSettingsEscalationsAccessRequests{}

	err = json.Unmarshal(data, &varOrgSettingsEscalationsAccessRequests)

	if err != nil {
		return err
	}

	*o = OrgSettingsEscalationsAccessRequests(varOrgSettingsEscalationsAccessRequests)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endUserCanEscalate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsEscalationsAccessRequests struct {
	value *OrgSettingsEscalationsAccessRequests
	isSet bool
}

func (v NullableOrgSettingsEscalationsAccessRequests) Get() *OrgSettingsEscalationsAccessRequests {
	return v.value
}

func (v *NullableOrgSettingsEscalationsAccessRequests) Set(val *OrgSettingsEscalationsAccessRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsEscalationsAccessRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsEscalationsAccessRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsEscalationsAccessRequests(val *OrgSettingsEscalationsAccessRequests) *NullableOrgSettingsEscalationsAccessRequests {
	return &NullableOrgSettingsEscalationsAccessRequests{value: val, isSet: true}
}

func (v NullableOrgSettingsEscalationsAccessRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsEscalationsAccessRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
