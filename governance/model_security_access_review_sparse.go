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
	"time"
)

// SecurityAccessReviewSparse struct for SecurityAccessReviewSparse
type SecurityAccessReviewSparse struct {
	// The ID of the security access review
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
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewSparse SecurityAccessReviewSparse

// NewSecurityAccessReviewSparse instantiates a new SecurityAccessReviewSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSparse(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status SecurityAccessReviewStatus, name string, endTime time.Time, reviewerSettings SecurityAccessReviewReviewerSettings) *SecurityAccessReviewSparse {
	this := SecurityAccessReviewSparse{}
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

// NewSecurityAccessReviewSparseWithDefaults instantiates a new SecurityAccessReviewSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSparseWithDefaults() *SecurityAccessReviewSparse {
	this := SecurityAccessReviewSparse{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SecurityAccessReviewSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SecurityAccessReviewSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *SecurityAccessReviewSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SecurityAccessReviewSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *SecurityAccessReviewSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *SecurityAccessReviewSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *SecurityAccessReviewSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *SecurityAccessReviewSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviewSparse) GetLinks() map[string]Link {
	if o == nil || o.Links == nil {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewSparse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *SecurityAccessReviewSparse) SetLinks(v map[string]Link) {
	o.Links = &v
}

// GetStatus returns the Status field value
func (o *SecurityAccessReviewSparse) GetStatus() SecurityAccessReviewStatus {
	if o == nil {
		var ret SecurityAccessReviewStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetStatusOk() (*SecurityAccessReviewStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SecurityAccessReviewSparse) SetStatus(v SecurityAccessReviewStatus) {
	o.Status = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewSparse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewSparse) SetName(v string) {
	o.Name = v
}

// GetEndTime returns the EndTime field value
func (o *SecurityAccessReviewSparse) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SecurityAccessReviewSparse) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *SecurityAccessReviewSparse) GetReviewerSettings() SecurityAccessReviewReviewerSettings {
	if o == nil {
		var ret SecurityAccessReviewReviewerSettings
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSparse) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *SecurityAccessReviewSparse) SetReviewerSettings(v SecurityAccessReviewReviewerSettings) {
	o.ReviewerSettings = v
}

func (o SecurityAccessReviewSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["reviewerSettings"] = o.ReviewerSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewSparse) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewSparse := _SecurityAccessReviewSparse{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewSparse)
	if err == nil {
		*o = SecurityAccessReviewSparse(varSecurityAccessReviewSparse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewSparse struct {
	value *SecurityAccessReviewSparse
	isSet bool
}

func (v NullableSecurityAccessReviewSparse) Get() *SecurityAccessReviewSparse {
	return v.value
}

func (v *NullableSecurityAccessReviewSparse) Set(val *SecurityAccessReviewSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSparse(val *SecurityAccessReviewSparse) *NullableSecurityAccessReviewSparse {
	return &NullableSecurityAccessReviewSparse{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
