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

// checks if the ReviewerLevelStartReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerLevelStartReview{}

// ReviewerLevelStartReview Indicates rules for starting reviews at this level
type ReviewerLevelStartReview struct {
	// The day when reviewer level starts: 1. For the first level, this value is always `0` since the first level starts when the campaign starts. 2. For the second level, specify a value that's greater than `0`. This indicates the day when the reviews move to the second level.
	OnDay                int32                        `json:"onDay"`
	When                 *ReviewerLowerLevelCondition `json:"when,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerLevelStartReview ReviewerLevelStartReview

// NewReviewerLevelStartReview instantiates a new ReviewerLevelStartReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerLevelStartReview(onDay int32) *ReviewerLevelStartReview {
	this := ReviewerLevelStartReview{}
	this.OnDay = onDay
	return &this
}

// NewReviewerLevelStartReviewWithDefaults instantiates a new ReviewerLevelStartReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerLevelStartReviewWithDefaults() *ReviewerLevelStartReview {
	this := ReviewerLevelStartReview{}
	var onDay int32 = 0
	this.OnDay = onDay
	return &this
}

// GetOnDay returns the OnDay field value
func (o *ReviewerLevelStartReview) GetOnDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OnDay
}

// GetOnDayOk returns a tuple with the OnDay field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelStartReview) GetOnDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnDay, true
}

// SetOnDay sets field value
func (o *ReviewerLevelStartReview) SetOnDay(v int32) {
	o.OnDay = v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *ReviewerLevelStartReview) GetWhen() ReviewerLowerLevelCondition {
	if o == nil || IsNil(o.When) {
		var ret ReviewerLowerLevelCondition
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelStartReview) GetWhenOk() (*ReviewerLowerLevelCondition, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ReviewerLevelStartReview) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given ReviewerLowerLevelCondition and assigns it to the When field.
func (o *ReviewerLevelStartReview) SetWhen(v ReviewerLowerLevelCondition) {
	o.When = &v
}

func (o ReviewerLevelStartReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerLevelStartReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["onDay"] = o.OnDay
	if !IsNil(o.When) {
		toSerialize["when"] = o.When
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerLevelStartReview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"onDay",
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

	varReviewerLevelStartReview := _ReviewerLevelStartReview{}

	err = json.Unmarshal(data, &varReviewerLevelStartReview)

	if err != nil {
		return err
	}

	*o = ReviewerLevelStartReview(varReviewerLevelStartReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "onDay")
		delete(additionalProperties, "when")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerLevelStartReview struct {
	value *ReviewerLevelStartReview
	isSet bool
}

func (v NullableReviewerLevelStartReview) Get() *ReviewerLevelStartReview {
	return v.value
}

func (v *NullableReviewerLevelStartReview) Set(val *ReviewerLevelStartReview) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerLevelStartReview) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerLevelStartReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerLevelStartReview(val *ReviewerLevelStartReview) *NullableReviewerLevelStartReview {
	return &NullableReviewerLevelStartReview{value: val, isSet: true}
}

func (v NullableReviewerLevelStartReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerLevelStartReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
