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

// checks if the ReviewsReassign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewsReassign{}

// ReviewsReassign struct for ReviewsReassign
type ReviewsReassign struct {
	// The Okta user `id` of the new reviewer
	ReviewerId string `json:"reviewerId"`
	// A list of reviews (review `id` values) that are reassigned to the new reviewer.
	ReviewIds     []string           `json:"reviewIds"`
	ReviewerLevel *ReviewerLevelType `json:"reviewerLevel,omitempty"`
	// A note to justify the reassignment decision for the specified review(s).
	Note                 string `json:"note"`
	AdditionalProperties map[string]interface{}
}

type _ReviewsReassign ReviewsReassign

// NewReviewsReassign instantiates a new ReviewsReassign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewsReassign(reviewerId string, reviewIds []string, note string) *ReviewsReassign {
	this := ReviewsReassign{}
	this.ReviewerId = reviewerId
	this.ReviewIds = reviewIds
	this.Note = note
	return &this
}

// NewReviewsReassignWithDefaults instantiates a new ReviewsReassign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewsReassignWithDefaults() *ReviewsReassign {
	this := ReviewsReassign{}
	return &this
}

// GetReviewerId returns the ReviewerId field value
func (o *ReviewsReassign) GetReviewerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewerId
}

// GetReviewerIdOk returns a tuple with the ReviewerId field value
// and a boolean to check if the value has been set.
func (o *ReviewsReassign) GetReviewerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerId, true
}

// SetReviewerId sets field value
func (o *ReviewsReassign) SetReviewerId(v string) {
	o.ReviewerId = v
}

// GetReviewIds returns the ReviewIds field value
func (o *ReviewsReassign) GetReviewIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReviewIds
}

// GetReviewIdsOk returns a tuple with the ReviewIds field value
// and a boolean to check if the value has been set.
func (o *ReviewsReassign) GetReviewIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewIds, true
}

// SetReviewIds sets field value
func (o *ReviewsReassign) SetReviewIds(v []string) {
	o.ReviewIds = v
}

// GetReviewerLevel returns the ReviewerLevel field value if set, zero value otherwise.
func (o *ReviewsReassign) GetReviewerLevel() ReviewerLevelType {
	if o == nil || IsNil(o.ReviewerLevel) {
		var ret ReviewerLevelType
		return ret
	}
	return *o.ReviewerLevel
}

// GetReviewerLevelOk returns a tuple with the ReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewsReassign) GetReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || IsNil(o.ReviewerLevel) {
		return nil, false
	}
	return o.ReviewerLevel, true
}

// HasReviewerLevel returns a boolean if a field has been set.
func (o *ReviewsReassign) HasReviewerLevel() bool {
	if o != nil && !IsNil(o.ReviewerLevel) {
		return true
	}

	return false
}

// SetReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the ReviewerLevel field.
func (o *ReviewsReassign) SetReviewerLevel(v ReviewerLevelType) {
	o.ReviewerLevel = &v
}

// GetNote returns the Note field value
func (o *ReviewsReassign) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *ReviewsReassign) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *ReviewsReassign) SetNote(v string) {
	o.Note = v
}

func (o ReviewsReassign) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewsReassign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviewerId"] = o.ReviewerId
	toSerialize["reviewIds"] = o.ReviewIds
	if !IsNil(o.ReviewerLevel) {
		toSerialize["reviewerLevel"] = o.ReviewerLevel
	}
	toSerialize["note"] = o.Note

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewsReassign) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviewerId",
		"reviewIds",
		"note",
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

	varReviewsReassign := _ReviewsReassign{}

	err = json.Unmarshal(data, &varReviewsReassign)

	if err != nil {
		return err
	}

	*o = ReviewsReassign(varReviewsReassign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewerId")
		delete(additionalProperties, "reviewIds")
		delete(additionalProperties, "reviewerLevel")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewsReassign struct {
	value *ReviewsReassign
	isSet bool
}

func (v NullableReviewsReassign) Get() *ReviewsReassign {
	return v.value
}

func (v *NullableReviewsReassign) Set(val *ReviewsReassign) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewsReassign) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewsReassign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewsReassign(val *ReviewsReassign) *NullableReviewsReassign {
	return &NullableReviewsReassign{value: val, isSet: true}
}

func (v NullableReviewsReassign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewsReassign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
