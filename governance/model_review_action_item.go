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

// checks if the ReviewActionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewActionItem{}

// ReviewActionItem A review with the action to apply to it
type ReviewActionItem struct {
	// Review ID
	ReviewId string       `json:"reviewId"`
	Action   ReviewAction `json:"action"`
	// The reassigned reviewer's Okta user ID Required when `action` is `REASSIGN`. All items with `action` set to `REASSIGN` in a single request must specify the same `reviewerId`.
	ReviewerId *string `json:"reviewerId,omitempty"`
	// Optional note to justify the action taken on this review
	Note                 NullableString `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewActionItem ReviewActionItem

// NewReviewActionItem instantiates a new ReviewActionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewActionItem(reviewId string, action ReviewAction) *ReviewActionItem {
	this := ReviewActionItem{}
	this.ReviewId = reviewId
	this.Action = action
	return &this
}

// NewReviewActionItemWithDefaults instantiates a new ReviewActionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewActionItemWithDefaults() *ReviewActionItem {
	this := ReviewActionItem{}
	return &this
}

// GetReviewId returns the ReviewId field value
func (o *ReviewActionItem) GetReviewId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewId
}

// GetReviewIdOk returns a tuple with the ReviewId field value
// and a boolean to check if the value has been set.
func (o *ReviewActionItem) GetReviewIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewId, true
}

// SetReviewId sets field value
func (o *ReviewActionItem) SetReviewId(v string) {
	o.ReviewId = v
}

// GetAction returns the Action field value
func (o *ReviewActionItem) GetAction() ReviewAction {
	if o == nil {
		var ret ReviewAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ReviewActionItem) GetActionOk() (*ReviewAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ReviewActionItem) SetAction(v ReviewAction) {
	o.Action = v
}

// GetReviewerId returns the ReviewerId field value if set, zero value otherwise.
func (o *ReviewActionItem) GetReviewerId() string {
	if o == nil || IsNil(o.ReviewerId) {
		var ret string
		return ret
	}
	return *o.ReviewerId
}

// GetReviewerIdOk returns a tuple with the ReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewActionItem) GetReviewerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerId) {
		return nil, false
	}
	return o.ReviewerId, true
}

// HasReviewerId returns a boolean if a field has been set.
func (o *ReviewActionItem) HasReviewerId() bool {
	if o != nil && !IsNil(o.ReviewerId) {
		return true
	}

	return false
}

// SetReviewerId gets a reference to the given string and assigns it to the ReviewerId field.
func (o *ReviewActionItem) SetReviewerId(v string) {
	o.ReviewerId = &v
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewActionItem) GetNote() string {
	if o == nil || IsNil(o.Note.Get()) {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewActionItem) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *ReviewActionItem) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *ReviewActionItem) SetNote(v string) {
	o.Note.Set(&v)
}

// SetNoteNil sets the value for Note to be an explicit nil
func (o *ReviewActionItem) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *ReviewActionItem) UnsetNote() {
	o.Note.Unset()
}

func (o ReviewActionItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewActionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviewId"] = o.ReviewId
	toSerialize["action"] = o.Action
	if !IsNil(o.ReviewerId) {
		toSerialize["reviewerId"] = o.ReviewerId
	}
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewActionItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviewId",
		"action",
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

	varReviewActionItem := _ReviewActionItem{}

	err = json.Unmarshal(data, &varReviewActionItem)

	if err != nil {
		return err
	}

	*o = ReviewActionItem(varReviewActionItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewId")
		delete(additionalProperties, "action")
		delete(additionalProperties, "reviewerId")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewActionItem struct {
	value *ReviewActionItem
	isSet bool
}

func (v NullableReviewActionItem) Get() *ReviewActionItem {
	return v.value
}

func (v *NullableReviewActionItem) Set(val *ReviewActionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewActionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewActionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewActionItem(val *ReviewActionItem) *NullableReviewActionItem {
	return &NullableReviewActionItem{value: val, isSet: true}
}

func (v NullableReviewActionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewActionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
