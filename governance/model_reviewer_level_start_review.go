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

// ReviewerLevelStartReview One can indicate, the rules for which the reviews can move to that level.
type ReviewerLevelStartReview struct {
	// The day on which that reviewer level will start.  It will be `0` for `FIRST` reviewer level, as the first level will start when the campaign starts.  For `SECOND` reviewer level specify a value greater than `0`. This will indicate the day, the reviews will be moved to second level.
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
	if o == nil || o.When == nil {
		var ret ReviewerLowerLevelCondition
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelStartReview) GetWhenOk() (*ReviewerLowerLevelCondition, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ReviewerLevelStartReview) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given ReviewerLowerLevelCondition and assigns it to the When field.
func (o *ReviewerLevelStartReview) SetWhen(v ReviewerLowerLevelCondition) {
	o.When = &v
}

func (o ReviewerLevelStartReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["onDay"] = o.OnDay
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewerLevelStartReview) UnmarshalJSON(bytes []byte) (err error) {
	varReviewerLevelStartReview := _ReviewerLevelStartReview{}

	err = json.Unmarshal(bytes, &varReviewerLevelStartReview)
	if err == nil {
		*o = ReviewerLevelStartReview(varReviewerLevelStartReview)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "onDay")
		delete(additionalProperties, "when")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
