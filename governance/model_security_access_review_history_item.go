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

// checks if the SecurityAccessReviewHistoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewHistoryItem{}

// SecurityAccessReviewHistoryItem struct for SecurityAccessReviewHistoryItem
type SecurityAccessReviewHistoryItem struct {
	Id string `json:"id"`
	// Indicates if the action or change was made by Okta
	SystemGenerated bool `json:"systemGenerated"`
	// The date and time of the action or change
	Timestamp            time.Time         `json:"timestamp"`
	Message              string            `json:"message"`
	PrincipalProfile     *PrincipalProfile `json:"principalProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewHistoryItem SecurityAccessReviewHistoryItem

// NewSecurityAccessReviewHistoryItem instantiates a new SecurityAccessReviewHistoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewHistoryItem(id string, systemGenerated bool, timestamp time.Time, message string) *SecurityAccessReviewHistoryItem {
	this := SecurityAccessReviewHistoryItem{}
	this.Id = id
	this.SystemGenerated = systemGenerated
	this.Timestamp = timestamp
	this.Message = message
	return &this
}

// NewSecurityAccessReviewHistoryItemWithDefaults instantiates a new SecurityAccessReviewHistoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewHistoryItemWithDefaults() *SecurityAccessReviewHistoryItem {
	this := SecurityAccessReviewHistoryItem{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewHistoryItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewHistoryItem) SetId(v string) {
	o.Id = v
}

// GetSystemGenerated returns the SystemGenerated field value
func (o *SecurityAccessReviewHistoryItem) GetSystemGenerated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SystemGenerated
}

// GetSystemGeneratedOk returns a tuple with the SystemGenerated field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItem) GetSystemGeneratedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemGenerated, true
}

// SetSystemGenerated sets field value
func (o *SecurityAccessReviewHistoryItem) SetSystemGenerated(v bool) {
	o.SystemGenerated = v
}

// GetTimestamp returns the Timestamp field value
func (o *SecurityAccessReviewHistoryItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *SecurityAccessReviewHistoryItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetMessage returns the Message field value
func (o *SecurityAccessReviewHistoryItem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SecurityAccessReviewHistoryItem) SetMessage(v string) {
	o.Message = v
}

// GetPrincipalProfile returns the PrincipalProfile field value if set, zero value otherwise.
func (o *SecurityAccessReviewHistoryItem) GetPrincipalProfile() PrincipalProfile {
	if o == nil || IsNil(o.PrincipalProfile) {
		var ret PrincipalProfile
		return ret
	}
	return *o.PrincipalProfile
}

// GetPrincipalProfileOk returns a tuple with the PrincipalProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItem) GetPrincipalProfileOk() (*PrincipalProfile, bool) {
	if o == nil || IsNil(o.PrincipalProfile) {
		return nil, false
	}
	return o.PrincipalProfile, true
}

// HasPrincipalProfile returns a boolean if a field has been set.
func (o *SecurityAccessReviewHistoryItem) HasPrincipalProfile() bool {
	if o != nil && !IsNil(o.PrincipalProfile) {
		return true
	}

	return false
}

// SetPrincipalProfile gets a reference to the given PrincipalProfile and assigns it to the PrincipalProfile field.
func (o *SecurityAccessReviewHistoryItem) SetPrincipalProfile(v PrincipalProfile) {
	o.PrincipalProfile = &v
}

func (o SecurityAccessReviewHistoryItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewHistoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["systemGenerated"] = o.SystemGenerated
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["message"] = o.Message
	if !IsNil(o.PrincipalProfile) {
		toSerialize["principalProfile"] = o.PrincipalProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewHistoryItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"systemGenerated",
		"timestamp",
		"message",
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

	varSecurityAccessReviewHistoryItem := _SecurityAccessReviewHistoryItem{}

	err = json.Unmarshal(data, &varSecurityAccessReviewHistoryItem)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewHistoryItem(varSecurityAccessReviewHistoryItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "systemGenerated")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "message")
		delete(additionalProperties, "principalProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewHistoryItem struct {
	value *SecurityAccessReviewHistoryItem
	isSet bool
}

func (v NullableSecurityAccessReviewHistoryItem) Get() *SecurityAccessReviewHistoryItem {
	return v.value
}

func (v *NullableSecurityAccessReviewHistoryItem) Set(val *SecurityAccessReviewHistoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewHistoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewHistoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewHistoryItem(val *SecurityAccessReviewHistoryItem) *NullableSecurityAccessReviewHistoryItem {
	return &NullableSecurityAccessReviewHistoryItem{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewHistoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewHistoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
