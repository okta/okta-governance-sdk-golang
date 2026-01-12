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
	"fmt"
	"time"
)

// checks if the SecurityAccessReviewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewRequest{}

// SecurityAccessReviewRequest The request to create a security access review. This includes the principal ID, the name of the security access review, the end time, and the reviewer settings. The end time defaults to seven days from the start time if not specified.
type SecurityAccessReviewRequest struct {
	// The Okta user ID in the security access review
	PrincipalId string `json:"principalId"`
	// The name of the security access review
	Name string `json:"name"`
	// The date and time when the security access review closes, defaulting to seven days after the creation of the security access review. It must be at least one day and less than six months after creation.
	EndTime              *time.Time                           `json:"endTime,omitempty"`
	ReviewerSettings     SecurityAccessReviewReviewerSettings `json:"reviewerSettings"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewRequest SecurityAccessReviewRequest

// NewSecurityAccessReviewRequest instantiates a new SecurityAccessReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewRequest(principalId string, name string, reviewerSettings SecurityAccessReviewReviewerSettings) *SecurityAccessReviewRequest {
	this := SecurityAccessReviewRequest{}
	this.PrincipalId = principalId
	this.Name = name
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

// GetName returns the Name field value
func (o *SecurityAccessReviewRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewRequest) SetName(v string) {
	o.Name = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SecurityAccessReviewRequest) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SecurityAccessReviewRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalId"] = o.PrincipalId
	toSerialize["name"] = o.Name
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	toSerialize["reviewerSettings"] = o.ReviewerSettings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalId",
		"name",
		"reviewerSettings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSecurityAccessReviewRequest := _SecurityAccessReviewRequest{}

	err = json.Unmarshal(data, &varSecurityAccessReviewRequest)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewRequest(varSecurityAccessReviewRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "reviewerSettings")
		o.AdditionalProperties = additionalProperties
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
