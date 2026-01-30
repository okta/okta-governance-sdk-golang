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
	"time"
)

// checks if the SecurityAccessReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReview{}

// SecurityAccessReview struct for SecurityAccessReview
type SecurityAccessReview struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string                     `json:"lastUpdatedBy"`
	Links         *map[string]Link           `json:"_links,omitempty"`
	Status        SecurityAccessReviewStatus `json:"status"`
	// The name of the security access review
	Name string `json:"name"`
	// The end time of the security access review
	EndTime              time.Time                            `json:"endTime"`
	ReviewerSettings     SecurityAccessReviewReviewerSettings `json:"reviewerSettings"`
	Summary              *ServerMessage                       `json:"summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReview SecurityAccessReview

// NewSecurityAccessReview instantiates a new SecurityAccessReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReview(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status SecurityAccessReviewStatus, name string, endTime time.Time, reviewerSettings SecurityAccessReviewReviewerSettings) *SecurityAccessReview {
	this := SecurityAccessReview{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Status = status
	this.Name = name
	this.EndTime = endTime
	this.ReviewerSettings = reviewerSettings
	return &this
}

// NewSecurityAccessReviewWithDefaults instantiates a new SecurityAccessReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewWithDefaults() *SecurityAccessReview {
	this := SecurityAccessReview{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReview) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SecurityAccessReview) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SecurityAccessReview) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *SecurityAccessReview) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SecurityAccessReview) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *SecurityAccessReview) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *SecurityAccessReview) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *SecurityAccessReview) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *SecurityAccessReview) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReview) GetLinks() map[string]Link {
	if o == nil || IsNil(o.Links) {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReview) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *SecurityAccessReview) SetLinks(v map[string]Link) {
	o.Links = &v
}

// GetStatus returns the Status field value
func (o *SecurityAccessReview) GetStatus() SecurityAccessReviewStatus {
	if o == nil {
		var ret SecurityAccessReviewStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetStatusOk() (*SecurityAccessReviewStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SecurityAccessReview) SetStatus(v SecurityAccessReviewStatus) {
	o.Status = v
}

// GetName returns the Name field value
func (o *SecurityAccessReview) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReview) SetName(v string) {
	o.Name = v
}

// GetEndTime returns the EndTime field value
func (o *SecurityAccessReview) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SecurityAccessReview) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *SecurityAccessReview) GetReviewerSettings() SecurityAccessReviewReviewerSettings {
	if o == nil {
		var ret SecurityAccessReviewReviewerSettings
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *SecurityAccessReview) SetReviewerSettings(v SecurityAccessReviewReviewerSettings) {
	o.ReviewerSettings = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *SecurityAccessReview) GetSummary() ServerMessage {
	if o == nil || IsNil(o.Summary) {
		var ret ServerMessage
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReview) GetSummaryOk() (*ServerMessage, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *SecurityAccessReview) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given ServerMessage and assigns it to the Summary field.
func (o *SecurityAccessReview) SetSummary(v ServerMessage) {
	o.Summary = &v
}

func (o SecurityAccessReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["status"] = o.Status
	toSerialize["name"] = o.Name
	toSerialize["endTime"] = o.EndTime
	toSerialize["reviewerSettings"] = o.ReviewerSettings
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"status",
		"name",
		"endTime",
		"reviewerSettings",
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

	varSecurityAccessReview := _SecurityAccessReview{}

	err = json.Unmarshal(data, &varSecurityAccessReview)

	if err != nil {
		return err
	}

	*o = SecurityAccessReview(varSecurityAccessReview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "status")
		delete(additionalProperties, "name")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "reviewerSettings")
		delete(additionalProperties, "summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReview struct {
	value *SecurityAccessReview
	isSet bool
}

func (v NullableSecurityAccessReview) Get() *SecurityAccessReview {
	return v.value
}

func (v *NullableSecurityAccessReview) Set(val *SecurityAccessReview) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReview) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReview(val *SecurityAccessReview) *NullableSecurityAccessReview {
	return &NullableSecurityAccessReview{value: val, isSet: true}
}

func (v NullableSecurityAccessReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
