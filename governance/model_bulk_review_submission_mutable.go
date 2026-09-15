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

// checks if the BulkReviewSubmissionMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReviewSubmissionMutable{}

// BulkReviewSubmissionMutable The request body for submitting bulk decisions for a campaign review. It holds the criteria for which campaign review items to target and the decision to apply to those items. The reviewer level and note are optional fields that can provide additional context for the bulk decision. The reviewer level comes into picture, when multi level campaign is in the picture.
type BulkReviewSubmissionMutable struct {
	Criteria       BulkReviewSubmissionCriteriaMutable `json:"criteria"`
	ReviewDecision BulkReviewDecision                  `json:"reviewDecision"`
	ReviewerLevel  *ReviewerLevelType                  `json:"reviewerLevel,omitempty"`
	// The reviewer's note to include with the bulk decision. This note is added to all reviews that match the bulk-review criteria.
	Notes                *string `json:"notes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkReviewSubmissionMutable BulkReviewSubmissionMutable

// NewBulkReviewSubmissionMutable instantiates a new BulkReviewSubmissionMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReviewSubmissionMutable(criteria BulkReviewSubmissionCriteriaMutable, reviewDecision BulkReviewDecision) *BulkReviewSubmissionMutable {
	this := BulkReviewSubmissionMutable{}
	this.Criteria = criteria
	this.ReviewDecision = reviewDecision
	return &this
}

// NewBulkReviewSubmissionMutableWithDefaults instantiates a new BulkReviewSubmissionMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReviewSubmissionMutableWithDefaults() *BulkReviewSubmissionMutable {
	this := BulkReviewSubmissionMutable{}
	return &this
}

// GetCriteria returns the Criteria field value
func (o *BulkReviewSubmissionMutable) GetCriteria() BulkReviewSubmissionCriteriaMutable {
	if o == nil {
		var ret BulkReviewSubmissionCriteriaMutable
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionMutable) GetCriteriaOk() (*BulkReviewSubmissionCriteriaMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// SetCriteria sets field value
func (o *BulkReviewSubmissionMutable) SetCriteria(v BulkReviewSubmissionCriteriaMutable) {
	o.Criteria = v
}

// GetReviewDecision returns the ReviewDecision field value
func (o *BulkReviewSubmissionMutable) GetReviewDecision() BulkReviewDecision {
	if o == nil {
		var ret BulkReviewDecision
		return ret
	}

	return o.ReviewDecision
}

// GetReviewDecisionOk returns a tuple with the ReviewDecision field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionMutable) GetReviewDecisionOk() (*BulkReviewDecision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewDecision, true
}

// SetReviewDecision sets field value
func (o *BulkReviewSubmissionMutable) SetReviewDecision(v BulkReviewDecision) {
	o.ReviewDecision = v
}

// GetReviewerLevel returns the ReviewerLevel field value if set, zero value otherwise.
func (o *BulkReviewSubmissionMutable) GetReviewerLevel() ReviewerLevelType {
	if o == nil || IsNil(o.ReviewerLevel) {
		var ret ReviewerLevelType
		return ret
	}
	return *o.ReviewerLevel
}

// GetReviewerLevelOk returns a tuple with the ReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionMutable) GetReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || IsNil(o.ReviewerLevel) {
		return nil, false
	}
	return o.ReviewerLevel, true
}

// HasReviewerLevel returns a boolean if a field has been set.
func (o *BulkReviewSubmissionMutable) HasReviewerLevel() bool {
	if o != nil && !IsNil(o.ReviewerLevel) {
		return true
	}

	return false
}

// SetReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the ReviewerLevel field.
func (o *BulkReviewSubmissionMutable) SetReviewerLevel(v ReviewerLevelType) {
	o.ReviewerLevel = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *BulkReviewSubmissionMutable) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionMutable) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *BulkReviewSubmissionMutable) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *BulkReviewSubmissionMutable) SetNotes(v string) {
	o.Notes = &v
}

func (o BulkReviewSubmissionMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReviewSubmissionMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["criteria"] = o.Criteria
	toSerialize["reviewDecision"] = o.ReviewDecision
	if !IsNil(o.ReviewerLevel) {
		toSerialize["reviewerLevel"] = o.ReviewerLevel
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkReviewSubmissionMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"criteria",
		"reviewDecision",
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

	varBulkReviewSubmissionMutable := _BulkReviewSubmissionMutable{}

	err = json.Unmarshal(data, &varBulkReviewSubmissionMutable)

	if err != nil {
		return err
	}

	*o = BulkReviewSubmissionMutable(varBulkReviewSubmissionMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "criteria")
		delete(additionalProperties, "reviewDecision")
		delete(additionalProperties, "reviewerLevel")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkReviewSubmissionMutable struct {
	value *BulkReviewSubmissionMutable
	isSet bool
}

func (v NullableBulkReviewSubmissionMutable) Get() *BulkReviewSubmissionMutable {
	return v.value
}

func (v *NullableBulkReviewSubmissionMutable) Set(val *BulkReviewSubmissionMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewSubmissionMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewSubmissionMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewSubmissionMutable(val *BulkReviewSubmissionMutable) *NullableBulkReviewSubmissionMutable {
	return &NullableBulkReviewSubmissionMutable{value: val, isSet: true}
}

func (v NullableBulkReviewSubmissionMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewSubmissionMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
