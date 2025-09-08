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
	"time"
)

// SecurityAccessReviewSubAccessItemGroupInfo struct for SecurityAccessReviewSubAccessItemGroupInfo
type SecurityAccessReviewSubAccessItemGroupInfo struct {
	// A brief description of the group
	Description *string `json:"description,omitempty"`
	// The date the group was assigned to the user
	AssignedDate *time.Time `json:"assignedDate,omitempty"`
	// The group owner name
	GroupOwner     *string         `json:"groupOwner,omitempty"`
	AssignmentType *AssignmentType `json:"assignmentType,omitempty"`
	// The last time the user/group pair was reviewed in an access certification campaign
	LastAccessCertificationReviewedDate *time.Time `json:"lastAccessCertificationReviewedDate,omitempty"`
	// The last time an action was taken on this group for the user in a security access review
	LastSecurityAccessReviewDate *time.Time `json:"lastSecurityAccessReviewDate,omitempty"`
	// All governance labels applied to the group
	GovernanceLabels     []GovernanceLabel `json:"governanceLabels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewSubAccessItemGroupInfo SecurityAccessReviewSubAccessItemGroupInfo

// NewSecurityAccessReviewSubAccessItemGroupInfo instantiates a new SecurityAccessReviewSubAccessItemGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSubAccessItemGroupInfo() *SecurityAccessReviewSubAccessItemGroupInfo {
	this := SecurityAccessReviewSubAccessItemGroupInfo{}
	return &this
}

// NewSecurityAccessReviewSubAccessItemGroupInfoWithDefaults instantiates a new SecurityAccessReviewSubAccessItemGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSubAccessItemGroupInfoWithDefaults() *SecurityAccessReviewSubAccessItemGroupInfo {
	this := SecurityAccessReviewSubAccessItemGroupInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetDescription(v string) {
	o.Description = &v
}

// GetAssignedDate returns the AssignedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignedDate() time.Time {
	if o == nil || o.AssignedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignedDate
}

// GetAssignedDateOk returns a tuple with the AssignedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignedDateOk() (*time.Time, bool) {
	if o == nil || o.AssignedDate == nil {
		return nil, false
	}
	return o.AssignedDate, true
}

// HasAssignedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasAssignedDate() bool {
	if o != nil && o.AssignedDate != nil {
		return true
	}

	return false
}

// SetAssignedDate gets a reference to the given time.Time and assigns it to the AssignedDate field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetAssignedDate(v time.Time) {
	o.AssignedDate = &v
}

// GetGroupOwner returns the GroupOwner field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGroupOwner() string {
	if o == nil || o.GroupOwner == nil {
		var ret string
		return ret
	}
	return *o.GroupOwner
}

// GetGroupOwnerOk returns a tuple with the GroupOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGroupOwnerOk() (*string, bool) {
	if o == nil || o.GroupOwner == nil {
		return nil, false
	}
	return o.GroupOwner, true
}

// HasGroupOwner returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasGroupOwner() bool {
	if o != nil && o.GroupOwner != nil {
		return true
	}

	return false
}

// SetGroupOwner gets a reference to the given string and assigns it to the GroupOwner field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetGroupOwner(v string) {
	o.GroupOwner = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignmentType() AssignmentType {
	if o == nil || o.AssignmentType == nil {
		var ret AssignmentType
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignmentTypeOk() (*AssignmentType, bool) {
	if o == nil || o.AssignmentType == nil {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasAssignmentType() bool {
	if o != nil && o.AssignmentType != nil {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given AssignmentType and assigns it to the AssignmentType field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetAssignmentType(v AssignmentType) {
	o.AssignmentType = &v
}

// GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastAccessCertificationReviewedDate() time.Time {
	if o == nil || o.LastAccessCertificationReviewedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessCertificationReviewedDate
}

// GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool) {
	if o == nil || o.LastAccessCertificationReviewedDate == nil {
		return nil, false
	}
	return o.LastAccessCertificationReviewedDate, true
}

// HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasLastAccessCertificationReviewedDate() bool {
	if o != nil && o.LastAccessCertificationReviewedDate != nil {
		return true
	}

	return false
}

// SetLastAccessCertificationReviewedDate gets a reference to the given time.Time and assigns it to the LastAccessCertificationReviewedDate field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetLastAccessCertificationReviewedDate(v time.Time) {
	o.LastAccessCertificationReviewedDate = &v
}

// GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastSecurityAccessReviewDate() time.Time {
	if o == nil || o.LastSecurityAccessReviewDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSecurityAccessReviewDate
}

// GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool) {
	if o == nil || o.LastSecurityAccessReviewDate == nil {
		return nil, false
	}
	return o.LastSecurityAccessReviewDate, true
}

// HasLastSecurityAccessReviewDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasLastSecurityAccessReviewDate() bool {
	if o != nil && o.LastSecurityAccessReviewDate != nil {
		return true
	}

	return false
}

// SetLastSecurityAccessReviewDate gets a reference to the given time.Time and assigns it to the LastSecurityAccessReviewDate field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetLastSecurityAccessReviewDate(v time.Time) {
	o.LastSecurityAccessReviewDate = &v
}

// GetGovernanceLabels returns the GovernanceLabels field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGovernanceLabels() []GovernanceLabel {
	if o == nil || o.GovernanceLabels == nil {
		var ret []GovernanceLabel
		return ret
	}
	return o.GovernanceLabels
}

// GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGovernanceLabelsOk() ([]GovernanceLabel, bool) {
	if o == nil || o.GovernanceLabels == nil {
		return nil, false
	}
	return o.GovernanceLabels, true
}

// HasGovernanceLabels returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasGovernanceLabels() bool {
	if o != nil && o.GovernanceLabels != nil {
		return true
	}

	return false
}

// SetGovernanceLabels gets a reference to the given []GovernanceLabel and assigns it to the GovernanceLabels field.
func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetGovernanceLabels(v []GovernanceLabel) {
	o.GovernanceLabels = v
}

func (o SecurityAccessReviewSubAccessItemGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AssignedDate != nil {
		toSerialize["assignedDate"] = o.AssignedDate
	}
	if o.GroupOwner != nil {
		toSerialize["groupOwner"] = o.GroupOwner
	}
	if o.AssignmentType != nil {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if o.LastAccessCertificationReviewedDate != nil {
		toSerialize["lastAccessCertificationReviewedDate"] = o.LastAccessCertificationReviewedDate
	}
	if o.LastSecurityAccessReviewDate != nil {
		toSerialize["lastSecurityAccessReviewDate"] = o.LastSecurityAccessReviewDate
	}
	if o.GovernanceLabels != nil {
		toSerialize["governanceLabels"] = o.GovernanceLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewSubAccessItemGroupInfo) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewSubAccessItemGroupInfo := _SecurityAccessReviewSubAccessItemGroupInfo{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewSubAccessItemGroupInfo)
	if err == nil {
		*o = SecurityAccessReviewSubAccessItemGroupInfo(varSecurityAccessReviewSubAccessItemGroupInfo)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "assignedDate")
		delete(additionalProperties, "groupOwner")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "lastAccessCertificationReviewedDate")
		delete(additionalProperties, "lastSecurityAccessReviewDate")
		delete(additionalProperties, "governanceLabels")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewSubAccessItemGroupInfo struct {
	value *SecurityAccessReviewSubAccessItemGroupInfo
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItemGroupInfo) Get() *SecurityAccessReviewSubAccessItemGroupInfo {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItemGroupInfo) Set(val *SecurityAccessReviewSubAccessItemGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItemGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItemGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItemGroupInfo(val *SecurityAccessReviewSubAccessItemGroupInfo) *NullableSecurityAccessReviewSubAccessItemGroupInfo {
	return &NullableSecurityAccessReviewSubAccessItemGroupInfo{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItemGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItemGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
