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

// checks if the BulkReviewSubmissionJobDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReviewSubmissionJobDetails{}

// BulkReviewSubmissionJobDetails The details of a bulk-review submission job resulting from a request
type BulkReviewSubmissionJobDetails struct {
	// The `id` of the job
	Id     string                       `json:"id"`
	Status BulkReviewDecisionStatusType `json:"status"`
	// A list of errors encountered while processing the bulk review submission job
	Errors []ModelError `json:"errors,omitempty"`
	// The number of eligible reviews in the submission
	EligibleReviewRecordCount int32                      `json:"eligibleReviewRecordCount"`
	Links                     *BulkReviewSubmissionLinks `json:"_links,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _BulkReviewSubmissionJobDetails BulkReviewSubmissionJobDetails

// NewBulkReviewSubmissionJobDetails instantiates a new BulkReviewSubmissionJobDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReviewSubmissionJobDetails(id string, status BulkReviewDecisionStatusType, eligibleReviewRecordCount int32) *BulkReviewSubmissionJobDetails {
	this := BulkReviewSubmissionJobDetails{}
	this.Id = id
	this.Status = status
	this.EligibleReviewRecordCount = eligibleReviewRecordCount
	return &this
}

// NewBulkReviewSubmissionJobDetailsWithDefaults instantiates a new BulkReviewSubmissionJobDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReviewSubmissionJobDetailsWithDefaults() *BulkReviewSubmissionJobDetails {
	this := BulkReviewSubmissionJobDetails{}
	return &this
}

// GetId returns the Id field value
func (o *BulkReviewSubmissionJobDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionJobDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkReviewSubmissionJobDetails) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *BulkReviewSubmissionJobDetails) GetStatus() BulkReviewDecisionStatusType {
	if o == nil {
		var ret BulkReviewDecisionStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionJobDetails) GetStatusOk() (*BulkReviewDecisionStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkReviewSubmissionJobDetails) SetStatus(v BulkReviewDecisionStatusType) {
	o.Status = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkReviewSubmissionJobDetails) GetErrors() []ModelError {
	if o == nil || IsNil(o.Errors) {
		var ret []ModelError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionJobDetails) GetErrorsOk() ([]ModelError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkReviewSubmissionJobDetails) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ModelError and assigns it to the Errors field.
func (o *BulkReviewSubmissionJobDetails) SetErrors(v []ModelError) {
	o.Errors = v
}

// GetEligibleReviewRecordCount returns the EligibleReviewRecordCount field value
func (o *BulkReviewSubmissionJobDetails) GetEligibleReviewRecordCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EligibleReviewRecordCount
}

// GetEligibleReviewRecordCountOk returns a tuple with the EligibleReviewRecordCount field value
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionJobDetails) GetEligibleReviewRecordCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EligibleReviewRecordCount, true
}

// SetEligibleReviewRecordCount sets field value
func (o *BulkReviewSubmissionJobDetails) SetEligibleReviewRecordCount(v int32) {
	o.EligibleReviewRecordCount = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BulkReviewSubmissionJobDetails) GetLinks() BulkReviewSubmissionLinks {
	if o == nil || IsNil(o.Links) {
		var ret BulkReviewSubmissionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReviewSubmissionJobDetails) GetLinksOk() (*BulkReviewSubmissionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BulkReviewSubmissionJobDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BulkReviewSubmissionLinks and assigns it to the Links field.
func (o *BulkReviewSubmissionJobDetails) SetLinks(v BulkReviewSubmissionLinks) {
	o.Links = &v
}

func (o BulkReviewSubmissionJobDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReviewSubmissionJobDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["eligibleReviewRecordCount"] = o.EligibleReviewRecordCount
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkReviewSubmissionJobDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"eligibleReviewRecordCount",
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

	varBulkReviewSubmissionJobDetails := _BulkReviewSubmissionJobDetails{}

	err = json.Unmarshal(data, &varBulkReviewSubmissionJobDetails)

	if err != nil {
		return err
	}

	*o = BulkReviewSubmissionJobDetails(varBulkReviewSubmissionJobDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "eligibleReviewRecordCount")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkReviewSubmissionJobDetails struct {
	value *BulkReviewSubmissionJobDetails
	isSet bool
}

func (v NullableBulkReviewSubmissionJobDetails) Get() *BulkReviewSubmissionJobDetails {
	return v.value
}

func (v *NullableBulkReviewSubmissionJobDetails) Set(val *BulkReviewSubmissionJobDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewSubmissionJobDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewSubmissionJobDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewSubmissionJobDetails(val *BulkReviewSubmissionJobDetails) *NullableBulkReviewSubmissionJobDetails {
	return &NullableBulkReviewSubmissionJobDetails{value: val, isSet: true}
}

func (v NullableBulkReviewSubmissionJobDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewSubmissionJobDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
