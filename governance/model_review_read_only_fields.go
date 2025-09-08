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

// ReviewReadOnlyFields Additional read-only attributes for reviews
type ReviewReadOnlyFields struct {
	Note                 *Note                 `json:"note,omitempty"`
	ReviewerGroupProfile *ReviewerGroupProfile `json:"reviewerGroupProfile,omitempty"`
	// Applicable only for multi level campaign. Provides details about the reviewer and decisions (if any) made at each reviewer level is captured here.
	AllReviewerLevels    []ReviewerLevelInfoFull `json:"allReviewerLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewReadOnlyFields ReviewReadOnlyFields

// NewReviewReadOnlyFields instantiates a new ReviewReadOnlyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewReadOnlyFields() *ReviewReadOnlyFields {
	this := ReviewReadOnlyFields{}
	return &this
}

// NewReviewReadOnlyFieldsWithDefaults instantiates a new ReviewReadOnlyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewReadOnlyFieldsWithDefaults() *ReviewReadOnlyFields {
	this := ReviewReadOnlyFields{}
	return &this
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ReviewReadOnlyFields) GetNote() Note {
	if o == nil || o.Note == nil {
		var ret Note
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewReadOnlyFields) GetNoteOk() (*Note, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ReviewReadOnlyFields) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given Note and assigns it to the Note field.
func (o *ReviewReadOnlyFields) SetNote(v Note) {
	o.Note = &v
}

// GetReviewerGroupProfile returns the ReviewerGroupProfile field value if set, zero value otherwise.
func (o *ReviewReadOnlyFields) GetReviewerGroupProfile() ReviewerGroupProfile {
	if o == nil || o.ReviewerGroupProfile == nil {
		var ret ReviewerGroupProfile
		return ret
	}
	return *o.ReviewerGroupProfile
}

// GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewReadOnlyFields) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool) {
	if o == nil || o.ReviewerGroupProfile == nil {
		return nil, false
	}
	return o.ReviewerGroupProfile, true
}

// HasReviewerGroupProfile returns a boolean if a field has been set.
func (o *ReviewReadOnlyFields) HasReviewerGroupProfile() bool {
	if o != nil && o.ReviewerGroupProfile != nil {
		return true
	}

	return false
}

// SetReviewerGroupProfile gets a reference to the given ReviewerGroupProfile and assigns it to the ReviewerGroupProfile field.
func (o *ReviewReadOnlyFields) SetReviewerGroupProfile(v ReviewerGroupProfile) {
	o.ReviewerGroupProfile = &v
}

// GetAllReviewerLevels returns the AllReviewerLevels field value if set, zero value otherwise.
func (o *ReviewReadOnlyFields) GetAllReviewerLevels() []ReviewerLevelInfoFull {
	if o == nil || o.AllReviewerLevels == nil {
		var ret []ReviewerLevelInfoFull
		return ret
	}
	return o.AllReviewerLevels
}

// GetAllReviewerLevelsOk returns a tuple with the AllReviewerLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewReadOnlyFields) GetAllReviewerLevelsOk() ([]ReviewerLevelInfoFull, bool) {
	if o == nil || o.AllReviewerLevels == nil {
		return nil, false
	}
	return o.AllReviewerLevels, true
}

// HasAllReviewerLevels returns a boolean if a field has been set.
func (o *ReviewReadOnlyFields) HasAllReviewerLevels() bool {
	if o != nil && o.AllReviewerLevels != nil {
		return true
	}

	return false
}

// SetAllReviewerLevels gets a reference to the given []ReviewerLevelInfoFull and assigns it to the AllReviewerLevels field.
func (o *ReviewReadOnlyFields) SetAllReviewerLevels(v []ReviewerLevelInfoFull) {
	o.AllReviewerLevels = v
}

func (o ReviewReadOnlyFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.ReviewerGroupProfile != nil {
		toSerialize["reviewerGroupProfile"] = o.ReviewerGroupProfile
	}
	if o.AllReviewerLevels != nil {
		toSerialize["allReviewerLevels"] = o.AllReviewerLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewReadOnlyFields) UnmarshalJSON(bytes []byte) (err error) {
	varReviewReadOnlyFields := _ReviewReadOnlyFields{}

	err = json.Unmarshal(bytes, &varReviewReadOnlyFields)
	if err == nil {
		*o = ReviewReadOnlyFields(varReviewReadOnlyFields)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "note")
		delete(additionalProperties, "reviewerGroupProfile")
		delete(additionalProperties, "allReviewerLevels")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewReadOnlyFields struct {
	value *ReviewReadOnlyFields
	isSet bool
}

func (v NullableReviewReadOnlyFields) Get() *ReviewReadOnlyFields {
	return v.value
}

func (v *NullableReviewReadOnlyFields) Set(val *ReviewReadOnlyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewReadOnlyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewReadOnlyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewReadOnlyFields(val *ReviewReadOnlyFields) *NullableReviewReadOnlyFields {
	return &NullableReviewReadOnlyFields{value: val, isSet: true}
}

func (v NullableReviewReadOnlyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewReadOnlyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
