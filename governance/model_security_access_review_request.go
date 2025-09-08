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

// SecurityAccessReviewRequest The request to create a security access review. This includes the principal ID, the name of the security access review, the end time, and the reviewer settings. The end time defaults to seven days from the start time if not specified.
type SecurityAccessReviewRequest struct {
	// The Okta user ID in the security access review
	PrincipalId string `json:"principalId"`
	// The name of the security access review
	SecurityAccessReviewName string `json:"securityAccessReviewName"`
	// The end time of the security access review. If not specified this defaults to seven days from the start time.
	EndTime              *time.Time                           `json:"endTime,omitempty"`
	ReviewerSettings     SecurityAccessReviewReviewerSettings `json:"reviewerSettings"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewRequest SecurityAccessReviewRequest

// NewSecurityAccessReviewRequest instantiates a new SecurityAccessReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewRequest(principalId string, securityAccessReviewName string, reviewerSettings SecurityAccessReviewReviewerSettings) *SecurityAccessReviewRequest {
	this := SecurityAccessReviewRequest{}
	this.PrincipalId = principalId
	this.SecurityAccessReviewName = securityAccessReviewName
	this.ReviewerSettings = reviewerSettings
	return &this
}

// NewSecurityAccessReviewRequestWithDefaults instantiates a new SecurityAccessReviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewRequestWithDefaults() *SecurityAccessReviewRequest {
	this := SecurityAccessReviewRequest{}
	return &this
}

// GetPrincipalId returns the PrincipalId field value
func (o *SecurityAccessReviewRequest) GetPrincipalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetPrincipalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalId, true
}

// SetPrincipalId sets field value
func (o *SecurityAccessReviewRequest) SetPrincipalId(v string) {
	o.PrincipalId = v
}

// GetSecurityAccessReviewName returns the SecurityAccessReviewName field value
func (o *SecurityAccessReviewRequest) GetSecurityAccessReviewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityAccessReviewName
}

// GetSecurityAccessReviewNameOk returns a tuple with the SecurityAccessReviewName field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetSecurityAccessReviewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityAccessReviewName, true
}

// SetSecurityAccessReviewName sets field value
func (o *SecurityAccessReviewRequest) SetSecurityAccessReviewName(v string) {
	o.SecurityAccessReviewName = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SecurityAccessReviewRequest) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SecurityAccessReviewRequest) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SecurityAccessReviewRequest) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *SecurityAccessReviewRequest) GetReviewerSettings() SecurityAccessReviewReviewerSettings {
	if o == nil {
		var ret SecurityAccessReviewReviewerSettings
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *SecurityAccessReviewRequest) SetReviewerSettings(v SecurityAccessReviewReviewerSettings) {
	o.ReviewerSettings = v
}

func (o SecurityAccessReviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["principalId"] = o.PrincipalId
	}
	if true {
		toSerialize["securityAccessReviewName"] = o.SecurityAccessReviewName
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["reviewerSettings"] = o.ReviewerSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewRequest := _SecurityAccessReviewRequest{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewRequest)
	if err == nil {
		*o = SecurityAccessReviewRequest(varSecurityAccessReviewRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "principalId")
		delete(additionalProperties, "securityAccessReviewName")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "reviewerSettings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewRequest struct {
	value *SecurityAccessReviewRequest
	isSet bool
}

func (v NullableSecurityAccessReviewRequest) Get() *SecurityAccessReviewRequest {
	return v.value
}

func (v *NullableSecurityAccessReviewRequest) Set(val *SecurityAccessReviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewRequest(val *SecurityAccessReviewRequest) *NullableSecurityAccessReviewRequest {
	return &NullableSecurityAccessReviewRequest{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
