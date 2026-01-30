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
)

// checks if the SecurityAccessReviewsStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewsStats{}

// SecurityAccessReviewsStats struct for SecurityAccessReviewsStats
type SecurityAccessReviewsStats struct {
	// The number of active security access reviews
	ActiveCount int32 `json:"activeCount"`
	// The number of pending security access reviews
	PendingCount int32 `json:"pendingCount"`
	// The number of errored security access reviews
	ErrorCount int32 `json:"errorCount"`
	// The number of closed security access reviews
	ClosedCount          int32 `json:"closedCount"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewsStats SecurityAccessReviewsStats

// NewSecurityAccessReviewsStats instantiates a new SecurityAccessReviewsStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewsStats(activeCount int32, pendingCount int32, errorCount int32, closedCount int32) *SecurityAccessReviewsStats {
	this := SecurityAccessReviewsStats{}
	this.ActiveCount = activeCount
	this.PendingCount = pendingCount
	this.ErrorCount = errorCount
	this.ClosedCount = closedCount
	return &this
}

// NewSecurityAccessReviewsStatsWithDefaults instantiates a new SecurityAccessReviewsStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewsStatsWithDefaults() *SecurityAccessReviewsStats {
	this := SecurityAccessReviewsStats{}
	return &this
}

// GetActiveCount returns the ActiveCount field value
func (o *SecurityAccessReviewsStats) GetActiveCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveCount
}

// GetActiveCountOk returns a tuple with the ActiveCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewsStats) GetActiveCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveCount, true
}

// SetActiveCount sets field value
func (o *SecurityAccessReviewsStats) SetActiveCount(v int32) {
	o.ActiveCount = v
}

// GetPendingCount returns the PendingCount field value
func (o *SecurityAccessReviewsStats) GetPendingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewsStats) GetPendingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingCount, true
}

// SetPendingCount sets field value
func (o *SecurityAccessReviewsStats) SetPendingCount(v int32) {
	o.PendingCount = v
}

// GetErrorCount returns the ErrorCount field value
func (o *SecurityAccessReviewsStats) GetErrorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewsStats) GetErrorCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCount, true
}

// SetErrorCount sets field value
func (o *SecurityAccessReviewsStats) SetErrorCount(v int32) {
	o.ErrorCount = v
}

// GetClosedCount returns the ClosedCount field value
func (o *SecurityAccessReviewsStats) GetClosedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClosedCount
}

// GetClosedCountOk returns a tuple with the ClosedCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewsStats) GetClosedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClosedCount, true
}

// SetClosedCount sets field value
func (o *SecurityAccessReviewsStats) SetClosedCount(v int32) {
	o.ClosedCount = v
}

func (o SecurityAccessReviewsStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewsStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeCount"] = o.ActiveCount
	toSerialize["pendingCount"] = o.PendingCount
	toSerialize["errorCount"] = o.ErrorCount
	toSerialize["closedCount"] = o.ClosedCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewsStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activeCount",
		"pendingCount",
		"errorCount",
		"closedCount",
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

	varSecurityAccessReviewsStats := _SecurityAccessReviewsStats{}

	err = json.Unmarshal(data, &varSecurityAccessReviewsStats)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewsStats(varSecurityAccessReviewsStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activeCount")
		delete(additionalProperties, "pendingCount")
		delete(additionalProperties, "errorCount")
		delete(additionalProperties, "closedCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewsStats struct {
	value *SecurityAccessReviewsStats
	isSet bool
}

func (v NullableSecurityAccessReviewsStats) Get() *SecurityAccessReviewsStats {
	return v.value
}

func (v *NullableSecurityAccessReviewsStats) Set(val *SecurityAccessReviewsStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewsStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewsStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewsStats(val *SecurityAccessReviewsStats) *NullableSecurityAccessReviewsStats {
	return &NullableSecurityAccessReviewsStats{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewsStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewsStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
