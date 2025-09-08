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

// ReviewerGroupProfile Applicable only when `reviewerType` is `GROUP` or `RESOURCE_OWNER`. For all other reviewer type, this property can be ignored. It represents the group in charge of reviews.
type ReviewerGroupProfile struct {
	Name                 string    `json:"name"`
	GroupId              string    `json:"groupId"`
	GroupType            GroupType `json:"groupType"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerGroupProfile ReviewerGroupProfile

// NewReviewerGroupProfile instantiates a new ReviewerGroupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerGroupProfile(name string, groupId string, groupType GroupType) *ReviewerGroupProfile {
	this := ReviewerGroupProfile{}
	this.Name = name
	this.GroupId = groupId
	this.GroupType = groupType
	return &this
}

// NewReviewerGroupProfileWithDefaults instantiates a new ReviewerGroupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerGroupProfileWithDefaults() *ReviewerGroupProfile {
	this := ReviewerGroupProfile{}
	return &this
}

// GetName returns the Name field value
func (o *ReviewerGroupProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReviewerGroupProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReviewerGroupProfile) SetName(v string) {
	o.Name = v
}

// GetGroupId returns the GroupId field value
func (o *ReviewerGroupProfile) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ReviewerGroupProfile) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ReviewerGroupProfile) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupType returns the GroupType field value
func (o *ReviewerGroupProfile) GetGroupType() GroupType {
	if o == nil {
		var ret GroupType
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *ReviewerGroupProfile) GetGroupTypeOk() (*GroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *ReviewerGroupProfile) SetGroupType(v GroupType) {
	o.GroupType = v
}

func (o ReviewerGroupProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["groupType"] = o.GroupType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewerGroupProfile) UnmarshalJSON(bytes []byte) (err error) {
	varReviewerGroupProfile := _ReviewerGroupProfile{}

	err = json.Unmarshal(bytes, &varReviewerGroupProfile)
	if err == nil {
		*o = ReviewerGroupProfile(varReviewerGroupProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "groupType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewerGroupProfile struct {
	value *ReviewerGroupProfile
	isSet bool
}

func (v NullableReviewerGroupProfile) Get() *ReviewerGroupProfile {
	return v.value
}

func (v *NullableReviewerGroupProfile) Set(val *ReviewerGroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerGroupProfile(val *ReviewerGroupProfile) *NullableReviewerGroupProfile {
	return &NullableReviewerGroupProfile{value: val, isSet: true}
}

func (v NullableReviewerGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
