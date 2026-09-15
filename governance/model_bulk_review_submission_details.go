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

// checks if the BulkReviewSubmissionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReviewSubmissionDetails{}

// BulkReviewSubmissionDetails The details of a bulk review submission resulting from a request
type BulkReviewSubmissionDetails struct {
	Status BulkReviewDecisionStatusType `json:"status"`
	// The number of eligible reviews in the submission
	EligibleReviewRecordCount int32 `json:"eligibleReviewRecordCount"`
	// The `id` of the job
	JobId                string                     `json:"jobId"`
	Links                *BulkReviewSubmissionLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkReviewSubmissionDetails BulkReviewSubmissionDetails

// NewBulkReviewSubmissionDetails instantiates a new BulkReviewSubmissionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReviewSubmissionDetails(status BulkReviewDecisionStatusType, eligibleReviewRecordCount int32, jobId string) *BulkReviewSubmissionDetails {
	this := BulkReviewSubmissionDetails{}
	this.Status = status
	this.EligibleReviewRecordCount = eligibleReviewRecordCount
	this.JobId = jobId
	return &this
}

// NewBulkReviewSubmissionDetailsWithDefaults instantiates a new BulkReviewSubmissionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReviewSubmissionDetailsWithDefaults() *BulkReviewSubmissionDetails {
	this := BulkReviewSubmissionDetails{}
	return &this
}

// GetStatus returns the Status field value
func (o *BulkReviewSubmissionDetails) GetStatus() BulkReviewDecisionStatusType {
	if o == nil {
		var ret BulkReviewDecisionStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionDetails) GetStatusOk() (*BulkReviewDecisionStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkReviewSubmissionDetails) SetStatus(v BulkReviewDecisionStatusType) {
	o.Status = v
}

// GetEligibleReviewRecordCount returns the EligibleReviewRecordCount field value
func (o *BulkReviewSubmissionDetails) GetEligibleReviewRecordCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EligibleReviewRecordCount
}

// GetEligibleReviewRecordCountOk returns a tuple with the EligibleReviewRecordCount field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionDetails) GetEligibleReviewRecordCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EligibleReviewRecordCount, true
}

// SetEligibleReviewRecordCount sets field value
func (o *BulkReviewSubmissionDetails) SetEligibleReviewRecordCount(v int32) {
	o.EligibleReviewRecordCount = v
}

// GetJobId returns the JobId field value
func (o *BulkReviewSubmissionDetails) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionDetails) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *BulkReviewSubmissionDetails) SetJobId(v string) {
	o.JobId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BulkReviewSubmissionDetails) GetLinks() BulkReviewSubmissionLinks {
	if o == nil || IsNil(o.Links) {
		var ret BulkReviewSubmissionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionDetails) GetLinksOk() (*BulkReviewSubmissionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BulkReviewSubmissionDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BulkReviewSubmissionLinks and assigns it to the Links field.
func (o *BulkReviewSubmissionDetails) SetLinks(v BulkReviewSubmissionLinks) {
	o.Links = &v
}

func (o BulkReviewSubmissionDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReviewSubmissionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["eligibleReviewRecordCount"] = o.EligibleReviewRecordCount
	toSerialize["jobId"] = o.JobId
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkReviewSubmissionDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"eligibleReviewRecordCount",
		"jobId",
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

	varBulkReviewSubmissionDetails := _BulkReviewSubmissionDetails{}

	err = json.Unmarshal(data, &varBulkReviewSubmissionDetails)

	if err != nil {
		return err
	}

	*o = BulkReviewSubmissionDetails(varBulkReviewSubmissionDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "eligibleReviewRecordCount")
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkReviewSubmissionDetails struct {
	value *BulkReviewSubmissionDetails
	isSet bool
}

func (v NullableBulkReviewSubmissionDetails) Get() *BulkReviewSubmissionDetails {
	return v.value
}

func (v *NullableBulkReviewSubmissionDetails) Set(val *BulkReviewSubmissionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewSubmissionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewSubmissionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewSubmissionDetails(val *BulkReviewSubmissionDetails) *NullableBulkReviewSubmissionDetails {
	return &NullableBulkReviewSubmissionDetails{value: val, isSet: true}
}

func (v NullableBulkReviewSubmissionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewSubmissionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
