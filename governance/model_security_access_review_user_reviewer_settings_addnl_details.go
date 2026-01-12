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

// checks if the SecurityAccessReviewUserReviewerSettingsAddnlDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewUserReviewerSettingsAddnlDetails{}

// SecurityAccessReviewUserReviewerSettingsAddnlDetails Holds reviewer settings in the created security access review and more detailed reviewer details involved in the security access review.
type SecurityAccessReviewUserReviewerSettingsAddnlDetails struct {
	// The list of reviewer user IDs for the security access review
	IncludedUserIds []string `json:"includedUserIds,omitempty"`
	// The list of user reviewers for the security access review
	IncludedUserProfiles []PrincipalProfileEnriched `json:"includedUserProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewUserReviewerSettingsAddnlDetails SecurityAccessReviewUserReviewerSettingsAddnlDetails

// NewSecurityAccessReviewUserReviewerSettingsAddnlDetails instantiates a new SecurityAccessReviewUserReviewerSettingsAddnlDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewUserReviewerSettingsAddnlDetails() *SecurityAccessReviewUserReviewerSettingsAddnlDetails {
	this := SecurityAccessReviewUserReviewerSettingsAddnlDetails{}
	return &this
}

// NewSecurityAccessReviewUserReviewerSettingsAddnlDetailsWithDefaults instantiates a new SecurityAccessReviewUserReviewerSettingsAddnlDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewUserReviewerSettingsAddnlDetailsWithDefaults() *SecurityAccessReviewUserReviewerSettingsAddnlDetails {
	this := SecurityAccessReviewUserReviewerSettingsAddnlDetails{}
	return &this
}

// GetIncludedUserIds returns the IncludedUserIds field value if set, zero value otherwise.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserIds() []string {
	if o == nil || IsNil(o.IncludedUserIds) {
		var ret []string
		return ret
	}
	return o.IncludedUserIds
}

// GetIncludedUserIdsOk returns a tuple with the IncludedUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUserIds) {
		return nil, false
	}
	return o.IncludedUserIds, true
}

// HasIncludedUserIds returns a boolean if a field has been set.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) HasIncludedUserIds() bool {
	if o != nil && !IsNil(o.IncludedUserIds) {
		return true
	}

	return false
}

// SetIncludedUserIds gets a reference to the given []string and assigns it to the IncludedUserIds field.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) SetIncludedUserIds(v []string) {
	o.IncludedUserIds = v
}

// GetIncludedUserProfiles returns the IncludedUserProfiles field value if set, zero value otherwise.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserProfiles() []PrincipalProfileEnriched {
	if o == nil || IsNil(o.IncludedUserProfiles) {
		var ret []PrincipalProfileEnriched
		return ret
	}
	return o.IncludedUserProfiles
}

// GetIncludedUserProfilesOk returns a tuple with the IncludedUserProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserProfilesOk() ([]PrincipalProfileEnriched, bool) {
	if o == nil || IsNil(o.IncludedUserProfiles) {
		return nil, false
	}
	return o.IncludedUserProfiles, true
}

// HasIncludedUserProfiles returns a boolean if a field has been set.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) HasIncludedUserProfiles() bool {
	if o != nil && !IsNil(o.IncludedUserProfiles) {
		return true
	}

	return false
}

// SetIncludedUserProfiles gets a reference to the given []PrincipalProfileEnriched and assigns it to the IncludedUserProfiles field.
func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) SetIncludedUserProfiles(v []PrincipalProfileEnriched) {
	o.IncludedUserProfiles = v
}

func (o SecurityAccessReviewUserReviewerSettingsAddnlDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewUserReviewerSettingsAddnlDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludedUserIds) {
		toSerialize["includedUserIds"] = o.IncludedUserIds
	}
	if !IsNil(o.IncludedUserProfiles) {
		toSerialize["includedUserProfiles"] = o.IncludedUserProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviewUserReviewerSettingsAddnlDetails := _SecurityAccessReviewUserReviewerSettingsAddnlDetails{}

	err = json.Unmarshal(data, &varSecurityAccessReviewUserReviewerSettingsAddnlDetails)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewUserReviewerSettingsAddnlDetails(varSecurityAccessReviewUserReviewerSettingsAddnlDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "includedUserIds")
		delete(additionalProperties, "includedUserProfiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails struct {
	value *SecurityAccessReviewUserReviewerSettingsAddnlDetails
	isSet bool
}

func (v NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) Get() *SecurityAccessReviewUserReviewerSettingsAddnlDetails {
	return v.value
}

func (v *NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) Set(val *SecurityAccessReviewUserReviewerSettingsAddnlDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewUserReviewerSettingsAddnlDetails(val *SecurityAccessReviewUserReviewerSettingsAddnlDetails) *NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails {
	return &NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewUserReviewerSettingsAddnlDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
