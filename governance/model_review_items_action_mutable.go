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

// checks if the ReviewItemsActionMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewItemsActionMutable{}

// ReviewItemsActionMutable The request body to submit actions for a collection of reviews assigned to the current reviewer in a campaign. You can submit up to 50 review actions (in a combination of `APPROVE`, `REVOKE`, or `REASSIGN`) in a single request. For `REASSIGN` actions, you must also include the new reviewer's ID (`reviewerId`). You can only reassign reviews to the same reviewer in a single request. That is, all `REASSIGN` actions must target the same `reviewerId` in a request.
type ReviewItemsActionMutable struct {
	// List of review actions
	Reviews              []ReviewActionItem `json:"reviews"`
	ReviewerLevel        *ReviewerLevelType `json:"reviewerLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewItemsActionMutable ReviewItemsActionMutable

// NewReviewItemsActionMutable instantiates a new ReviewItemsActionMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewItemsActionMutable(reviews []ReviewActionItem) *ReviewItemsActionMutable {
	this := ReviewItemsActionMutable{}
	this.Reviews = reviews
	return &this
}

// NewReviewItemsActionMutableWithDefaults instantiates a new ReviewItemsActionMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewItemsActionMutableWithDefaults() *ReviewItemsActionMutable {
	this := ReviewItemsActionMutable{}
	return &this
}

// GetReviews returns the Reviews field value
func (o *ReviewItemsActionMutable) GetReviews() []ReviewActionItem {
	if o == nil {
		var ret []ReviewActionItem
		return ret
	}

	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value
// and a boolean to check if the value has been set.
func (o *ReviewItemsActionMutable) GetReviewsOk() ([]ReviewActionItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reviews, true
}

// SetReviews sets field value
func (o *ReviewItemsActionMutable) SetReviews(v []ReviewActionItem) {
	o.Reviews = v
}

// GetReviewerLevel returns the ReviewerLevel field value if set, zero value otherwise.
func (o *ReviewItemsActionMutable) GetReviewerLevel() ReviewerLevelType {
	if o == nil || IsNil(o.ReviewerLevel) {
		var ret ReviewerLevelType
		return ret
	}
	return *o.ReviewerLevel
}

// GetReviewerLevelOk returns a tuple with the ReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewItemsActionMutable) GetReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || IsNil(o.ReviewerLevel) {
		return nil, false
	}
	return o.ReviewerLevel, true
}

// HasReviewerLevel returns a boolean if a field has been set.
func (o *ReviewItemsActionMutable) HasReviewerLevel() bool {
	if o != nil && !IsNil(o.ReviewerLevel) {
		return true
	}

	return false
}

// SetReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the ReviewerLevel field.
func (o *ReviewItemsActionMutable) SetReviewerLevel(v ReviewerLevelType) {
	o.ReviewerLevel = &v
}

func (o ReviewItemsActionMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewItemsActionMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviews"] = o.Reviews
	if !IsNil(o.ReviewerLevel) {
		toSerialize["reviewerLevel"] = o.ReviewerLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewItemsActionMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviews",
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

	varReviewItemsActionMutable := _ReviewItemsActionMutable{}

	err = json.Unmarshal(data, &varReviewItemsActionMutable)

	if err != nil {
		return err
	}

	*o = ReviewItemsActionMutable(varReviewItemsActionMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reviews")
		delete(additionalProperties, "reviewerLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewItemsActionMutable struct {
	value *ReviewItemsActionMutable
	isSet bool
}

func (v NullableReviewItemsActionMutable) Get() *ReviewItemsActionMutable {
	return v.value
}

func (v *NullableReviewItemsActionMutable) Set(val *ReviewItemsActionMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewItemsActionMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewItemsActionMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewItemsActionMutable(val *ReviewItemsActionMutable) *NullableReviewItemsActionMutable {
	return &NullableReviewItemsActionMutable{value: val, isSet: true}
}

func (v NullableReviewItemsActionMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewItemsActionMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
