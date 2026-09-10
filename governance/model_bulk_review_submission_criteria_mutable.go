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

// checks if the BulkReviewSubmissionCriteriaMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReviewSubmissionCriteriaMutable{}

// BulkReviewSubmissionCriteriaMutable The criteria for selecting reviews
type BulkReviewSubmissionCriteriaMutable struct {
	Type                 BulkReviewCriteriaType                `json:"type"`
	GovAnalyzerCriteria  *BulkReviewCriteriaGovAnalyzerMutable `json:"govAnalyzerCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkReviewSubmissionCriteriaMutable BulkReviewSubmissionCriteriaMutable

// NewBulkReviewSubmissionCriteriaMutable instantiates a new BulkReviewSubmissionCriteriaMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReviewSubmissionCriteriaMutable(type_ BulkReviewCriteriaType) *BulkReviewSubmissionCriteriaMutable {
	this := BulkReviewSubmissionCriteriaMutable{}
	this.Type = type_
	return &this
}

// NewBulkReviewSubmissionCriteriaMutableWithDefaults instantiates a new BulkReviewSubmissionCriteriaMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReviewSubmissionCriteriaMutableWithDefaults() *BulkReviewSubmissionCriteriaMutable {
	this := BulkReviewSubmissionCriteriaMutable{}
	return &this
}

// GetType returns the Type field value
func (o *BulkReviewSubmissionCriteriaMutable) GetType() BulkReviewCriteriaType {
	if o == nil {
		var ret BulkReviewCriteriaType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionCriteriaMutable) GetTypeOk() (*BulkReviewCriteriaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BulkReviewSubmissionCriteriaMutable) SetType(v BulkReviewCriteriaType) {
	o.Type = v
}

// GetGovAnalyzerCriteria returns the GovAnalyzerCriteria field value if set, zero value otherwise.
func (o *BulkReviewSubmissionCriteriaMutable) GetGovAnalyzerCriteria() BulkReviewCriteriaGovAnalyzerMutable {
	if o == nil || IsNil(o.GovAnalyzerCriteria) {
		var ret BulkReviewCriteriaGovAnalyzerMutable
		return ret
	}
	return *o.GovAnalyzerCriteria
}

// GetGovAnalyzerCriteriaOk returns a tuple with the GovAnalyzerCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionCriteriaMutable) GetGovAnalyzerCriteriaOk() (*BulkReviewCriteriaGovAnalyzerMutable, bool) {
	if o == nil || IsNil(o.GovAnalyzerCriteria) {
		return nil, false
	}
	return o.GovAnalyzerCriteria, true
}

// HasGovAnalyzerCriteria returns a boolean if a field has been set.
func (o *BulkReviewSubmissionCriteriaMutable) HasGovAnalyzerCriteria() bool {
	if o != nil && !IsNil(o.GovAnalyzerCriteria) {
		return true
	}

	return false
}

// SetGovAnalyzerCriteria gets a reference to the given BulkReviewCriteriaGovAnalyzerMutable and assigns it to the GovAnalyzerCriteria field.
func (o *BulkReviewSubmissionCriteriaMutable) SetGovAnalyzerCriteria(v BulkReviewCriteriaGovAnalyzerMutable) {
	o.GovAnalyzerCriteria = &v
}

func (o BulkReviewSubmissionCriteriaMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReviewSubmissionCriteriaMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.GovAnalyzerCriteria) {
		toSerialize["govAnalyzerCriteria"] = o.GovAnalyzerCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkReviewSubmissionCriteriaMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varBulkReviewSubmissionCriteriaMutable := _BulkReviewSubmissionCriteriaMutable{}

	err = json.Unmarshal(data, &varBulkReviewSubmissionCriteriaMutable)

	if err != nil {
		return err
	}

	*o = BulkReviewSubmissionCriteriaMutable(varBulkReviewSubmissionCriteriaMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "govAnalyzerCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkReviewSubmissionCriteriaMutable struct {
	value *BulkReviewSubmissionCriteriaMutable
	isSet bool
}

func (v NullableBulkReviewSubmissionCriteriaMutable) Get() *BulkReviewSubmissionCriteriaMutable {
	return v.value
}

func (v *NullableBulkReviewSubmissionCriteriaMutable) Set(val *BulkReviewSubmissionCriteriaMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewSubmissionCriteriaMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewSubmissionCriteriaMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewSubmissionCriteriaMutable(val *BulkReviewSubmissionCriteriaMutable) *NullableBulkReviewSubmissionCriteriaMutable {
	return &NullableBulkReviewSubmissionCriteriaMutable{value: val, isSet: true}
}

func (v NullableBulkReviewSubmissionCriteriaMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewSubmissionCriteriaMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
