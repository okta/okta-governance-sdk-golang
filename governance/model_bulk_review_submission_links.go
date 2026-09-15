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

// checks if the BulkReviewSubmissionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReviewSubmissionLinks{}

// BulkReviewSubmissionLinks Links available for a bulk-review submission
type BulkReviewSubmissionLinks struct {
	Self                     Link  `json:"self"`
	BulkReviewDecisionStatus *Link `json:"bulkReviewDecisionStatus,omitempty"`
	ReviewRecordsProcessed   *Link `json:"reviewRecordsProcessed,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _BulkReviewSubmissionLinks BulkReviewSubmissionLinks

// NewBulkReviewSubmissionLinks instantiates a new BulkReviewSubmissionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReviewSubmissionLinks(self Link) *BulkReviewSubmissionLinks {
	this := BulkReviewSubmissionLinks{}
	this.Self = self
	return &this
}

// NewBulkReviewSubmissionLinksWithDefaults instantiates a new BulkReviewSubmissionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReviewSubmissionLinksWithDefaults() *BulkReviewSubmissionLinks {
	this := BulkReviewSubmissionLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *BulkReviewSubmissionLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *BulkReviewSubmissionLinks) SetSelf(v Link) {
	o.Self = v
}

// GetBulkReviewDecisionStatus returns the BulkReviewDecisionStatus field value if set, zero value otherwise.
func (o *BulkReviewSubmissionLinks) GetBulkReviewDecisionStatus() Link {
	if o == nil || IsNil(o.BulkReviewDecisionStatus) {
		var ret Link
		return ret
	}
	return *o.BulkReviewDecisionStatus
}

// GetBulkReviewDecisionStatusOk returns a tuple with the BulkReviewDecisionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionLinks) GetBulkReviewDecisionStatusOk() (*Link, bool) {
	if o == nil || IsNil(o.BulkReviewDecisionStatus) {
		return nil, false
	}
	return o.BulkReviewDecisionStatus, true
}

// HasBulkReviewDecisionStatus returns a boolean if a field has been set.
func (o *BulkReviewSubmissionLinks) HasBulkReviewDecisionStatus() bool {
	if o != nil && !IsNil(o.BulkReviewDecisionStatus) {
		return true
	}

	return false
}

// SetBulkReviewDecisionStatus gets a reference to the given Link and assigns it to the BulkReviewDecisionStatus field.
func (o *BulkReviewSubmissionLinks) SetBulkReviewDecisionStatus(v Link) {
	o.BulkReviewDecisionStatus = &v
}

// GetReviewRecordsProcessed returns the ReviewRecordsProcessed field value if set, zero value otherwise.
func (o *BulkReviewSubmissionLinks) GetReviewRecordsProcessed() Link {
	if o == nil || IsNil(o.ReviewRecordsProcessed) {
		var ret Link
		return ret
	}
	return *o.ReviewRecordsProcessed
}

// GetReviewRecordsProcessedOk returns a tuple with the ReviewRecordsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionLinks) GetReviewRecordsProcessedOk() (*Link, bool) {
	if o == nil || IsNil(o.ReviewRecordsProcessed) {
		return nil, false
	}
	return o.ReviewRecordsProcessed, true
}

// HasReviewRecordsProcessed returns a boolean if a field has been set.
func (o *BulkReviewSubmissionLinks) HasReviewRecordsProcessed() bool {
	if o != nil && !IsNil(o.ReviewRecordsProcessed) {
		return true
	}

	return false
}

// SetReviewRecordsProcessed gets a reference to the given Link and assigns it to the ReviewRecordsProcessed field.
func (o *BulkReviewSubmissionLinks) SetReviewRecordsProcessed(v Link) {
	o.ReviewRecordsProcessed = &v
}

func (o BulkReviewSubmissionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReviewSubmissionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.BulkReviewDecisionStatus) {
		toSerialize["bulkReviewDecisionStatus"] = o.BulkReviewDecisionStatus
	}
	if !IsNil(o.ReviewRecordsProcessed) {
		toSerialize["reviewRecordsProcessed"] = o.ReviewRecordsProcessed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkReviewSubmissionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varBulkReviewSubmissionLinks := _BulkReviewSubmissionLinks{}

	err = json.Unmarshal(data, &varBulkReviewSubmissionLinks)

	if err != nil {
		return err
	}

	*o = BulkReviewSubmissionLinks(varBulkReviewSubmissionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "bulkReviewDecisionStatus")
		delete(additionalProperties, "reviewRecordsProcessed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkReviewSubmissionLinks struct {
	value *BulkReviewSubmissionLinks
	isSet bool
}

func (v NullableBulkReviewSubmissionLinks) Get() *BulkReviewSubmissionLinks {
	return v.value
}

func (v *NullableBulkReviewSubmissionLinks) Set(val *BulkReviewSubmissionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewSubmissionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewSubmissionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewSubmissionLinks(val *BulkReviewSubmissionLinks) *NullableBulkReviewSubmissionLinks {
	return &NullableBulkReviewSubmissionLinks{value: val, isSet: true}
}

func (v NullableBulkReviewSubmissionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewSubmissionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
