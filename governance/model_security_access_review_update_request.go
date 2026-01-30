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
	"time"
)

// checks if the SecurityAccessReviewUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewUpdateRequest{}

// SecurityAccessReviewUpdateRequest The request to update a security access review. This includes the end time and the reviewer settings.
type SecurityAccessReviewUpdateRequest struct {
	// The date and time when the security access review closes.
	EndTime              *time.Time                            `json:"endTime,omitempty"`
	ReviewerSettings     *SecurityAccessReviewReviewerSettings `json:"reviewerSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewUpdateRequest SecurityAccessReviewUpdateRequest

// NewSecurityAccessReviewUpdateRequest instantiates a new SecurityAccessReviewUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewUpdateRequest() *SecurityAccessReviewUpdateRequest {
	this := SecurityAccessReviewUpdateRequest{}
	return &this
}

// NewSecurityAccessReviewUpdateRequestWithDefaults instantiates a new SecurityAccessReviewUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewUpdateRequestWithDefaults() *SecurityAccessReviewUpdateRequest {
	this := SecurityAccessReviewUpdateRequest{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SecurityAccessReviewUpdateRequest) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewUpdateRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SecurityAccessReviewUpdateRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SecurityAccessReviewUpdateRequest) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetReviewerSettings returns the ReviewerSettings field value if set, zero value otherwise.
func (o *SecurityAccessReviewUpdateRequest) GetReviewerSettings() SecurityAccessReviewReviewerSettings {
	if o == nil || IsNil(o.ReviewerSettings) {
		var ret SecurityAccessReviewReviewerSettings
		return ret
	}
	return *o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewUpdateRequest) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool) {
	if o == nil || IsNil(o.ReviewerSettings) {
		return nil, false
	}
	return o.ReviewerSettings, true
}

// HasReviewerSettings returns a boolean if a field has been set.
func (o *SecurityAccessReviewUpdateRequest) HasReviewerSettings() bool {
	if o != nil && !IsNil(o.ReviewerSettings) {
		return true
	}

	return false
}

// SetReviewerSettings gets a reference to the given SecurityAccessReviewReviewerSettings and assigns it to the ReviewerSettings field.
func (o *SecurityAccessReviewUpdateRequest) SetReviewerSettings(v SecurityAccessReviewReviewerSettings) {
	o.ReviewerSettings = &v
}

func (o SecurityAccessReviewUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.ReviewerSettings) {
		toSerialize["reviewerSettings"] = o.ReviewerSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviewUpdateRequest := _SecurityAccessReviewUpdateRequest{}

	err = json.Unmarshal(data, &varSecurityAccessReviewUpdateRequest)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewUpdateRequest(varSecurityAccessReviewUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "reviewerSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewUpdateRequest struct {
	value *SecurityAccessReviewUpdateRequest
	isSet bool
}

func (v NullableSecurityAccessReviewUpdateRequest) Get() *SecurityAccessReviewUpdateRequest {
	return v.value
}

func (v *NullableSecurityAccessReviewUpdateRequest) Set(val *SecurityAccessReviewUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewUpdateRequest(val *SecurityAccessReviewUpdateRequest) *NullableSecurityAccessReviewUpdateRequest {
	return &NullableSecurityAccessReviewUpdateRequest{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
